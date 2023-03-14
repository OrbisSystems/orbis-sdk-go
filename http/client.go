package http

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"

	sdk "github.com/OrbisSystems/orbis-sdk-go/interface"
)

const (
	accessTokenHeader    = "Authorization"
	contentTypeHeaderKey = "Content-Type"
	applicationJson      = "application/json"
	defaultTimeout       = time.Minute
)

// OrbisClient is a wrapper for http client with additional functionality.
type OrbisClient struct {
	http.Client

	auth sdk.Auth
}

// New returns new OrbisClient instance with default http client.
func New(auth sdk.Auth) sdk.HTTPClient {
	var httpCli = http.Client{
		Timeout: defaultTimeout,
	}

	cli := OrbisClient{
		Client: httpCli,
		auth:   auth,
	}

	return &cli
}

// NewWithHttp returns new OrbisClient instance with http client user defined.
func NewWithHttp(auth sdk.Auth, client http.Client) sdk.HTTPClient {
	orbisClient := OrbisClient{
		auth:   auth,
		Client: client,
	}

	return &orbisClient
}

// Get makes an HTTP GET request to provided URL
func (c *OrbisClient) Get(ctx context.Context, url string, headers http.Header) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return response, errors.Wrap(err, "GET - request creation failed")
	}

	request.Header = headers

	return c.do(ctx, request)
}

// Post makes an HTTP POST request to provided URL and requestBody
func (c *OrbisClient) Post(ctx context.Context, url string, body io.Reader, headers http.Header) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return response, errors.Wrap(err, "POST - request creation failed")
	}

	request.Header = headers

	return c.do(ctx, request)
}

// do makes an HTTP request with the native `http.do` interface
func (c *OrbisClient) do(ctx context.Context, request *http.Request) (*http.Response, error) {
	request.Header = c.getTokenHeader(ctx, request.Header)
	if request.Header.Get(contentTypeHeaderKey) == "" {
		request.Header.Add(contentTypeHeaderKey, applicationJson)
	}

	response, err := c.Client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// getTokenHeader sets access token to request header.
func (c *OrbisClient) getTokenHeader(ctx context.Context, header http.Header) http.Header {
	if header == nil {
		header = http.Header{}
	}
	for c.auth.GetTokenRefreshingState() {
	} // lock getting token from auth system if there is a process for token refresh

	token, err := c.auth.GetToken(ctx)
	if err != nil {
		return http.Header{}
	}

	header.Add(accessTokenHeader, token.AccessToken)

	return header
}
