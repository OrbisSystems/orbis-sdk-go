package auth

import (
	"fmt"
	"net/http"
)

func (a *Auth) GetTokenHeader(header http.Header) http.Header {
	if header == nil {
		header = http.Header{}
	}
	token, err := a.GetToken()
	if err != nil {
		return nil
	}

	header.Add(AccessTokenHeader, fmt.Sprintf("%s %s", BearerSchema, token))

	return header
}
