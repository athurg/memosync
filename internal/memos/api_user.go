package memos

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/usememos/memos/api"
)

type User api.User

// UserList fetch user list from memos server
func (c *Client) UserList() ([]User, error) {
	users := []User{}
	err := c.request("GET", "/api/user", nil, nil, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// CreateUser create a user with username and random password
func (c *Client) CreateUser(username string) (*User, error) {
	param := map[string]any{
		"role":     api.NormalUser,
		"username": username,
		"password": uuid.New().String(),
	}

	var user User
	err := c.request("POST", "/api/user", nil, param, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UserList reset user's password into a random UUID
func (c *Client) ResetUserOpenId(userId int) (*User, error) {
	param := map[string]any{"resetOpenId": true}

	var user User
	err := c.request("PATCH", fmt.Sprintf("/api/user/%d", userId), nil, param, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
