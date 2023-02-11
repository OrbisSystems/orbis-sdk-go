package news

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	sdk "github.com/OrbisSystems/orbis-sdk-go"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

// News service provides API for getting news.
type News struct {
	sdk.Auth

	cfg config.Config
	cli sdk.HTTPClient
}

func New(cfg config.Config, auth sdk.Auth, cli sdk.HTTPClient) *News {
	return &News{
		Auth: auth,
		cfg:  cfg,
		cli:  cli,
	}
}

// GetByFilter returns news by filters.
func (c *News) GetByFilter(ctx context.Context, req model.NewsFilterRequest) ([]model.NewsResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := c.cli.Post(ctx, c.cfg.AuthHost+model.URLInsightBase+model.URLInsightNewsFilter, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get news by filter")
	}

	var resp []model.NewsResponse
	err = c.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetByID returns news by ID.
func (c *News) GetByID(ctx context.Context, req model.NewsByIDRequest) (model.NewsResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return model.NewsResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := c.cli.Post(ctx, c.cfg.AuthHost+model.URLInsightBase+model.URLInsightNewsByID, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.NewsResponse{}, errors.Wrap(err, "couldn't get news by id")
	}

	var resp model.NewsResponse
	err = c.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.NewsResponse{}, err
	}

	return resp, err
}

// GetSymbolSubjects returns available subjects for symbol.
func (c *News) GetSymbolSubjects(ctx context.Context, req model.SymbolSubjectsRequest) ([]model.SymbolSubjectsResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := c.cli.Post(ctx, c.cfg.AuthHost+model.URLInsightBase+model.URLInsightNewsSymbolSubjects, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get news symbol subject")
	}

	var resp []model.SymbolSubjectsResponse
	err = c.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetAvailableTaxonomy returns all available taxonomy codes for news. You can use it for GetByFilter filter.
func (c *News) GetAvailableTaxonomy(ctx context.Context) ([]model.TaxonomyCode, error) {
	r, err := c.cli.Get(ctx, c.cfg.AuthHost+model.URLInsightBase+model.URLInsightNewsTaxonomy, nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get news available taxonomy codes")
	}

	var resp []model.TaxonomyCode
	err = c.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetAvailableSubjects returns all available news subjects.
func (c *News) GetAvailableSubjects(ctx context.Context) ([]string, error) {
	r, err := c.cli.Get(ctx, c.cfg.AuthHost+model.URLInsightBase+model.URLInsightNewsRelevance, nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get news available subjects")
	}

	var resp []string
	err = c.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetAvailableSymbols returns all available symbols of news.
func (c *News) GetAvailableSymbols(ctx context.Context) ([]string, error) {
	r, err := c.cli.Get(ctx, c.cfg.AuthHost+model.URLInsightBase+model.URLInsightNewsSymbols, nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get news available symbols")
	}

	var resp []string
	err = c.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetAvailableSources returns all available news sources.
func (c *News) GetAvailableSources(ctx context.Context) ([]string, error) {
	r, err := c.cli.Get(ctx, c.cfg.AuthHost+model.URLInsightBase+model.URLInsightNewsSources, nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get news available sources")
	}

	var resp []string
	err = c.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetAvailableLanguages returns all available news languages.
func (c *News) GetAvailableLanguages(ctx context.Context) ([]string, error) {
	r, err := c.cli.Get(ctx, c.cfg.AuthHost+model.URLInsightBase+model.URLInsightNewsLang, nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get news available languages")
	}

	var resp []string
	err = c.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
