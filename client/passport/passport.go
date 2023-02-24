package passport

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

type Passport struct {
	sdk.Auth

	cfg config.Config
	cli sdk.HTTPClient
}

func New(cfg config.Config, auth sdk.Auth, cli sdk.HTTPClient) *Passport {
	return &Passport{
		Auth: auth,
		cfg:  cfg,
		cli:  cli,
	}
}

func (p *Passport) Articles(ctx context.Context, req model.ArticlesRequest) ([]model.Article, error) {
	log.Trace("Articles called")

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := p.cli.Post(ctx, p.cfg.AuthHost+model.URLInsightBase+model.URLInsightPassportArticles, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get articles")
	}

	var resp []model.Article
	err = p.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (p *Passport) Newsfeed(ctx context.Context, req model.NewsfeedRequest) ([]model.Newsfeed, error) {
	log.Trace("Newsfeed called")

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := p.cli.Post(ctx, p.cfg.AuthHost+model.URLInsightBase+model.URLInsightPassportNewsFeed, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get article news feed")
	}

	var resp []model.Newsfeed
	err = p.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (p *Passport) ArticleByID(ctx context.Context, req model.ArticleByIDRequest) (model.Article, error) {
	log.Trace("ArticleByID called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.Article{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := p.cli.Post(ctx, p.cfg.AuthHost+model.URLInsightBase+model.URLInsightPassportArticleByID, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.Article{}, errors.Wrap(err, "couldn't get article by id")
	}

	var resp model.Article
	err = p.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.Article{}, err
	}

	return resp, err
}

func (p *Passport) SearchArticle(ctx context.Context, req model.SearchArticleRequest) ([]model.Article, error) {
	log.Trace("SearchArticle called")

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := p.cli.Post(ctx, p.cfg.AuthHost+model.URLInsightBase+model.URLInsightPassportSearchArticle, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get articles by search")
	}

	var resp []model.Article
	err = p.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (p *Passport) AuthorProfile(ctx context.Context, req model.AuthorProfileRequest) ([]model.AuthorProfileResponse, error) {
	log.Trace("AuthorProfile called")

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := p.cli.Post(ctx, p.cfg.AuthHost+model.URLInsightBase+model.URLInsightPassportAuthorProfile, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get author profiles")
	}

	var resp []model.AuthorProfileResponse
	err = p.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (p *Passport) MostPopularTags(ctx context.Context, req model.MostPopularTagsRequest) ([]model.TagShortInfo, error) {
	log.Trace("MostPopularTags called")

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := p.cli.Post(ctx, p.cfg.AuthHost+model.URLInsightBase+model.URLInsightPassportMostPopularTags, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get most popular tags")
	}

	var resp []model.TagShortInfo
	err = p.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}
