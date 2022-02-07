package orbis

import (
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	getTypes               = "/research/actions/types"
	getPaymentTypes        = "/research/actions/paymentTypes"
	searchCorporateActions = "/research/actions/search"
)

func (c *Client) GetTypes() ([]string, error) {
	r, err := c.client.Get(c.config.ApiUrl+getTypes, nil)
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

func (c *Client) GetPaymentTypes() ([]string, error) {
	r, err := c.client.Get(c.config.ApiUrl+getPaymentTypes, nil)
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

func (c *Client) SearchCorporateActions(symbol *string, dateFrom, dateTo *time.Time, types []string, usePayDate *bool) ([]models.CorporatesActionsSearch, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+searchCorporateActions, nil)
	if err != nil {
		return nil, err
	}

	q := httpReq.URL.Query()
	if symbol != nil {
		q.Add("symbol", *symbol)
	}
	if dateFrom != nil {
		q.Add("dateFrom", dateFrom.Format("01/02/2006"))
	}
	if dateTo != nil {
		q.Add("dateTo", dateTo.Format("01/02/2006"))
	}
	if len(types) > 0 {
		q.Add("types", strings.Join(types, ","))
	}
	if usePayDate != nil {
		q.Add("usePayDate", strconv.FormatBool(*usePayDate))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp []models.CorporatesActionsSearch
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
