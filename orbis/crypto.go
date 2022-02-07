package orbis

import (
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
)

var (
	orderPreviewCrypto = "/orders/v2/crypto/preview"
	orderPlaceCrypto   = "/orders/v2/crypto/place"
)

func (c *Client) OrderPreviewCrypto(req models.OrderCryptoReq) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+orderPreviewCrypto)
}

func (c *Client) OrderPlaceCrypto(req models.OrderCryptoReq) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+orderPlaceCrypto)
}
