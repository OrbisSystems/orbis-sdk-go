package orbis

import (
	"bytes"
	"encoding/json"
	"fmt"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var (
	getCompanyResearch              = "/research/%s"
	getOptionRatios                 = "/research/option/ratios"
	getOptionRadar                  = "/api/research/option/radar"
	getIndexComponents              = "/research/index/components/%s?loadQuotes=$s"
	getStockUpgrades                = "/research/upgrades"
	getOwnershipDetails             = "/research/ownerships/details"
	getOwnerDetails                 = "/research/owner/details"
	getEarningsCalendar             = "/research/earnings"
	getEarningsCalendarForPortfolio = "/research/earnings"
	getFundamentals                 = "/research/fundamentals/%s/%s"
	getFundamentalTypes             = "/research/fundamentals/types"
	researchScreener                = "/research/screener"
	researchScreenerV2              = "/research/screener"
	getIndustryHeatMap              = "/research/heatMap/industryRanking"
	getIndustryList                 = "/research/industries/list"
	getSymbolHeatMap                = "/research/heatMap/symbols"
	getIndexOverview                = "/research/marketOverview"
	getHistoricalData               = "/research/historicalPrices"
	getTopQuotesByCategory          = "/quotes/top10/%s"
)

func (c *Client) GetCompanyResearch(symbol string) (models.CompanyResearch, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(getCompanyResearch, symbol), nil)
	if err != nil {
		return models.CompanyResearch{}, err
	}

	var resp models.CompanyResearch
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.CompanyResearch{}, err
	}

	return resp, nil
}

func (c *Client) GetOptionRatios(symbols []string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getOptionRatios, nil)
	if err != nil {
		return nil, err
	}

	q := httpReq.URL.Query()
	if len(symbols) > 0 {
		q.Add("symbols", strings.Join(symbols, ","))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}
func (c *Client) GetOptionRadar(filter string) (models.OptionRadar, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getOptionRadar, nil)
	if err != nil {
		return models.OptionRadar{}, err
	}

	q := httpReq.URL.Query()
	q.Add("filter", filter)

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return models.OptionRadar{}, err
	}

	var resp models.OptionRadar
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.OptionRadar{}, err
	}

	return resp, nil
}
func (c *Client) GetIndexComponents(index string, loadQuotes *bool) (models.IndexComponents, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+fmt.Sprintf(getIndexComponents, index), nil)
	if err != nil {
		return models.IndexComponents{}, err
	}

	q := httpReq.URL.Query()
	if loadQuotes != nil {
		q.Add("loadQuotes", strconv.FormatBool(*loadQuotes))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return models.IndexComponents{}, err
	}

	var resp models.IndexComponents
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.IndexComponents{}, err
	}

	return resp, nil
}
func (c *Client) GetStockUpgrades(symbols []string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getStockUpgrades, nil)
	if err != nil {
		return nil, err
	}

	q := httpReq.URL.Query()
	if len(symbols) > 0 {
		q.Add("symbols", strings.Join(symbols, ","))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetOwnershipDetails(symbols []string, filter, sort string, count, page int64) (models.OwnershipDetails, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getOwnershipDetails, nil)
	if err != nil {
		return nil, err
	}

	q := httpReq.URL.Query()
	q.Add("filter", filter)
	q.Add("sort", sort)
	q.Add("count", strconv.FormatInt(count, 10))
	q.Add("page", strconv.FormatInt(page, 10))
	if len(symbols) > 0 {
		q.Add("symbols", strings.Join(symbols, ","))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp models.OwnershipDetails
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetOwnerDetails(ids []string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getOwnerDetails, nil)
	if err != nil {
		return nil, err
	}

	q := httpReq.URL.Query()
	if len(ids) > 0 {
		q.Add("ids", strings.Join(ids, ","))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}
func (c *Client) GetEarningCalendar(excludeQuotes, loadQuotes *bool, days *int64, symbols []string) (models.EarningsCalendar, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getEarningsCalendar, nil)
	if err != nil {
		return nil, err
	}

	q := httpReq.URL.Query()
	if excludeQuotes != nil {
		q.Add("excludeQuotes", strconv.FormatBool(*excludeQuotes))
	}
	if loadQuotes != nil {
		q.Add("loadQuotes", strconv.FormatBool(*loadQuotes))
	}
	if days != nil {
		q.Add("days", strconv.FormatInt(*days, 10))
	}
	if len(symbols) > 0 {
		q.Add("symbols", strings.Join(symbols, ","))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp models.EarningsCalendar
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (c *Client) GetEarningCalendarForPortfolio(excludeQuotes, loadQuotes *bool, days *int64) (models.EarningsCalendarForPortfolio, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getEarningsCalendarForPortfolio, nil)
	if err != nil {
		return nil, err
	}

	q := httpReq.URL.Query()
	if excludeQuotes != nil {
		q.Add("excludeQuotes", strconv.FormatBool(*excludeQuotes))
	}
	if loadQuotes != nil {
		q.Add("loadQuotes", strconv.FormatBool(*loadQuotes))
	}
	if days != nil {
		q.Add("days", strconv.FormatInt(*days, 10))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp models.EarningsCalendarForPortfolio
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (c *Client) GetFundamentals(typeFund, symbol string) (models.Fundamentals, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(getFundamentals, typeFund, symbol), nil)
	if err != nil {
		return models.Fundamentals{}, err
	}

	var resp models.Fundamentals
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.Fundamentals{}, err
	}

	return resp, nil
}

func (c *Client) GetFundamentalTypes() ([]string, error) {
	r, err := c.client.Get(c.config.ApiUrl+getFundamentalTypes, nil)
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

func (c *Client) Screener(req models.ScreenerResearch) (models.Screener, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return models.Screener{}, err
	}
	r, err := c.client.Post(c.config.ApiUrl+researchScreener, bytes.NewBuffer(b), nil)
	if err != nil {
		return models.Screener{}, err
	}

	var resp models.Screener
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.Screener{}, err
	}

	return resp, nil
}
func (c *Client) ScreenerV2(req models.ScreenerResearch) (models.Screener, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return models.Screener{}, err
	}
	r, err := c.client.Post(c.config.ApiUrl+researchScreenerV2, bytes.NewBuffer(b), nil)
	if err != nil {
		return models.Screener{}, err
	}

	var resp models.Screener
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.Screener{}, err
	}

	return resp, nil
}
func (c *Client) GetIndustryHeatMap(industryType string) (models.IndustryHeatMap, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getIndustryHeatMap, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("industryType", industryType)
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp models.IndustryHeatMap
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.IndustryHeatMap{}, err
	}

	return resp, nil
}
func (c *Client) GetIndustryList(industry string) (models.IndustryList, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getIndustryList, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("industry", industry)
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp models.IndustryList
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (c *Client) GetSymbolHeatMap(amountLimit *float64, numSymbols, industryCode *int64, exchangeFilter string) (models.SymbolHeatMap, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getSymbolHeatMap, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("exchangeFilter", exchangeFilter)
	if amountLimit != nil {
		q.Add("amountLimit", fmt.Sprintf("%v", *amountLimit))
	}
	if numSymbols != nil {
		q.Add("numSymbols", strconv.FormatInt(*numSymbols, 10))
	}
	if industryCode != nil {
		q.Add("industryCode", strconv.FormatInt(*industryCode, 10))
	}
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp models.SymbolHeatMap
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (c *Client) GetIndexOverview() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+getIndexOverview, nil)
	if err != nil {
		return nil, err
	}
	return c.client.GetBodyAndCheckOK(r)
}
func (c *Client) GetHistoricalData(symbol string, dates []string) (models.HistoricalData, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getHistoricalData, nil)
	if err != nil {
		return models.HistoricalData{}, err
	}
	q := httpReq.URL.Query()
	q.Add("symbol", symbol)
	if len(dates) > 0 {
		q.Add("dates", strings.Join(dates, ","))
	}
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return models.HistoricalData{}, err
	}

	var resp models.HistoricalData
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.HistoricalData{}, err
	}

	return resp, nil
}
func (c *Client) GetTopQuotesByCategory(category, market, order string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getTopQuotesByCategory, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("category", category)
	q.Add("market", market)
	q.Add("order", order)

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}
