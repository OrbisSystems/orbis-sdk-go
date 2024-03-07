package storage

import (
	"context"
	"sync"
)

// InMemoryStorage is just a simple in-memory storage.
// It implements interface Storage.
type InMemoryStorage struct {
	mx   sync.RWMutex
	data []byte
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{}
}

func (s *InMemoryStorage) Store(_ context.Context, data []byte) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.data = data

	return nil
}

func (s *InMemoryStorage) Get(_ context.Context) ([]byte, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	return s.data, nil
}
