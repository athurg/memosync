package memos

type Client struct {
	OpenId string
	Addr   string
}

func New(addr, openId string) *Client {
	return &Client{Addr: addr, OpenId: openId}
}

