package account

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	sdk "github.com/OrbisSystems/orbis-sdk-go"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/model"
	"github.com/OrbisSystems/orbis-sdk-go/utils"
)

// Account service provides all API for account managing.
type Account struct {
	sdk.Auth

	cfg       config.Config
	cli       sdk.HTTPClient
	validator sdk.Validator
}

func New(cfg config.Config, auth sdk.Auth, cli sdk.HTTPClient, validator sdk.Validator) *Account {
	return &Account{
		Auth:      auth,
		cfg:       cfg,
		cli:       cli,
		validator: validator,
	}
}

// LoginByEmail uses for getting token if you are a client with the biggest access in you company.
// After log in, you can create api keys with limited access for you users using CreateAPIKey method.
// With these api keys your users can log in via SDK using LoginByAPIKey method.
func (c *Account) LoginByEmail(ctx context.Context, req model.LoginByEmailRequest) error {
	if err := c.validator.ValidateLoginByEmailRequest(req); err != nil {
		return err
	}

	if req.DeviceID == "" {
		req.DeviceID = utils.GenerateRandomString(10)
	}

	body, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := c.cli.Post(ctx, c.cfg.AuthHost+model.URLB2BLoginByEmail, bytes.NewBuffer(body), nil)
	if err != nil {
		return errors.Wrap(err, "account couldn't login by email")
	}

	var resp model.LoginByEmailResponse
	err = c.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return err
	}

	return c.SetToken(ctx, model.Token{
		AccessToken:  resp.LoginBasic.Tokens.AccessToken,
		RefreshToken: resp.LoginBasic.Tokens.RefreshToken,
	})
}

// LoginByAPIKey allows lo log in into the system for users with limited access.
// See LoginByEmail and CreateAPIKey for more information.
func (c *Account) LoginByAPIKey(ctx context.Context, req model.LoginByAPIKeyRequest) error {
	if err := c.validator.ValidateLoginByAPIKeyRequest(req); err != nil {
		return err
	}

	body, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := c.cli.Post(ctx, c.cfg.AuthHost+model.URLB2BLoginByAPIKey, bytes.NewBuffer(body), nil)
	if err != nil {
		return errors.Wrap(err, "account couldn't login by api key")
	}

	var resp model.LoginByAPIKeyResponse
	err = c.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return err
	}

	return c.SetToken(ctx, model.Token{
		AccessToken:  resp.ApiKeysLogin.Tokens.AccessToken,
		RefreshToken: resp.ApiKeysLogin.Tokens.RefreshToken,
	})
}

// CreateAPIKey provides creating api keys with limited access to branches, microservices, specific roles and permissions.
func (c *Account) CreateAPIKey(ctx context.Context, req model.CreateAPIKeyRequest) (model.CreateAPIKeyResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return model.CreateAPIKeyResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := c.cli.Post(ctx, c.cfg.AuthHost+model.URLB2BCreateAPIKey, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.CreateAPIKeyResponse{}, errors.Wrap(err, "account couldn't login")
	}

	var resp model.CreateAPIKeyResponse
	err = c.cli.UnmarshalAndCheckOk(&resp, r)

	return resp, err
}

// TODO token refresh
