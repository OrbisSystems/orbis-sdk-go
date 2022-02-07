package orbis

import (
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
	"net/http"
	"time"
)

var (
	confirmDocument                   = "/documents/confirm"
	confirmListDocuments              = "/documents/confirm/list"
	getServiceDocs                    = "/documents/service"
	getAvailableDocumentTypes         = "/documents/list"
	getAvailableDocumentTypesAndCodes = "documents/typesAndCodes"
	download                          = "/documents/download"
)

func (c *Client) ConfirmDocument(account string, date *time.Time) (models.ConfirmDocumentResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+confirmDocument, nil)
	if err != nil {
		return models.ConfirmDocumentResponse{}, err
	}
	q := httpReq.URL.Query()
	q.Add("account", account)
	if date != nil {
		q.Add("date", date.Format("01-02-2006"))
	}
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return models.ConfirmDocumentResponse{}, err
	}

	var resp models.ConfirmDocumentResponse
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.ConfirmDocumentResponse{}, err
	}

	return resp, nil
}

func (c *Client) ConfirmListDocuments(account string, beginDate, endDate *time.Time) (models.ListOfConfirms, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+confirmListDocuments, nil)
	if err != nil {
		return models.ListOfConfirms{}, err
	}
	q := httpReq.URL.Query()
	q.Add("account", account)
	if beginDate != nil {
		q.Add("beginDate", beginDate.Format("01-02-2006"))
	}
	if endDate != nil {
		q.Add("endDate", endDate.Format("01-02-2006"))
	}
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return models.ListOfConfirms{}, err
	}

	var resp models.ListOfConfirms
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.ListOfConfirms{}, err
	}

	return resp, nil
}

func (c *Client) GetServiceDocs(req models.ServiceDocsRequest) (models.ServiceDocsResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getServiceDocs, nil)
	if err != nil {
		return models.ServiceDocsResponse{}, err
	}
	q := httpReq.URL.Query()
	if req.AccountNumber != nil {
		q.Add("accountNumber", *req.AccountNumber)
	}
	if req.DocType != nil {
		q.Add("docType", *req.DocType)
	}
	if req.BeginMonth != nil {
		q.Add("beginMonth", req.BeginMonth.Format("01"))
	}
	if req.BeginYear != nil {
		q.Add("beginYear", req.BeginYear.Format("2006"))
	}
	if req.EndMonth != nil {
		q.Add("endMonth", req.EndMonth.Format("01"))
	}
	if req.EndYear != nil {
		q.Add("endYear", req.EndYear.Format("2006"))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return models.ServiceDocsResponse{}, err
	}

	var resp models.ServiceDocsResponse
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.ServiceDocsResponse{}, err
	}

	return resp, nil
}

func (c *Client) GetAvailableDocumentTypes() ([]string, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getAvailableDocumentTypes, nil)
	if err != nil {
		return nil, err
	}
	r, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	var resp []string
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetAvailableDocumentTypesAndCodes() (models.AvailableDocumentTypesAndCodes, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getAvailableDocumentTypesAndCodes, nil)
	if err != nil {
		return models.AvailableDocumentTypesAndCodes{}, err
	}
	r, err := c.client.Do(httpReq)
	if err != nil {
		return models.AvailableDocumentTypesAndCodes{}, err
	}

	var resp models.AvailableDocumentTypesAndCodes
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.AvailableDocumentTypesAndCodes{}, err
	}

	return resp, nil
}

func (c *Client) Download(account, key string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+download, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("account", account)
	q.Add("key", key)

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}
