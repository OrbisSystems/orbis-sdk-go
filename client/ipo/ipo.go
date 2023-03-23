package ipo

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

type IPO struct {
	sdk.Auth

	url string
	cli sdk.HTTPClient
}

func New(url string, auth sdk.Auth, cli sdk.HTTPClient) *IPO {
	return &IPO{
		Auth: auth,
		url:  url,
		cli:  cli,
	}
}

func (i *IPO) GetUpcomingIPOs(ctx context.Context, limit, offset int) (model.IPOResponse, error) {
	log.Trace("GetUpcomingIPOs called")

	r, err := i.cli.Get(ctx, fmt.Sprintf("%s?limit=%d&offset=%d", i.url+model.URLInsightBase+model.URLInsightIPOsUpcoming, limit, offset), nil)
	if err != nil {
		return model.IPOResponse{}, errors.Wrap(err, "couldn't get upcoming ipo")
	}

	var resp model.IPOResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.IPOResponse{}, err
	}

	return resp, err
}

func (i *IPO) GetRecentIPOs(ctx context.Context, req model.RecentIPORequest) (model.IPOResponse, error) {
	log.Trace("GetRecentIPOs called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.IPOResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := i.cli.Post(ctx, i.url+model.URLInsightBase+model.URLInsightIPOsRecent, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.IPOResponse{}, errors.Wrap(err, "couldn't get recent ipo")
	}

	var resp model.IPOResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.IPOResponse{}, err
	}

	return resp, err
}
