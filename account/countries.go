package account

const apiUrlCountiesList = "/countries/list"

func (c *Client) GetListCountries(with []string) (interface{}, error) {
	body := map[string][]string{
		"with": with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlCountiesList)
}
