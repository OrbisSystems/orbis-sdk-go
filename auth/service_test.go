package auth

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/OrbisSystems/orbis-sdk-go/interfaces/mock"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New(nil))
}

func TestAuth_SetToken(t *testing.T) {
	testCases := []struct {
		name   string
		expErr error
		input  model.Token
		fn     func(ctx context.Context, token model.Token) *Auth
	}{
		{
			name:   "success",
			expErr: nil,
			input: model.Token{
				AccessToken:      "AccessToken",
				RefreshToken:     "RefreshToken",
				AccessExpiresAt:  1679588406,
				RefreshExpiresAt: 1679588406,
				PairId:           "123123",
			},
			fn: func(ctx context.Context, token model.Token) *Auth {
				ctrl := gomock.NewController(t)
				storage := mock.NewMockStorage(ctrl)

				bb := `{"access_token":"AccessToken","refresh_token":"RefreshToken","access_expires_at":1679588406,"refresh_expires_at":1679588406,"pair_id":"123123"}`

				storage.EXPECT().Store(ctx, []byte(bb)).Return(nil)

				return New(storage)
			},
		},
		{
			name:   "err/storage",
			expErr: errors.New("failed to store"),
			input: model.Token{
				AccessToken:      "AccessToken",
				RefreshToken:     "RefreshToken",
				AccessExpiresAt:  1679588406,
				RefreshExpiresAt: 1679588406,
				PairId:           "123123",
			},
			fn: func(ctx context.Context, token model.Token) *Auth {
				ctrl := gomock.NewController(t)
				storage := mock.NewMockStorage(ctrl)

				bb := `{"access_token":"AccessToken","refresh_token":"RefreshToken","access_expires_at":1679588406,"refresh_expires_at":1679588406,"pair_id":"123123"}`

				storage.EXPECT().Store(ctx, []byte(bb)).Return(errors.New("failed to store"))

				return New(storage)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			a := tc.fn(ctx, tc.input)
			err := a.SetToken(ctx, tc.input)
			if tc.expErr != nil {
				assert.EqualError(t, tc.expErr, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAuth_GetToken(t *testing.T) {
	testCases := []struct {
		name     string
		expErr   error
		expValue model.Token
		fn       func(ctx context.Context) *Auth
	}{
		{
			name:   "success",
			expErr: nil,
			expValue: model.Token{
				AccessToken:      "AccessToken",
				RefreshToken:     "RefreshToken",
				AccessExpiresAt:  1679588406,
				RefreshExpiresAt: 1679588406,
				PairId:           "123123",
			},
			fn: func(ctx context.Context) *Auth {
				ctrl := gomock.NewController(t)
				storage := mock.NewMockStorage(ctrl)

				bb := `{"access_token":"AccessToken","refresh_token":"RefreshToken","access_expires_at":1679588406,"refresh_expires_at":1679588406,"pair_id":"123123"}`

				storage.EXPECT().Get(ctx).Return([]byte(bb), nil)

				return New(storage)
			},
		},
		{
			name:     "err/storage",
			expErr:   errors.New("failed to get"),
			expValue: model.Token{},
			fn: func(ctx context.Context) *Auth {
				ctrl := gomock.NewController(t)
				storage := mock.NewMockStorage(ctrl)

				storage.EXPECT().Get(ctx).Return(nil, errors.New("failed to get"))

				return New(storage)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			a := tc.fn(ctx)
			tkn, err := a.GetToken(ctx)
			if tc.expErr != nil {
				assert.EqualError(t, tc.expErr, err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expValue, tkn)
		})
	}
}
