package market

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	sdk "github.com/OrbisSystems/orbis-sdk-go/interface"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

type WorldMarket struct {
	sdk.Auth

	url string
	cli sdk.HTTPClient
}

func New(url string, auth sdk.Auth, cli sdk.HTTPClient) *WorldMarket {
	return &WorldMarket{
		Auth: auth,
		url:  url,
		cli:  cli,
	}
}

func (wm *WorldMarket) GetContinents(ctx context.Context, limit, offset int) ([]model.Continent, error) {
	log.Trace("GetContinents called")

	r, err := wm.cli.Get(ctx, fmt.Sprintf("%s?limit=%d&offset=%d", wm.url+model.URLInsightBase+model.URLInsightContinents, limit, offset), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get continents")
	}

	var resp []model.Continent
	err = wm.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (wm *WorldMarket) GetRegions(ctx context.Context, limit, offset int) ([]model.Region, error) {
	log.Trace("GetRegions called")

	r, err := wm.cli.Get(ctx, fmt.Sprintf("%s?limit=%d&offset=%d", wm.url+model.URLInsightBase+model.URLInsightRegions, limit, offset), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get regions")
	}

	var resp []model.Region
	err = wm.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (wm *WorldMarket) GetCountryCodes(ctx context.Context, limit, offset int) ([]model.CountryCode, error) {
	log.Trace("GetCountryCodes called")

	r, err := wm.cli.Get(ctx, fmt.Sprintf("%s?limit=%d&offset=%d", wm.url+model.URLInsightBase+model.URLInsightCountryCodes, limit, offset), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get country codes")
	}

	var resp []model.CountryCode
	err = wm.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (wm *WorldMarket) GetGlobalIndexes(ctx context.Context, limit, offset int, continent, quoteType string) ([]model.GlobalIndexFull, error) {
	log.Trace("GetGlobalIndexes called")

	r, err := wm.cli.Get(ctx, fmt.Sprintf("%s?limit=%d&offset=%d&continent=%s&quoteType=%s", wm.url+model.URLInsightBase+model.URLInsightGlobalIndexes, limit, offset, continent, quoteType), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get global indexes")
	}

	var resp []model.GlobalIndexFull
	err = wm.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
