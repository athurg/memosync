package memos

import (
	"fmt"
)

type User struct {
	ID        int32  `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	RowStatus string `json:"rowStatus"`
}

// UserList fetch user list from memos server
func (c *Client) UserList() ([]User, error) {
	users := []User{}
	err := c.request("GET", "/api/v1/user", nil, nil, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// FetchUserInfo fetch user info
func (c *Client) FetchUserInfo(userId int32) (*User, error) {
	var user User
	err := c.request("GET", fmt.Sprintf("/api/v1/user/%d", userId), nil, nil, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FetchUserInfo fetch user info
func (c *Client) UpdateUserNickname(userId int32, nickname string) (*User, error) {
	param := map[string]any{"nickname": nickname}

	var user User
	err := c.request("PATCH", fmt.Sprintf("/api/v1/user/%d", userId), nil, param, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
