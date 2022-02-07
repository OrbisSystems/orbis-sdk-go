package account

const apiUrlStates = "/states/list"

func (c *Client) GetListStates() (interface{}, error) {
	return c.client.MarshallAndSendPost(nil, c.config.ApiUrl+apiUrlStates)
}
