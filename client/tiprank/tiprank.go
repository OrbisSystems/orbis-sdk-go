package tiprank

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

type TipRank struct {
	sdk.Auth

	cfg config.Config
	cli sdk.HTTPClient
}

func New(cfg config.Config, auth sdk.Auth, cli sdk.HTTPClient) *TipRank {
	return &TipRank{
		Auth: auth,
		cfg:  cfg,
		cli:  cli,
	}
}

func (t *TipRank) AnalystConsensus(ctx context.Context, req model.AnalystConsensusRequest) ([]model.AnalystConsensusResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := t.cli.Post(ctx, t.cfg.AuthHost+model.URLInsightBase+model.URLInsightTipRankAnalystConsensus, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get analyst consensus")
	}

	var resp []model.AnalystConsensusResponse
	err = t.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (t *TipRank) LatestAnalystRatingsOnStock(ctx context.Context, req model.LatestAnalystRatingsOnStockRequest) ([]model.LatestAnalystRatingsOnStockResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := t.cli.Post(ctx, t.cfg.AuthHost+model.URLInsightBase+model.URLInsightTipRankAnalystMulti, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get latest analyst ratings on stock")
	}

	var resp []model.LatestAnalystRatingsOnStockResponse
	err = t.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (t *TipRank) LiveFeed(ctx context.Context, req model.LiveFeedRequest) ([]model.LiveFeedResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := t.cli.Post(ctx, t.cfg.AuthHost+model.URLInsightBase+model.URLInsightTipRankLiveFeed, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get live feed")
	}

	var resp []model.LiveFeedResponse
	err = t.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (t *TipRank) TrendingStocks(ctx context.Context, req model.TrendingStocksRequest) ([]model.TrendingStocksResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := t.cli.Post(ctx, t.cfg.AuthHost+model.URLInsightBase+model.URLInsightTipRankTrendingStocks, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get trending stocks")
	}

	var resp []model.TrendingStocksResponse
	err = t.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (t *TipRank) AnalystPortfolios(ctx context.Context, req model.PortfoliosRequest) ([]model.PortfoliosResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := t.cli.Post(ctx, t.cfg.AuthHost+model.URLInsightBase+model.URLInsightTipRankAnalystPortfolio, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get analyst portfolios")
	}

	var resp []model.PortfoliosResponse
	err = t.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (t *TipRank) AnalystProfile(ctx context.Context, id string) (model.AnalystProfileResponse, error) {
	r, err := t.cli.Get(ctx, fmt.Sprintf("%s?id=%s", t.cfg.AuthHost+model.URLInsightBase+model.URLInsightTipRankAnalystProfile, id), nil)
	if err != nil {
		return model.AnalystProfileResponse{}, errors.Wrap(err, "couldn't get analyst profile")
	}

	var resp model.AnalystProfileResponse
	err = t.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.AnalystProfileResponse{}, err
	}

	return resp, err
}

func (t *TipRank) SectorConsensus(ctx context.Context) (model.SectorConsensusResponse, error) {
	r, err := t.cli.Get(ctx, t.cfg.AuthHost+model.URLInsightBase+model.URLInsightTipRankSectorConsensus, nil)
	if err != nil {
		return model.SectorConsensusResponse{}, errors.Wrap(err, "couldn't get sector consensus")
	}

	var resp model.SectorConsensusResponse
	err = t.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.SectorConsensusResponse{}, err
	}

	return resp, err
}

func (t *TipRank) BestPerformingExperts(ctx context.Context, num int) ([]model.BestPerformingExpertsResponse, error) {
	r, err := t.cli.Get(ctx, fmt.Sprintf("%s?num=%d", t.cfg.AuthHost+model.URLInsightBase+model.URLInsightTipRankBestPerformingExperts, num), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get best performing experts")
	}

	var resp []model.BestPerformingExpertsResponse
	err = t.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (t *TipRank) StocksSimilarStocks(ctx context.Context, ticker string) ([]string, error) {
	r, err := t.cli.Get(ctx, fmt.Sprintf("%s?ticker=%s", t.cfg.AuthHost+model.URLInsightBase+model.URLInsightTipRankStocksSimilarStocks, ticker), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get stocks similar stocks")
	}

	var resp []string
	err = t.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (t *TipRank) AnalystsExpertPictureStore(ctx context.Context) (model.AnalystsExpertPictureStoreResponse, error) {
	r, err := t.cli.Get(ctx, t.cfg.AuthHost+model.URLInsightBase+model.URLInsightTipRankAnalystsExpertPictureStore, nil)
	if err != nil {
		return model.AnalystsExpertPictureStoreResponse{}, errors.Wrap(err, "couldn't get analysts expert picture store")
	}

	var resp model.AnalystsExpertPictureStoreResponse
	err = t.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.AnalystsExpertPictureStoreResponse{}, err
	}

	return resp, err
}

func (t *TipRank) SupportedTickers(ctx context.Context) (model.SupportedTickersResponse, error) {
	r, err := t.cli.Get(ctx, t.cfg.AuthHost+model.URLInsightBase+model.URLInsightTipRankSupportedTickers, nil)
	if err != nil {
		return model.SupportedTickersResponse{}, errors.Wrap(err, "couldn't get supported tickers")
	}

	var resp model.SupportedTickersResponse
	err = t.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.SupportedTickersResponse{}, err
	}

	return resp, err
}

func (t *TipRank) GeneralStockUpdates(ctx context.Context, utcTime, details string) (model.GeneralStockUpdatesResponse, error) {
	r, err := t.cli.Get(ctx, fmt.Sprintf("%s?utc_time=%s&details=%s", t.cfg.AuthHost+model.URLInsightBase+model.URLInsightTipRankGeneralStockUpdates, utcTime, details), nil)
	if err != nil {
		return model.GeneralStockUpdatesResponse{}, errors.Wrap(err, "couldn't get general stock updates")
	}

	var resp model.GeneralStockUpdatesResponse
	err = t.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.GeneralStockUpdatesResponse{}, err
	}

	return resp, err
}

func (t *TipRank) InsidersOverview(ctx context.Context, expertUID string) (model.InsidersOverviewResponse, error) {
	r, err := t.cli.Get(ctx, fmt.Sprintf("%s?expert_uid=%s", t.cfg.AuthHost+model.URLInsightBase+model.URLInsightTipRankInsidersOverview, expertUID), nil)
	if err != nil {
		return model.InsidersOverviewResponse{}, errors.Wrap(err, "couldn't get insiders overview")
	}

	var resp model.InsidersOverviewResponse
	err = t.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.InsidersOverviewResponse{}, err
	}

	return resp, err
}

func (t *TipRank) InsidersBestPerformingExperts(ctx context.Context, num int) ([]model.InsidersBestPerformingExpertsResponse, error) {
	r, err := t.cli.Get(ctx, fmt.Sprintf("%s?num=%d", t.cfg.AuthHost+model.URLInsightBase+model.URLInsightTipRankInsidersBestPerformingExperts, num), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get insiders best performing experts")
	}

	var resp []model.InsidersBestPerformingExpertsResponse
	err = t.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (t *TipRank) InsidersLiveFeed(ctx context.Context, num int, sort string) ([]model.InsidersLiveFeedResponse, error) {
	r, err := t.cli.Get(ctx, fmt.Sprintf("%s?num=%d&sort=%s", t.cfg.AuthHost+model.URLInsightBase+model.URLInsightTipRankInsidersLiveFeed, num, sort), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get insiders live feed")
	}

	var resp []model.InsidersLiveFeedResponse
	err = t.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (t *TipRank) HedgeFundsBestPerformingExperts(ctx context.Context, num int) ([]model.HedgeFundsBestPerformingExpertsResponse, error) {
	r, err := t.cli.Get(ctx, fmt.Sprintf("%s?num=%d", t.cfg.AuthHost+model.URLInsightBase+model.URLInsightTipRankHedgefundsBestPerformingExperts, num), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get hedgefunds best performing experts")
	}

	var resp []model.HedgeFundsBestPerformingExpertsResponse
	err = t.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
