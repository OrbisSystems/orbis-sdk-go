package auth

import "github.com/OrbisSystems/orbis-sdk-go/storage"

const (
	AccessTokenHeader = "Authorization"
	BearerSchema      = "Bearer"
)

type Auth struct {
	storage storage.Storage
	key     string
}

func NewAuth(storage storage.Storage, key string) API {
	return &Auth{storage: storage, key: key}
}
