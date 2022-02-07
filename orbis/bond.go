package orbis

import (
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
)

var (
	placementBonds       = "/orders/v2/bond/place"
	quoteRequestBonds    = "/orders/v2/bond/quote/request"
	quoteAcceptanceBonds = "/orders/v2/bond/quote/accept"
)

func (c *Client) PlacementBonds(req models.BondPlacementReq) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+placementBonds)
}

func (c *Client) QuoteRequestBonds(req models.QuoteBondPlacementReq) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+quoteRequestBonds)
}

func (c *Client) QuoteAcceptanceBonds(req models.BondQuoteAcceptanceReq) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+quoteAcceptanceBonds)
}
