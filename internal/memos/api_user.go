package memos

import (
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
