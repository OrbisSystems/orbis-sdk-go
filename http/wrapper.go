package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *OrbisClient) GetBodyAndCheckOK(r *http.Response) (io.ReadCloser, error) {
	return c.getBodyAndCheckStatus(r, []int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent})
}

func (c *OrbisClient) UnmarshalAndCheckOk(v interface{}, r *http.Response) error {
	body, err := c.GetBodyAndCheckOK(r)
	if err != nil {
		return err
	}

	defer body.Close()

	return json.NewDecoder(body).Decode(v)
}

func (c *OrbisClient) getBodyAndCheckStatus(r *http.Response, expectedStatus []int) (io.ReadCloser, error) {
	for _, v := range expectedStatus {
		if v == r.StatusCode {
			return r.Body, nil
		}
	}

	var bodyMap interface{}
	json.NewDecoder(r.Body).Decode(&bodyMap)

	return nil, fmt.Errorf("wrong status code %d. Body: %v", r.StatusCode, bodyMap)
}
