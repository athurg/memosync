package main

import (
	"fmt"
	"log"
	"memosync/internal/memos"
)

// TODO:
//
//	Replace maxPerLoop with start memo.ID until the API support it
//
// Max sync memos count
const maxPerLoop int = 10

var lastCheckTs int64

func syncMemos() {
	users, err := resetAndGetUsers(addr, openid)
	if err != nil {
		log.Printf("fail to resetOpenIdAndGetUsers: %s", err)
		return
	}

	for _, user := range users {
		err := syncTargetToUser(user)
		if err != nil {
			log.Printf("fail to syncTargetToUser: %s", err)
		}
	}
}

// syncTargetToUser sync target memos server into user's memos
func syncTargetToUser(u memos.User) error {
	log.Printf("Sync UserID=%d for %s", u.ID, u.Username)

	hostSvr := memos.New(addr, u.OpenID)
	srcSvr := memos.New(u.Username, "")

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
				return fmt.Errorf("fail to create resource for memo %v of %s: %s", memo, u.Username, err)
			}

			resourceIds = append(resourceIds, newResource.ID)
		}

		resourceCount += len(resourceIds)

		_, err = hostSvr.CreateMemo(memo.Content, memo.CreatedTs, resourceIds)
		if err != nil {
			return fmt.Errorf("fail to create memo %v of %s: %s", memo, u.Username, err)
		}
	}

	log.Printf("Total %d memos, create %d memos with %d resources", len(memos), len(memos)-skipCount, resourceCount)

	return nil
}
