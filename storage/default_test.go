package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInMemoryStorage(t *testing.T) {
	assert.NotNil(t, NewInMemoryStorage())
}

func TestInMemoryStorage_Store(t *testing.T) {
	data := []byte{1, 2, 3, 4, 5}
	s := NewInMemoryStorage()
	err := s.Store(nil, data)
	assert.NoError(t, err)
	assert.EqualValues(t, data, s.data)
}

func TestInMemoryStorage_Get(t *testing.T) {
	data := []byte{1, 2, 3, 4, 5}
	s := NewInMemoryStorage()
	err := s.Store(nil, data)
	assert.NoError(t, err)

	returnedData, err := s.Get(nil)
	assert.NoError(t, err)
	assert.EqualValues(t, data, returnedData)
}
