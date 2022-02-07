package account

const apiUrlRegisterDevice = "/devices/register"

func (c *Client) RegisterDevice(name, token string) (interface{}, error) {
	body := map[string]string{
		"name":  name,
		"token": token,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlDocumentsGenerateQRCode)
}
