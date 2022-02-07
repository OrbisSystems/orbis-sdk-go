package account

const apiUrlSuperAdminGetServicesList = "/super-admin/services/list"

func (c *Client) GetServicesList(with []string) (interface{}, error) {
	body := map[string][]string{
		"with": with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminGetServicesList)
}
