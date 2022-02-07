package orbis

import (
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
)

var updateTradingLimit = "/user/balance/adjust"

func (c *Client) UpdateTradingLimit(req models.TradingLimitManagementRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+updateTradingLimit)
}
