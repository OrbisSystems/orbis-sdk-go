package equity

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	sdk "github.com/OrbisSystems/orbis-sdk-go"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

// Equity service returns equity quotes.
type Equity struct {
	sdk.Auth

	cfg config.Config
	cli sdk.HTTPClient
}

func New(cfg config.Config, auth sdk.Auth, cli sdk.HTTPClient) *Equity {
	return &Equity{
		Auth: auth,
		cfg:  cfg,
		cli:  cli,
	}
}

func (e *Equity) GetQuotesEquityData(ctx context.Context, symbols, quoteType string) ([]model.EquityDataResponse, error) {
	r, err := e.cli.Get(ctx, fmt.Sprintf("%s?symbols=%s&quote_type=%s", e.cfg.AuthHost+model.URLInsightBase+model.URLInsightQuotesEquity, symbols, quoteType), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get quotes equity")
	}

	var resp []model.EquityDataResponse
	err = e.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
