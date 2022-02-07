package client

import (
	"github.com/OrbisSystems/orbis-sdk-go/account"
	"github.com/OrbisSystems/orbis-sdk-go/avatar"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/http"
	"github.com/OrbisSystems/orbis-sdk-go/logos"
	"github.com/OrbisSystems/orbis-sdk-go/orbis"
	"github.com/OrbisSystems/orbis-sdk-go/passport"
	"github.com/OrbisSystems/orbis-sdk-go/storage"
)

type API struct {
	storage storage.Storage
	client  http.Client

	Account  account.API
	Logos    logos.API
	Avatar   avatar.Avatar
	Orbis    orbis.API
	Passport passport.API
}

// New returns Default client for orbis API's
func New() *API {
	return NewConfigured(storage.NewInMemory(), http.DefaultClient, config.OrbisConfig{}, http.DefaultAuthKey)
}

// New returns configured client for orbis API's
func NewConfigured(storage storage.Storage, client http.Client, config config.OrbisConfig, authKey string) *API {
	api := API{}
	api.client = client
	api.storage = storage

	api.Account = account.NewConfiguredClient(api.storage, api.client, config, authKey)
	api.Logos = logos.NewConfiguredClient(api.storage, api.client, config, authKey)
	api.Avatar = avatar.NewConfiguredClient(api.storage, api.client, config, authKey)
	api.Orbis = orbis.NewConfiguredClient(api.storage, api.client, config, authKey)
	api.Passport = passport.NewConfiguredClient(api.storage, api.client, config, authKey)

	return &api
}
