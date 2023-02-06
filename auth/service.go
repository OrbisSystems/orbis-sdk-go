package auth

import (
	"context"

	"github.com/OrbisSystems/orbis-sdk-go/model"
)

type Auth struct {
	storage Storage
}

func New(storage Storage) *Auth {
	return &Auth{storage: storage}
}

func (a *Auth) SetToken(ctx context.Context, token model.Token) error {
	return a.storage.Store(ctx, token)
}

func (a *Auth) GetToken(ctx context.Context) (model.Token, error) {
	return a.storage.Get(ctx)
}
