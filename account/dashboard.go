package account

const apiUrlGetDashboardData = "/dashboard/get"

func (c *Client) GetDashboardData(interval string, with []string) (interface{}, error) {
	body := map[string]interface{}{
		"with":     with,
		"interval": interval,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlGetDashboardData)
}
