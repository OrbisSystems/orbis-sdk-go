package news

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	sdk "github.com/OrbisSystems/orbis-sdk-go/interface"
	"github.com/OrbisSystems/orbis-sdk-go/model"
	"github.com/OrbisSystems/orbis-sdk-go/utils"
)

// News service provides API for getting news.
type News struct {
	url string
	cli sdk.HTTPClient
}

func New(url string, cli sdk.HTTPClient) *News {
	return &News{
		url: url,
		cli: cli,
	}
}

// GetByFilter returns news by filters.
func (c *News) GetByFilter(ctx context.Context, req model.NewsFilterRequest) (model.NewsResponse, error) {
	log.Trace("GetByFilter called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.NewsResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := c.cli.Post(ctx, c.url+model.URLInsightBase+model.URLInsightNewsFilter, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.NewsResponse{}, errors.Wrap(err, "couldn't get news by filter")
	}

	var resp model.NewsResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.NewsResponse{}, err
	}

	return resp, err
}

// GetByID returns news by ID.
func (c *News) GetByID(ctx context.Context, req model.NewsRequest) (model.News, error) {
	log.Trace("GetByID called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.News{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := c.cli.Post(ctx, c.url+model.URLInsightBase+model.URLInsightNewsByID, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.News{}, errors.Wrap(err, "couldn't get news by id")
	}

	var resp model.News
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.News{}, err
	}

	return resp, err
}

// GetTagsForSymbol returns available subjects for symbol.
func (c *News) GetTagsForSymbol(ctx context.Context, req model.SymbolTagsRequest) ([]string, error) {
	log.Trace("GetTagsForSymbol called")

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := c.cli.Post(ctx, c.url+model.URLInsightBase+model.URLInsightNewsSymbolTags, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get news symbol tags")
	}

	var resp []string
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetChannelsForSymbol returns available subjects for symbol.
func (c *News) GetChannelsForSymbol(ctx context.Context, req model.SymbolChannelsRequest) ([]string, error) {
	log.Trace("GetChannelsForSymbol called")

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := c.cli.Post(ctx, c.url+model.URLInsightBase+model.URLInsightNewsSymbolChannels, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get news symbol channels")
	}

	var resp []string
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetAvailableAuthors returns all available taxonomy codes for news. You can use it for GetByFilter filter.
func (c *News) GetAvailableAuthors(ctx context.Context) ([]string, error) {
	log.Trace("GetAvailableAuthors called")

	r, err := c.cli.Get(ctx, c.url+model.URLInsightBase+model.URLInsightNewsAuthors, nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get news available authors")
	}

	var resp []string
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetAvailableChannels returns all available news subjects.
func (c *News) GetAvailableChannels(ctx context.Context) ([]string, error) {
	log.Trace("GetAvailableChannels called")

	r, err := c.cli.Get(ctx, c.url+model.URLInsightBase+model.URLInsightNewsChannels, nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get news available channels")
	}

	var resp []string
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetAvailableSymbols returns all available symbols of news.
func (c *News) GetAvailableSymbols(ctx context.Context) ([]string, error) {
	log.Trace("GetAvailableSymbols called")

	r, err := c.cli.Get(ctx, c.url+model.URLInsightBase+model.URLInsightNewsSymbols, nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get news available symbols")
	}

	var resp []string
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetAvailableTags returns all available news sources.
func (c *News) GetAvailableTags(ctx context.Context) ([]string, error) {
	log.Trace("GetAvailableTags called")

	r, err := c.cli.Get(ctx, c.url+model.URLInsightBase+model.URLInsightNewsTags, nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get news available tags")
	}

	var resp []string
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
