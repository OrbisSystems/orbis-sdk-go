package funds

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

// Funds service returns different funds info such as screener, top, etc.
type Funds struct {
	cli sdk.HTTPClient
}

func New(cli sdk.HTTPClient) *Funds {
	return &Funds{
		cli: cli,
	}
}

func (f *Funds) GetFundDetails(ctx context.Context, symbol string) (model.GetFundDetailsResponse, error) {
	log.Trace("GetFundDetails called")

	r, err := f.cli.Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightFundsDetails, symbol), nil)
	if err != nil {
		return model.GetFundDetailsResponse{}, errors.Wrap(err, "couldn't get funds details")
	}

	var resp model.GetFundDetailsResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.GetFundDetailsResponse{}, err
	}

	return resp, err
}

func (f *Funds) GetFundScreenerFilters(ctx context.Context) (model.GetFundScreenerFiltersResponse, error) {
	log.Trace("GetFundScreenerFilters called")

	r, err := f.cli.Get(ctx, model.URLInsightBase+model.URLInsightFundsScreenerFilters, nil)
	if err != nil {
		return model.GetFundScreenerFiltersResponse{}, errors.Wrap(err, "couldn't get funds screener filters")
	}

	var resp model.GetFundScreenerFiltersResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.GetFundScreenerFiltersResponse{}, err
	}

	return resp, err
}

func (f *Funds) ScreenFunds(ctx context.Context, req model.FundScreenerRequest) (model.FundScreenerResponse, error) {
	log.Trace("ScreenFunds called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.FundScreenerResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := f.cli.Post(ctx, model.URLInsightBase+model.URLInsightFundsScreener, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.FundScreenerResponse{}, errors.Wrap(err, "couldn't get screen funds")
	}

	var resp model.FundScreenerResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.FundScreenerResponse{}, err
	}

	return resp, err
}

func (f *Funds) ScreenSectorFunds(ctx context.Context, req model.FundSectorScreenerRequest) (model.FundScreenerResponse, error) {
	log.Trace("ScreenSectorFunds called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.FundScreenerResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := f.cli.Post(ctx, model.URLInsightBase+model.URLInsightFundsSectorScreener, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.FundScreenerResponse{}, errors.Wrap(err, "couldn't get screen sector funds")
	}

	var resp model.FundScreenerResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.FundScreenerResponse{}, err
	}

	return resp, err
}

func (f *Funds) GetTopFunds(ctx context.Context, req model.GetTopFundsRequest) ([]string, error) {
	log.Trace("GetTopFunds called")

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := f.cli.Post(ctx, model.URLInsightBase+model.URLInsightFundsTop, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get top funds")
	}

	var resp []string
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
