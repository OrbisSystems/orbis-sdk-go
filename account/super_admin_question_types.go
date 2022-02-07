package account

const apiUrlSuperAdminGetServiceQuestionsList = "/super-admin/question-types/list"

func (c *Client) GetServiceQuestionsList(with []string) (interface{}, error) {
	body := map[string][]string{
		"with": with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminGetServiceQuestionsList)
}
