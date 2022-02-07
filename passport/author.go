package passport

import (
	"fmt"
	"github.com/OrbisSystems/orbis-sdk-go/models/passport"
)

var apiUrlAuthorProfile = "/%s/author/%s"

func (c *Client) GetAuthorProfile(lang, ID string) (passport.Author, error) {
	if lang == "" {
		return passport.Author{}, errorEmptyLang
	}

	if ID == "" {
		return passport.Author{}, errorEmptyID
	}

	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(apiUrlAuthorProfile, lang, ID), nil)
	if err != nil {
		return passport.Author{}, err
	}

	var resp passport.Author

	errUnmarshal := c.client.UnmarshalAndCheckOk(&resp, r)
	if errUnmarshal != nil {
		return passport.Author{}, errUnmarshal
	}

	return resp, nil
}

func (c *Client) GetAuthorProfileBySecret(lang, ID string, secret passport.Secret) (passport.Author, error) {
	if lang == "" {
		return passport.Author{}, errorEmptyLang
	}

	if ID == "" {
		return passport.Author{}, errorEmptyID
	}

	url := fmt.Sprintf("%s?api_key=%s&api_secret=%s", fmt.Sprintf(apiUrlAuthorProfile, lang, ID),
		secret.APIKey, secret.APISecret)

	r, err := c.client.Get(c.config.ApiUrl+url, nil)
	if err != nil {
		return passport.Author{}, err
	}

	var resp passport.Author

	errUnmarshal := c.client.UnmarshalAndCheckOk(&resp, r)
	if errUnmarshal != nil {
		return passport.Author{}, errUnmarshal
	}

	return resp, nil
}
