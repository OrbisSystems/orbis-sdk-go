package orbis

import (
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"net/http"
	"strconv"
)

var (
	getRecent      = "/research/ipo/recent"
	getUpcoming    = "/research/ipo/upcoming"
	getPerformance = "/research/ipo/performance"
)

func (c *Client) GetRecent(loadQuotes *bool, time, sort *string) (models.GetIPOS, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getRecent, nil)
	if err != nil {
		return nil, err
	}

	q := httpReq.URL.Query()
	if loadQuotes != nil {
		q.Add("loadQuotes", strconv.FormatBool(*loadQuotes))
	}
	if time != nil {
		q.Add("time", *time)
	}
	if sort != nil {
		q.Add("sort", *sort)
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp models.GetIPOS
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetUpcoming() (models.GetIPOS, error) {
	r, err := c.client.Get(c.config.ApiUrl+getUpcoming, nil)
	if err != nil {
		return nil, err
	}

	var resp models.GetIPOS
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetPerformance(loadQuotes *bool, time, sort *string) (models.PerformanceIPOS, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getPerformance, nil)
	if err != nil {
		return nil, err
	}

	q := httpReq.URL.Query()
	if loadQuotes != nil {
		q.Add("loadQuotes", strconv.FormatBool(*loadQuotes))
	}
	if time != nil {
		q.Add("time", *time)
	}
	if sort != nil {
		q.Add("sort", *sort)
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp models.PerformanceIPOS
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
