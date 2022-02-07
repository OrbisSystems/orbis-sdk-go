package orbis

import (
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
)

var (
	previewOptions           = "/orders/v2/preview/option"
	placementOptions         = "/orders/v2/place/option"
	multilegStrategyOptions  = "/orders/v2/preview/options/multileg/strategy"
	previewMultilegOptions   = "/orders/v2/preview/options/multileg"
	placementMultilegOptions = "/orders/v2/place/options/multileg"
	replacePreviewOptions    = "/orders/v2/multileg/replace/preview"
	replaceMultilegOptions   = "/orders/v2/multileg/replace"
)

func (c *Client) PreviewOptions(req models.OptionsPreviewReq) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+previewEquities)
}

func (c *Client) PlacementOptions(req models.OptionsPreviewReq) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+placementOptions)
}

func (c *Client) MultilegStrategyOptions(req models.MultilegStrategyReq) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+multilegStrategyOptions)
}

func (c *Client) PreviewMultilegOptions(req models.PreviewMultilegReq) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+previewMultilegOptions)
}

func (c *Client) PlacementMultilegOptions(req models.MultilegStrategyReq) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+placementMultilegOptions)
}

func (c *Client) ReplacePreviewOptions(req models.ReplacePreviewReq) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+replacePreviewOptions)
}

func (c *Client) ReplaceMultilegOptions(req models.ReplacePreviewReq) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+replaceMultilegOptions)
}
