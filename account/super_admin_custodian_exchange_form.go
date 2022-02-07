package account

const apiUrlSuperAdminGetCustodianExchangeFormList = "/custodian-exchange-form-keys/list"

func (c *Client) GetCustodianExchangeFormList(with []string) (interface{}, error) {
	body := map[string][]string{
		"with": with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminGetCustodianExchangeFormList)
}
