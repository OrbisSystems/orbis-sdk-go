package account

func (c *Client) SendPhoneVerification(typeVerification string) (interface{}, error) {
	header := map[string]string{
		"type": typeVerification,
	}
	return c.client.MarshallAndSendPost(header, c.config.ApiUrl+apiKeysListUrl)
}

func (c *Client) VerifyPhone(code string) (interface{}, error) {
	header := map[string]string{
		"type": code,
	}
	return c.client.MarshallAndSendPost(header, c.config.ApiUrl+apiKeysListUrl)
}
