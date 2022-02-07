package account

const (
	apiUrlStockLendingGetPosition         = "/stock-lending/positions"
	apiUrlStockLendingGetAccountPortfolio = "/stock-lending/account-portfolio"
)

func (c *Client) GetPositions() (interface{}, error) {
	return c.client.MarshallAndSendPost(nil, c.config.ApiUrl+apiUrlStockLendingGetPosition)
}
func (c *Client) GetAccountPortfolio() (interface{}, error) {
	return c.client.MarshallAndSendPost(nil, c.config.ApiUrl+apiUrlStockLendingGetAccountPortfolio)
}
