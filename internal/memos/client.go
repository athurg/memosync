package memos

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	Token string
	Addr  string
}

func New(addr string) *Client {
	return &Client{Addr: addr}
}

func NewWithUser(addr, user, pass string) (*Client, error) {
	token, err := SignIn(addr, user, pass)
	if err != nil {
		return nil, err
	}

	return &Client{Addr: addr, Token: token}, nil
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignIn(addr, user, pass string) (string, error) {
	apiUrl := addr + "/api/v1/auth/signin"
	buff, err := json.Marshal(SignInRequest{Username: user, Password: pass})
	if err != nil {
		return "", fmt.Errorf("fail to marshal param %v: %s", user, err)
	}

	resp, err := http.Post(apiUrl, "application/json", bytes.NewReader(buff))
	if err != nil {
		return "", fmt.Errorf("fail to do http.Request: %s", err)
	}
	defer resp.Body.Close()

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "memos.access-token" {
			return cookie.Value, nil
		}
	}

	return "", errors.New("no access-token")
}

func (c *Client) request(method, apiPath string, query url.Values, param, result any) error {
	apiUrl := c.Addr + apiPath
	if query != nil {
		apiUrl += "?" + query.Encode()
	}

	buff, err := json.Marshal(param)
	if err != nil {
		return fmt.Errorf("fail to marshal param %v: %s", param, err)
	}

	req, err := http.NewRequest(method, apiUrl, bytes.NewReader(buff))
	if err != nil {
		return fmt.Errorf("fail to create http.Request: %s", err)
	}

	req.Header.Set("User-Agent", "Memos Discovery")
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	httpClient := &http.Client{Timeout: 5 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("fail to do http.Request: %s", err)
	}
	defer resp.Body.Close()

	buff, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("fail to read http.Response: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errInfo struct {
			Error   string
			Message string
		}

		err = json.Unmarshal(buff, &errInfo)
		if err != nil {
			return fmt.Errorf("fail to unmarshal http.Response %s: %s", string(buff), err)
		}

		return errors.New(errInfo.Error)
	}

	if result == nil {
		return nil
	}

	err = json.Unmarshal(buff, result)
	if err != nil {
		return fmt.Errorf("fail to unmarshal http.Response %s: %s", string(buff), err)
	}

	return nil
}
