package main

import (
	"fmt"
	"log"
	"memosync/internal/memos"
	"net/url"
	"strconv"
	"strings"
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
	srcUrl, err := url.Parse(u.Username)
	if err != nil {
		return fmt.Errorf("fail to url.Parse %s: %s", u.Username, err)
	}

	// only one user
	var srcUserID int
	if srcUrl.Path != "/" && srcUrl.Path != "" {
		if !strings.HasPrefix(srcUrl.Path, "/u/") {
			return fmt.Errorf("invalid url %s: path should empty or start with /u/", u.Username)
		}
		userIDStr := strings.TrimPrefix(srcUrl.Path, "/u/")
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return fmt.Errorf("fail to parse userid from %s:%s", userIDStr, err)
		}
		srcUserID = userID
	}

	srcAddr := srcUrl.Scheme + "://" + srcUrl.Host
	srcSvr := memos.New(srcAddr, "")

	var allMemos []memos.Memo
	if srcUserID > 0 {
		log.Printf("Sync UserID=%d for user=%d of %s", u.ID, srcUserID, srcAddr)
		allMemos, err = srcSvr.UserMemoList(srcUserID, 0, maxPerLoop)
	} else {
		log.Printf("Sync UserID=%d for all user of %s", u.ID, srcAddr)
		allMemos, err = srcSvr.MemoList(0, maxPerLoop)
	}
	if err != nil {
		return fmt.Errorf("fail to fetch memos: %s", err)
	}

	var skipCount, resourceCount int

	hostSvr := memos.New(addr, u.OpenID)
	for _, memo := range allMemos {
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

	log.Printf("Total %d memos, create %d memos with %d resources", len(allMemos), len(allMemos)-skipCount, resourceCount)

	return nil
}
