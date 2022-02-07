package orbis

import (
	"bytes"
	"encoding/json"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"net/http"
)

var (
	initiatePost             = "/iav/init"
	completePlaid            = "/iav/complete"
	getListBankAccountsPlaid = "/iav/list"
	getListBankAccountPlaid  = "/iav/link"
)

func (c *Client) InitiatePost(newValidation bool) (models.InitiatePost, error) {
	reqBody, err := json.Marshal(struct {
		NewValidation bool `json:"newValidation"`
	}{NewValidation: newValidation})
	if err != nil {
		return models.InitiatePost{}, err
	}
	header := http.Header{}
	header.Add("Content-Type", "application/x-www-form-urlencoded")

	r, err := c.client.Post(c.config.ApiUrl+initiatePost, bytes.NewBuffer(reqBody), header)
	if err != nil {
		return models.InitiatePost{}, err
	}

	var resp models.InitiatePost
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.InitiatePost{}, err
	}

	return resp, nil
}

func (c *Client) CompletePlaid(req models.PlaidRequest) (models.Complete, error) {
	var resp models.Complete
	err := c.client.DoPostAndUnmarshall(req, &resp, c.config.ApiUrl+completePlaid)
	if err != nil {
		return models.Complete{}, err
	}
	return resp, nil
}

func (c *Client) GetListBankAccountsPlaid(req models.PlaidRequest) (models.ListBankAccounts, error) {
	var resp models.ListBankAccounts
	err := c.client.DoPostAndUnmarshall(req, &resp, c.config.ApiUrl+getListBankAccountsPlaid)
	if err != nil {
		return models.ListBankAccounts{}, err
	}
	return resp, nil
}

func (c *Client) GetListBankAccountPlaid(req models.PlaidRequest) (models.ListBankAccount, error) {
	var resp models.ListBankAccount
	err := c.client.DoPostAndUnmarshall(req, &resp, c.config.ApiUrl+getListBankAccountsPlaid)
	if err != nil {
		return models.ListBankAccount{}, err
	}
	return resp, nil
}
