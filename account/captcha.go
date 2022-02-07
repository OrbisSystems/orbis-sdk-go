package account

func (c *Client) GetCaptcha() (interface{}, error) {
	return c.client.MarshallAndSendPost(nil, c.config.ApiUrl+apiKeysListUrl)
}
