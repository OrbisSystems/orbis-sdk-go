package orbis

import (
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"net/http"
	"strconv"
	"strings"
)

var (
	getInitScreener      = "/research/etf/init"
	getScreener          = "/research/etf/screener"
	getTopTenETF         = "/research/etf/top10"
	getDetails           = "/research/etf/details"
	getBreakDownBySector = "/research/etf/breakdown"
)

func (c *Client) GetInitScreener() (models.InitScreener, error) {
	r, err := c.client.Get(c.config.ApiUrl+getInitScreener, nil)
	if err != nil {
		return models.InitScreener{}, err
	}

	var resp models.InitScreener
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.InitScreener{}, err
	}

	return resp, nil
}

func (c *Client) GetScreener(req models.ScreenerRequest) (models.ScrenerResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getScreener, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("asc", strconv.FormatBool(req.Asc))
	q.Add("loadIndustries", strconv.FormatBool(req.LoadIndustries))
	q.Add("loadUpgrades", strconv.FormatBool(req.LoadUpgrades))
	q.Add("loadEarnings", strconv.FormatBool(req.LoadEarnings))

	if req.InverseType != nil {
		q.Add("inverseType", *req.InverseType)
	}
	if req.LeverageType != nil {
		q.Add("leverageType", *req.LeverageType)
	}
	if req.AssetID != nil {
		q.Add("assetId", *req.AssetID)
	}
	if req.CategoryID != nil {
		q.Add("categoryId", *req.CategoryID)
	}
	if req.SponsorID != nil {
		q.Add("sponsorId", *req.SponsorID)
	}
	if req.Start != nil {
		q.Add("sponsorId", strconv.FormatInt(*req.Start, 10))
	}
	if req.Num != nil {
		q.Add("num", strconv.FormatInt(*req.Num, 10))
	}
	if req.Sort != nil {
		q.Add("sort", *req.Sort)
	}
	if len(req.Filters) > 0 {
		q.Add("filters", strings.Join(req.Filters, ","))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp models.ScrenerResponse
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetTopTenETF(inverseType, leverageType *string) (models.Top10, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getTopTenETF, nil)
	if err != nil {
		return models.Top10{}, err
	}
	q := httpReq.URL.Query()
	if inverseType != nil {
		q.Add("inverseType", *inverseType)
	}
	if leverageType != nil {
		q.Add("leverageType", *leverageType)
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

func (c *Client) GetDetails(symbol *string, loadQuotes *bool) (models.Details, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getDetails, nil)
	if err != nil {
		return models.Details{}, err
	}
	q := httpReq.URL.Query()
	if symbol != nil {
		q.Add("symbol", *symbol)
	}
	if loadQuotes != nil {
		q.Add("loadQuotes", strconv.FormatBool(*loadQuotes))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return models.Details{}, err
	}

	var resp models.Details
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.Details{}, err
	}

	return resp, nil
}

func (c *Client) GetBreakDownBySector(duration *string, displaySize, exposure *int64) (models.BreakdownBySector, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getBreakDownBySector, nil)
	if err != nil {
		return models.BreakdownBySector{}, err
	}
	q := httpReq.URL.Query()
	if duration != nil {
		q.Add("duration", *duration)
	}
	if displaySize != nil {
		q.Add("displaySize", strconv.FormatInt(*displaySize, 10))
	}
	if exposure != nil {
		q.Add("exposure", strconv.FormatInt(*exposure, 10))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return models.BreakdownBySector{}, err
	}

	var resp models.BreakdownBySector
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.BreakdownBySector{}, err
	}

	return resp, nil
}
