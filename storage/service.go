package storage

import (
	"github.com/pkg/errors"
)

var (
	emptyValueErr = errors.New("value must be not empty")
	emptyKeyErr   = errors.New("value must be not empty")
)

type InMemory struct {
	values map[string][]byte
}

func NewInMemory() Storage {
	values := make(map[string][]byte)
	return &InMemory{values: values}
}

func (i *InMemory) Save(key string, value []byte) error {
	if len(value) == 0 {
		return emptyValueErr
	}
	if len(key) == 0 {
		return emptyKeyErr
	}

	i.values[key] = value

	return nil
}

func (i *InMemory) Get(key string) ([]byte, bool, error) {
	if len(key) == 0 {
		return nil, false, emptyKeyErr
	}
	v, ok := i.values[key]
	return v, ok, nil
}
