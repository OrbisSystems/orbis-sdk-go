package account

const apiUrlCustodianSend = "/applications/custodian/send"

func (c *Client) SendToCustodian(applicationID string, with []string) (interface{}, error) {
	body := map[string]interface{}{
		"application_id": applicationID,
		"with":           with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlCustodianSend)
}
