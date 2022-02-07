package mock

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
)

type ExpectedResp struct {
	Ok bool `json:"ok"`
}

var DefaultExpectedResp = ExpectedResp{Ok: true}

func GetReaderCloser(v interface{}) io.ReadCloser {
	resp, _ := json.Marshal(v)
	return ioutil.NopCloser(bytes.NewReader(resp))
}
