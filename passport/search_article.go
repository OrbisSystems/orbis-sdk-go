package passport

import (
	"fmt"
	"github.com/OrbisSystems/orbis-sdk-go/models/passport"
	"github.com/pkg/errors"
)

var apiUrlPassportSearchArticle = "/api/%s/search?s=%s"

var (
	errorEmptyS    = errors.New("s must be not empty")
	errorEmptyLang = errors.New("lang must be not empty")
	errorEmptyID   = errors.New("ID must be not empty")
)

func (c *Client) SearchArticle(lang, s string, searchType passport.TypeArticle) ([]passport.Article, error) {
	url, err := getSearchArticleParamsAndValidate(lang, s, searchType)
	if err != nil {
		return nil, err
	}

	r, err := c.client.Get(c.config.ApiUrl+url, nil)
	if err != nil {
		return nil, err
	}

	var resp []passport.Article

	errUnmarshal := c.client.UnmarshalAndCheckOk(&resp, r)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return resp, nil
}

func (c *Client) SearchArticleSecretBySecret(lang, s string, searchType passport.TypeArticle, secret passport.Secret) ([]passport.Article, error) {
	url, err := getSearchArticleParamsAndValidate(lang, s, searchType)
	if err != nil {
		return nil, err
	}

	url = fmt.Sprintf("%s&api_secret=%s&api_key=%s", url, secret.APIKey, secret.APIKey)
	r, err := c.client.Get(c.config.ApiUrl+url, nil)
	if err != nil {
		return nil, err
	}

	var resp []passport.Article
	errUnmarshal := c.client.UnmarshalAndCheckOk(&resp, r)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return resp, nil
}

func getSearchArticleParamsAndValidate(lang string, s string, searchType passport.TypeArticle) (string, error) {
	if s == "" {
		return "", errorEmptyS
	}

	if lang == "" {
		return "", errorEmptyLang
	}

	url := fmt.Sprintf(apiUrlPassportSearchArticle, lang, s)

	if searchType != "" {
		url = fmt.Sprintf("%s?type=%s", url, searchType)
	}
	return url, nil
}
