package memos

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID        int32  `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	OpenID    string `json:"openId"`
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

// CreateUser create a user with username and random password
func (c *Client) CreateUser(username string) (*User, error) {
	param := map[string]any{
		"role":     "USER",
		"username": username,
		"nickname": username,
		"password": uuid.New().String(),
	}

	var user User
	err := c.request("POST", "/api/v1/user", nil, param, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UserList reset user's password into a random UUID
func (c *Client) ResetUserOpenId(userId int32) (*User, error) {
	param := map[string]any{"resetOpenId": true}

	var user User
	err := c.request("PATCH", fmt.Sprintf("/api/v1/user/%d", userId), nil, param, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
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
