package account

const (
	apiUrlListBranches = "/branches/list"
	apiUrlGetBranch    = "/branches/get"
)

func (c *Client) GetListBranches(with []string) (interface{}, error) {
	body := map[string][]string{
		"with": with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlListBranches)
}

func (c *Client) GetBranch(ID string) (interface{}, error) {
	body := map[string]string{
		"id": ID,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlGetBranch)
}
