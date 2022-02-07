package passport

import (
	"fmt"
	"github.com/OrbisSystems/orbis-sdk-go/models/passport"
	"github.com/pkg/errors"
)

var apiUrlPassportAuth = "/jwt/authenticate?email=%s&password=%s"

func (c *Client) Auth(email, password string) (passport.JWTAuthenticate, error) {
	if email == "" || password == "" {
		return passport.JWTAuthenticate{}, errors.New("email and password should be not empty")
	}
	r, err := c.client.Post(c.config.ApiUrl+fmt.Sprintf(apiUrlPassportAuth, email, password), nil, nil)
	if err != nil {
		return passport.JWTAuthenticate{}, err
	}

	var resp passport.JWTAuthenticate

	errUnmarshal := c.client.UnmarshalAndCheckOk(&resp, r)
	if errUnmarshal != nil {
		return passport.JWTAuthenticate{}, errUnmarshal
	}

	return resp, nil
}
