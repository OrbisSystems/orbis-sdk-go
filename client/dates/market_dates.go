package dates

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

type MarketDates struct {
	url string
	cli sdk.HTTPClient
}

func New(url string, cli sdk.HTTPClient) *MarketDates {
	return &MarketDates{
		url: url,
		cli: cli,
	}
}

func (m *MarketDates) GetMarketDatesHistory(ctx context.Context, req model.GetMarketDatesRequest) (model.GetMarketDatesResponse, error) {
	log.Trace("GetMarketDatesHistory called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.GetMarketDatesResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := m.cli.Post(ctx, m.url+model.URLInsightBase+model.URLInsightMarketDatesHistory, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.GetMarketDatesResponse{}, errors.Wrap(err, "couldn't get market dates history")
	}

	var resp model.GetMarketDatesResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.GetMarketDatesResponse{}, err
	}

	return resp, err
}

func (m *MarketDates) GetTodayMarketHours(ctx context.Context, market string) (model.GetMarketHoursResponse, error) {
	log.Trace("GetTodayMarketHours called")

	r, err := m.cli.Get(ctx, fmt.Sprintf("%s?market=%s", m.url+model.URLInsightBase+model.URLInsightMarketDatesToday, market), nil)
	if err != nil {
		return model.GetMarketHoursResponse{}, errors.Wrap(err, "couldn't get today's market hours")
	}

	var resp model.GetMarketHoursResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.GetMarketHoursResponse{}, err
	}

	return resp, err
}
