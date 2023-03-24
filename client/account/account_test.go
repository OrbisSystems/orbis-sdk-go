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

func TestAccount_LoginByAPIKey(t *testing.T) {
	var (
		token = model.LoginByAPIKeyResponse{
			Status: 201,
			ApiKeysLogin: struct {
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

		req = model.LoginByAPIKeyRequest{
			APIKey:    "11111111",
			APISecret: "22222222",
		}
		rawToken = `{"status":201,"api_keys_login":{"tokens":{"access_token":"assda","refresh_token":"fewfsdf","access_expires_at":1679588406,"refresh_expires_at":1679588406,"pair_id":"111"}}}`
		rawReq   = `{"api_key":"11111111","api_secret":"22222222"}`
		url      = "http://localhost"
		testErr  = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.LoginByAPIKeyRequest
		hasErr bool
		fn     func(ctx context.Context, req model.LoginByAPIKeyRequest) *Account
	}{
		{
			name:   "success",
			input:  req,
			hasErr: false,
			fn: func(ctx context.Context, req model.LoginByAPIKeyRequest) *Account {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)
				auth := mock.NewMockAuth(ctrl)

				r := io.NopCloser(strings.NewReader(rawToken))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, url+model.URLB2BLoginByAPIKey, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)
				auth.EXPECT().SetToken(ctx, token.ApiKeysLogin.Tokens).Return(nil)

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
			fn: func(ctx context.Context, req model.LoginByAPIKeyRequest) *Account {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)
				auth := mock.NewMockAuth(ctrl)

				r := io.NopCloser(strings.NewReader(rawToken))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, url+model.URLB2BLoginByAPIKey, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)
				auth.EXPECT().SetToken(ctx, token.ApiKeysLogin.Tokens).Return(testErr)

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
			fn: func(ctx context.Context, req model.LoginByAPIKeyRequest) *Account {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawToken))

				httpResponse := &http.Response{
					StatusCode: http.StatusBadGateway,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, url+model.URLB2BLoginByAPIKey, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

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
			fn: func(ctx context.Context, req model.LoginByAPIKeyRequest) *Account {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, url+model.URLB2BLoginByAPIKey, bytes.NewBuffer(bb), nil).Return(nil, testErr)

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
			err := a.LoginByAPIKey(ctx, tc.input)

			if tc.hasErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAccount_CreateAPIKey(t *testing.T) {
	var (
		apiKey model.CreateAPIKeyResponse

		req = model.CreateAPIKeyRequest{
			AcSameAsForUser: true,
			Name:            "key",
		}
		rawApiKeys = `{"api_keys":{"ac_same_as_for_user":false,"api_key":"123123","api_secret":"qqqwww","branches":null,"created_at":"","id":10,"name":"","permissions":null,"roles":null,"updated_at":""},"status":201}`
		rawReq     = `{"ac_same_as_for_user":true,"branches":null,"firms":null,"has_access_to_all_firm_branches":false,"name":"key","permissions":null,"roles":null,"updated_at":""}`
		url        = "http://localhost"
		testErr    = errors.New("process error")
	)

	apiKey.Status = 201
	apiKey.ApiKeys.ApiKey = "123123"
	apiKey.ApiKeys.ApiSecret = "qqqwww"
	apiKey.ApiKeys.Id = 10

	testCases := []struct {
		name   string
		input  model.CreateAPIKeyRequest
		hasErr bool
		fn     func(ctx context.Context, req model.CreateAPIKeyRequest) *Account
	}{
		{
			name:   "success",
			input:  req,
			hasErr: false,
			fn: func(ctx context.Context, req model.CreateAPIKeyRequest) *Account {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawApiKeys))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, url+model.URLB2BCreateAPIKey, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

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
			name:   "err/unmarshal",
			input:  req,
			hasErr: true,
			fn: func(ctx context.Context, req model.CreateAPIKeyRequest) *Account {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawApiKeys))

				httpResponse := &http.Response{
					StatusCode: http.StatusBadGateway,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, url+model.URLB2BCreateAPIKey, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

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
			fn: func(ctx context.Context, req model.CreateAPIKeyRequest) *Account {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, url+model.URLB2BCreateAPIKey, bytes.NewBuffer(bb), nil).Return(nil, testErr)

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
			keys, err := a.CreateAPIKey(ctx, tc.input)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, keys)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, apiKey, keys)
			}
		})
	}
}

func TestAccount_RefreshToken(t *testing.T) {
	var (
		tkn = model.Token{
			AccessToken:      "assda",
			RefreshToken:     "fewfsdf",
			AccessExpiresAt:  1679588406,
			RefreshExpiresAt: 1679588406,
			PairId:           "111",
		}
		token = model.LoginByEmailResponse{
			Status: 201,
			LoginBasic: struct {
				Tokens model.Token `json:"tokens"`
			}{
				Tokens: tkn},
		}
		url = "http://localhost"

		rawToken = `{"status":201,"login_basic":{"tokens":{"access_token":"assda","refresh_token":"fewfsdf","access_expires_at":1679588406,"refresh_expires_at":1679588406,"pair_id":"111"}}}`
		rawReq   = `{"refresh_token":"fewfsdf"}`
		testErr  = errors.New("process error")
	)

	testCases := []struct {
		name   string
		hasErr bool
		fn     func(ctx context.Context) *Account
	}{
		{
			name:   "success",
			hasErr: false,
			fn: func(ctx context.Context) *Account {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)
				auth := mock.NewMockAuth(ctrl)

				auth.EXPECT().SetTokenRefreshingState(true)
				auth.EXPECT().SetTokenRefreshingState(false)
				auth.EXPECT().GetToken(ctx).Return(tkn, nil).MaxTimes(2)
				auth.EXPECT().SetToken(ctx, token.LoginBasic.Tokens).Return(nil)

				bb := []byte(rawReq)

				r := io.NopCloser(strings.NewReader(rawToken))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, url+model.URLB2BRefreshToken, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

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
			hasErr: true,
			fn: func(ctx context.Context) *Account {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)
				auth := mock.NewMockAuth(ctrl)

				auth.EXPECT().SetTokenRefreshingState(true)
				auth.EXPECT().SetTokenRefreshingState(false)
				auth.EXPECT().GetToken(ctx).Return(tkn, nil).MaxTimes(2)
				auth.EXPECT().SetToken(ctx, token.LoginBasic.Tokens).Return(testErr)

				bb := []byte(rawReq)

				r := io.NopCloser(strings.NewReader(rawToken))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, url+model.URLB2BRefreshToken, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

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
			name:   "err/unmarshalling",
			hasErr: true,
			fn: func(ctx context.Context) *Account {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)
				auth := mock.NewMockAuth(ctrl)

				auth.EXPECT().SetTokenRefreshingState(true)
				auth.EXPECT().SetTokenRefreshingState(false)
				auth.EXPECT().GetToken(ctx).Return(tkn, nil)

				bb := []byte(rawReq)

				r := io.NopCloser(strings.NewReader("empty"))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, url+model.URLB2BRefreshToken, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

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
			name:   "err/client",
			hasErr: true,
			fn: func(ctx context.Context) *Account {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)
				auth := mock.NewMockAuth(ctrl)

				auth.EXPECT().SetTokenRefreshingState(true)
				auth.EXPECT().SetTokenRefreshingState(false)
				auth.EXPECT().GetToken(ctx).Return(tkn, nil)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, url+model.URLB2BRefreshToken, bytes.NewBuffer(bb), nil).Return(nil, testErr)

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
			name:   "err/get_token",
			hasErr: true,
			fn: func(ctx context.Context) *Account {
				ctrl := gomock.NewController(t)
				auth := mock.NewMockAuth(ctrl)

				auth.EXPECT().SetTokenRefreshingState(true)
				auth.EXPECT().SetTokenRefreshingState(false)
				auth.EXPECT().GetToken(ctx).Return(model.Token{}, testErr)

				return &Account{
					Auth:                   auth,
					url:                    url,
					resetTokenRefreshCh:    make(chan int64),
					watchTokenRefreshState: true,
					refreshTicker:          time.NewTicker(time.Hour * 100),
				}
			},
		},
		{
			name:   "err/no_refresh_token",
			hasErr: true,
			fn: func(ctx context.Context) *Account {
				ctrl := gomock.NewController(t)
				auth := mock.NewMockAuth(ctrl)

				auth.EXPECT().SetTokenRefreshingState(true)
				auth.EXPECT().SetTokenRefreshingState(false)
				auth.EXPECT().GetToken(ctx).Return(model.Token{}, nil)

				return &Account{
					Auth:                   auth,
					url:                    url,
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
			a := tc.fn(ctx)
			err := a.RefreshToken(ctx)

			if tc.hasErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
