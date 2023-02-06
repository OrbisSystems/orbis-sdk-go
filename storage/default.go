package storage

import (
	"context"
	"sync"

	"github.com/OrbisSystems/orbis-sdk-go/model"
)

type DefaultStorage struct {
	mx    sync.RWMutex
	token model.Token
}

func NewDefaultStorage() *DefaultStorage {
	return &DefaultStorage{}
}

func (s *DefaultStorage) Store(_ context.Context, token model.Token) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.token = token

	return nil
}

func (s *DefaultStorage) Get(ctx context.Context) (model.Token, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	return s.token, nil
}
