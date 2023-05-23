package memos

import (
	"net/url"
	"strconv"

	"github.com/usememos/memos/api"
)

type (
	Memo       = api.Memo
	MemoCreate = api.MemoCreate
)

// MemoList fetch memo list from memos server
func (c *Client) MemoList(offset, limit int) ([]Memo, error) {
	query := url.Values{
		"offset": {strconv.Itoa(offset)},
		"limit":  {strconv.Itoa(limit)},
	}

	memos := []Memo{}
	err := c.request("GET", "/api/memo/all", query, nil, &memos)
	if err != nil {
		return nil, err
	}

	return memos, nil
}

// CreateMemo create memo from exists one
func (c *Client) CreateMemo(content string, createdTs int64, resourceIds []int) (*Memo, error) {
	param := MemoCreate{
		Content:        content,
		CreatedTs:      &createdTs,
		Visibility:     api.Protected,
		ResourceIDList: resourceIds,
	}

	var result Memo
	err := c.request("POST", "/api/memo", nil, param, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
