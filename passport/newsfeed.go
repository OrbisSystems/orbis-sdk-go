package passport

import (
	"fmt"
	"github.com/OrbisSystems/orbis-sdk-go/models/passport"
	"github.com/pkg/errors"
)

var apiUrlPassportNewsfeed = "/newsfeed/%s/all"

func (c *Client) Newsfeed(lang, search string) ([]passport.Article, error) {
	if lang == "" {
		return nil, errors.New("lang must be not empty")
	}

	var queryParams string
	if search != "" {
		queryParams = "search=" + search
	}

	url := fmt.Sprintf(apiUrlPassportNewsfeed, lang)
	if queryParams != "" {
		url = url + "?" + queryParams
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
