package main

import (
	"discovery/internal/memos"
	"log"
	"strings"
)

// registerOrGetUser register user in our *memos*, who's username is the targetUrl. Then fetch the user's OpenID
func registerTargets(addr, openid string, targetUrls []string) error {
	svr := memos.New(addr, openid)

	users, err := svr.UserList()
	if err != nil {
		return err
	}

	userMap := make(map[string]struct{}, len(users))
	for _, u := range users {
		userMap[u.Username] = struct{}{}
	}

	// At last, fill all user with OpenID
	for _, targetUrl := range targetUrls {
		_, ok := userMap[targetUrl]
		if ok {
			log.Printf("Skip for %s", targetUrl)
			continue
		}

		// For user not exists, just register a new user
		log.Printf("Create for %s", targetUrl)
		_, err := svr.CreateUser(targetUrl)
		if err != nil {
			return err
		}
	}

	return nil
}

// resetOpenIdAndFetchTargetUsers retrive all users who's username start with *http* and reset OpenID of the user
func resetOpenIdAndFetchTargetUsers(addr, openid string) ([]memos.User, error) {
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

		log.Printf("Regenerate openID for %d (%s)", u.ID, u.Username)
		u, err := svr.ResetUserOpenId(u.ID)
		if err != nil {
			return nil, err
		}

		targets = append(targets, *u)
	}

	return targets, nil
}
