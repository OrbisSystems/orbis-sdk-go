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

func (i *HoopsAI) DailySummary(ctx context.Context, req model.HSDailySummaryRequest) (model.HSDailySummaryResponse, error) {
	log.Trace("HS DailySummary called")

	urlQuery, err := utils.BuildURLQueryParams(req)
	if err != nil {
		return model.HSDailySummaryResponse{}, err
	}

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s/%s?%s", model.URLInsightBase+model.URLInsightHSDailySummary, req.Asset, urlQuery), nil)
	if err != nil {
		return model.HSDailySummaryResponse{}, errors.Wrap(err, "couldn't get hs daily summery")
	}

	var resp model.HSDailySummaryResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.HSDailySummaryResponse{}, err
	}

	return resp, err
}
