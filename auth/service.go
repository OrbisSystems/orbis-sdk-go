package auth

import (
	"context"
	"encoding/json"
	"sync"

	log "github.com/sirupsen/logrus"

	sdk "github.com/OrbisSystems/orbis-sdk-go/interfaces"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

// Auth is a service for token processing. Now it's just a wrapping for calling storage service.
type Auth struct {
	storage sdk.Storage

	mx                   sync.Mutex
	tokenRefreshingState bool
}

func New(storage sdk.Storage) *Auth {
	return &Auth{storage: storage}
}

// SetToken store token to the storage.
func (a *Auth) SetToken(ctx context.Context, token model.Token) error {
	log.Trace("setting token to storage")
	bb, err := json.Marshal(token)
	if err != nil {
		return err
	}
	return a.storage.Store(ctx, bb)
}

// GetToken returns token from the storage.
func (a *Auth) GetToken(ctx context.Context) (model.Token, error) {
	log.Tracef("getting token from storage")
	data, err := a.storage.Get(ctx)
	if err != nil {
		return model.Token{}, err
	}

	var tkn model.Token
	err = json.Unmarshal(data, &tkn)

	return tkn, err
}

func (a *Auth) SetTokenRefreshingState(state bool) {
	a.mx.Lock()
	defer a.mx.Unlock()
	a.tokenRefreshingState = state
}

func (a *Auth) GetTokenRefreshingState() bool {
	return a.tokenRefreshingState
}
