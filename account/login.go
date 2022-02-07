package account

import (
	"bytes"
	"encoding/json"
	account "github.com/OrbisSystems/orbis-sdk-go/models/account"
	"github.com/pkg/errors"
)

const (
	authLoginUrl  = `/auth/login`
	authLogoutUrl = `/auth/logout`
	authRefresh   = `/api/auth/refresh`
)

// Login - provides login to the account system
func (c *Client) Login(req account.LoginRequest) (account.Login, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return account.Login{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := c.client.Post(c.config.ApiUrl+authLoginUrl, bytes.NewBuffer(body), nil)
	if err != nil {
		return account.Login{}, errors.Wrap(err, "account couldn't login")
	}

	var resp account.Login
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return account.Login{}, err
	}

	c.SetToken(resp.Login.Token.AccessToken)

	return resp, nil
}

// Logout - provides logout from the account system
func (c *Client) Logout() error {
	//todo remove token
	_, err := c.client.Post(c.config.ApiUrl+authLogoutUrl, nil, nil)
	if err != nil {
		return errors.Wrap(err, "couldn't logout")
	}

	return nil
}

func (c *Client) Refresh(errChan chan error) {
	for {
		select {
		case <-c.tickerRefresh.C:
			r, err := c.client.Post(c.config.ApiUrl+authRefresh, nil, nil)
			if err != nil {
				errChan <- err
			}

			var resp account.Login
			err = c.client.UnmarshalAndCheckOk(&resp, r)
			if err != nil {
				errChan <- err
			}

			c.SetToken(resp.Login.Token.AccessToken)
		}
	}
}
