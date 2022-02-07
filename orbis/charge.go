package orbis

import (
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
	"net/http"
	"strconv"
)

var (
	searchCharge = "/v2/charges/search"
	applyCharge  = "/v2/charges/apply"
	cancelCharge = "/v2/charges/cancel"
)

func (c *Client) SearchCharge(req models.SearchChargesRequest) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+searchCharge, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("AccountId", strconv.FormatInt(req.AccountID, 10))
	q.Add("AccountNumber", req.AccountNumber)
	if req.DateFrom != nil {
		q.Add("DateFrom", req.DateFrom.Format("01-02-2006"))
	}
	if req.DateTo != nil {
		q.Add("DateTo", req.DateTo.Format("01-02-2006"))
	}
	if req.ChargeRef != nil {
		q.Add("ChargeRef", *req.ChargeRef)
	}
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) ApplyCharge(req models.ApplyChargeRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+applyCharge)
}

func (c *Client) CancelCharge(chargeRef string) (io.ReadCloser, error) {
	req := struct {
		ChargeRef string `json:"chargeRef"`
	}{
		ChargeRef: chargeRef,
	}
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+cancelCharge)
}
