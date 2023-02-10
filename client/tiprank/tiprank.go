package tiprank

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	sdk "github.com/OrbisSystems/orbis-sdk-go"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

type TipRank struct {
	sdk.Auth

	cfg       config.Config
	cli       sdk.HTTPClient
	validator sdk.Validator
}

func New(cfg config.Config, auth sdk.Auth, cli sdk.HTTPClient, validator sdk.Validator) *TipRank {
	return &TipRank{
		Auth:      auth,
		cfg:       cfg,
		cli:       cli,
		validator: validator,
	}
}

// TODO add validator

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
