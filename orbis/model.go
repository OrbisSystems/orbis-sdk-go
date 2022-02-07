package orbis

import (
	"fmt"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
	"net/http"
	"strconv"
	"time"
)

var (
	getModelListing           = "/v2/advisory/models"
	getComponentTags          = "/v2/advisory/models/components/tags"
	getBuyingPower            = "/v2/advisory/model/balance/%s/%s"
	getRealtimeBalance        = "/v2/advisory/model/rtb/%s"
	getRealtimeBalanceHistory = "v2/advisory/model/rtb/history/%s"
	getAdjustments            = "/v2/advisory/model/adjustments/%s"
	getAdjustmentPreview      = "/v2/advisory/model/adjustments/preview/%s"
	getMemberAccounts         = "/v2/advisory/model/accounts/%s"
	orphanedAccounts          = "/v2/advisory/model/accounts/orphaned"
	getMemberAccountStatistic = "/v2/advisory/model/accounts/stats/%s"
	getMemberRealtimeBalances = "/v2/advisory/model/rtbs/%s"
	getPortfolio              = "/v2/advisory/model/portfolios/%s"
	getAllModelBuyingPower    = "/v2/advisory/model/balances/%s"
	getPositionLookup         = "/v2/advisory/model/position/equity/%s"
	getPositionLookupOptions  = "/v2/advisory/model/position/option/%s"
	updateAdjustment          = "/v2/advisory/model/adjustments/modify/%s"
	createModel               = "/v2/advisory/model/create"
	updateModel               = "/v2/advisory/model/update"
	deleteModel               = "/v2/advisory/model/delete"
	updateComponent           = "/v2/advisory/model/component/update"
	linkOfMembersAccounts     = "/v2/advisory/model/account/link"
	unlinkMemberAccount       = "/v2/advisory/model/account/unlink"
	allocations               = "/v2/advisory/model/orders/allocate"
)

func (c *Client) GetModelListing(loadQuotes, loadAccounts bool) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getModelListing, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("loadQuotes", strconv.FormatBool(loadQuotes))
	q.Add("loadAccounts", strconv.FormatBool(loadAccounts))
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetComponentTags() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+getComponentTags, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetBuyingPower(currency, modelID string) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(getBuyingPower, currency, modelID), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetRealtimeBalance(modelID string) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(getRealtimeBalance, modelID), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetRealtimeBalanceHistory(modelID string, dateFrom, dateTo *time.Time, page, count int64) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+fmt.Sprintf(getRealtimeBalanceHistory, modelID), nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("page", strconv.FormatInt(page, 10))
	q.Add("count", strconv.FormatInt(count, 10))
	if dateFrom != nil {
		q.Add("dateFrom", dateFrom.Format("01-02-2006"))
	}
	if dateTo != nil {
		q.Add("dateTo", dateTo.Format("01-02-2006"))
	}
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetAdjustments(modelID, status string, loadQuotes bool) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+fmt.Sprintf(getAdjustments, modelID), nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("status", status)
	q.Add("loadQuotes", strconv.FormatBool(loadQuotes))
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetAdjustmentPreview(adjustmentID string) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(getAdjustmentPreview, adjustmentID), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetMemberAccounts(modelID string, loadRtb, loadRtbHistory bool) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+fmt.Sprintf(getMemberAccounts, modelID), nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("loadRtb", strconv.FormatBool(loadRtb))
	q.Add("loadRtbHistory", strconv.FormatBool(loadRtbHistory))
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) OrphanedAccounts() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+orphanedAccounts, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetMemberAccountStatistic(modelID string) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(getMemberAccountStatistic, modelID), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetMemberRealtimeBalances(modelID string) (models.MemberRealtimeBalances, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(getMemberRealtimeBalances, modelID), nil)
	if err != nil {
		return nil, err
	}

	var resp models.MemberRealtimeBalances
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetPortfolio(modelID string, req models.ModelPortfolioRequest) (models.Portfolio, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+fmt.Sprintf(getPortfolio, modelID), nil)
	if err != nil {
		return models.Portfolio{}, err
	}
	q := httpReq.URL.Query()
	if req.LoadQuotes != nil {
		q.Add("loadQuotes", strconv.FormatBool(*req.LoadQuotes))
	}
	if req.Combined != nil {
		q.Add("combined", strconv.FormatBool(*req.Combined))
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
	if req.LoadCost != nil {
		q.Add("loadCost", strconv.FormatBool(*req.LoadCost))
	}
	if req.MarketCapBreakDown != nil {
		q.Add("marketCapBreakDown", strconv.FormatBool(*req.MarketCapBreakDown))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return models.Portfolio{}, err
	}

	var resp models.Portfolio
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.Portfolio{}, err
	}

	return resp, nil
}

func (c *Client) GetAllModelBuyingPower(fCurrency, sCurrencyQuery string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+fmt.Sprintf(getAllModelBuyingPower, fCurrency), nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("currency", sCurrencyQuery)
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetPositionLookup(modelID, symbol string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+fmt.Sprintf(getPositionLookup, modelID), nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("symbol", symbol)
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetPositionLookupOptions(modelID, symbol string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+fmt.Sprintf(getPositionLookupOptions, modelID), nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("symbol", symbol)
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) UpdateAdjustment(action string, req models.UpdateAdjustmentRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+fmt.Sprintf(updateAdjustment, action))
}

func (c *Client) CreateModel(req models.CreateModelRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+createModel)
}

func (c *Client) UpdateModel(req models.UpdateModelRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+updateModel)
}

func (c *Client) DeleteModel(ID int64) (io.ReadCloser, error) {
	req := struct {
		ID int64 `json:"id"`
	}{ID: ID}
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+deleteModel)
}

func (c *Client) UpdateComponent(req models.UpdateComponentRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+updateComponent)
}

func (c *Client) LinkOfMembersAccounts(req models.LinkOfMembersRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+linkOfMembersAccounts)
}

func (c *Client) UnlinkMemberAccount(req models.LinkOfMembersRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+unlinkMemberAccount)
}

func (c *Client) Allocations(req models.AllocationModelRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+allocations)
}
