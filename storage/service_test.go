package storage

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	inMemoryInstance = NewInMemory()
	defaultValue     = []byte{1, 2, 4}
)

func TestInMemory_Save(t *testing.T) {
	key := strconv.Itoa(time.Now().Nanosecond())
	err := inMemoryInstance.Save(key, defaultValue)

	assert.NoError(t, err, "couldn't save data into storage")

	v, ok, err := inMemoryInstance.Get(key)
	assert.NoError(t, err, "err must be nil")
	assert.True(t, ok, "value must be returned")
	assert.Equal(t, defaultValue, v, "returned wrong value")
}

func TestInMemory_SaveEmptyValue(t *testing.T) {
	err := inMemoryInstance.Save(strconv.Itoa(time.Now().Nanosecond()), []byte{})
	assert.Error(t, err, "expected err")
}

func TestInMemory_SaveEmptyKey(t *testing.T) {
	err := inMemoryInstance.Save("", defaultValue)
	assert.Error(t, err, "expected err")
}

func TestInMemory_Get(t *testing.T) {
	key := strconv.Itoa(time.Now().Nanosecond())
	err := inMemoryInstance.Save(key, defaultValue)

	assert.NoError(t, err, "couldn't save data into storage")

	v, ok, err := inMemoryInstance.Get(key)
	assert.NoError(t, err, "err must be nil")
	assert.True(t, ok, "value must be returned")
	assert.Equal(t, defaultValue, v, "returned wrong value")
}

func TestInMemory_GetEmptyKey(t *testing.T) {
	v, ok, err := inMemoryInstance.Get("")
	assert.Error(t, err, "expected err")
	assert.False(t, ok, "value must be returned")
	assert.Nil(t, v, "returned wrong value")
}
