package funds

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	sdk "github.com/OrbisSystems/orbis-sdk-go"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

// Funds service returns different funds info such as screener, top, etc.
type Funds struct {
	sdk.Auth

	cfg config.Config
	cli sdk.HTTPClient
}

func New(cfg config.Config, auth sdk.Auth, cli sdk.HTTPClient) *Funds {
	return &Funds{
		Auth: auth,
		cfg:  cfg,
		cli:  cli,
	}
}

func (f *Funds) GetFundDetails(ctx context.Context, symbol string) (model.GetFundDetailsResponse, error) {
	r, err := f.cli.Get(ctx, fmt.Sprintf("%s?symbol=%s", f.cfg.AuthHost+model.URLInsightBase+model.URLInsightFundsDetails, symbol), nil)
	if err != nil {
		return model.GetFundDetailsResponse{}, errors.Wrap(err, "couldn't get funds details")
	}

	var resp model.GetFundDetailsResponse
	err = f.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.GetFundDetailsResponse{}, err
	}

	return resp, err
}

func (f *Funds) GetFundScreenerFilters(ctx context.Context) (model.GetFundScreenerFiltersResponse, error) {
	r, err := f.cli.Get(ctx, f.cfg.AuthHost+model.URLInsightBase+model.URLInsightFundsScreenerFilters, nil)
	if err != nil {
		return model.GetFundScreenerFiltersResponse{}, errors.Wrap(err, "couldn't get funds screener filters")
	}

	var resp model.GetFundScreenerFiltersResponse
	err = f.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.GetFundScreenerFiltersResponse{}, err
	}

	return resp, err
}

func (f *Funds) ScreenFunds(ctx context.Context, req model.FundScreenerRequest) (model.FundScreenerResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return model.FundScreenerResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := f.cli.Post(ctx, f.cfg.AuthHost+model.URLInsightBase+model.URLInsightFundsScreener, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.FundScreenerResponse{}, errors.Wrap(err, "couldn't get screen funds")
	}

	var resp model.FundScreenerResponse
	err = f.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.FundScreenerResponse{}, err
	}

	return resp, err
}

func (f *Funds) ScreenSectorFunds(ctx context.Context, req model.FundSectorScreenerRequest) (model.FundScreenerResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return model.FundScreenerResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := f.cli.Post(ctx, f.cfg.AuthHost+model.URLInsightBase+model.URLInsightFundsSectorScreener, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.FundScreenerResponse{}, errors.Wrap(err, "couldn't get screen sector funds")
	}

	var resp model.FundScreenerResponse
	err = f.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.FundScreenerResponse{}, err
	}

	return resp, err
}

func (f *Funds) GetTopFunds(ctx context.Context, req model.GetTopFundsRequest) ([]string, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := f.cli.Post(ctx, f.cfg.AuthHost+model.URLInsightBase+model.URLInsightFundsTop, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get top funds")
	}

	var resp []string
	err = f.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
