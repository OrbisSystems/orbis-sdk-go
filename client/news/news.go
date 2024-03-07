package news

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	sdk "github.com/OrbisSystems/orbis-sdk-go/interfaces"
	"github.com/OrbisSystems/orbis-sdk-go/model"
	"github.com/OrbisSystems/orbis-sdk-go/utils"
)

// News service provides API for getting news.
type News struct {
	cli    sdk.HTTPClient
	logger *log.Logger
}

func New(cli sdk.HTTPClient, logger *log.Logger) *News {
	return &News{
		cli:    cli,
		logger: logger,
	}
}

// GetByFilter returns news by filters.
func (c *News) GetByFilter(ctx context.Context, req model.NewsFilterRequest) (model.NewsResponse, error) {
	c.logger.Trace("GetByFilter called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.NewsResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := c.cli.Post(ctx, model.URLInsightBase+model.URLInsightNewsFilter, bytes.NewBuffer(body), nil)
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
	c.logger.Trace("GetByID called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.News{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := c.cli.Post(ctx, model.URLInsightBase+model.URLInsightNewsByID, bytes.NewBuffer(body), nil)
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

// GetAvailableSymbols returns all available symbols of news.
func (c *News) GetAvailableSymbols(ctx context.Context) ([]string, error) {
	c.logger.Trace("GetAvailableSymbols called")

	r, err := c.cli.Get(ctx, model.URLInsightBase+model.URLInsightNewsSymbols, nil)
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

// GetAvailableAuthors returns all available authors for news. You can use it for GetByFilter filter.
func (c *News) GetAvailableAuthors(ctx context.Context, symbol *string) ([]string, error) {
	c.logger.Trace("GetAvailableAuthors called")

	body, err := prepareBodyWithSymbol(symbol)
	if err != nil {
		return nil, err
	}

	r, err := c.cli.Post(ctx, model.URLInsightBase+model.URLInsightNewsAuthors, body, nil)
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

// GetAvailableChannels returns all available news channels.
func (c *News) GetAvailableChannels(ctx context.Context, symbol *string) ([]string, error) {
	c.logger.Trace("GetAvailableChannels called")

	body, err := prepareBodyWithSymbol(symbol)
	if err != nil {
		return nil, err
	}

	r, err := c.cli.Post(ctx, model.URLInsightBase+model.URLInsightNewsChannels, body, nil)
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

// GetAvailableTags returns all available news tags.
func (c *News) GetAvailableTags(ctx context.Context, symbol *string) ([]string, error) {
	c.logger.Trace("GetAvailableTags called")

	body, err := prepareBodyWithSymbol(symbol)
	if err != nil {
		return nil, err
	}

	r, err := c.cli.Post(ctx, model.URLInsightBase+model.URLInsightNewsTags, body, nil)
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

func prepareBodyWithSymbol(symbol *string) (io.Reader, error) {
	var (
		body io.Reader
	)
	if symbol != nil {
		req := model.HelpRequestWithSymbol{Symbol: symbol}
		b, err := json.Marshal(req)
		if err != nil {
			return nil, errors.Wrap(err, "couldn't marshal input parameters")
		}
		body = bytes.NewBuffer(b)
	}

	return body, nil
}
