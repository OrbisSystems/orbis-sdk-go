package avatar

import (
	"fmt"
	"io"
	"os"
)

const (
	apiUrlGetAccountAvatar    = "/account/%s?"
	apiUrlUploadAccountAvatar = "/account/%s?"
	apiUrlDeleteAccountAvatar = "/account/%s?"
)

func (c *Client) GetAccountAvatar(accountID, token string) (io.ReadCloser, error) {
	queryParams, err := getQueryParamsToken(accountID, token)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(apiUrlGetAccountAvatar, accountID) + queryParams
	fullUrl := c.config.ApiUrl + url

	r, err := c.client.Get(fullUrl, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}
func (c *Client) GetAccountAvatarByOrbisToken(accountID, orbisToken, env string) (io.ReadCloser, error) {
	queryParams, err := getQueryParamsOrbisToken(accountID, orbisToken, env)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(apiUrlGetAccountAvatar, accountID) + queryParams
	fullUrl := c.config.ApiUrl + url

	r, err := c.client.Get(fullUrl, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) UploadAccountAvatar(accountID, token string, avatar os.File) (io.ReadCloser, error) {
	queryParams, err := getQueryParamsToken(accountID, token)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(apiUrlUploadAccountAvatar, accountID) + queryParams
	fullUrl := c.config.ApiUrl + url

	//todo add file to body
	r, err := c.client.Post(fullUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}
func (c *Client) UploadAccountAvatarByOrbisToken(accountID, orbisToken, env string, avatar os.File) (io.ReadCloser, error) {
	queryParams, err := getQueryParamsOrbisToken(accountID, orbisToken, env)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(apiUrlUploadAccountAvatar, accountID) + queryParams
	fullUrl := c.config.ApiUrl + url

	//todo add file to body
	r, err := c.client.Post(fullUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) RemoveAccountAvatar(accountID, token string) (io.ReadCloser, error) {
	queryParams, err := getQueryParamsToken(accountID, token)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(apiUrlDeleteAccountAvatar, accountID) + queryParams
	fullUrl := c.config.ApiUrl + url

	r, err := c.client.Delete(fullUrl, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}
func (c *Client) RemoveAccountAvatarByOrbisToken(accountID, orbisToken, env string) (io.ReadCloser, error) {
	queryParams, err := getQueryParamsOrbisToken(accountID, orbisToken, env)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(apiUrlDeleteAccountAvatar, accountID) + queryParams
	fullUrl := c.config.ApiUrl + url

	r, err := c.client.Delete(fullUrl, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}
