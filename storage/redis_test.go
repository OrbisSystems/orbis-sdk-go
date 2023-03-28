package storage

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"

	"github.com/OrbisSystems/orbis-sdk-go/interfaces/mock"
)

func TestNewRedisStorageWithOwnCli(t *testing.T) {
	cli, err := NewRedisStorageWithOwnCli(Config{}, nil)
	assert.NoError(t, err)
	assert.NotNil(t, cli)
}

func TestClient_Store(t *testing.T) {
	var (
		cfg = Config{KeyForStore: "qwerty"}

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		hasErr bool
		input  []byte
		fn     func(ctx context.Context, data []byte) *Client
	}{
		{
			name:   "success",
			hasErr: false,
			input:  []byte{1, 2, 3, 4, 5},
			fn: func(ctx context.Context, data []byte) *Client {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockRedisRepo(ctrl)

				cli.EXPECT().Set(ctx, cfg.KeyForStore, data, time.Duration(0)).Return(&redis.StatusCmd{})

				return &Client{
					cli:         cli,
					keyForStore: cfg.KeyForStore,
				}
			},
		},
		{
			name:   "err",
			hasErr: true,
			input:  []byte{1, 2, 3, 4, 5},
			fn: func(ctx context.Context, data []byte) *Client {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockRedisRepo(ctrl)

				cmdSt := &redis.StatusCmd{}
				cmdSt.SetErr(testErr)

				cli.EXPECT().Set(ctx, cfg.KeyForStore, data, time.Duration(0)).Return(cmdSt)

				return &Client{
					cli:         cli,
					keyForStore: cfg.KeyForStore,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			f := tc.fn(ctx, tc.input)
			err := f.Store(ctx, tc.input)

			if tc.hasErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestClient_Get(t *testing.T) {
	var (
		cfg = Config{KeyForStore: "qwerty"}

		res = []byte{1, 2, 3, 4, 5, 6}

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		hasErr bool
		fn     func(ctx context.Context) *Client
	}{
		{
			name:   "success",
			hasErr: false,
			fn: func(ctx context.Context) *Client {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockRedisRepo(ctrl)

				cmdSt := &redis.StringCmd{}
				cmdSt.SetVal(string(res))

				cli.EXPECT().Get(ctx, cfg.KeyForStore).Return(cmdSt)

				return &Client{
					cli:         cli,
					keyForStore: cfg.KeyForStore,
				}
			},
		},
		{
			name:   "err",
			hasErr: true,
			fn: func(ctx context.Context) *Client {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockRedisRepo(ctrl)

				cmdSt := &redis.StringCmd{}
				cmdSt.SetErr(testErr)

				cli.EXPECT().Get(ctx, cfg.KeyForStore).Return(cmdSt)

				return &Client{
					cli:         cli,
					keyForStore: cfg.KeyForStore,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			f := tc.fn(ctx)
			r, err := f.Get(ctx)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Nil(t, r)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, res, r)
			}
		})
	}
}
