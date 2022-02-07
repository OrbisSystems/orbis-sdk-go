package orbis

import (
	"fmt"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"net/http"
	"strconv"
	"strings"
)

var (
	liveFeedUrl          = "/research/tipranks/livefeed"
	getAnalystPortfolio  = "/research/tipranks/analyst/portfolio/%s"
	getAnalystProfile    = "/research/tipranks/analyst/profile/%s"
	topAnalyst           = "/research/tipranks/analyst/top"
	getRatingsBySymbol   = "/research/tipranks/analyst/%s"
	insiders             = "/research/tipranks/insiders"
	insiderTransactions  = "/research/tipranks/insiders/transactions"
	getSimilarStocks     = "/research/tipranks/stocks/similar/%s"
	getOverview          = "/research/tipranks/stocks/overview/%s"
	getTopStocksTipRanks = "/research/tipranks/stocks/overview/%s"
	consensusBySector    = "/research/tipranks/consensus/sector"
	consensusByBlogger   = "/research/tipranks/consensus/blogger/%s"
	trending             = "/research/tipranks/trending"
	hedgefunds           = "/research/tipranks/hedgefunds/%s"
	stockSentiment       = "/research/tipranks/rating/blogger/%s"
	bloggerRatings       = "/research/tipranks/rating/blogger/%s"
)

func (c *Client) GetLiveFeed(count int64) (models.LiveFeedResponse, error) {
	url := fmt.Sprintf("%s%s?count=%d", c.config.ApiUrl, liveFeedUrl, count)
	r, err := c.client.Get(url, nil)
	if err != nil {
		return nil, err
	}

	var resp models.LiveFeedResponse

	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetAnalystPortfolio(expertUUID string, count int64, loadQuotes bool) (models.GetAnalystResponse, error) {
	req, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+fmt.Sprintf(getAnalystPortfolio, expertUUID), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("loadQuotes", strconv.FormatBool(loadQuotes))
	q.Add("count", strconv.FormatInt(count, 10))

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp models.GetAnalystResponse
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetAnalystProfile(expertUUID string) (models.AnalystProfileResponse, error) {
	req, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+fmt.Sprintf(getAnalystProfile, expertUUID), nil)
	if err != nil {
		return models.AnalystProfileResponse{}, err
	}
	q := req.URL.Query()

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return models.AnalystProfileResponse{}, err
	}

	var resp models.AnalystProfileResponse
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.AnalystProfileResponse{}, err
	}

	return resp, nil
}

func (c *Client) GetTopAnalyst(count int64) (models.GetTopAnalystResponse, error) {
	url := fmt.Sprintf("%s%s?count=%d", c.config.ApiUrl, topAnalyst, count)
	r, err := c.client.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var resp models.GetTopAnalystResponse
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetRatingsBySymbol(symbol string, count int64) (models.RatingBySymbolResponse, error) {
	req, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+fmt.Sprintf(getRatingsBySymbol, symbol), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("count", strconv.FormatInt(count, 10))

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp models.RatingBySymbolResponse
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetInsider(symbols []string) (models.InsiderResponse, error) {
	url := fmt.Sprintf("%s%s?symbols=%s", c.config.ApiUrl, insiders, strings.Join(symbols[:], ","))
	r, err := c.client.Get(url, nil)
	if err != nil {
		return models.InsiderResponse{}, err
	}

	var resp models.LiveFeedResponse

	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.InsiderResponse{}, err
	}

	return models.InsiderResponse{}, nil
}

func (c *Client) GetInsiderTransactions(symbols []string, count int64) (models.InsiderTransactionsResponse, error) {
	url := fmt.Sprintf("%s%s?symbols=%s&count=%d", c.config.ApiUrl, insiderTransactions, strings.Join(symbols[:], ","), count)
	r, err := c.client.Get(url, nil)
	if err != nil {
		return nil, err
	}

	var resp models.InsiderTransactionsResponse

	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetSimilarStocks(symbol string) ([]string, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(getSimilarStocks, symbol), nil)
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

func (c *Client) Overview(symbols []string) (models.OverviewResponse, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(getOverview, strings.Join(symbols, ",")), nil)
	if err != nil {
		return nil, err
	}

	var resp models.OverviewResponse
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetTopStocksTipRanks(symbols []string) (models.GetTopStocksResponse, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(getTopStocksTipRanks, strings.Join(symbols, ",")), nil)
	if err != nil {
		return nil, err
	}

	var resp models.GetTopStocksResponse
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetConsensusBySector() (models.ConsensusBySector, error) {
	url := fmt.Sprintf("%s%s", c.config.ApiUrl, consensusBySector)
	r, err := c.client.Get(url, nil)
	if err != nil {
		return nil, err
	}

	var resp models.ConsensusBySector

	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetConsensusByBlogger(symbol string) (models.ConsensusByBloggerResponse, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(consensusByBlogger, symbol), nil)
	if err != nil {
		return models.ConsensusByBloggerResponse{}, err
	}

	var resp models.ConsensusByBloggerResponse
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.ConsensusByBloggerResponse{}, err
	}

	return resp, nil
}

func (c *Client) GetTrending(count, period int64) (models.TrendingResponse, error) {
	url := fmt.Sprintf("%s%s?count=%d&period=%d", c.config.ApiUrl, trending, count, period)
	r, err := c.client.Get(url, nil)
	if err != nil {
		return nil, err
	}

	var resp models.TrendingResponse

	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetHedgeFunds(symbol string) (models.HedgeFundsResponse, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(hedgefunds, symbol), nil)
	if err != nil {
		return models.HedgeFundsResponse{}, err
	}

	var resp models.HedgeFundsResponse
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.HedgeFundsResponse{}, err
	}

	return resp, nil
}

func (c *Client) GetStockSentiment(symbol string) (models.StockSentimentResponse, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(stockSentiment, symbol), nil)
	if err != nil {
		return models.StockSentimentResponse{}, err
	}

	var resp models.StockSentimentResponse
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.StockSentimentResponse{}, err
	}

	return resp, nil
}

func (c *Client) GetBloggerRatings(symbol string) (models.BloggerRatingsResponse, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(bloggerRatings, symbol), nil)
	if err != nil {
		return models.BloggerRatingsResponse{}, err
	}

	var resp models.BloggerRatingsResponse
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.BloggerRatingsResponse{}, err
	}

	return resp, nil
}
