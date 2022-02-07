package orbis

import (
	"fmt"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
)

var (
	getAvailableAgreements = "/user/agreements/all/active"
	getAgreement           = "/user/agreements/contents/%s/%s/body.html"
	unsignedAgreements     = "/user/agreements/unsigned"
	attestStatus           = "/user/professional/attest"
	signAgreement          = "/user/agreements/sign"
	getCryptoAgreement     = "/user/agreements/crypto/Agreement"
	signCryptoAgreement    = "/user/agreements/crypto/sign"
)

func (c *Client) GetAvailableAgreements() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+getAvailableAgreements, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetAgreement(version, code string) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(getAgreement, version, code), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) UnsignedAgreements() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+unsignedAgreements, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) AttestStatus(req models.AttestStatusRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+attestStatus)
}

func (c *Client) SignAgreement(req models.SignAgreementRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+signAgreement)
}

func (c *Client) SignCryptoAgreement(req models.SignCryptoAgreementRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+signCryptoAgreement)
}

func (c *Client) GetCryptoAgreement() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+getCryptoAgreement, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}
