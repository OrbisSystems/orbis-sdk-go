package hs

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	sdk "github.com/OrbisSystems/orbis-sdk-go/interfaces"
	"github.com/OrbisSystems/orbis-sdk-go/model"
	"github.com/OrbisSystems/orbis-sdk-go/utils"
)

type HoopsAI struct {
	cli sdk.HTTPClient
}

func New(cli sdk.HTTPClient) *HoopsAI {
	return &HoopsAI{
		cli: cli,
	}
}

func (i *HoopsAI) DailySummary(ctx context.Context, req model.HSDailySummaryRequest) (map[string]interface{}, error) {
	log.Trace("HS DailySummary called")

	urlQuery, err := utils.BuildURLQueryParams(req)
	if err != nil {
		return nil, err
	}

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s/%s?%s", model.URLInsightBase+model.URLInsightHSDailySummary, req.Asset, urlQuery), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs daily summery")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) WeeklySummary(ctx context.Context, req model.HSWeeklySummaryRequest) (map[string]interface{}, error) {
	log.Trace("HS WeeklySummary called")

	urlQuery, err := utils.BuildURLQueryParams(req)
	if err != nil {
		return nil, err
	}

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s/%s?%s", model.URLInsightBase+model.URLInsightHSWeeklySummary, req.Asset, urlQuery), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs weekly summery")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) Portfolio(ctx context.Context, req model.HSPortfolioRequest) (map[string]interface{}, error) {
	log.Trace("HS Portfolio called")

	urlQuery, err := utils.BuildURLQueryParams(req)
	if err != nil {
		return nil, err
	}

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s?%s", model.URLInsightBase+model.URLInsightHSPortfolio, urlQuery), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs portfolio")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) Watchlist(ctx context.Context, req model.HWatchlistRequest) (map[string]interface{}, error) {
	log.Trace("HS Watchlist called")

	urlQuery, err := utils.BuildURLQueryParams(req)
	if err != nil {
		return nil, err
	}

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s/%s?%s", model.URLInsightBase+model.URLInsightHSWatchlist, req.Asset, urlQuery), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs watchlist")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
