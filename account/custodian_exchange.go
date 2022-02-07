package account

import "github.com/OrbisSystems/orbis-sdk-go/models/account"

const (
	apiUrlUpdateWatchmenCompleteStatus = "/webhooks/custodian-exchange/watchmen-complete"
	apiUrlPushALEs                     = "/webhooks/custodian-exchange/ales"
	apiUrlPushSketches                 = "/webhooks/custodian-exchange/sketches"
)

func (c *Client) UpdateWatchmenCompleteStatus(req account.WatchmenCompleteReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlUpdateWatchmenCompleteStatus)
}
func (c *Client) PushALEs(req account.PushALEsReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlPushALEs)
}
func (c *Client) PushSketches(req account.PushSketchesReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlPushSketches)
}
