package account

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/OrbisSystems/orbis-sdk-go/interface/mock"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	cli := mock.NewMockHTTPClient(ctrl)
	auth := mock.NewMockAuth(ctrl)

	auth.EXPECT().GetToken(gomock.Any()).Return(model.Token{}, nil).AnyTimes()

	assert.NotNil(t, New("", auth, cli))
}

// nolint:gosec
func TestAccount_LoginByEmail(t *testing.T) {
	var (
		token = model.LoginByEmailResponse{
			Status: 201,
			LoginBasic: struct {
				Tokens model.Token `json:"tokens"`
			}{
				Tokens: model.Token{
					AccessToken:      "assda",
					RefreshToken:     "fewfsdf",
					AccessExpiresAt:  1679588406,
					RefreshExpiresAt: 1679588406,
					PairId:           "111",
				}},
		}

		req = model.LoginByEmailRequest{
			Email:    "local@local.com",
			Password: "pass",
			DeviceID: "123",
		}
		rawToken = `{"status":201,"login_basic":{"tokens":{"access_token":"assda","refresh_token":"fewfsdf","access_expires_at":1679588406,"refresh_expires_at":1679588406,"pair_id":"111"}}}`
		rawReq   = `{"email":"local@local.com","password":"pass","device_id":"123","remember_me":false}`
		url      = "http://localhost"
		testErr  = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.LoginByEmailRequest
		hasErr bool
		fn     func(ctx context.Context, req model.LoginByEmailRequest) *Account
	}{
		{
			name:   "success",
			input:  req,
			hasErr: false,
			fn: func(ctx context.Context, req model.LoginByEmailRequest) *Account {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)
				auth := mock.NewMockAuth(ctrl)

				r := io.NopCloser(strings.NewReader(rawToken))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, url+model.URLB2BLoginByEmail, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)
				auth.EXPECT().SetToken(ctx, token.LoginBasic.Tokens).Return(nil)

				return &Account{
					Auth:                   auth,
					url:                    url,
					cli:                    cli,
					resetTokenRefreshCh:    make(chan int64),
					watchTokenRefreshState: true,
					refreshTicker:          time.NewTicker(time.Hour * 100),
				}
			},
		},
		{
			name:   "err/set_token",
			input:  req,
			hasErr: true,
			fn: func(ctx context.Context, req model.LoginByEmailRequest) *Account {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)
				auth := mock.NewMockAuth(ctrl)

				r := io.NopCloser(strings.NewReader(rawToken))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, url+model.URLB2BLoginByEmail, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)
				auth.EXPECT().SetToken(ctx, token.LoginBasic.Tokens).Return(testErr)

				return &Account{
					Auth:                   auth,
					url:                    url,
					cli:                    cli,
					resetTokenRefreshCh:    make(chan int64),
					watchTokenRefreshState: true,
					refreshTicker:          time.NewTicker(time.Hour * 100),
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  req,
			hasErr: true,
			fn: func(ctx context.Context, req model.LoginByEmailRequest) *Account {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawToken))

				httpResponse := &http.Response{
					StatusCode: http.StatusBadGateway,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, url+model.URLB2BLoginByEmail, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Account{
					url:                    url,
					cli:                    cli,
					resetTokenRefreshCh:    make(chan int64),
					watchTokenRefreshState: true,
					refreshTicker:          time.NewTicker(time.Hour * 100),
				}
			},
		},
		{
			name:   "err/api_call",
			input:  req,
			hasErr: true,
			fn: func(ctx context.Context, req model.LoginByEmailRequest) *Account {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, url+model.URLB2BLoginByEmail, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Account{
					url:                    url,
					cli:                    cli,
					resetTokenRefreshCh:    make(chan int64),
					watchTokenRefreshState: true,
					refreshTicker:          time.NewTicker(time.Hour * 100),
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			a := tc.fn(ctx, tc.input)
			err := a.LoginByEmail(ctx, tc.input)

			if tc.hasErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
