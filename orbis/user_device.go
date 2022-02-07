package orbis

import (
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
	"net/http"
)

var (
	registerUserDevice   = "/user/device/register"
	deRegisterUserDevice = "/user/device/deregister"
	firebaseRegister     = "/user/device/firebase/register"
)

func (c *Client) RegisterUserDevice() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+registerUserDevice, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) DeRegisterUserDevice(token string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+deRegisterUserDevice, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("token", token)
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) FirebaseRegister(req models.FirebaseRegisterRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+firebaseRegister)
}
