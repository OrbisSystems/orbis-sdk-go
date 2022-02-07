package account

const apiUrlSuperAdminGetOptionsList = "/super-admin/options/list"

func (c *Client) GetOptionsList(with []string) (interface{}, error) {
	body := map[string][]string{
		"with": with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminGetOptionsList)
}
