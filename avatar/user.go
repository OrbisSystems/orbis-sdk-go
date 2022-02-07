package avatar

import (
	"fmt"
	"io"
	"os"
)

const (
	apiUrlGetUserAvatar    = "/user/%s?"
	apiUrlUploadUserAvatar = "/user/%s?"
	apiUrlDeleteUserAvatar = "/user/%s?"
)

// GetUserAvatar - get user avatar
// userID - ID of user
// token - Generated token by `/authorize` endpoint. Required when `orbisToken` is not provided.
// orbisToken - Orbis token. Required when `token` is not provided.
// env - Orbis environment. Required when `orbisToken` is provided.
func (c *Client) GetUserAvatar(userID, token string) (io.ReadCloser, error) {
	queryParams, err := getQueryParamsToken(userID, token)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(apiUrlGetUserAvatar, userID) + queryParams
	fullUrl := c.config.ApiUrl + url

	r, err := c.client.Get(fullUrl, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetUserAvatarByOrbisToken(userID, orbisToken, env string) (io.ReadCloser, error) {
	queryParams, err := getQueryParamsOrbisToken(userID, orbisToken, env)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(apiUrlGetUserAvatar, userID) + queryParams
	fullUrl := c.config.ApiUrl + url

	r, err := c.client.Get(fullUrl, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) UploadUserAvatar(userID, token string, avatar os.File) (io.ReadCloser, error) {
	queryParams, err := getQueryParamsToken(userID, token)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(apiUrlUploadUserAvatar, userID) + queryParams
	fullUrl := c.config.ApiUrl + url

	r, err := c.client.Post(fullUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}
func (c *Client) UploadUserAvatarByOrbisToken(userID, orbisToken, env string, avatar os.File) (io.ReadCloser, error) {
	queryParams, err := getQueryParamsOrbisToken(userID, orbisToken, env)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(apiUrlUploadUserAvatar, userID) + queryParams
	fullUrl := c.config.ApiUrl + url

	r, err := c.client.Post(fullUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) RemoveUserAvatar(userID, token string) (io.ReadCloser, error) {
	queryParams, err := getQueryParamsToken(userID, token)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(apiUrlDeleteUserAvatar, userID) + queryParams
	fullUrl := c.config.ApiUrl + url

	r, err := c.client.Delete(fullUrl, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) RemoveUserAvatarByOrbisToken(userID, orbisToken, env string) (io.ReadCloser, error) {
	queryParams, err := getQueryParamsOrbisToken(userID, orbisToken, env)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(apiUrlDeleteUserAvatar, userID) + queryParams
	fullUrl := c.config.ApiUrl + url

	r, err := c.client.Delete(fullUrl, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}
