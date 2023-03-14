package quotes

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	sdk "github.com/OrbisSystems/orbis-sdk-go/interface"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

// Quotes service returns quotes data.
type Quotes struct {
	sdk.Auth

	url string
	cli sdk.HTTPClient
}

func New(url string, auth sdk.Auth, cli sdk.HTTPClient) *Quotes {
	return &Quotes{
		Auth: auth,
		url:  url,
		cli:  cli,
	}
}

func (q *Quotes) GetQuotesEquityData(ctx context.Context, symbols, quoteType string) ([]model.QuoteEquityDataResponse, error) {
	log.Trace("GetQuotesEquityData called")

	r, err := q.cli.Get(ctx, fmt.Sprintf("%s?symbols=%s&quote_type=%s", q.url+model.URLInsightBase+model.URLInsightQuotesEquity, symbols, quoteType), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get quotes equity")
	}

	var resp []model.QuoteEquityDataResponse
	err = q.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (q *Quotes) GetQuoteHistory(ctx context.Context, req model.QuoteHistoryRequest) (model.QuoteHistoryResponse, error) {
	log.Trace("GetQuoteHistory called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.QuoteHistoryResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := q.cli.Post(ctx, q.url+model.URLInsightBase+model.URLInsightQuoteHistory, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.QuoteHistoryResponse{}, errors.Wrap(err, "couldn't get quote history")
	}

	var resp model.QuoteHistoryResponse
	err = q.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.QuoteHistoryResponse{}, err
	}

	return resp, err
}

func (q *Quotes) GetIntradayQuotes(ctx context.Context, req model.IntradayRequest) ([]model.IntradayResponse, error) {
	log.Trace("GetIntradayQuotes called")

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := q.cli.Post(ctx, q.url+model.URLInsightBase+model.URLInsightIntradayQuotes, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get intraday quotes")
	}

	var resp []model.IntradayResponse
	err = q.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
