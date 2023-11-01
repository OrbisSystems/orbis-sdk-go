package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	sdk "github.com/OrbisSystems/orbis-sdk-go/interfaces"
	"github.com/OrbisSystems/orbis-sdk-go/interfaces/mock"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New("", nil))
}

func TestNewWithHttp(t *testing.T) {
	assert.NotNil(t, NewWithHttp("", nil, nil))
}

func TestOrbisClient_Get(t *testing.T) {
	baseURL := "https://localhost.com"
	url := "/path"
	tkn := model.Token{
		AccessToken:  "asdasdasd",
		RefreshToken: "cxzczxczx",
	}
	r := io.NopCloser(strings.NewReader(`{"data": "data"}`))
	httpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}
	testErr := errors.New("error")

	testCases := []struct {
		name   string
		hasErr bool
		fn     func(ctx context.Context) sdk.HTTPClient
	}{
		{
			name:   "success/with_token",
			hasErr: false,
			fn: func(ctx context.Context) sdk.HTTPClient {
				ctrl := gomock.NewController(t)
				httpCli := mock.NewMockHTTPExecutor(ctrl)
				auth := mock.NewMockAuth(ctrl)

				req, _ := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", baseURL, url), http.NoBody)
				header := http.Header{}
				header.Add(accessTokenHeader, tkn.AccessToken)
				header.Add(contentTypeHeaderKey, applicationJson)
				req.Header = header

				auth.EXPECT().GetToken(ctx).Return(tkn, nil)
				httpCli.EXPECT().Do(req).Return(httpResponse, nil)

				return NewWithHttp(baseURL, auth, httpCli)
			},
		},
		{
			name:   "success/no_token",
			hasErr: false,
			fn: func(ctx context.Context) sdk.HTTPClient {
				ctrl := gomock.NewController(t)
				httpCli := mock.NewMockHTTPExecutor(ctrl)
				auth := mock.NewMockAuth(ctrl)

				req, _ := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", baseURL, url), http.NoBody)
				header := http.Header{}
				header.Add(contentTypeHeaderKey, applicationJson)
				req.Header = header

				auth.EXPECT().GetToken(ctx).Return(model.Token{}, testErr)
				httpCli.EXPECT().Do(req).Return(httpResponse, nil)

				return NewWithHttp(baseURL, auth, httpCli)
			},
		},
		{
			name:   "err/cli",
			hasErr: true,
			fn: func(ctx context.Context) sdk.HTTPClient {
				ctrl := gomock.NewController(t)
				httpCli := mock.NewMockHTTPExecutor(ctrl)
				auth := mock.NewMockAuth(ctrl)

				req, _ := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", baseURL, url), http.NoBody)
				header := http.Header{}
				header.Add(accessTokenHeader, tkn.AccessToken)
				header.Add(contentTypeHeaderKey, applicationJson)
				req.Header = header

				auth.EXPECT().GetToken(ctx).Return(tkn, nil)
				httpCli.EXPECT().Do(req).Return(nil, testErr)

				return NewWithHttp(baseURL, auth, httpCli)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			cli := tc.fn(ctx)

			resp, err := cli.Get(ctx, url, nil)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.EqualValues(t, httpResponse.Body, resp.Body)
			}
		})
	}
}

func TestOrbisClient_Post(t *testing.T) {
	baseURL := "https://localhost.com"
	url := "/path"
	tkn := model.Token{
		AccessToken:  "asdasdasd",
		RefreshToken: "cxzczxczx",
	}
	r := io.NopCloser(strings.NewReader(`{"data": "data"}`))
	httpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}
	testErr := errors.New("error")

	testCases := []struct {
		name   string
		hasErr bool
		fn     func(ctx context.Context) sdk.HTTPClient
	}{
		{
			name:   "success/with_token",
			hasErr: false,
			fn: func(ctx context.Context) sdk.HTTPClient {
				ctrl := gomock.NewController(t)
				httpCli := mock.NewMockHTTPExecutor(ctrl)
				auth := mock.NewMockAuth(ctrl)

				req, _ := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s%s", baseURL, url), nil) // nolint
				header := http.Header{}
				header.Add(accessTokenHeader, tkn.AccessToken)
				header.Add(contentTypeHeaderKey, applicationJson)
				req.Header = header

				auth.EXPECT().GetToken(ctx).Return(tkn, nil)
				httpCli.EXPECT().Do(req).Return(httpResponse, nil)

				return NewWithHttp(baseURL, auth, httpCli)
			},
		},
		{
			name:   "success/no_token",
			hasErr: false,
			fn: func(ctx context.Context) sdk.HTTPClient {
				ctrl := gomock.NewController(t)
				httpCli := mock.NewMockHTTPExecutor(ctrl)
				auth := mock.NewMockAuth(ctrl)

				req, _ := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s%s", baseURL, url), nil) // nolint
				header := http.Header{}
				header.Add(contentTypeHeaderKey, applicationJson)
				req.Header = header

				auth.EXPECT().GetToken(ctx).Return(model.Token{}, testErr)
				httpCli.EXPECT().Do(req).Return(httpResponse, nil)

				return NewWithHttp(baseURL, auth, httpCli)
			},
		},
		{
			name:   "err/cli",
			hasErr: true,
			fn: func(ctx context.Context) sdk.HTTPClient {
				ctrl := gomock.NewController(t)
				httpCli := mock.NewMockHTTPExecutor(ctrl)
				auth := mock.NewMockAuth(ctrl)

				req, _ := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s%s", baseURL, url), nil) // nolint
				header := http.Header{}
				header.Add(accessTokenHeader, tkn.AccessToken)
				header.Add(contentTypeHeaderKey, applicationJson)
				req.Header = header

				auth.EXPECT().GetToken(ctx).Return(tkn, nil)
				httpCli.EXPECT().Do(req).Return(nil, testErr)

				return NewWithHttp(baseURL, auth, httpCli)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			cli := tc.fn(ctx)

			resp, err := cli.Post(ctx, url, nil, nil)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.EqualValues(t, httpResponse.Body, resp.Body)
			}
		})
	}
}
