package client

import (
	sdk "github.com/OrbisSystems/orbis-sdk-go"
	"github.com/OrbisSystems/orbis-sdk-go/client/account"
	"github.com/OrbisSystems/orbis-sdk-go/client/logos"
	"github.com/OrbisSystems/orbis-sdk-go/client/news"
	"github.com/OrbisSystems/orbis-sdk-go/client/passport"
	"github.com/OrbisSystems/orbis-sdk-go/client/tiprank"
	"github.com/OrbisSystems/orbis-sdk-go/config"
)

// Client top-level client for this SDK. Use it for calling all available API we provide for you.
type Client struct {
	AccountService  sdk.AccountService
	NewsService     sdk.NewsService
	LogosService    sdk.LogosService
	PassportService sdk.PassportService
	TipRankService  sdk.TipRankService
}

// SDKBuilder provides building Orbis Client with custom parts.
// If you wish to mock some parts of SDK (like Logos service or News service), you can set it via builder's set functions.
type SDKBuilder struct {
	cli *Client

	cfg        config.Config
	httpClient sdk.HTTPClient
	auth       sdk.Auth
}

func NewSDKBuilder(cfg config.Config, httpClient sdk.HTTPClient, auth sdk.Auth) *SDKBuilder {
	return &SDKBuilder{
		cli:        &Client{},
		cfg:        cfg,
		httpClient: httpClient,
		auth:       auth,
	}
}

func (b *SDKBuilder) SetAccountService(srv sdk.AccountService) *SDKBuilder {
	b.cli.AccountService = srv
	return b
}

func (b *SDKBuilder) SetNewsService(srv sdk.NewsService) *SDKBuilder {
	b.cli.NewsService = srv
	return b
}

func (b *SDKBuilder) SetLogosService(srv sdk.LogosService) *SDKBuilder {
	b.cli.LogosService = srv
	return b
}

func (b *SDKBuilder) SetPassportService(srv sdk.PassportService) *SDKBuilder {
	b.cli.PassportService = srv
	return b
}

func (b *SDKBuilder) SetTipRankService(srv sdk.TipRankService) *SDKBuilder {
	b.cli.TipRankService = srv
	return b
}

func (b *SDKBuilder) Build() *Client {
	if b.cli.AccountService == nil {
		b.cli.AccountService = account.New(b.cfg, b.auth, b.httpClient) // default
	}
	if b.cli.NewsService == nil {
		b.cli.NewsService = news.New(b.cfg, b.auth, b.httpClient) // default
	}

	if b.cli.LogosService == nil {
		b.cli.LogosService = logos.New(b.cfg, b.auth, b.httpClient) // default
	}

	if b.cli.PassportService == nil {
		b.cli.PassportService = passport.New(b.cfg, b.auth, b.httpClient) // default
	}

	if b.cli.TipRankService == nil {
		b.cli.TipRankService = tiprank.New(b.cfg, b.auth, b.httpClient) // default
	}

	return b.cli
}
