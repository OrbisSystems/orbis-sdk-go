package account

import (
	token "github.com/OrbisSystems/orbis-sdk-go/models/token"
	user "github.com/OrbisSystems/orbis-sdk-go/models/user"
)

type LoginResponse struct {
	Status bool      `json:"status"`
	Login  LoginData `json:"login"`
}

type LoginData struct {
	Token token.Token `json:"token"`
	User  user.User   `json:"user"`
}

type LoginRequest struct {
	Email    string   `json:"email"`
	Password string   `json:"password"`
	With     []string `json:"with,omitempty"`
}
