package account

import "fmt"

const (
	apiUrlCustodianStatusUpdate = "/webhooks/custodian-status-update"
	apiUrlGenerateTwiMLForVoice = "/webhooks/custodian-status-update"
)

func (c *Client) CustodianStatusUpdate(accountNumber, status string) (interface{}, error) {
	body := map[string]string{
		"account_number": accountNumber,
		"status":         status,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlCustodianStatusUpdate)
}

func (c *Client) GenerateTwiMLForVoice(message string) (interface{}, error) {
	url := fmt.Sprintf("%s%s/say=%s", c.config.ApiUrl, apiUrlGenerateTwiMLForVoice, message)
	return c.client.MarshallAndSendPost(nil, url)
}
