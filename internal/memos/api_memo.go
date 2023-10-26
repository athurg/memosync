package memos

import (
	"net/url"
	"strconv"
)

type Memo struct {
	Content        string     `json:"content,omitempty"`
	CreatedTs      int64      `json:"createdTs,omitempty"`
	Visibility     string     `json:"visibility,omitempty"`
	ResourceIDList []int32    `json:"resourceIdList,omitempty"`
	ResourceList   []Resource `json:"resourceList,omitempty"`
}

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
func (c *Client) CreateMemo(content string, createdTs int64, resourceIds []int32) error {
	param := Memo{
		Content:        content,
		CreatedTs:      createdTs,
		Visibility:     "PUBLIC",
		ResourceIDList: resourceIds,
	}

	err := c.request("POST", "/api/v1/memo", nil, param, nil)
	if err != nil {
		return err
	}

	return nil
}
