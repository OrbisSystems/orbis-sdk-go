package orbis

import (
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
	"net/http"
	"strconv"
)

var (
	getRtd                     = "/v1/branch/rtb"
	getRtds                    = "/v1/branch/rtbs"
	getInventory               = "/v1/branch/inventory"
	getRtdsTotal               = "/v1/branch/total"
	getBranchPortfolio         = "/v1/branch/portfolio"
	getOptionExpirations       = "/v1/branch/portfolio/options/upcoming"
	getOptionExpirationDetails = "/v1//branch/portfolio/options/upcoming/details"
)

func (c *Client) GetRtd() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+getRtd, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetRtds() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+getRtds, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetInventory() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+getInventory, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetRtdsTotal() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+getRtdsTotal, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetBranchPortfolio(req models.PortfolioRequest) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getBranchPortfolio, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	if req.LoadQuotes != nil {
		q.Add("loadQuotes", strconv.FormatBool(*req.LoadQuotes))
	}
	if req.LoadIndustries != nil {
		q.Add("loadIndustries", strconv.FormatBool(*req.LoadIndustries))
	}
	if req.LoadUpgrades != nil {
		q.Add("loadUpgrades", strconv.FormatBool(*req.LoadUpgrades))
	}
	if req.LoadEarnings != nil {
		q.Add("loadEarnings", strconv.FormatBool(*req.LoadEarnings))
	}
	if req.LoadRtb != nil {
		q.Add("loadRtb", strconv.FormatBool(*req.LoadRtb))
	}
	if req.MarketCapBreakDown != nil {
		q.Add("marketCapBreakDown", strconv.FormatBool(*req.MarketCapBreakDown))
	}
	if req.Combined != nil {
		q.Add("combined", strconv.FormatBool(*req.Combined))
	}
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetOptionExpirations() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+getOptionExpirations, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetOptionExpirationDetails(req models.OptionExpirationsDetailsRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, getOptionExpirationDetails)
}
