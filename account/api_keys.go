package account

const (
	apiKeysListUrl   = "/auth/api-keys/list"
	apiKeysCreateUrl = "/auth/api-keys/create"
	apiKeysDeleteUrl = "/auth/api-keys/delete"
)

func (c *Client) List(with []string) (interface{}, error) {
	header := map[string][]string{
		"with": with,
	}
	return c.client.MarshallAndSendPost(header, c.config.ApiUrl+apiKeysListUrl)

}
func (c *Client) Create(with []string) (interface{}, error) {
	header := map[string][]string{
		"with": with,
	}
	return c.client.MarshallAndSendPost(header, c.config.ApiUrl+apiKeysCreateUrl)
}
func (c *Client) Delete(ID string, with []string) (interface{}, error) {
	header := map[string]interface{}{
		"with": with,
		"id":   ID,
	}
	return c.client.MarshallAndSendPost(header, c.config.ApiUrl+apiKeysDeleteUrl)
}
