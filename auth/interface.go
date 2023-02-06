package auth

import (
	"context"

	"github.com/OrbisSystems/orbis-sdk-go/model"
)

type Storage interface {
	Store(ctx context.Context, token model.Token) error
	Get(ctx context.Context) (model.Token, error)
}
