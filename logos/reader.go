package logos

import (
	"fmt"
	"io"
)

const (
	apiUrlLogosGetAll            = "/symbol?page=&%d"
	apiUrlLogosSymbol            = "/symbol/%s"
	apiUrlLogosGetSymbols        = "/symbols/%s"
	apiUrlGetSymbolsSocials      = "/social/%s"
	apiUrlGetSymbolLogo          = "/logos/%s.png"
	apiUrlGetSymbolLogoConverted = "/logos/%s/%s"
)

func (c *Client) GetAllLogos(page int) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(apiUrlLogosGetAll, page), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetSymbolLogos(symbol string) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(apiUrlLogosSymbol, symbol), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetSymbolsLogos(symbols []string) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(apiUrlLogosGetSymbols, symbols), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetSymbolsSocials(symbol string) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(apiUrlGetSymbolsSocials, symbol), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetSymbolLogo(symbol string) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(apiUrlGetSymbolLogo, symbol), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetSymbolLogoConverted(symbol, conversion string) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(apiUrlGetSymbolLogoConverted, symbol, conversion), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}
