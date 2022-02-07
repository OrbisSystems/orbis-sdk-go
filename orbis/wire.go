package orbis

import (
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
)

var outgoing = "/transfer/wire/outgoing"

func (c *Client) Outgoing(req models.WireRequest) (io.ReadCloser, error) {
	url := c.config.ApiUrl + outgoing
	return c.client.DoPostAndReturnBody(req, url)
}
