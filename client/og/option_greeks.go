package og

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	sdk "github.com/OrbisSystems/orbis-sdk-go"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

type OptionGreeks struct {
	sdk.Auth

	cfg config.Config
	cli sdk.HTTPClient
}

func New(cfg config.Config, auth sdk.Auth, cli sdk.HTTPClient) *OptionGreeks {
	return &OptionGreeks{
		Auth: auth,
		cfg:  cfg,
		cli:  cli,
	}
}

func (og *OptionGreeks) CalculateParams(ctx context.Context, req model.CalculateParamsRequest) (model.CalculateParamsResponse, error) {
	log.Trace("CalculateParams called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.CalculateParamsResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := og.cli.Post(ctx, og.cfg.AuthHost+model.URLInsightBase+model.URLInsightOptionGreeksCalculateParams, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.CalculateParamsResponse{}, errors.Wrap(err, "couldn't get og calculated params")
	}

	var resp model.CalculateParamsResponse
	err = og.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.CalculateParamsResponse{}, err
	}

	return resp, err
}

func (og *OptionGreeks) CalculateMatrix(ctx context.Context, req model.CalculateParamsRequest) (model.CalculateMatrixParamsRequest, error) {
	log.Trace("CalculateMatrix called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.CalculateMatrixParamsRequest{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := og.cli.Post(ctx, og.cfg.AuthHost+model.URLInsightBase+model.URLInsightOptionGreeksCalculateMatrix, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.CalculateMatrixParamsRequest{}, errors.Wrap(err, "couldn't get og calculated matrix")
	}

	var resp model.CalculateMatrixParamsRequest
	err = og.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.CalculateMatrixParamsRequest{}, err
	}

	return resp, err
}
