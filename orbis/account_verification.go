package orbis

import (
	"bytes"
	"encoding/json"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
	"net/http"
)

var (
	initiateAccountVerification = "/iav/init"
	completeAccount             = "/iav/complete"
	getListBankAccounts         = "/iav/list"
	getListBankAccount          = "/iav/link"
	importAccountVerification   = "iav/import"
)

func (c *Client) InitiateAccountVerification(subject string) (models.InitiatePost, error) {
	reqBody, err := json.Marshal(struct {
		Subject string `json:"subject"`
	}{Subject: subject})
	if err != nil {
		return models.InitiatePost{}, err
	}
	header := http.Header{}
	header.Add("Content-Type", "application/x-www-form-urlencoded")

	r, err := c.client.Post(c.config.ApiUrl+initiateAccountVerification, bytes.NewBuffer(reqBody), header)
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

func (c *Client) CompleteAccount(req models.PlaidRequest) (models.Complete, error) {
	var resp models.Complete
	err := c.client.DoPostAndUnmarshall(req, &resp, c.config.ApiUrl+completeAccount)
	if err != nil {
		return models.Complete{}, err
	}
	return resp, nil
}

func (c *Client) GetListBankAccounts(req models.PlaidRequest) (models.ListBankAccounts, error) {
	var resp models.ListBankAccounts
	err := c.client.DoPostAndUnmarshall(req, &resp, c.config.ApiUrl+getListBankAccounts)
	if err != nil {
		return models.ListBankAccounts{}, err
	}
	return resp, nil
}

func (c *Client) GetListBankAccount(req models.PlaidRequest) (models.ListBankAccount, error) {
	var resp models.ListBankAccount
	err := c.client.DoPostAndUnmarshall(req, &resp, c.config.ApiUrl+getListBankAccount)
	if err != nil {
		return models.ListBankAccount{}, err
	}
	return resp, nil
}

func (c *Client) Import(req models.ImportInstantRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+getListBankAccounts)
}
