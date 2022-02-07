package orbis

import (
	"fmt"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
)

const (
	notificationLink     = "/v1/branch/feed/queue/%s/link"
	notificationUnLink   = "/v1/branch/feed/queue/%s/unlink"
	getNotificationLinks = "/v1/branch/feed/queues/%s"
)

func (c *Client) NotificationLink(typeLink string, req models.NotificationLinkRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+fmt.Sprintf(notificationLink, typeLink))
}

func (c *Client) NotificationUnLink(typeLink string, req models.NotificationUnLinkRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+fmt.Sprintf(notificationUnLink, typeLink))
}

func (c *Client) GetNotificationLinks(typeLink string) (io.ReadCloser, error) {
	url := fmt.Sprintf(getNotificationLinks, typeLink)
	r, err := c.client.Get(c.config.ApiUrl+url, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}
