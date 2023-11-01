package fi

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

// FixedIncome service provides information about fixed income
type FixedIncome struct {
	cli sdk.HTTPClient
}

func New(cli sdk.HTTPClient) *FixedIncome {
	return &FixedIncome{
		cli: cli,
	}
}

func (f *FixedIncome) GetFixedIncomeEntryByID(ctx context.Context, id string) (model.FixedIncome, error) {
	log.Trace("GetFixedIncomeEntryByID called")

	r, err := f.cli.Get(ctx, fmt.Sprintf("%s/%s", model.URLInsightBase+model.URLInsightFixedIncome, id), nil)
	if err != nil {
		return model.FixedIncome{}, errors.Wrap(err, "couldn't get fixed income by id")
	}

	var resp model.FixedIncome
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.FixedIncome{}, err
	}

	return resp, err
}

func (f *FixedIncome) GetFixedIncomeEntries(ctx context.Context, req model.GetFixedIncomeEntriesRequest) (model.GetFixedIncomeEntriesResponse, error) {
	log.Trace("GetFixedIncomeEntries called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.GetFixedIncomeEntriesResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := f.cli.Post(ctx, model.URLInsightBase+model.URLInsightFixedIncome, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.GetFixedIncomeEntriesResponse{}, errors.Wrap(err, "couldn't get fixed income entries")
	}

	var resp model.GetFixedIncomeEntriesResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.GetFixedIncomeEntriesResponse{}, err
	}

	return resp, err
}

func (f *FixedIncome) GetFixedIncomeHistorical(ctx context.Context, req model.GetFixedIncomeHistoricalRequest) (model.GetFixedIncomeHistoricalResponse, error) {
	log.Trace("GetFixedIncomeHistorical called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.GetFixedIncomeHistoricalResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := f.cli.Post(ctx, model.URLInsightBase+model.URLInsightFixedIncomeHistorical, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.GetFixedIncomeHistoricalResponse{}, errors.Wrap(err, "couldn't get fixed income historical")
	}

	var resp model.GetFixedIncomeHistoricalResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.GetFixedIncomeHistoricalResponse{}, err
	}

	return resp, err
}
