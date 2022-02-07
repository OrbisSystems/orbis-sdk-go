package logos

import (
	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/http"
	"github.com/OrbisSystems/orbis-sdk-go/storage"
	"io"
)

type API interface {
	GetAllLogos(page int) (io.ReadCloser, error)
	GetSymbolLogos(symbol string) (io.ReadCloser, error)
	GetSymbolsLogos(symbols []string) (io.ReadCloser, error)
	GetSymbolsSocials(symbol string) (io.ReadCloser, error)
	GetSymbolLogo(symbol string) (io.ReadCloser, error)
	GetSymbolLogoConverted(symbol, conversion string) (io.ReadCloser, error)
}

type Client struct {
	auth.API
	client http.Client

	config config.OrbisConfig
}

// NewClient returns new API with defaults storage.Storage and http.OrbisClient
func NewClient(config config.OrbisConfig) API {
	return NewConfiguredClient(storage.NewInMemory(), http.DefaultClient, config, http.DefaultAuthKey)
}

// NewConfiguredClient returns new API
func NewConfiguredClient(storage storage.Storage,
	client http.Client, config config.OrbisConfig, key string) API {
	return &Client{
		API:    auth.NewAuth(storage, key),
		client: client,
		config: config,
	}
}
