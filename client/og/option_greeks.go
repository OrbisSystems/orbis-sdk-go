package og

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	sdk "github.com/OrbisSystems/orbis-sdk-go/interfaces"
	"github.com/OrbisSystems/orbis-sdk-go/model"
	"github.com/OrbisSystems/orbis-sdk-go/utils"
)

type OptionGreeks struct {
	cli sdk.HTTPClient
}

func New(cli sdk.HTTPClient) *OptionGreeks {
	return &OptionGreeks{
		cli: cli,
	}
}

func (og *OptionGreeks) CalculateParams(ctx context.Context, req model.CalculateParamsRequest) (model.CalculateParamsResponse, error) {
	log.Trace("CalculateParams called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.CalculateParamsResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := og.cli.Post(ctx, model.URLInsightBase+model.URLInsightOptionGreeksCalculateParams, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.CalculateParamsResponse{}, errors.Wrap(err, "couldn't get og calculated params")
	}

	var resp model.CalculateParamsResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.CalculateParamsResponse{}, err
	}

	return resp, err
}

func (og *OptionGreeks) CalculateMatrix(ctx context.Context, req model.CalculateMatrixParamsRequest) (model.CalculateMatrixResponse, error) {
	log.Trace("CalculateMatrix called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.CalculateMatrixResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := og.cli.Post(ctx, model.URLInsightBase+model.URLInsightOptionGreeksCalculateMatrix, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.CalculateMatrixResponse{}, errors.Wrap(err, "couldn't get og calculated matrix")
	}

	var resp model.CalculateMatrixResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.CalculateMatrixResponse{}, err
	}

	return resp, err
}
