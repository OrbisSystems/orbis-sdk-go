package orbis

import (
	"fmt"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
	"net/http"
)

var (
	getListRelationship         = "/transfer/relationship/list/%s"
	getPaymentMethods           = "/transfer/bank/paymentMethods"
	getPaymentMethodsViaAddress = "/transfer/bank/paymentMethods"
	getBankCodes                = "/transfer/bank/bankCodes"
	getBankCodesViaAddress      = "/transfer/bank/bankCodes"
	create                      = "/transfer/relationship/create"
	getWireInformation          = "/transfer/relationship/wire/create"
	bankingRelationship         = "/transfer/relationship/bank/create"
	approve                     = "/transfer/relationship/approve"
	cancel                      = "/transfer/relationship/cancel/%s"
	rename                      = "/transfer/relationship/rename/%s"
)

func (c *Client) GetListRelationship(account string) (io.ReadCloser, error) {
	url := fmt.Sprintf(getListRelationship, account)
	r, err := c.client.Get(c.config.ApiUrl+url, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetPaymentMethods(country string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getPaymentMethods, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("country", country)
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetPaymentMethodsViaAddress(req models.BankCodesViaAddressRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+getPaymentMethodsViaAddress)
}

func (c *Client) GetBankCodes(country string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getBankCodes, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("country", country)
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetBankCodesViaAddress(req models.BankCodesViaAddressRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+getBankCodesViaAddress)
}

func (c *Client) Create(req models.CreateTransferRequest) (models.Create, error) {
	var resp models.Create
	err := c.client.DoPostAndUnmarshall(req, &resp, c.config.ApiUrl+create)
	if err != nil {
		return models.Create{}, err
	}
	return resp, nil
}

func (c *Client) GetWireInformation(req models.RelationshipInformationRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+getWireInformation)
}

func (c *Client) BankingRelationship(req models.RelationshipInformationRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+bankingRelationship)
}

func (c *Client) Approve(req models.ApproveTransferRequest) (models.Approve, error) {
	var resp models.Approve
	err := c.client.DoPostAndUnmarshall(req, &resp, c.config.ApiUrl+approve)
	if err != nil {
		return models.Approve{}, err
	}
	return resp, nil
}

func (c *Client) Cancel(entryID string, req models.CancelRelationshipTransferRequest) (models.Cancel, error) {
	var resp models.Cancel
	err := c.client.DoPostAndUnmarshall(req, &resp, c.config.ApiUrl+fmt.Sprintf(cancel, entryID))
	if err != nil {
		return models.Cancel{}, err
	}
	return resp, nil
}

func (c *Client) Rename(entryID, nickname string) (models.Rename, error) {
	var resp models.Rename
	req := struct {
		Nickname string `json:"nickname"`
	}{
		Nickname: nickname,
	}
	err := c.client.DoPostAndUnmarshall(req, &resp, c.config.ApiUrl+fmt.Sprintf(rename, entryID))
	if err != nil {
		return models.Rename{}, err
	}
	return resp, nil
}
