package account

const apiUrlGetDocumentTypesList = "/super-admin/document-types/list"

func (c *Client) GetDocumentTypesList(with []string) (interface{}, error) {
	body := map[string][]string{
		"with": with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlGetDocumentTypesList)
}
