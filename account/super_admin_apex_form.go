package account

const apiUrlGetApexFormList = "/super-admin/apex-form-keys/list"

func (c *Client) GetApexFormList(with []string) (interface{}, error) {
	body := map[string][]string{
		"with": with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlGetApexFormList)
}
