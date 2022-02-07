package orbis

import models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"

var deactivateDevice = "/api/device/deactivate"

func (c *Client) DeactivateDevice(req models.DeactivateDeviceBody) (models.DeactivateDeviceResponse, error) {
	var resp models.DeactivateDeviceResponse
	err := c.client.DoPostAndUnmarshall(req, &resp, c.config.ApiUrl+deactivateDevice)
	if err != nil {
		return models.DeactivateDeviceResponse{}, err
	}
	return resp, err
}
