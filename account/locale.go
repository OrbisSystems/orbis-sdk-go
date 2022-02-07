package account

import "github.com/OrbisSystems/orbis-sdk-go/models/account"

const (
	apiUrlGetListLocales = "/locales/list"
	apiUrlChangeLocale   = "/locales/change"
)

func (c *Client) GetListLocales() (interface{}, error) {
	return c.client.MarshallAndSendPost(nil, c.config.ApiUrl+apiUrlGetListLocales)
}
func (c *Client) ChangeLocale(locale account.Locale) (interface{}, error) {
	return c.client.MarshallAndSendPost(locale, c.config.ApiUrl+apiUrlChangeLocale)
}
