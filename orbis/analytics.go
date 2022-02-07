package orbis

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var (
	getOverexposureAnalysis     = "/v2/advisory/analytics/overexposure/%s/%s/%s"
	getRealtimeBalanceBreakDown = "/v2/advisory/analytics/breakdown/rtbs/%s"
	getModelPerformance         = "/v2/advisory/analytics/model/performance/%s/%d"
	getModelDriftReport         = "/v2/advisory/analytics/model/drift/%s"
)

func (c *Client) GetOverexposureAnalysis(overexposureType, scope, ID string, threshold string) (io.ReadCloser, error) {
	url := fmt.Sprintf(getOverexposureAnalysis, overexposureType, scope, ID)
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+url, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("threshold", threshold)
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetRealtimeBalanceBreakDown(modelID string, bucketSize int64) (io.ReadCloser, error) {
	url := fmt.Sprintf(getRealtimeBalanceBreakDown, modelID)
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+url, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("bucketSize", strconv.FormatInt(bucketSize, 10))
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetModelPerformance(modelID string, performanceRange int64) (io.ReadCloser, error) {
	url := fmt.Sprintf(getModelPerformance, modelID, performanceRange)
	r, err := c.client.Get(c.config.ApiUrl+url, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetModelDriftReport(modelID, threshold string) (io.ReadCloser, error) {
	url := fmt.Sprintf(getModelDriftReport, modelID)
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+url, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("threshold", threshold)
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}
