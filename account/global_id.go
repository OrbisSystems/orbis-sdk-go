package account

import "github.com/OrbisSystems/orbis-sdk-go/models/account"

const apiUrlGlobalIDUpdateService = "/webhooks/global-id/update-services"

func (c *Client) UpdateBackgroundCheckService(req account.UpdateServicesReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlGlobalIDUpdateService)
}
