package client

import (
	"os"

	log "github.com/sirupsen/logrus"

	sdk "github.com/OrbisSystems/orbis-sdk-go"
	"github.com/OrbisSystems/orbis-sdk-go/client/account"
	"github.com/OrbisSystems/orbis-sdk-go/client/dates"
	"github.com/OrbisSystems/orbis-sdk-go/client/funds"
	"github.com/OrbisSystems/orbis-sdk-go/client/ipo"
	"github.com/OrbisSystems/orbis-sdk-go/client/logos"
	"github.com/OrbisSystems/orbis-sdk-go/client/market"
	"github.com/OrbisSystems/orbis-sdk-go/client/news"
	"github.com/OrbisSystems/orbis-sdk-go/client/og"
	"github.com/OrbisSystems/orbis-sdk-go/client/passport"
	"github.com/OrbisSystems/orbis-sdk-go/client/quotes"
	"github.com/OrbisSystems/orbis-sdk-go/client/research"
	"github.com/OrbisSystems/orbis-sdk-go/client/tiprank"
	"github.com/OrbisSystems/orbis-sdk-go/config"
)

var logLevelMap = map[config.Level]log.Level{
	config.PanicLogLevel: log.PanicLevel,
	config.FatalLogLevel: log.FatalLevel,
	config.ErrorLogLevel: log.ErrorLevel,
	config.WarnLogLevel:  log.WarnLevel,
	config.InfoLogLevel:  log.InfoLevel,
	config.DebugLogLevel: log.DebugLevel,
	config.TraceLogLevel: log.TraceLevel,
}

// Client top-level client for this SDK. Use it for calling all available API we provide for you.
type Client struct {
	Account      sdk.AccountService
	News         sdk.NewsService
	Logos        sdk.LogosService
	Passport     sdk.PassportService
	TipRank      sdk.TipRankService
	Quote        sdk.QuoteService
	Funds        sdk.FundsService
	Research     sdk.ResearchService
	IPO          sdk.IPOService
	WorldMarket  sdk.WorldMarketService
	MarketDates  sdk.MarketDatesService
	OptionGreeks sdk.OptionGreeksService
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
	log.SetLevel(logLevelMap[cfg.LogLevel])
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)

	return &SDKBuilder{
		cli:        &Client{},
		cfg:        cfg,
		httpClient: httpClient,
		auth:       auth,
	}
}

func (b *SDKBuilder) SetAccountService(srv sdk.AccountService) *SDKBuilder {
	b.cli.Account = srv
	return b
}

func (b *SDKBuilder) SetNewsService(srv sdk.NewsService) *SDKBuilder {
	b.cli.News = srv
	return b
}

func (b *SDKBuilder) SetLogosService(srv sdk.LogosService) *SDKBuilder {
	b.cli.Logos = srv
	return b
}

func (b *SDKBuilder) SetPassportService(srv sdk.PassportService) *SDKBuilder {
	b.cli.Passport = srv
	return b
}

func (b *SDKBuilder) SetTipRankService(srv sdk.TipRankService) *SDKBuilder {
	b.cli.TipRank = srv
	return b
}

func (b *SDKBuilder) SetQuoteService(srv sdk.QuoteService) *SDKBuilder {
	b.cli.Quote = srv
	return b
}

func (b *SDKBuilder) SetFundsService(srv sdk.FundsService) *SDKBuilder {
	b.cli.Funds = srv
	return b
}

func (b *SDKBuilder) SetResearchService(srv sdk.ResearchService) *SDKBuilder {
	b.cli.Research = srv
	return b
}

func (b *SDKBuilder) SetIPOService(srv sdk.IPOService) *SDKBuilder {
	b.cli.IPO = srv
	return b
}

func (b *SDKBuilder) SetWorldMarketService(srv sdk.WorldMarketService) *SDKBuilder {
	b.cli.WorldMarket = srv
	return b
}

func (b *SDKBuilder) SetMarketDatesService(srv sdk.MarketDatesService) *SDKBuilder {
	b.cli.MarketDates = srv
	return b
}

func (b *SDKBuilder) SetOptionGreeksService(srv sdk.OptionGreeksService) *SDKBuilder {
	b.cli.OptionGreeks = srv
	return b
}

// Build builds Orbis Client.
// It uses default services except you set some service manually.
func (b *SDKBuilder) Build() *Client {
	if b.cli.Account == nil {
		b.cli.Account = account.New(b.cfg, b.auth, b.httpClient) // default
	}
	if b.cli.News == nil {
		b.cli.News = news.New(b.cfg, b.auth, b.httpClient) // default
	}

	if b.cli.Logos == nil {
		b.cli.Logos = logos.New(b.cfg, b.auth, b.httpClient) // default
	}

	if b.cli.Passport == nil {
		b.cli.Passport = passport.New(b.cfg, b.auth, b.httpClient) // default
	}

	if b.cli.TipRank == nil {
		b.cli.TipRank = tiprank.New(b.cfg, b.auth, b.httpClient) // default
	}

	if b.cli.Quote == nil {
		b.cli.Quote = quotes.New(b.cfg, b.auth, b.httpClient) // default
	}

	if b.cli.Funds == nil {
		b.cli.Funds = funds.New(b.cfg, b.auth, b.httpClient) // default
	}

	if b.cli.Research == nil {
		b.cli.Research = research.New(b.cfg, b.auth, b.httpClient) // default
	}

	if b.cli.IPO == nil {
		b.cli.IPO = ipo.New(b.cfg, b.auth, b.httpClient) // default
	}

	if b.cli.WorldMarket == nil {
		b.cli.WorldMarket = market.New(b.cfg, b.auth, b.httpClient) // default
	}

	if b.cli.MarketDates == nil {
		b.cli.MarketDates = dates.New(b.cfg, b.auth, b.httpClient) // default
	}

	if b.cli.OptionGreeks == nil {
		b.cli.OptionGreeks = og.New(b.cfg, b.auth, b.httpClient) // default
	}

	return b.cli
}
