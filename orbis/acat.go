package orbis

import (
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
	"net/http"
	"strconv"
)

var (
	searchACAT   = "/transfer/acat/search"
	cancelACAT   = "/transfer/acat/cancel"
	initiateACAT = "/transfer/acat/initiate"
	bulkSchedule = "/transfer/acat/schedule"
)

func (c *Client) SearchACAT(req models.SearchRequest) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+searchACAT, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	if req.Status != nil {
		q.Add("status", *req.Status)
	}
	if req.Account != nil {
		q.Add("account", *req.Account)
	}
	if req.AccountID != nil {
		q.Add("accountId", strconv.FormatInt(*req.AccountID, 10))
	}
	if req.RequestID != nil {
		q.Add("requestId", *req.RequestID)
	}
	if req.BatchID != nil {
		q.Add("batchId", *req.BatchID)
	}
	if req.LoadAccounts != nil {
		q.Add("loadAccounts", strconv.FormatBool(*req.LoadAccounts))
	}
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) CancelACAT(req models.CancelACATRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+cancelACAT)
}

func (c *Client) InitiateACAT(req models.InitiateACATRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+initiateACAT)
}

func (c *Client) BulkSchedule(req models.InitiateACATScheduleRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+bulkSchedule)
}
