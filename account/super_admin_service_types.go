package account

const apiUrlSuperAdminGetServiceTypesList = "/super-admin/service-types/list"

func (c *Client) GetServiceTypesList(with []string) (interface{}, error) {
	body := map[string][]string{
		"with": with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminGetServiceTypesList)
}
