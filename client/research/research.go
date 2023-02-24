package research

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	sdk "github.com/OrbisSystems/orbis-sdk-go"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

type Research struct {
	sdk.Auth

	cfg config.Config
	cli sdk.HTTPClient
}

func New(cfg config.Config, auth sdk.Auth, cli sdk.HTTPClient) *Research {
	return &Research{
		Auth: auth,
		cfg:  cfg,
		cli:  cli,
	}
}

func (s *Research) GetCompanyProfile(ctx context.Context, symbol string) (model.CompanyProfile, error) {
	log.Trace("GetCompanyProfile called")

	r, err := s.cli.Get(ctx, fmt.Sprintf("%s?symbol=%s", s.cfg.AuthHost+model.URLInsightBase+model.URLInsightCompanyProfile, symbol), nil)
	if err != nil {
		return model.CompanyProfile{}, errors.Wrap(err, "couldn't get company profile")
	}

	var resp model.CompanyProfile
	err = s.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.CompanyProfile{}, err
	}

	return resp, err
}

func (s *Research) GetCombinedProfile(ctx context.Context, symbol string) (model.CompanyProfile, error) {
	log.Trace("GetCombinedProfile called")

	r, err := s.cli.Get(ctx, fmt.Sprintf("%s?symbol=%s", s.cfg.AuthHost+model.URLInsightBase+model.URLInsightQuoteProfile, symbol), nil)
	if err != nil {
		return model.CompanyProfile{}, errors.Wrap(err, "couldn't get combined company profile")
	}

	var resp model.CompanyProfile
	err = s.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.CompanyProfile{}, err
	}

	return resp, err
}

func (s *Research) GetOwnershipsBySymbol(ctx context.Context, req model.GetOwnershipsBySymbolRequest) (model.GetOwnershipsBySymbolResponse, error) {
	log.Trace("GetOwnershipsBySymbol called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.GetOwnershipsBySymbolResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := s.cli.Post(ctx, s.cfg.AuthHost+model.URLInsightBase+model.URLInsightSymbolOwnerships, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.GetOwnershipsBySymbolResponse{}, errors.Wrap(err, "couldn't get ownership by symbol")
	}

	var resp model.GetOwnershipsBySymbolResponse
	err = s.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.GetOwnershipsBySymbolResponse{}, err
	}

	return resp, err
}

func (s *Research) GetOwnershipsByID(ctx context.Context, req model.GetOwnershipsByIDRequest) (model.GetOwnershipsBySymbolResponse, error) {
	log.Trace("GetOwnershipsByID called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.GetOwnershipsBySymbolResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := s.cli.Post(ctx, s.cfg.AuthHost+model.URLInsightBase+model.URLInsightOwnerships, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.GetOwnershipsBySymbolResponse{}, errors.Wrap(err, "couldn't get ownership by id")
	}

	var resp model.GetOwnershipsBySymbolResponse
	err = s.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.GetOwnershipsBySymbolResponse{}, err
	}

	return resp, err
}

func (s *Research) GetEarningReleases(ctx context.Context, req model.EarningReleasesRequest) (model.EarningReleasesResponse, error) {
	log.Trace("GetEarningReleases called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.EarningReleasesResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := s.cli.Post(ctx, s.cfg.AuthHost+model.URLInsightBase+model.URLInsightEarningsCalendar, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.EarningReleasesResponse{}, errors.Wrap(err, "couldn't get earning release")
	}

	var resp model.EarningReleasesResponse
	err = s.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.EarningReleasesResponse{}, err
	}

	return resp, err
}

func (s *Research) GetSymbolFundamentals(ctx context.Context, req model.EarningReleasesRequest) (model.GetSymbolFundamentalsResponse, error) {
	log.Trace("GetSymbolFundamentals called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.GetSymbolFundamentalsResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := s.cli.Post(ctx, s.cfg.AuthHost+model.URLInsightBase+model.URLInsightFundamentals, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.GetSymbolFundamentalsResponse{}, errors.Wrap(err, "couldn't get fundamentals")
	}

	var resp model.GetSymbolFundamentalsResponse
	err = s.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.GetSymbolFundamentalsResponse{}, err
	}

	return resp, err
}

func (s *Research) Screener(ctx context.Context, req model.StockScreenerRequest) (model.StockScreenerResponse, error) {
	log.Trace("Screener called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.StockScreenerResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := s.cli.Post(ctx, s.cfg.AuthHost+model.URLInsightBase+model.URLInsightStockScreener, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.StockScreenerResponse{}, errors.Wrap(err, "couldn't get fundamentals")
	}

	var resp model.StockScreenerResponse
	err = s.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.StockScreenerResponse{}, err
	}

	return resp, err
}

func (s *Research) StockMarketHeatmap(ctx context.Context, heatmapName, quoteType string) (model.StockMarketHeatmapResponse, error) {
	log.Trace("StockMarketHeatmap called")

	r, err := s.cli.Get(ctx, fmt.Sprintf("%s?heatmapName=%s&quoteType=%s", s.cfg.AuthHost+model.URLInsightBase+model.URLInsightHeatmaps, heatmapName, quoteType), nil)
	if err != nil {
		return model.StockMarketHeatmapResponse{}, errors.Wrap(err, "couldn't get combined company profile")
	}

	var resp model.StockMarketHeatmapResponse
	err = s.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.StockMarketHeatmapResponse{}, err
	}

	return resp, err
}

func (s *Research) GetIndustriesPerformance(ctx context.Context, req model.GetIndustriesPerformanceRequest) (model.GetIndustriesPerformanceResponse, error) {
	log.Trace("GetIndustriesPerformance called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.GetIndustriesPerformanceResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := s.cli.Post(ctx, s.cfg.AuthHost+model.URLInsightBase+model.URLInsightIndustriesPerformance, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.GetIndustriesPerformanceResponse{}, errors.Wrap(err, "couldn't get fundamentals")
	}

	var resp model.GetIndustriesPerformanceResponse
	err = s.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.GetIndustriesPerformanceResponse{}, err
	}

	return resp, err
}
