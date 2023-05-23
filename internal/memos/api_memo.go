package memos

import (
	"net/url"
	"strconv"

	"github.com/usememos/memos/api"
)

type Memo api.Memo

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
