package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	netUrl "net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func (c *OrbisClient) GetBodyAndCheckStatus(r *http.Response, expectedStatus []int) (io.ReadCloser, error) {
	for _, v := range expectedStatus {
		if v == r.StatusCode {
			return r.Body, nil
		}
	}
	return nil, fmt.Errorf("wrong status code %d", r.StatusCode)
}

func (c *OrbisClient) GetBodyAndCheckOK(r *http.Response) (io.ReadCloser, error) {
	return c.GetBodyAndCheckStatus(r, []int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent})
}

func (c *OrbisClient) MarshallAndSendPost(ctx context.Context, v interface{}, url string) (io.ReadCloser, error) {
	body, err := c.GetMarshallingBody(v)
	if err != nil {
		return nil, err
	}

	r, err := c.Post(ctx, url, body, nil)
	if err != nil {
		return nil, errors.Wrap(err, "account couldn't login")
	}

	return c.GetBodyAndCheckOK(r)
}

func (c *OrbisClient) GetMarshallingBody(v interface{}) (*bytes.Buffer, error) {
	if v == nil {
		return nil, nil
	}
	body, err := json.Marshal(v)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	return bytes.NewBuffer(body), err
}

func (c *OrbisClient) GetFormCodeRequest(data map[string]string, url string) (*http.Request, error) {
	dataHeader := netUrl.Values{}
	for s := range data {
		dataHeader.Add(s, data[s])
	}

	encodedData := dataHeader.Encode()
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(encodedData))
	if err != nil {
		return nil, fmt.Errorf("couldn't create request %s", err.Error())
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(encodedData)))

	return req, err
}

func (c *OrbisClient) UnmarshalAndCheckOk(v interface{}, r *http.Response) error {
	body, err := c.GetBodyAndCheckOK(r)
	if err != nil {
		return err
	}

	defer body.Close()

	return json.NewDecoder(body).Decode(v)
}

func (c *OrbisClient) DoPostAndReturnBody(ctx context.Context, req interface{}, url string) (io.ReadCloser, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	r, err := c.Post(ctx, url, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, err
	}

	return c.GetBodyAndCheckOK(r)
}

func (c *OrbisClient) DoPostAndUnmarshall(ctx context.Context, req, v interface{}, url string) error {
	b, err := c.DoPostAndReturnBody(ctx, req, url)
	if err != nil {
		return err
	}

	defer b.Close()

	return json.NewDecoder(b).Decode(v)
}
