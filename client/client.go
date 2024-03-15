package client

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	"github.com/OrbisSystems/orbis-sdk-go/client/account"
	"github.com/OrbisSystems/orbis-sdk-go/client/dates"
	"github.com/OrbisSystems/orbis-sdk-go/client/fi"
	"github.com/OrbisSystems/orbis-sdk-go/client/funds"
	"github.com/OrbisSystems/orbis-sdk-go/client/hs"
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
	"github.com/OrbisSystems/orbis-sdk-go/http"
	sdk "github.com/OrbisSystems/orbis-sdk-go/interfaces"
	"github.com/OrbisSystems/orbis-sdk-go/ws"
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
	HoopsAI      sdk.HoopsAIService
	FixedIncome  sdk.FixedIncomeService
	WS           sdk.WS
}

func NewSDK(cfg config.Config, auth sdk.Auth) *Client {
	httpClient := http.New(wrapHTTPS(cfg.Host), auth)

	return newCli(cfg, auth, httpClient)
}

func newCli(cfg config.Config, auth sdk.Auth, httpClient sdk.HTTPClient) *Client {
	logger := logrus.New()
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetOutput(os.Stderr)
	logger.SetLevel(getLogLevel(cfg.LogLevel))

	return &Client{
		Account:      account.New(auth, httpClient, logger, cfg.ManualTokenRefresh),
		News:         news.New(httpClient, logger),
		Logos:        logos.New(httpClient, logger),
		Passport:     passport.New(httpClient, logger),
		TipRank:      tiprank.New(httpClient, logger),
		Quote:        quotes.New(httpClient, logger),
		Funds:        funds.New(httpClient, logger),
		Research:     research.New(httpClient, logger),
		IPO:          ipo.New(httpClient, logger),
		WorldMarket:  market.New(httpClient, logger),
		MarketDates:  dates.New(httpClient, logger),
		OptionGreeks: og.New(httpClient, logger),
		HoopsAI:      hs.New(httpClient, logger),
		FixedIncome:  fi.New(httpClient, logger),
		WS:           ws.New(cfg, auth),
	}
}

func (c *Client) Close() error {
	return c.WS.Close()
}

func wrapHTTPS(hostname string) string {
	return fmt.Sprintf("https://%s", hostname)
}

func getLogLevel(lvl config.Level) log.Level {
	if v, ok := logLevelMap[lvl]; ok {
		return v
	}
	return log.InfoLevel
}
