package account

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	sdk "github.com/OrbisSystems/orbis-sdk-go"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

var (
	ErrEmptyRefreshToken = errors.New("empty refresh token from storage")
)

// Account service provides all API for account managing.
type Account struct {
	sdk.Auth

	cfg config.Config
	cli sdk.HTTPClient

	resetTokenRefreshCh    chan int64
	watchTokenRefreshState bool
}

func New(cfg config.Config, auth sdk.Auth, cli sdk.HTTPClient) *Account {
	a := &Account{
		Auth: auth,
		cfg:  cfg,
		cli:  cli,

		resetTokenRefreshCh: make(chan int64),
	}

	a.watchTokenRefresh()

	return a
}

func (c *Account) watchTokenRefresh() {
	// run background routing for token refreshing
	go func() {
		log.Trace("watchTokenRefresh started")
		c.watchTokenRefreshState = true

		defer func() {
			log.Trace("watchTokenRefresh stopped")
			c.watchTokenRefreshState = false
		}()

		var refreshDuration time.Duration

		tkn, err := c.Auth.GetToken(context.Background())

		// if got an error from storage
		// or token is not presented in storage
		// or access token was expired
		// -> exit, need to call Login flow
		if err != nil || tkn.AccessExpiresAt == 0 || tkn.AccessExpiresAt < time.Now().Unix() {
			log.Tracef("reason for stoping watchTokenRefresh. Err %v, tkn.AccessExpiresAt %d", err, tkn.AccessExpiresAt)
			return
		}

		refreshDuration = getRefreshDuration(tkn.AccessExpiresAt)

		ticker := time.NewTicker(refreshDuration)

		for {
			select {
			case <-ticker.C:
				log.Trace("refreshing token...")
				rCtx, _ := context.WithTimeout(context.Background(), time.Minute)
				err := c.RefreshToken(rCtx)
				if err != nil {
					log.Errorf("error while updating token by refresh token: %v", err)
				}
			case v := <-c.resetTokenRefreshCh:
				log.Tracef("resetting token refresh duration... new unix time: %d", v)
				ticker.Reset(getRefreshDuration(v))
			}
		}
	}()
}

// LoginByEmail uses for getting token if you are a client with the biggest access in you company.
// After log in, you can create api keys with limited access for you users using CreateAPIKey method.
// With these api keys your users can log in via SDK using LoginByAPIKey method.
func (c *Account) LoginByEmail(ctx context.Context, req model.LoginByEmailRequest) error {
	log.Trace("LoginByEmail called")

	if req.DeviceID == "" {
		req.DeviceID = uuid.New().String()
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

	c.processNewToken(resp.LoginBasic.Tokens)

	return c.SetToken(ctx, resp.LoginBasic.Tokens)
}

// LoginByAPIKey allows lo log in into the system for users with limited access.
// See LoginByEmail and CreateAPIKey for more information.
func (c *Account) LoginByAPIKey(ctx context.Context, req model.LoginByAPIKeyRequest) error {
	log.Trace("LoginByAPIKey called")

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

	c.processNewToken(resp.ApiKeysLogin.Tokens)

	return c.SetToken(ctx, resp.ApiKeysLogin.Tokens)
}

// CreateAPIKey provides creating api keys with limited access to branches, microservices, specific roles and permissions.
func (c *Account) CreateAPIKey(ctx context.Context, req model.CreateAPIKeyRequest) (model.CreateAPIKeyResponse, error) {
	log.Trace("CreateAPIKey called")

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

func (c *Account) RefreshToken(ctx context.Context) error {
	log.Trace("RefreshToken called")

	tkn, err := c.GetToken(ctx)
	if err != nil {
		return err
	}
	if tkn.RefreshToken == "" {
		return ErrEmptyRefreshToken
	}

	req := model.RefreshTokenRequest{RefreshToken: tkn.RefreshToken}
	body, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := c.cli.Post(ctx, c.cfg.AuthHost+model.URLB2BRefreshToken, bytes.NewBuffer(body), nil)
	if err != nil {
		return errors.Wrap(err, "account couldn't refresh token")
	}

	var resp model.LoginByEmailResponse
	err = c.cli.UnmarshalAndCheckOk(&resp, r)

	if err != nil {
		return err
	}

	c.processNewToken(resp.LoginBasic.Tokens)

	return c.SetToken(ctx, resp.LoginBasic.Tokens)
}

func getRefreshDuration(v int64) time.Duration {
	expTime := time.Unix(v, 0)
	return expTime.UTC().Sub(time.Now())
}

func (c *Account) processNewToken(tkn model.Token) {
	if !c.watchTokenRefreshState {
		c.watchTokenRefresh()
	}

	time.Sleep(time.Millisecond * 100)

	go func() {
		c.resetTokenRefreshCh <- tkn.AccessExpiresAt
	}()
}
