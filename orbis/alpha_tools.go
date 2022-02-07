package orbis

import (
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"net/http"
	"strconv"
	"strings"
)

var (
	getAlphaTracker      = "/quotes/rst"
	getSeasonality       = "/quotes/seasonality"
	getWeaklySeasonality = "/quotes/seasonality/weekly"
)

func (c *Client) GetAlphaTracker(benchmark, rangE *string, symbols []string, period *int64, setEMA *bool) ([]models.AlphaTracker, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getAlphaTracker, nil)
	if err != nil {
		return nil, err
	}

	q := httpReq.URL.Query()
	if benchmark != nil {
		q.Add("benchmark", *benchmark)
	}
	if rangE != nil {
		q.Add("range", *rangE)
	}
	if len(symbols) > 0 {
		q.Add("symbols", strings.Join(symbols, ","))
	}
	if period != nil {
		q.Add("period", strconv.FormatInt(*period, 10))
	}
	if setEMA != nil {
		q.Add("count", strconv.FormatBool(*setEMA))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp []models.AlphaTracker
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetSeasonality(benchmark *string, symbols []string, years int64) ([]models.AlphaTracker, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getSeasonality, nil)
	if err != nil {
		return nil, err
	}

	q := httpReq.URL.Query()
	q.Add("years", strconv.FormatInt(years, 10))
	if benchmark != nil {
		q.Add("benchmark", *benchmark)
	}
	if len(symbols) > 0 {
		q.Add("symbols", strings.Join(symbols, ","))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp []models.AlphaTracker
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetWeaklySeasonality(benchmark *string, symbols []string, years int64) ([]models.AlphaTracker, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getWeaklySeasonality, nil)
	if err != nil {
		return nil, err
	}

	q := httpReq.URL.Query()
	q.Add("years", strconv.FormatInt(years, 10))
	if benchmark != nil {
		q.Add("benchmark", *benchmark)
	}
	if len(symbols) > 0 {
		q.Add("symbols", strings.Join(symbols, ","))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp []models.AlphaTracker
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
