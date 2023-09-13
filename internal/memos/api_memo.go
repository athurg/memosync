package memos

import (
	"net/url"
	"strconv"

	"github.com/usememos/memos/api"
)

type (
	Memo       = api.MemoResponse
	MemoCreate = api.CreateMemoRequest
)

// MemoList fetch memo list from memos server
func (c *Client) MemoList(offset, limit int) ([]Memo, error) {
	query := url.Values{
		"offset": {strconv.Itoa(offset)},
		"limit":  {strconv.Itoa(limit)},
	}

	var respInfo struct {
		Data []Memo
	}
	err := c.request("GET", "/api/memo/all", query, nil, &respInfo)
	if err != nil {
		return nil, err
	}

	return respInfo.Data, nil
}

// UserMemoList fetch memo list of specified user from memos server
func (c *Client) UserMemoList(userId, offset, limit int) ([]Memo, error) {
	query := url.Values{
		"creatorId": {strconv.Itoa(userId)},
		"offset":    {strconv.Itoa(offset)},
		"limit":     {strconv.Itoa(limit)},
	}

	var respInfo []Memo
	err := c.request("GET", "/api/v1/memo", query, nil, &respInfo)
	if err != nil {
		return nil, err
	}

	return respInfo, nil
}

// CreateMemo create memo from exists one
func (c *Client) CreateMemo(content string, createdTs int64, resourceIds []int) (*Memo, error) {
	param := MemoCreate{
		Content:        content,
		CreatedTs:      &createdTs,
		Visibility:     api.Public,
		ResourceIDList: resourceIds,
	}

	var result struct {
		Data Memo
	}
	err := c.request("POST", "/api/v1/memo", nil, param, &result)
	if err != nil {
		return nil, err
	}

	return &result.Data, nil
}
