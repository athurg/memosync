package memos

import (
	"fmt"
	"net/url"

	"github.com/usememos/memos/api"
)

type Resource = api.Resource

// ResourceLink generate a external link for specified resource
func (c *Client) ResourceLink(r Resource) string {
	return c.Addr + fmt.Sprintf("/o/r/%d/%s/%s", r.ID, r.PublicID, url.PathEscape(r.Filename))
}

// CreateExternalLinkResource create resource with external link
func (c *Client) CreateExternalLinkResource(link, filename, mime string) (*Resource, error) {
	param := api.ResourceCreate{
		ExternalLink: link,
		Filename:     filename,
		Type:         mime,
	}

	var result Resource
	err := c.request("POST", "/api/resource", nil, param, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
