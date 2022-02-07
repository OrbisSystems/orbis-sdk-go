package orbis

import (
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
)

var (
	previewEquities          = "/orders/v2/preview/equity"
	placementEquities        = "/orders/v2/place/equity"
	basketPreviewEquities    = "/orders/v2/basket/preview/equity"
	basketPlacementEquities  = "/orders/v2/basket/place/equity"
	advisorPlacementEquities = "/orders/v2/advisory/equity/preview"
	advisorPreviewEquities   = "/orders/v2/advisory/equity/place"
)

func (c *Client) PreviewEquities(req models.EquitiesRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+previewEquities)
}

func (c *Client) PlacementEquities(req models.EquitiesRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+placementEquities)
}

func (c *Client) BasketPreviewEquities(req models.EquitiesRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+basketPreviewEquities)
}

func (c *Client) BasketPlacementEquities(req models.EquitiesRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+basketPlacementEquities)
}

func (c *Client) AdvisorPlacementEquities(req models.AdvisorRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+advisorPlacementEquities)
}

func (c *Client) AdvisorPreviewEquities(req models.AdvisorRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+advisorPreviewEquities)
}
