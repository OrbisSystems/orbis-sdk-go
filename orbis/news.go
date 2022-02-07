package orbis

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var (
	getList         = "/research/news"
	getNewsBySymbol = "/research/news/ticker/%s"
	getSEC          = "/research/news/sec/%s"
	lookupNews      = "/research/news/%s"
)

func (c *Client) GetList(start, count int64, filter string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getList, nil)
	if err != nil {
		return nil, err
	}

	q := httpReq.URL.Query()
	q.Add("start", strconv.FormatInt(start, 10))
	q.Add("count", strconv.FormatInt(count, 10))
	q.Add("filter", filter)

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetNewsBySymbol(start, count int64, symbol string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+fmt.Sprintf(getNewsBySymbol, symbol), nil)
	if err != nil {
		return nil, err
	}

	q := httpReq.URL.Query()
	q.Add("start", strconv.FormatInt(start, 10))
	q.Add("count", strconv.FormatInt(count, 10))

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetSEC(start, count int64, symbol string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+fmt.Sprintf(getSEC, symbol), nil)
	if err != nil {
		return nil, err
	}

	q := httpReq.URL.Query()
	q.Add("start", strconv.FormatInt(start, 10))
	q.Add("count", strconv.FormatInt(count, 10))

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) LookupNews(newsID string) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(lookupNews, newsID), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}
