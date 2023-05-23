package main

import (
	"log"
	"memosync/internal/memos"
	"strings"
)

// resetAndGetUsers retrive all users who's username start with *http* and reset OpenID of the them
func resetAndGetUsers(addr, openid string) ([]memos.User, error) {
	svr := memos.New(addr, openid)

	users, err := svr.UserList()
	if err != nil {
		return nil, err
	}

	targets := make([]memos.User, 0, len(users))
	for _, u := range users {
		if !strings.HasPrefix(u.Username, "http") {
			continue
		}

		log.Printf("Regenerate OpenID for %d (%s)", u.ID, u.Username)
		u, err := svr.ResetUserOpenId(u.ID)
		if err != nil {
			return nil, err
		}

		targets = append(targets, *u)
	}

	return targets, nil
}
