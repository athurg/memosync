package memos

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	OpenId string
	Addr   string
}

func New(addr, openId string) *Client {
	return &Client{Addr: addr, OpenId: openId}
}

func (c *Client) request(method, apiPath string, query url.Values, param, result any) error {
	apiUrl := c.Addr + apiPath
	if query == nil {
		query = url.Values{}
	}
	if c.OpenId != "" {
		query.Set("openId", c.OpenId)
	}
	if q := query.Encode(); q != "" {
		apiUrl += "?" + q
	}

	buff, err := json.Marshal(param)
	if err != nil {
		return fmt.Errorf("fail to marshal param %v: %s", param, err)
	}

	req, err := http.NewRequest(method, apiUrl, bytes.NewReader(buff))
	if err != nil {
		return fmt.Errorf("fail to create http.Request: %s", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("fail to do http.Request: %s", err)
	}
	defer resp.Body.Close()

	buff, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("fail to read http.Response: %s", err)
	}

	var respInfo struct {
		Error   string
		Message string
		Data    any
	}
	respInfo.Data = result

	err = json.Unmarshal(buff, &respInfo)
	if err != nil {
		return fmt.Errorf("fail to unmarshal http.Response %s: %s", string(buff), err)
	}

	if respInfo.Error != "" {
		return errors.New(respInfo.Error)
	}

	return nil
}
