package orbis

import (
	"fmt"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"github.com/pkg/errors"
	"io"
	"strings"
	"time"
)

var (
	getEquity             = "/quotes/equity"
	searchEquity          = "/quotes/search"
	getOption             = "/quotes/option"
	getOptionsPricingType = "/quotes/option/pricing/type"
	getBonds              = "/quotes/bonds"
	getHistorical         = "/quotes/equity/historical"
	getIntradayChart      = "/quotes/equity/intraday"
	getEquityTicks        = "/quotes/equity/ticks"
	getShortability       = "/quotes/equity/shortability/%s"
	getTopStocks          = "/quotes/top10/%s"
	getOptionChainDates   = "/quotes/option/chain/dates/%s"
	getOptionChain        = "/quotes/option/chain/%s/%s"
	getOptionChainGreeks  = "/quotes/option/greeks/%s"
)

func (c *Client) GetEquity(symbol string) (models.EquityElement, error) {
	if symbol == "" {
		return models.EquityElement{}, errors.New("symbol must be present")
	}

	var url = c.config.ApiUrl + getEquity + "?symbol=" + symbol
	r, err := c.client.Get(url, nil)
	if err != nil {
		return models.EquityElement{}, err
	}

	var resp models.EquityElement
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.EquityElement{}, err
	}

	return resp, nil
}

func (c *Client) SearchQuote(criteria string, limit int, loadQuotes bool) (models.TopStockQuote, error) {
	if criteria == "" {
		return models.TopStockQuote{}, errors.New("criteria must be present")
	}
	if limit < 1 {
		return models.TopStockQuote{}, errors.New("limit must be greater then zero")
	}

	var url = fmt.Sprintf("%s?criteria=%s&limit=%d&loadQuotes=?%t", searchEquity, criteria, limit, loadQuotes)

	r, err := c.client.Get(c.config.ApiUrl+url, nil)
	if err != nil {
		return models.TopStockQuote{}, err
	}

	var resp models.TopStockQuote
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.TopStockQuote{}, err
	}

	return resp, nil
}

func (c *Client) GetOptions(symbols []string, greeks bool) (io.ReadCloser, error) {
	if len(symbols) == 0 {
		return nil, errors.New("symbols is empty")
	}

	var url = fmt.Sprintf("%s?symbosl=%s&greeks=%t", getOption, strings.Join(symbols, ","), greeks)

	r, err := c.client.Get(c.config.ApiUrl+url, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

// similar with GetOption
func (c *Client) GetOptionsGreeks(symbols []string, greeks bool) (io.ReadCloser, error) {
	if len(symbols) == 0 {
		return nil, errors.New("symbols is empty")
	}

	var url = fmt.Sprintf("%s?symbosl=%s&greeks=%t", getOption, strings.Join(symbols, ","), greeks)

	r, err := c.client.Get(c.config.ApiUrl+url, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetOptionsPricingType(underlying, symbol *string, expectedPx *float64) (io.ReadCloser, error) {
	var url = getOptionsPricingType + "?"
	if underlying != nil {
		url += fmt.Sprintf("underlying=%s&", *underlying)
	}
	if symbol != nil {
		url += fmt.Sprintf("symbol=%s", *symbol)
	}

	if expectedPx != nil {
		url += fmt.Sprintf("expectedPx=%d", expectedPx)
	}

	url = url[0 : len(url)-1]

	r, err := c.client.Get(c.config.ApiUrl+url, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetBonds(cusip []string) (io.ReadCloser, error) {
	var url = getBonds
	if len(cusip) != 0 {
		url += fmt.Sprintf("?cusip=%s&", strings.Join(cusip, ","))
	}

	r, err := c.client.Get(c.config.ApiUrl+url, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetHistorical(req models.HistoricalRequest) (models.Historical, error) {
	var url = getHistorical + "?"
	if req.Symbol != "" {
		url += fmt.Sprintf("symbol=%s&", req.Symbol)
	}
	if req.Start != nil && req.End != nil {
		url += fmt.Sprintf("start=%s&", req.Start.Format("01/02/2006"))
		url += fmt.Sprintf("end=%s&", req.End.Format("01/02/2006"))
	} else {
		if req.Range == nil {
			return nil, errors.New("range or start/end must be provided")
		}
		url += fmt.Sprintf("range=%s&", *req.Range)
	}

	if req.RsiSpan < 0.0 {
		return nil, errors.New("rsiSpan must me greater then zero")
	}

	url += fmt.Sprintf("rsiSpan=%f&", req.RsiSpan)
	url += fmt.Sprintf("setWeekly=%t&", req.SetWeekly)

	url = url[0 : len(url)-1]

	r, err := c.client.Get(c.config.ApiUrl+url, nil)
	if err != nil {
		return nil, err
	}

	var resp models.Historical
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetIntradayChartData(symbol *string, rsiSpa *float64, from, to *int64, rowVolumes *bool) (
	models.IntradayChartData, error) {
	if symbol == nil {
		return nil, errors.New("symbol must be not nil")
	}

	if rsiSpa == nil {
		return nil, errors.New("rsiSpan must be not nil")
	}

	if *rsiSpa > 0.0 {
		return nil, errors.New("rsiSpan must be grater then zero")
	}

	var url = getIntradayChart + "?"
	url += fmt.Sprintf("symbol=%s&", *symbol)
	url += fmt.Sprintf("rsiSpan=%v&", *rsiSpa)

	if from != nil && to != nil {
		url += fmt.Sprintf("start=%d&", from)
		url += fmt.Sprintf("end=%d&", to)
	}

	if rowVolumes != nil {
		url += fmt.Sprintf("rowVolumes=%t&", *rowVolumes)
	}

	url = url[0 : len(url)-1]

	r, err := c.client.Get(c.config.ApiUrl+url, nil)
	if err != nil {
		return nil, err
	}

	var resp models.IntradayChartData
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetEquityTicks(symbol *string, startTime, endTime *time.Time) (io.ReadCloser, error) {
	if symbol == nil {
		return nil, errors.New("symbol must be not nil")
	}
	if startTime == nil || endTime == nil {
		return nil, errors.New("startTime or endTime must be not nil")
	}
	var url = getEquityTicks + "?"

	url += fmt.Sprintf("symbol=%s&", *symbol)
	url += fmt.Sprintf("startTime=%s&", startTime.Format("01/02/2006"))
	url += fmt.Sprintf("endTime=%s", endTime.Format("01/02/2006"))

	r, err := c.client.Get(c.config.ApiUrl+url, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetShortability(symbol string) (models.GetShortability, error) {
	if symbol == "" {
		return models.GetShortability{}, errors.New("symbol must be not empty")
	}
	url := c.config.ApiUrl + fmt.Sprintf(getShortability, symbol)

	r, err := c.client.Get(url, nil)
	if err != nil {
		return models.GetShortability{}, err
	}

	var resp models.GetShortability
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.GetShortability{}, err
	}

	return resp, nil
}

func (c *Client) GetTopStocksQuote(category, market, order *string, count *int64) (models.TopStocks, error) {
	if category == nil {
		return nil, errors.New("category must be not nil")
	}
	url := c.config.ApiUrl + fmt.Sprintf(getTopStocks, *category) + "?"

	if market != nil {
		url += fmt.Sprintf("market=%s&", *market)
	}
	if order != nil {
		url += fmt.Sprintf("order=%s&", *order)
	}

	if count != nil {
		url += fmt.Sprintf("count=%d&", *count)
	}

	url = url[0 : len(url)-1]

	r, err := c.client.Get(url, nil)
	if err != nil {
		return nil, err
	}

	var resp models.TopStocks
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (c *Client) GetOptionChainDates(symbol string, symbols []string) (io.ReadCloser, error) {
	if symbol == "" {
		return nil, errors.New("symbol must be not empty")
	}

	var url = c.config.ApiUrl + fmt.Sprintf(getOptionChainDates, symbol)

	if len(symbols) > 0 {
		url += fmt.Sprintf("?symbols=%s", strings.Join(symbols, ","))
	}

	r, err := c.client.Get(url, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetOptionChain(symbol, date string) (io.ReadCloser, error) {
	if symbol == "" {
		return nil, errors.New("symbol must be not empty")
	}

	if date == "" {
		return nil, errors.New("date must be not empty")
	}

	var url = c.config.ApiUrl + fmt.Sprintf(getOptionChain, symbol, date)
	r, err := c.client.Get(url, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetOptionChainGreek(symbol, date string, symbols []string) (io.ReadCloser, error) {
	if symbol == "" {
		return nil, errors.New("symbol must be not empty")
	}

	if date == "" {
		return nil, errors.New("date must be not empty")
	}

	var url = c.config.ApiUrl + fmt.Sprintf(getOptionChainGreeks, symbol, date)

	if len(symbols) > 0 {
		url += fmt.Sprintf("?symbols=%s", strings.Join(symbols, ","))
	}

	r, err := c.client.Get(url, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}
