package tiprank

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	sdk "github.com/OrbisSystems/orbis-sdk-go/interface"
	"github.com/OrbisSystems/orbis-sdk-go/model"
	"github.com/OrbisSystems/orbis-sdk-go/utils"
)

type TipRank struct {
	sdk.Auth

	url string
	cli sdk.HTTPClient
}

func New(url string, auth sdk.Auth, cli sdk.HTTPClient) *TipRank {
	return &TipRank{
		Auth: auth,
		url:  url,
		cli:  cli,
	}
}

func (t *TipRank) AnalystConsensus(ctx context.Context, req model.AnalystConsensusRequest) ([]model.AnalystConsensusResponse, error) {
	log.Trace("AnalystConsensus called")

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := t.cli.Post(ctx, t.url+model.URLInsightBase+model.URLInsightTipRankAnalystConsensus, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get analyst consensus")
	}

	var resp []model.AnalystConsensusResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (t *TipRank) LatestAnalystRatingsOnStock(ctx context.Context, req model.LatestAnalystRatingsOnStockRequest) ([]model.LatestAnalystRatingsOnStockResponse, error) {
	log.Trace("LatestAnalystRatingsOnStock called")

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := t.cli.Post(ctx, t.url+model.URLInsightBase+model.URLInsightTipRankAnalystMulti, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get latest analyst ratings on stock")
	}

	var resp []model.LatestAnalystRatingsOnStockResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (t *TipRank) LiveFeed(ctx context.Context, req model.LiveFeedRequest) ([]model.LiveFeedResponse, error) {
	log.Trace("LiveFeed called")

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := t.cli.Post(ctx, t.url+model.URLInsightBase+model.URLInsightTipRankLiveFeed, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get live feed")
	}

	var resp []model.LiveFeedResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (t *TipRank) TrendingStocks(ctx context.Context, req model.TrendingStocksRequest) ([]model.TrendingStocksResponse, error) {
	log.Trace("TrendingStocks called")

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := t.cli.Post(ctx, t.url+model.URLInsightBase+model.URLInsightTipRankTrendingStocks, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get trending stocks")
	}

	var resp []model.TrendingStocksResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (t *TipRank) AnalystPortfolios(ctx context.Context, req model.PortfoliosRequest) ([]model.PortfoliosResponse, error) {
	log.Trace("AnalystPortfolios called")

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := t.cli.Post(ctx, t.url+model.URLInsightBase+model.URLInsightTipRankAnalystPortfolio, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get analyst portfolios")
	}

	var resp []model.PortfoliosResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (t *TipRank) AnalystProfile(ctx context.Context, id string) (model.AnalystProfileResponse, error) {
	log.Trace("AnalystProfile called")

	r, err := t.cli.Get(ctx, fmt.Sprintf("%s?id=%s", t.url+model.URLInsightBase+model.URLInsightTipRankAnalystProfile, id), nil)
	if err != nil {
		return model.AnalystProfileResponse{}, errors.Wrap(err, "couldn't get analyst profile")
	}

	var resp model.AnalystProfileResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.AnalystProfileResponse{}, err
	}

	return resp, err
}

func (t *TipRank) SectorConsensus(ctx context.Context) (model.SectorConsensusResponse, error) {
	log.Trace("SectorConsensus called")

	r, err := t.cli.Get(ctx, t.url+model.URLInsightBase+model.URLInsightTipRankSectorConsensus, nil)
	if err != nil {
		return model.SectorConsensusResponse{}, errors.Wrap(err, "couldn't get sector consensus")
	}

	var resp model.SectorConsensusResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.SectorConsensusResponse{}, err
	}

	return resp, err
}

func (t *TipRank) BestPerformingExperts(ctx context.Context, num int) ([]model.BestPerformingExpertsResponse, error) {
	log.Trace("BestPerformingExperts called")

	r, err := t.cli.Get(ctx, fmt.Sprintf("%s?num=%d", t.url+model.URLInsightBase+model.URLInsightTipRankBestPerformingExperts, num), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get best performing experts")
	}

	var resp []model.BestPerformingExpertsResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (t *TipRank) StocksSimilarStocks(ctx context.Context, ticker string) ([]string, error) {
	log.Trace("StocksSimilarStocks called")

	r, err := t.cli.Get(ctx, fmt.Sprintf("%s?ticker=%s", t.url+model.URLInsightBase+model.URLInsightTipRankStocksSimilarStocks, ticker), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get stocks similar stocks")
	}

	var resp []string
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (t *TipRank) AnalystsExpertPictureStore(ctx context.Context) (model.AnalystsExpertPictureStoreResponse, error) {
	log.Trace("AnalystsExpertPictureStore called")

	r, err := t.cli.Get(ctx, t.url+model.URLInsightBase+model.URLInsightTipRankAnalystsExpertPictureStore, nil)
	if err != nil {
		return model.AnalystsExpertPictureStoreResponse{}, errors.Wrap(err, "couldn't get analysts expert picture store")
	}

	var resp model.AnalystsExpertPictureStoreResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.AnalystsExpertPictureStoreResponse{}, err
	}

	return resp, err
}

func (t *TipRank) SupportedTickers(ctx context.Context) (model.SupportedTickersResponse, error) {
	log.Trace("SupportedTickers called")

	r, err := t.cli.Get(ctx, t.url+model.URLInsightBase+model.URLInsightTipRankSupportedTickers, nil)
	if err != nil {
		return model.SupportedTickersResponse{}, errors.Wrap(err, "couldn't get supported tickers")
	}

	var resp model.SupportedTickersResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.SupportedTickersResponse{}, err
	}

	return resp, err
}

func (t *TipRank) GeneralStockUpdates(ctx context.Context, utcTime, details string) (model.GeneralStockUpdatesResponse, error) {
	log.Trace("GeneralStockUpdates called")

	r, err := t.cli.Get(ctx, fmt.Sprintf("%s?utc_time=%s&details=%s", t.url+model.URLInsightBase+model.URLInsightTipRankGeneralStockUpdates, utcTime, details), nil)
	if err != nil {
		return model.GeneralStockUpdatesResponse{}, errors.Wrap(err, "couldn't get general stock updates")
	}

	var resp model.GeneralStockUpdatesResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.GeneralStockUpdatesResponse{}, err
	}

	return resp, err
}

func (t *TipRank) InsidersOverview(ctx context.Context, expertUID string) (model.InsidersOverviewResponse, error) {
	log.Trace("InsidersOverview called")

	r, err := t.cli.Get(ctx, fmt.Sprintf("%s?expert_uid=%s", t.url+model.URLInsightBase+model.URLInsightTipRankInsidersOverview, expertUID), nil)
	if err != nil {
		return model.InsidersOverviewResponse{}, errors.Wrap(err, "couldn't get insiders overview")
	}

	var resp model.InsidersOverviewResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.InsidersOverviewResponse{}, err
	}

	return resp, err
}

func (t *TipRank) InsidersBestPerformingExperts(ctx context.Context, num int) ([]model.InsidersBestPerformingExpertsResponse, error) {
	log.Trace("InsidersBestPerformingExperts called")

	r, err := t.cli.Get(ctx, fmt.Sprintf("%s?num=%d", t.url+model.URLInsightBase+model.URLInsightTipRankInsidersBestPerformingExperts, num), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get insiders best performing experts")
	}

	var resp []model.InsidersBestPerformingExpertsResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (t *TipRank) InsidersLiveFeed(ctx context.Context, num int, sort string) ([]model.InsidersLiveFeedResponse, error) {
	log.Trace("InsidersLiveFeed called")

	r, err := t.cli.Get(ctx, fmt.Sprintf("%s?num=%d&sort=%s", t.url+model.URLInsightBase+model.URLInsightTipRankInsidersLiveFeed, num, sort), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get insiders live feed")
	}

	var resp []model.InsidersLiveFeedResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (t *TipRank) HedgeFundsBestPerformingExperts(ctx context.Context, num int) ([]model.HedgeFundsBestPerformingExpertsResponse, error) {
	log.Trace("HedgeFundsBestPerformingExperts called")

	r, err := t.cli.Get(ctx, fmt.Sprintf("%s?num=%d", t.url+model.URLInsightBase+model.URLInsightTipRankHedgefundsBestPerformingExperts, num), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hedgefunds best performing experts")
	}

	var resp []model.HedgeFundsBestPerformingExpertsResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
