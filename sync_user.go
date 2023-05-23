package main

import (
	"discovery/internal/memos"
	"log"
	"strings"
)

// registerOrGetUser register user in our *memos*, who's username is the targetUrl. Then fetch the user's OpenID
func registerOrGetUsers(addr, openid string, targetUrls []string) (map[string]memos.User, error) {
	svr := memos.New(addr, openid)

	// First, fetch all users
	allUsers, err := svr.UserList()
	if err != nil {
		return nil, err
	}

	// Secondly, filter user with *http* as username's prefix
	usersForTargets := make(map[string]memos.User, len(targetUrls))
	for _, u := range allUsers {
		if !strings.HasPrefix(u.Username, "http") {
			continue
		}
		usersForTargets[u.Username] = u
	}

	// At last, fill all user with OpenID
	for _, targetUrl := range targetUrls {
		// For exists user, there's no API to fetch OpenID.
		// So just regenerate and get it.
		if u, ok := usersForTargets[targetUrl]; ok {
			log.Printf("Regenerate openID for %s", targetUrl)
			u, err := svr.ResetUserOpenId(u.ID)
			if err != nil {
				return nil, err
			}

			// replace user with OpenID filled
			usersForTargets[targetUrl] = *u
			continue
		}

		// For user not exists, just register a new user
		log.Printf("Create user for %s", targetUrl)
		u, err := svr.CreateUser(targetUrl)
		if err != nil {
			return nil, err
		}

		// New user already has OpenID filled
		usersForTargets[targetUrl] = *u
	}

	return usersForTargets, nil
}
