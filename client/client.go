package client

import (
	sdk "github.com/OrbisSystems/orbis-sdk-go"
	"github.com/OrbisSystems/orbis-sdk-go/client/account"
	"github.com/OrbisSystems/orbis-sdk-go/client/logos"
	"github.com/OrbisSystems/orbis-sdk-go/client/news"
	"github.com/OrbisSystems/orbis-sdk-go/config"
)

// Client top-level client for this SDK. Use it for calling all available API we provide for you.
type Client struct {
	AccountService sdk.AccountService
	NewsService    sdk.NewsService
	LogosService   sdk.LogosService
}

func New(cfg config.Config, cli sdk.HTTPClient, auth sdk.Auth, validator sdk.Validator) *Client {
	return &Client{
		AccountService: account.New(cfg, auth, cli, validator),
		NewsService:    news.New(cfg, auth, cli, validator),
		LogosService:   logos.New(cfg, auth, cli, validator),
	}
}
