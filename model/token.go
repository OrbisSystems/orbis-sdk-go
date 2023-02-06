package model

type Token struct {
	AccessToken  string
	RefreshToken string
}

type LoginRequest struct {
	Email    string
	Password string
}
