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

func New(cfg config.Config, cli sdk.HTTPClient, auth sdk.Auth) *Client {
	return &Client{
		AccountService:  account.New(cfg, auth, cli),
		NewsService:     news.New(cfg, auth, cli),
		LogosService:    logos.New(cfg, auth, cli),
		PassportService: passport.New(cfg, auth, cli),
		TipRankService:  tiprank.New(cfg, auth, cli),
	}
}
