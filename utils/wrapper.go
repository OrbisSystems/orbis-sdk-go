package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBodyAndCheckOK(r *http.Response) (io.ReadCloser, error) {
	return getBodyAndCheckStatus(r, []int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent})
}

func UnmarshalAndCheckOk(v interface{}, r *http.Response) error {
	body, err := GetBodyAndCheckOK(r)
	if err != nil {
		return err
	}

	defer body.Close()

	return json.NewDecoder(body).Decode(v)
}

func getBodyAndCheckStatus(r *http.Response, expectedStatus []int) (io.ReadCloser, error) {
	for _, v := range expectedStatus {
		if v == r.StatusCode {
			return r.Body, nil
		}
	}

	var bodyMap interface{}
	json.NewDecoder(r.Body).Decode(&bodyMap)

	return nil, fmt.Errorf("wrong status code %d. Body: %v", r.StatusCode, bodyMap)
}
