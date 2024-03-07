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
	cli    sdk.HTTPClient
	logger *log.Logger
}

func New(cli sdk.HTTPClient, logger *log.Logger) *HoopsAI {
	return &HoopsAI{
		cli:    cli,
		logger: logger,
	}
}

func (i *HoopsAI) DailySummary(ctx context.Context, req model.HSDailySummaryRequest) (map[string]interface{}, error) {
	i.logger.Trace("HS DailySummary called")

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
	i.logger.Trace("HS WeeklySummary called")

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
	i.logger.Trace("HS Portfolio called")

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
	i.logger.Trace("HS Watchlist called")

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
	i.logger.Trace("HS WatchlistByUserAndName called")

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
	i.logger.Trace("HS TopGainers called")

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
	i.logger.Trace("HS TopLosers called")

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
	i.logger.Trace("HS TopMovers called")

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
	i.logger.Trace("HS DownTrend called")

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
	i.logger.Trace("HS UpTrend called")

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
	i.logger.Trace("HS MarketOverview called")

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
	i.logger.Trace("HS PriceTarget called")

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
	i.logger.Trace("HS UpcomingEarnings called")

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
	i.logger.Trace("HS RecentEarnings called")

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
	i.logger.Trace("HS RecordHigh called")

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
	i.logger.Trace("HS UnusualHighVolume called")

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
	i.logger.Trace("HS Data called")

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
	i.logger.Trace("HS CustomerAssets called")

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
	i.logger.Trace("HS GetUsers called")

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s/%s", model.URLInsightBase+model.URLInsightHSUsers, customer), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't hs get users")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) CreateUser(ctx context.Context, customer string) (map[string]interface{}, error) {
	i.logger.Trace("HS CreateUser called")

	r, err := i.cli.Post(ctx, fmt.Sprintf("%s/%s", model.URLInsightBase+model.URLInsightHSUsers, customer), nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't hs create user")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) DeleteUser(ctx context.Context, customer, userID string) (map[string]interface{}, error) {
	i.logger.Trace("HS DeleteUser called")

	r, err := i.cli.Delete(ctx, fmt.Sprintf("%s/%s/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID), nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't hs delete user")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) CreateWatchlistByUser(ctx context.Context, customer, userID string) (map[string]interface{}, error) {
	i.logger.Trace("HS CreateWatchlistByUser called")

	r, err := i.cli.Post(ctx, fmt.Sprintf("%s/%s/%s/watchlists", model.URLInsightBase+model.URLInsightHSUsers, customer, userID), nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't hs create watchlist by user")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) GetWatchlistByUser(ctx context.Context, customer, userID, watchlistName string) (map[string]interface{}, error) {
	i.logger.Trace("HS GetWatchlistByUser called")

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, watchlistName), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't hs get watchlist by user")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) AddSymbolToWatchlist(ctx context.Context, customer, userID, watchlistName, symbol string) (map[string]interface{}, error) {
	i.logger.Trace("HS AddSymbolToWatchlist called")

	r, err := i.cli.Post(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, watchlistName, symbol), nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't hs add symbol to watchlist")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) DeleteSymbolFromWatchlist(ctx context.Context, customer, userID, watchlistName, symbol string) (map[string]interface{}, error) {
	i.logger.Trace("HS DeleteSymbolFromWatchlist called")

	r, err := i.cli.Delete(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, watchlistName, symbol), nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't hs delete symbol from watchlist")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) AddSymbolsToWatchlist(ctx context.Context, customer, userID, watchlistName string, symbols []string) (map[string]interface{}, error) {
	i.logger.Trace("HS AddSymbolToWatchlist called")

	body, err := json.Marshal(struct {
		SymbolList string `json:"symbol_list"`
	}{
		SymbolList: utils.ArrayToJoinedString(symbols),
	})
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := i.cli.Post(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, watchlistName), bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't hs add symbols to watchlist")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) DeleteWatchlistByName(ctx context.Context, customer, userID, watchlistName string) (map[string]interface{}, error) {
	i.logger.Trace("HS DeleteWatchlistByName called")

	r, err := i.cli.Delete(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, watchlistName), nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't hs delete watchlist by name")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (i *HoopsAI) RenameWatchlist(ctx context.Context, customer, userID, oldWatchlistName, newWatchlistName string) (map[string]interface{}, error) {
	i.logger.Trace("HS RenameWatchlist called")

	r, err := i.cli.Put(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s/rename/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, oldWatchlistName, newWatchlistName), nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't hs rename watchlist")
	}

	var resp map[string]interface{}
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
