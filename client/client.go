package client

import (
	"context"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

const (
	AccessTokenHeader = "Authorization"
)

type OrbisClient struct {
	http.Client

	auth Auth
}

func NewClient(auth Auth) Client {
	var httpCli = http.Client{}

	cli := OrbisClient{
		Client: httpCli,
		auth:   auth,
	}

	return &cli
}

func NewClientWithHttp(auth Auth, client http.Client) Client {
	orbisClient := OrbisClient{
		auth:   auth,
		Client: client,
	}

	return &orbisClient
}

// Get makes an HTTP GET request to provided URL
func (c *OrbisClient) Get(ctx context.Context, url string, headers http.Header) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return response, errors.Wrap(err, "GET - request creation failed")
	}

	request.Header = headers

	return c.Do(ctx, request)
}

// Post makes an HTTP POST request to provided URL and requestBody
func (c *OrbisClient) Post(ctx context.Context, url string, body io.Reader, headers http.Header) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return response, errors.Wrap(err, "POST - request creation failed")
	}

	request.Header = headers

	return c.Do(ctx, request)
}

// Put makes an HTTP PUT request to provided URL and requestBody
func (c *OrbisClient) Put(ctx context.Context, url string, body io.Reader, headers http.Header) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return response, errors.Wrap(err, "PUT - request creation failed")
	}

	request.Header = headers

	return c.Do(ctx, request)
}

// Patch makes an HTTP PATCH request to provided URL and requestBody
func (c *OrbisClient) Patch(ctx context.Context, url string, body io.Reader, headers http.Header) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequest(http.MethodPatch, url, body)
	if err != nil {
		return response, errors.Wrap(err, "PATCH - request creation failed")
	}

	request.Header = headers

	return c.Do(ctx, request)
}

// Delete makes an HTTP DELETE request with provided URL
func (c *OrbisClient) Delete(ctx context.Context, url string, headers http.Header) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequest(http.MethodDelete, url, http.NoBody)
	if err != nil {
		return response, errors.Wrap(err, "DELETE - request creation failed")
	}

	request.Header = headers

	return c.Do(ctx, request)
}

// Do makes an HTTP request with the native `http.Do` interface
func (c *OrbisClient) Do(ctx context.Context, request *http.Request) (*http.Response, error) {
	request.Header = c.getTokenHeader(ctx, request.Header)
	request.Header.Add("Content-Type", "application/json")

	response, err := c.Client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *OrbisClient) getTokenHeader(ctx context.Context, header http.Header) http.Header {
	if header == nil {
		header = http.Header{}
	}
	token, err := c.auth.GetToken(ctx)
	if err != nil {
		return nil
	}

	header.Add(AccessTokenHeader, token.AccessToken)

	return header
}
