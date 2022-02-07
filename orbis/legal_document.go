package orbis

import (
	"fmt"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
	"net/http"
)

var (
	getDocumentClassifications = "/user/legal/classifications"
	getDocumentListing         = "/user/legal/documents"
	getPolicyValues            = "/user/legal/policies/listing/%s"
	getDocumentContent         = "/user/legal/document/content"
	updatePolicy               = "/user/legal/policy/%s"
)

func (c *Client) GetDocumentClassifications() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+getDocumentClassifications, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetDocumentListing(classification *string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getDocumentListing, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	if classification != nil {
		q.Add("classification", *classification)
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetPolicyValues(policyType string) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(getPolicyValues, policyType), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetDocumentContent(ID *string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getDocumentContent, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	if ID != nil {
		q.Add("id", *ID)
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) UpdatePolicy(policyType string, req models.UpdatePolicyRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+fmt.Sprintf(updatePolicy, policyType))
}
