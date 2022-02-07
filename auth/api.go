package auth

import "net/http"

type API interface {
	Token
	BearerToken
	RefreshToken
}

type Token interface {
	GetToken() (string, error)
	SetToken(token string)
}

type BearerToken interface {
	GetTokenHeader(header http.Header) http.Header
}

type RefreshToken interface {
	StartRefresh() chan error
}
