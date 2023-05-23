package main

import (
	"discovery/internal/memos"
	"fmt"
	"log"
)

// TODO:
//
//	Replace maxPerLoop with start memo.ID until the API support it
//
// Max sync memos count
var lastCheckTs int64

const maxPerLoop int = 10

// syncTargetToUser sync target memos server into user's memos
func syncTargetToUser(targetUrl string, u memos.User) error {
	log.Printf("Sync %s to user %d with %s", targetUrl, u.ID, u.OpenID)

	hostSvr := memos.New(addr, u.OpenID)
	srcSvr := memos.New(targetUrl, "")

	memos, err := srcSvr.MemoList(0, maxPerLoop)
	if err != nil {
		return fmt.Errorf("fail to fetch memos: %s", err)
	}

	var skipCount, resourceCount int

	for _, memo := range memos {
		// Skip memos already synced
		if memo.CreatedTs < lastCheckTs {
			skipCount += 1
			continue
		}

		// Create resources first
		resourceIds := make([]int, 0, len(memo.ResourceList))
		for _, resource := range memo.ResourceList {
			link := srcSvr.ResourceLink(*resource)
			newResource, err := hostSvr.CreateExternalLinkResource(link, resource.Filename, resource.Type)
			if err != nil {
				return fmt.Errorf("fail to create resource for memo %v of %s: %s", memo, targetUrl, err)
			}

			resourceIds = append(resourceIds, newResource.ID)
		}

		resourceCount += len(resourceIds)

		_, err = hostSvr.CreateMemo(memo.Content, memo.CreatedTs, resourceIds)
		if err != nil {
			return fmt.Errorf("fail to create memo %v of %s: %s", memo, targetUrl, err)
		}
	}

	log.Printf("Total %d memos synced %d with %d resources", len(memos), len(memos)-skipCount, resourceCount)

	return nil
}
