package avatar

import (
	"fmt"
	"io"
)

const (
	apiUrlAvatarAuth       = "/authorize?env=%s"
	apiUrlAvatarRenew      = "/renew"
	apiUrlAvatarInvalidate = "/invalidate"
)

// AvatarAuth provides authorization ability for avatar APIs
// env - Orbis environment for checking authorization.
// orbisToken - Orbis token for checking authorization
func (c *Client) AvatarAuth(env, orbisToken string) (io.ReadCloser, error) {
	data := map[string]string{
		"orbisToken": orbisToken,
	}
	request, err := c.client.GetFormCodeRequest(data, c.config.ApiUrl+fmt.Sprintf(apiUrlAvatarAuth, env))
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(response)
}

// AvatarRenewToken provides renewal of token
// token - The token for renewal
func (c *Client) AvatarRenewToken(token string) (io.ReadCloser, error) {
	data := map[string]string{
		"token": token,
	}
	request, err := c.client.GetFormCodeRequest(data, c.config.ApiUrl+apiUrlAvatarRenew)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(response)
}

// AvatarTokenInvalidate provides invalidation of token
// token - The token for invalidation
func (c *Client) AvatarTokenInvalidate(token string) (io.ReadCloser, error) {
	data := map[string]string{
		"token": token,
	}
	request, err := c.client.GetFormCodeRequest(data, c.config.ApiUrl+apiUrlAvatarInvalidate)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(response)
}
