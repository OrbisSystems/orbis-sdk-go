package orbis

import models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"

var (
	depositACH = "/transfer/ach/deposit"
	depositIRA = "/transfer/ach/deposit/ira"
	withdraw   = "/transfer/ach/withdraw"
)

func (c *Client) DepositACH(req models.ACHDepositRequest) (models.Deposit, error) {
	var resp models.Deposit
	err := c.client.DoPostAndUnmarshall(req, &resp, c.config.ApiUrl+depositACH)
	if err != nil {
		return models.Deposit{}, err
	}
	return resp, nil
}

func (c *Client) DepositIRA(req models.DepositIRARequest) (models.Deposit, error) {
	var resp models.Deposit
	err := c.client.DoPostAndUnmarshall(req, &resp, c.config.ApiUrl+depositIRA)
	if err != nil {
		return models.Deposit{}, err
	}
	return resp, nil
}

func (c *Client) Withdraw(req models.WithdrawRequest) (models.Withdraw, error) {
	var resp models.Withdraw
	err := c.client.DoPostAndUnmarshall(req, &resp, c.config.ApiUrl+withdraw)
	if err != nil {
		return models.Withdraw{}, err
	}
	return resp, nil
}
