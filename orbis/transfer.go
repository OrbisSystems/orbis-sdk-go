package orbis

import (
	"fmt"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
	"net/http"
	"strconv"
)

var (
	getTransferDetails      = "/transfer/details/%s/%s"
	getTransferStatus       = "/transfer/status/%s"
	getTransferStatusModel  = "/transfer/status/model/%s"
	getTransferStatusBranch = "/transfer/status/branch"
	getAvailableAmount      = "/transfer/available/%s"
	getLegalDocumentTypes   = "/transfer/legal/types"
	getLegalDocument        = "/transfer/legal/content"
	cancelTransfer          = "/transfer/ach/cancel"
)

func (c *Client) GetTransferDetails(account, transferID string) (io.ReadCloser, error) {
	url := fmt.Sprintf(getTransferDetails, account, transferID)
	r, err := c.client.Get(c.config.ApiUrl+url, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetTransferStatus(account string, req models.TransferStatusRequest) (models.TransferStatus, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+fmt.Sprintf(getTransferStatus, account), nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	if req.Direction != nil {
		q.Add("direction", *req.Direction)
	}
	if req.Mechanism != nil {
		q.Add("mechanism", *req.Mechanism)
	}
	if req.DateFrom != nil {
		q.Add("dateFrom", req.DateFrom.Format("01/02/2006"))
	}
	if req.DateTo != nil {
		q.Add("dateTo", req.DateTo.Format("01/02/2006"))
	}
	if req.Range != nil {
		q.Add("range", *req.Range)
	}
	if req.Start != nil {
		q.Add("start", *req.Start)
	}
	if req.PerPage != nil {
		q.Add("direction", strconv.FormatInt(*req.PerPage, 10))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp models.TransferStatus
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetTransferStatusModel(model int64) (models.TransferStatus, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+fmt.Sprintf(getTransferStatusModel,
		strconv.FormatInt(model, 10)), nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp models.TransferStatus
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetTransferStatusBranch() (models.TransferStatus, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getTransferStatusBranch, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp models.TransferStatus
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetAvailableAmount(account string) (models.AmountAvailable, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+fmt.Sprintf(getAvailableAmount, account), nil)
	if err != nil {
		return models.AmountAvailable{}, err
	}
	q := httpReq.URL.Query()
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return models.AmountAvailable{}, err
	}

	var resp models.AmountAvailable
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.AmountAvailable{}, err
	}

	return resp, nil
}

func (c *Client) GetLegalDocumentTypes() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+getLegalDocumentTypes, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetLegalDocument(ID string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getLegalDocument, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("id", ID)
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) CancelTransfer(req models.CancelTransferRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+cancelTransfer)
}
