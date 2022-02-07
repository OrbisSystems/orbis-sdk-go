package mock

import (
	httpClient "github.com/OrbisSystems/orbis-sdk-go/http"

	"io"
	"net/http"
)

type ClientHttpMock struct {
	resp   func() (*http.Response, error)
	client httpClient.Client
}

func NewMock(f func() (*http.Response, error), client httpClient.Client) httpClient.Client {
	return ClientHttpMock{resp: f, client: client}
}

func (c ClientHttpMock) Get(url string, headers http.Header) (*http.Response, error) {
	return c.resp()
}

func (c ClientHttpMock) Post(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	return c.resp()
}

func (c ClientHttpMock) Put(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	return c.resp()
}

func (c ClientHttpMock) Patch(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	return c.resp()
}

func (c ClientHttpMock) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.resp()
}

func (c ClientHttpMock) Do(request *http.Request) (*http.Response, error) {
	return c.resp()
}

func (c ClientHttpMock) GetBodyAndCheckStatus(r *http.Response, expectedStatus int) (io.ReadCloser, error) {
	return c.client.GetBodyAndCheckStatus(r, expectedStatus)
}

func (c ClientHttpMock) GetBodyAndCheckOK(r *http.Response) (io.ReadCloser, error) {
	return c.client.GetBodyAndCheckOK(r)
}

func (c ClientHttpMock) UnmarshalAndCheckOk(v interface{}, r *http.Response) error {
	return c.client.UnmarshalAndCheckOk(v, r)
}

func (c ClientHttpMock) MarshallAndSendPost(body interface{}, url string) (io.ReadCloser, error) {
	return c.client.MarshallAndSendPost(body, url)
}

func (c ClientHttpMock) GetFormCodeRequest(data map[string]string, url string) (*http.Request, error) {
	return c.client.GetFormCodeRequest(data, url)
}

func (c ClientHttpMock) DoPostAndReturnBody(req interface{}, url string) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, url)
}

func (c ClientHttpMock) DoPostAndUnmarshall(req, v interface{}, url string) error {
	return c.client.DoPostAndUnmarshall(req, v, url)
}
