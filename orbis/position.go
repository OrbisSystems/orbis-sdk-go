package orbis

import (
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
)

var transferRequest = "/transfer/position/request"

func (c *Client) TransferRequest(req models.TransferRequest) (io.ReadCloser, error) {
	url := c.config.ApiUrl + transferRequest
	return c.client.DoPostAndReturnBody(req, url)
}
