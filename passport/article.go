package passport

import (
	"fmt"
	"github.com/OrbisSystems/orbis-sdk-go/models/passport"
	"github.com/pkg/errors"
)

const (
	apiUrlGetArticles    = "/eng/all"
	apiUrlGetArticleByID = "/%s/%s"
	apiUrlTopTags        = "/%s/top/symbols"
)

func (c *Client) GetArticles(req passport.ArticleRequest) ([]passport.Article, error) {
	var queryParams string
	var isAdded = false

	if req.Type != "" {
		queryParams = queryParams + "type=" + string(req.Type)
		isAdded = true
	}

	if req.ReleasedAfter != "" {
		if isAdded {
			queryParams += "&"
		}
		queryParams = queryParams + "released_after=" + req.ReleasedAfter
		isAdded = true
	}
	if req.ReleasedBefore != "" {
		if isAdded {
			queryParams += "&"
		}
		queryParams = queryParams + "released_before=" + req.ReleasedAfter
		isAdded = true
	}
	if req.WithTag != "" {
		if isAdded {
			queryParams += "&"
		}
		queryParams = queryParams + "with_tag=" + req.WithTag
		isAdded = true
	}
	if req.WithSymbol != "" {
		if isAdded {
			queryParams += "&"
		}
		queryParams = queryParams + "with_symbol=" + req.WithSymbol
		isAdded = true
	}
	var url = apiUrlGetArticles
	if queryParams != "" {
		url = url + "?" + queryParams
	}
	r, err := c.client.Get(c.config.ApiUrl+url, nil)
	if err != nil {
		return nil, err
	}

	var resp []passport.Article
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetArticlesBySecret(req passport.ArticleWithSecretRequest) ([]passport.Article, error) {
	if req.APISecret == "" || req.APIKey == "" {
		return nil, errors.New("secret params must be set")
	}

	var queryParams string
	var isAdded = false

	if req.Type != "" {
		queryParams = queryParams + "type=" + string(req.Type)
		isAdded = true
	}

	if req.ReleasedAfter != "" {
		if isAdded {
			queryParams += "&"
		}
		queryParams = queryParams + "released_after=" + req.ReleasedAfter
		isAdded = true
	}
	if req.ReleasedBefore != "" {
		if isAdded {
			queryParams += "&"
		}
		queryParams = queryParams + "released_before=" + req.ReleasedAfter
		isAdded = true
	}
	if req.WithTag != "" {
		if isAdded {
			queryParams += "&"
		}
		queryParams = queryParams + "with_tag=" + req.WithTag
		isAdded = true
	}
	if req.WithSymbol != "" {
		if isAdded {
			queryParams += "&"
		}
		queryParams = queryParams + "with_symbol=" + req.WithSymbol
		isAdded = true
	}
	if isAdded {
		queryParams += "&"
	}
	queryParams += fmt.Sprintf("api_key=%s", req.APIKey)
	queryParams += fmt.Sprintf("&api_secret=%s", req.APISecret)

	url := apiUrlGetArticles + "?" + queryParams
	r, err := c.client.Get(c.config.ApiUrl+url, nil)
	if err != nil {
		return nil, err
	}

	var resp []passport.Article
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetArticleByID(ID, lang string) (passport.Article, error) {
	if ID == "" {
		return passport.Article{}, errorEmptyID
	}
	if lang == "" {
		return passport.Article{}, errorEmptyLang
	}

	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(apiUrlGetArticleByID, lang, ID), nil)
	if err != nil {
		return passport.Article{}, err
	}
	var resp passport.Article

	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return passport.Article{}, err
	}

	return resp, nil
}

func (c *Client) GetArticleByIDSecret(ID, lang string, secret passport.Secret) (passport.Article, error) {
	errSecret := checkSecret(secret)
	if errSecret != nil {
		return passport.Article{}, errSecret
	}

	if ID == "" {
		return passport.Article{}, errorEmptyID
	}
	if lang == "" {
		return passport.Article{}, errorEmptyLang
	}

	queryParams := fmt.Sprintf("?api_key=%s&api_secret=%s", secret.APIKey, secret.APIKey)
	url := fmt.Sprintf(apiUrlGetArticleByID, lang, ID) + queryParams
	r, err := c.client.Get(c.config.ApiUrl+url, nil)
	if err != nil {
		return passport.Article{}, err
	}
	var resp passport.Article

	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return passport.Article{}, err
	}

	return resp, nil
}

func (c *Client) GetMostPopular(typeArticle passport.TypeArticle, time, lang string) (
	passport.MostPopularTagsKeywords, error) {
	if typeArticle == "" {
		return nil, errors.New("type is required parameter")
	}
	if lang == "" {
		return nil, errorEmptyLang
	}

	queryParams := "?type=" + string(typeArticle)
	if time != "" {
		queryParams = queryParams + "&" + "time=" + time
	}

	url := fmt.Sprintf(apiUrlTopTags, lang) + queryParams
	r, err := c.client.Get(c.config.ApiUrl+url, nil)
	if err != nil {
		return nil, err
	}
	var resp passport.MostPopularTagsKeywords

	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

func (c *Client) GetMostPopularBySecret(typeArticle passport.TypeArticle, time, lang string, secret passport.Secret) (
	passport.MostPopularTagsKeywords, error) {
	errSecret := checkSecret(secret)
	if errSecret != nil {
		return nil, errSecret
	}

	if typeArticle == "" {
		return nil, errors.New("type is required parameter")
	}
	if lang == "" {
		return nil, errorEmptyLang
	}

	queryParams := "?type=" + string(typeArticle)
	if time != "" {
		queryParams = queryParams + "&" + "time=" + time
	}
	queryParams += fmt.Sprintf("&api_key=%s&api_secret=%s", secret.APIKey, secret.APIKey)
	url := fmt.Sprintf(apiUrlTopTags, lang) + queryParams
	r, err := c.client.Get(c.config.ApiUrl+url, nil)
	if err != nil {
		return nil, err
	}
	var resp passport.MostPopularTagsKeywords

	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
