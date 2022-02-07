package http

import (
	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/storage"
	"github.com/pkg/errors"

	"io"
	"net/http"
	"time"
)

const defaultHTTPTimeout = 30 * time.Second

type OrbisClient struct {
	http.Client

	auth    auth.API
	timeout time.Duration
}

var DefaultAuthKey = "defaultOrbisKey"
var DefaultClient = getDefaultClient(auth.NewAuth(storage.NewInMemory(), DefaultAuthKey))

func getDefaultClient(api auth.API) Client {
	return NewClient(defaultHTTPTimeout, api)
}

func NewClient(timeout time.Duration, api auth.API) Client {
	var a = http.Client{
		Timeout: timeout,
	}

	orbisClient := OrbisClient{
		Client:  a,
		timeout: timeout,
		auth:    api,
	}

	return &orbisClient
}

func NewClientWithHttp(timeout time.Duration, client http.Client) Client {
	orbisClient := OrbisClient{
		Client:  client,
		timeout: timeout,
	}

	return &orbisClient
}

// Get makes a HTTP GET request to provided URL
func (c *OrbisClient) Get(url string, headers http.Header) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return response, errors.Wrap(err, "GET - request creation failed")
	}

	request.Header = headers

	return c.Do(request)
}

// Post makes a HTTP POST request to provided URL and requestBody
func (c *OrbisClient) Post(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return response, errors.Wrap(err, "POST - request creation failed")
	}

	request.Header = headers

	return c.Do(request)
}

// Put makes a HTTP PUT request to provided URL and requestBody
func (c *OrbisClient) Put(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return response, errors.Wrap(err, "PUT - request creation failed")
	}

	request.Header = headers

	return c.Do(request)
}

// Patch makes a HTTP PATCH request to provided URL and requestBody
func (c *OrbisClient) Patch(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequest(http.MethodPatch, url, body)
	if err != nil {
		return response, errors.Wrap(err, "PATCH - request creation failed")
	}

	request.Header = headers

	return c.Do(request)
}

// Delete makes a HTTP DELETE request with provided URL
func (c *OrbisClient) Delete(url string, headers http.Header) (*http.Response, error) {
	var response *http.Response
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return response, errors.Wrap(err, "DELETE - request creation failed")
	}

	request.Header = headers

	return c.Do(request)
}

// Do makes an HTTP request with the native `http.Do` interface
func (c *OrbisClient) Do(request *http.Request) (*http.Response, error) {
	request.Header = c.auth.GetTokenHeader(request.Header)
	request.Header.Add("Content-Type", "application/json")

	response, err := c.Client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
