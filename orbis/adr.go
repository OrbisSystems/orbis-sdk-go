package orbis

import (
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"net/http"
	"strconv"
)

var (
	getTopTenADRS = "/research/adrs/top10"
	getDefaults   = "/research/adrs/top10/defaults"
	searchAdrS    = "/research/adrs"
)

func (c *Client) GetTopTenADRS(country *string, loadQuotes, loadEarningReleases, loadUpgradesDowngrades bool) (models.Top10, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getTopTenADRS, nil)
	if err != nil {
		return models.Top10{}, err
	}

	q := httpReq.URL.Query()
	q.Add("loadQuotes", strconv.FormatBool(loadQuotes))
	q.Add("loadEarningReleases", strconv.FormatBool(loadEarningReleases))
	q.Add("loadUpgradesDowngrades", strconv.FormatBool(loadUpgradesDowngrades))
	if country != nil {
		q.Add("country", *country)
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return models.Top10{}, err
	}

	var resp models.Top10
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.Top10{}, err
	}

	return resp, nil
}

func (c *Client) GetDefaults() ([]string, error) {
	r, err := c.client.Get(c.config.ApiUrl+getDefaults, nil)
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

func (c *Client) SearchAdrS(country *string, loadQuotes, loadEarningReleases, loadUpgradesDowngrades bool) (models.AddrSearch, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+searchAdrS, nil)
	if err != nil {
		return models.AddrSearch{}, err
	}

	q := httpReq.URL.Query()
	q.Add("loadQuotes", strconv.FormatBool(loadQuotes))
	q.Add("loadEarningReleases", strconv.FormatBool(loadEarningReleases))
	q.Add("loadUpgradesDowngrades", strconv.FormatBool(loadUpgradesDowngrades))
	if country != nil {
		q.Add("country", *country)
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return models.AddrSearch{}, err
	}

	var resp models.AddrSearch
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.AddrSearch{}, err
	}

	return resp, nil
}
