package orbis

import (
	"fmt"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
	"net/http"
	"strconv"
)

var (
	listing          = "/orders/%s"
	listingForModel  = "/orders/model/%s/%s"
	listingForBranch = "/orders/branch/%s"
	status           = "/orders/status/%s"
	cancelRequest    = "/orders/v2/replace"
	cancellation     = "/orders/v2/cancel"
)

func (c *Client) Listing(listingType string, loadQuotes bool, accountID, account, symbol *string) (
	io.ReadCloser, error) {
	req, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+fmt.Sprintf(listing, listingType), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("loadQuotes", strconv.FormatBool(loadQuotes))
	if accountID != nil {
		q.Add("accountId", *accountID)
	}
	if account != nil {
		q.Add("account", *account)
	}
	if symbol != nil {
		q.Add("symbol", *symbol)
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) ListingForModel(listingType, listingModel string, loadQuotes bool) (
	io.ReadCloser, error) {
	req, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+fmt.Sprintf(listingForModel, listingModel, listingType), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("loadQuotes", strconv.FormatBool(loadQuotes))

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) ListingForBranch(listingType string) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(listingForBranch, listingType), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) Status(orderRef string, loadQuotes bool, accountID, account *string) (io.ReadCloser, error) {
	req, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+fmt.Sprintf(status, orderRef), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("loadQuotes", strconv.FormatBool(loadQuotes))
	if accountID != nil {
		q.Add("accountId", *accountID)
	}
	if account != nil {
		q.Add("account", *account)
	}
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) CancelRequest(req models.CancelReplaceReq) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+cancelRequest)
}

func (c *Client) Cancellation(req models.CancellationReq) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+cancellation)
}
