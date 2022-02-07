package orbis

import (
	"fmt"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
)

var fundTransfer = "/transfer/bank/%s"

func (c *Client) FundTransfer(direction string, req models.FundTransferRequest) (io.ReadCloser, error) {
	url := c.config.ApiUrl + fmt.Sprintf(fundTransfer, direction)
	return c.client.DoPostAndReturnBody(req, url)
}
