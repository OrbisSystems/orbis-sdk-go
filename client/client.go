package client

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

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
	"github.com/OrbisSystems/orbis-sdk-go/http"
	sdk "github.com/OrbisSystems/orbis-sdk-go/interface"
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
	WS           sdk.WS
}

func NewSDK(cfg config.Config, auth sdk.Auth) *Client {
	log.SetLevel(getLogLevel(cfg.LogLevel))
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)

	httpsURL := wrapHTTPS(cfg.Host)
	httpClient := http.New(auth)

	return newCli(cfg, httpsURL, auth, httpClient)
}

func NewSDKWithHTTPClient(cfg config.Config, auth sdk.Auth, httpClient sdk.HTTPClient) *Client {
	log.SetLevel(getLogLevel(cfg.LogLevel))
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)

	httpsURL := wrapHTTPS(cfg.Host)

	return newCli(cfg, httpsURL, auth, httpClient)
}

func newCli(cfg config.Config, httpsURL string, auth sdk.Auth, httpClient sdk.HTTPClient) *Client {
	return &Client{
		Account:      account.New(httpsURL, auth, httpClient),
		News:         news.New(httpsURL, httpClient),
		Logos:        logos.New(httpsURL, httpClient),
		Passport:     passport.New(httpsURL, httpClient),
		TipRank:      tiprank.New(httpsURL, httpClient),
		Quote:        quotes.New(httpsURL, httpClient),
		Funds:        funds.New(httpsURL, httpClient),
		Research:     research.New(httpsURL, httpClient),
		IPO:          ipo.New(httpsURL, httpClient),
		WorldMarket:  market.New(httpsURL, httpClient),
		MarketDates:  dates.New(httpsURL, httpClient),
		OptionGreeks: og.New(httpsURL, httpClient),
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
