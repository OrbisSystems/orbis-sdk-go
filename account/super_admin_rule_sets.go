package account

const apiUrlSuperAdminGetRules = "/super-admin/rule-sets/list"

func (c *Client) GetRules(with []string) (interface{}, error) {
	body := map[string][]string{
		"with": with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminGetRules)
}
