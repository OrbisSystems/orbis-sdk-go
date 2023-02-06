package client

import (
	"context"
	"io"
	"net/http"

	"github.com/OrbisSystems/orbis-sdk-go/model"
)

type Client interface {
	Get(ctx context.Context, url string, headers http.Header) (*http.Response, error)
	Post(ctx context.Context, url string, body io.Reader, headers http.Header) (*http.Response, error)
	Put(ctx context.Context, url string, body io.Reader, headers http.Header) (*http.Response, error)
	Patch(ctx context.Context, url string, body io.Reader, headers http.Header) (*http.Response, error)
	Delete(ctx context.Context, url string, headers http.Header) (*http.Response, error)
	Do(ctx context.Context, request *http.Request) (*http.Response, error)
}

type Auth interface {
	SetToken(ctx context.Context, token model.Token) error
	GetToken(ctx context.Context) (model.Token, error)
}
