package quotes

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	sdk "github.com/OrbisSystems/orbis-sdk-go/interfaces"
	"github.com/OrbisSystems/orbis-sdk-go/model"
	"github.com/OrbisSystems/orbis-sdk-go/utils"
)

// Quotes service returns quotes data.
type Quotes struct {
	cli sdk.HTTPClient
}

func New(cli sdk.HTTPClient) *Quotes {
	return &Quotes{
		cli: cli,
	}
}

func (q *Quotes) GetQuotesEquityData(ctx context.Context, symbols, quoteType string) ([]model.QuoteEquityDataResponse, error) {
	log.Trace("GetQuotesEquityData called")

	r, err := q.cli.Get(ctx, fmt.Sprintf("%s?symbols=%s&quote_type=%s", model.URLInsightBase+model.URLInsightQuotesEquity, symbols, quoteType), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get quotes equity")
	}

	var resp []model.QuoteEquityDataResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
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

	r, err := q.cli.Post(ctx, model.URLInsightBase+model.URLInsightQuoteHistory, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.QuoteHistoryResponse{}, errors.Wrap(err, "couldn't get quote history")
	}

	var resp model.QuoteHistoryResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
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

	r, err := q.cli.Post(ctx, model.URLInsightBase+model.URLInsightIntradayQuotes, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get intraday quotes")
	}

	var resp []model.IntradayResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
