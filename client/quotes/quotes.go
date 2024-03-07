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
	cli    sdk.HTTPClient
	logger *log.Logger
}

func New(cli sdk.HTTPClient, logger *log.Logger) *Quotes {
	return &Quotes{
		cli:    cli,
		logger: logger,
	}
}

func (q *Quotes) GetQuotesEquityData(ctx context.Context, symbols, quoteType string) ([]model.QuoteEquityDataResponse, error) {
	q.logger.Trace("GetQuotesEquityData called")

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
	q.logger.Trace("GetQuoteHistory called")

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
	q.logger.Trace("GetIntradayQuotes called")

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

func (q *Quotes) GetSingleHistoricalQuote(ctx context.Context, symbols, date string) (model.HistoricalQuote, error) {
	q.logger.Trace("GetSingleHistoricalQuote called")

	r, err := q.cli.Get(ctx, fmt.Sprintf("%s?symbols=%s&date=%s", model.URLInsightBase+model.URLInsightHistoricalQuote, symbols, date), nil)
	if err != nil {
		return model.HistoricalQuote{}, errors.Wrap(err, "couldn't get historical quote")
	}

	var resp model.HistoricalQuote
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.HistoricalQuote{}, err
	}

	return resp, err
}
