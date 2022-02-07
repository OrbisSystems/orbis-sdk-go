package avatar

import (
	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/http"
	"github.com/OrbisSystems/orbis-sdk-go/storage"
	"io"
	"os"
)

type API interface {
	Authorization
	Avatar
}

type Authorization interface {
	AvatarAuth(env, orbisToken string) (io.ReadCloser, error)
	AvatarRenewToken(token string) (io.ReadCloser, error)
	AvatarTokenInvalidate(token string) (io.ReadCloser, error)
}

type Avatar interface {
	User
}

type User interface {
	GetUserAvatar(userID, token string) (io.ReadCloser, error)
	UploadUserAvatar(userID, token string, avatar os.File) (io.ReadCloser, error)
	RemoveUserAvatar(userID, token string) (io.ReadCloser, error)
	GetUserAvatarByOrbisToken(userID, orbisToken, env string) (io.ReadCloser, error)
	UploadUserAvatarByOrbisToken(userID, orbisToken, env string, avatar os.File) (io.ReadCloser, error)
	RemoveUserAvatarByOrbisToken(userID, orbisToken, env string) (io.ReadCloser, error)
}

type Account interface {
	GetAccountAvatar(accountID, token string) (io.ReadCloser, error)
	UploadAccountAvatar(accountID, token string, avatar os.File) (io.ReadCloser, error)
	RemoveAccountAvatar(accountID, token string) (io.ReadCloser, error)
	GetAccountAvatarByOrbisToken(accountID, orbisToken, env string) (io.ReadCloser, error)
	UploadAccountAvatarByOrbisToken(accountID, orbisToken, env string, avatar os.File) (io.ReadCloser, error)
	RemoveAccountAvatarByOrbisToken(accountID, orbisToken, env string) (io.ReadCloser, error)
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
