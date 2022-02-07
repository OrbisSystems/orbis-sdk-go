package account

const (
	apiUrlGetListSettings = "/settings/list"
	apiUrlGetSettings     = "/settings/get"
	apiUrlSetSettings     = "/settings/set"
)

func (c *Client) GetListSettings() (interface{}, error) {
	return c.client.MarshallAndSendPost(nil, c.config.ApiUrl+apiUrlGetListSettings)
}

func (c *Client) GetSetting(key string) (interface{}, error) {
	body := map[string]string{
		"key": key,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlGetSettings)
}
func (c *Client) SetSetting(key, value string) (interface{}, error) {
	body := map[string]interface{}{
		"key": key,
		"value": struct {
			Key string `json:"key"`
		}{Key: value},
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSetSettings)
}
