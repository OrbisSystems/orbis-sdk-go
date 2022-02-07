package orbis

import (
	"fmt"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
)

var (
	getMutualFundsList   = "/quotes/mf/allowed"
	getMutualFundDetails = "/quotes/mf/allowed/details/%s"
	checkMutualFund      = "/quotes/mf/check/%s"
	mutualFundsPlacement = "/orders/v2/mf/place"
	mutualFundPreview    = "/orders/v2/mf/preview"
)

func (c *Client) GetMutualFundsList() ([]models.QuoteACAT, error) {
	r, err := c.client.Get(c.config.ApiUrl+getMutualFundsList, nil)
	if err != nil {
		return nil, err
	}

	var resp []models.QuoteACAT
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetMutualFundDetails(symbol string) (models.MutualFundDetails, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(getMutualFundDetails, symbol), nil)
	if err != nil {
		return nil, err
	}

	var resp models.MutualFundDetails
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) CheckMutualFund(symbol string) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(checkMutualFund, symbol), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) MutualFundsPlacement(req models.MutualFundsRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+mutualFundsPlacement)
}

func (c *Client) MutualFundPreview(req models.MutualFundsRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+mutualFundPreview)
}
