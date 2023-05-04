package hs

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

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := i.cli.Post(ctx, model.URLInsightBase+model.URLInsightHSPortfolio, bytes.NewBuffer(body), nil)
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

func (i *HoopsAI) Watchlist(ctx context.Context, req model.HSWatchlistRequest) (map[string]interface{}, error) {
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

func (i *HoopsAI) WatchlistByUserAndName(ctx context.Context, req model.HSWatchlistByUserAndNameRequest) (map[string]interface{}, error) {
	log.Trace("HS WatchlistByUserAndName called")

	urlQuery, err := utils.BuildURLQueryParams(req)
	if err != nil {
		return nil, err
	}

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s/%s/%s?%s", model.URLInsightBase+model.URLInsightHSWatchlist, req.UserID, req.WatchlistName, urlQuery), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs watchlist by user and name")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) TopGainers(ctx context.Context, req model.HSMarketResearchRequest) (map[string]interface{}, error) {
	log.Trace("HS TopGainers called")

	urlQuery, err := utils.BuildURLQueryParams(req)
	if err != nil {
		return nil, err
	}

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s?%s", model.URLInsightBase+model.URLInsightHSTopGainers, urlQuery), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs top gainers")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) TopLosers(ctx context.Context, req model.HSMarketResearchRequest) (map[string]interface{}, error) {
	log.Trace("HS TopLosers called")

	urlQuery, err := utils.BuildURLQueryParams(req)
	if err != nil {
		return nil, err
	}

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s?%s", model.URLInsightBase+model.URLInsightHSTopLosers, urlQuery), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs top losers")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) TopMovers(ctx context.Context, req model.HSMarketResearchRequest) (map[string]interface{}, error) {
	log.Trace("HS TopMovers called")

	urlQuery, err := utils.BuildURLQueryParams(req)
	if err != nil {
		return nil, err
	}

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s?%s", model.URLInsightBase+model.URLInsightHSTopMovers, urlQuery), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs top movers")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) DownTrend(ctx context.Context, req model.HSMarketResearchRequest) (map[string]interface{}, error) {
	log.Trace("HS DownTrend called")

	urlQuery, err := utils.BuildURLQueryParams(req)
	if err != nil {
		return nil, err
	}

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s?%s", model.URLInsightBase+model.URLInsightHSDownTrend, urlQuery), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs down trend")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) UpTrend(ctx context.Context, req model.HSMarketResearchRequest) (map[string]interface{}, error) {
	log.Trace("HS UpTrend called")

	urlQuery, err := utils.BuildURLQueryParams(req)
	if err != nil {
		return nil, err
	}

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s?%s", model.URLInsightBase+model.URLInsightHSUpTrend, urlQuery), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs up trend")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) MarketOverview(ctx context.Context, req model.HSMarketResearchRequest) (map[string]interface{}, error) {
	log.Trace("HS MarketOverview called")

	urlQuery, err := utils.BuildURLQueryParams(req)
	if err != nil {
		return nil, err
	}

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s?%s", model.URLInsightBase+model.URLInsightHSMarketOverview, urlQuery), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs market overview")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) PriceTarget(ctx context.Context, req model.HSMarketResearchRequest) (map[string]interface{}, error) {
	log.Trace("HS PriceTarget called")

	urlQuery, err := utils.BuildURLQueryParams(req)
	if err != nil {
		return nil, err
	}

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s?%s", model.URLInsightBase+model.URLInsightHSPriceTarget, urlQuery), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs price target")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) UpcomingEarnings(ctx context.Context, req model.HSMarketResearchRequest) (map[string]interface{}, error) {
	log.Trace("HS UpcomingEarnings called")

	urlQuery, err := utils.BuildURLQueryParams(req)
	if err != nil {
		return nil, err
	}

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s?%s", model.URLInsightBase+model.URLInsightHSUpcomingEarnings, urlQuery), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs upcoming earnings")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) RecentEarnings(ctx context.Context, req model.HSMarketResearchRequest) (map[string]interface{}, error) {
	log.Trace("HS RecentEarnings called")

	urlQuery, err := utils.BuildURLQueryParams(req)
	if err != nil {
		return nil, err
	}

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s?%s", model.URLInsightBase+model.URLInsightHSRecentEarnings, urlQuery), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs recent earnings")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) RecordHigh(ctx context.Context, req model.HSMarketResearchRequest) (map[string]interface{}, error) {
	log.Trace("HS RecordHigh called")

	urlQuery, err := utils.BuildURLQueryParams(req)
	if err != nil {
		return nil, err
	}

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s?%s", model.URLInsightBase+model.URLInsightHSRecordHigh, urlQuery), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs record high")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) UnusualHighVolume(ctx context.Context, req model.HSMarketResearchRequest) (map[string]interface{}, error) {
	log.Trace("HS UnusualHighVolume called")

	urlQuery, err := utils.BuildURLQueryParams(req)
	if err != nil {
		return nil, err
	}

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s?%s", model.URLInsightBase+model.URLInsightHSUnusualHighVolume, urlQuery), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs unusual high volume")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) Data(ctx context.Context, req model.HSMarketResearchRequest) (map[string]interface{}, error) {
	log.Trace("HS Data called")

	urlQuery, err := utils.BuildURLQueryParams(req)
	if err != nil {
		return nil, err
	}

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s?%s", model.URLInsightBase+model.URLInsightHSAssets, urlQuery), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs assets")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) CustomerAssets(ctx context.Context, req model.HSMarketResearchRequest) (map[string]interface{}, error) {
	log.Trace("HS CustomerAssets called")

	urlQuery, err := utils.BuildURLQueryParams(req)
	if err != nil {
		return nil, err
	}

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s?%s", model.URLInsightBase+model.URLInsightHSCustomerAssets, urlQuery), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs customer assets")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) GetUsers(ctx context.Context, customer string) (map[string]interface{}, error) {
	log.Trace("HS GetUsers called")

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s/%s", model.URLInsightBase+model.URLInsightHSUsers, customer), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs get users")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) CreateUser(ctx context.Context, customer string) (map[string]interface{}, error) {
	log.Trace("HS CreateUser called")

	r, err := i.cli.Post(ctx, fmt.Sprintf("%s/%s", model.URLInsightBase+model.URLInsightHSUsers, customer), nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs create user")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) DeleteUser(ctx context.Context, customer, userID string) (map[string]interface{}, error) {
	log.Trace("HS DeleteUser called")

	r, err := i.cli.Delete(ctx, fmt.Sprintf("%s/%s/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID), nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hs delete user")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
