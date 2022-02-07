package orbis

import (
	"fmt"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
)

var (
	searchAllocation     = "/v2/allocations/search"
	getDetailsAllocation = "/v2/allocations/details/%s"
	cancelAllocation     = "/v2/allocations/cancel"
	validateAllocation   = "/v2/allocations/validate"
	scheduleAllocation   = "/v2/allocations/preallocate"
	triggerAllocation    = "/v2/allocations/trigger"
)

func (c *Client) SearchAllocation() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+searchAllocation, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetDetailsAllocation(allocationRef string) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(getDetailsAllocation, allocationRef), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) CancelAllocation(allocationRef string) (io.ReadCloser, error) {
	req := struct {
		AllocationRef string `json:"allocationRef"`
	}{
		AllocationRef: allocationRef,
	}
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+cancelAllocation)
}

func (c *Client) ValidateAllocation(allocationRef string) (io.ReadCloser, error) {
	req := struct {
		AllocationRef string `json:"allocationRef"`
	}{
		AllocationRef: allocationRef,
	}
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+validateAllocation)
}

func (c *Client) ScheduleAllocation(req models.ScheduleAllocationRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+scheduleAllocation)
}

func (c *Client) TriggerAllocation(allocationRef string) (io.ReadCloser, error) {
	req := struct {
		AllocationRef string `json:"allocationRef"`
	}{
		AllocationRef: allocationRef,
	}
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+triggerAllocation)
}
