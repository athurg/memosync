package memos

import (
	"fmt"
	"net/url"
)

type Resource struct {
	ID           int32  `json:"id"`
	PublicID     string `json:"publicId"`
	Filename     string `json:"filename"`
	Type         string `json:"type"`
	ExternalLink string `json:"externalLink"`
}

// ResourceLink generate a external link for specified resource
func (c *Client) ResourceLink(r Resource) string {
	return c.Addr + fmt.Sprintf("/o/r/%d/%s/%s", r.ID, r.PublicID, url.PathEscape(r.Filename))
}

// CreateExternalLinkResource create resource with external link
func (c *Client) CreateExternalLinkResource(link, filename, mime string) (*Resource, error) {
	param := Resource{
		ExternalLink: link,
		Filename:     filename,
		Type:         mime,
	}

	var result Resource
	err := c.request("POST", "/api/v1/resource", nil, param, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
