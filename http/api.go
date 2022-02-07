package http

import (
	"io"
	"net/http"
)

type Client interface {
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, body io.Reader, headers http.Header) (*http.Response, error)
	Put(url string, body io.Reader, headers http.Header) (*http.Response, error)
	Patch(url string, body io.Reader, headers http.Header) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
	Do(request *http.Request) (*http.Response, error)

	Wrapper
}

type Wrapper interface {
	GetBodyAndCheckStatus(r *http.Response, expectedStatus int) (io.ReadCloser, error)
	GetBodyAndCheckOK(r *http.Response) (io.ReadCloser, error)
	UnmarshalAndCheckOk(v interface{}, r *http.Response) error
	MarshallAndSendPost(body interface{}, url string) (io.ReadCloser, error)
	GetFormCodeRequest(data map[string]string, url string) (*http.Request, error)
	DoPostAndReturnBody(req interface{}, url string) (io.ReadCloser, error)
	DoPostAndUnmarshall(req, v interface{}, url string) error
}
