package account

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	sdk "github.com/OrbisSystems/orbis-sdk-go/interface"
	"github.com/OrbisSystems/orbis-sdk-go/model"
	"github.com/OrbisSystems/orbis-sdk-go/utils"
)

var (
	ErrEmptyRefreshToken = errors.New("empty refresh token from storage")
)

// Account service provides all API for account managing.
type Account struct {
	sdk.Auth

	url string
	cli sdk.HTTPClient

	resetTokenRefreshCh    chan int64
	watchTokenRefreshState bool

	refreshTicker *time.Ticker
}

func New(url string, auth sdk.Auth, cli sdk.HTTPClient) *Account {
	a := &Account{
		Auth: auth,
		url:  url,
		cli:  cli,

		resetTokenRefreshCh: make(chan int64),
		refreshTicker:       time.NewTicker(time.Hour * 100), // default
	}

	a.watchTokenRefresh()

	return a
}

func (a *Account) watchTokenRefresh() {
	// run background routing for token refreshing
	go func() {
		log.Trace("watchTokenRefresh started")
		a.watchTokenRefreshState = true

		defer func() {
			log.Trace("watchTokenRefresh stopped")
			a.watchTokenRefreshState = false
		}()

		var refreshDuration time.Duration

		tkn, err := a.Auth.GetToken(context.Background())

		// if no error from storage
		// and token is presented in storage
		// and access token is not expired
		// -> make a ticker
		if err == nil && tkn.AccessExpiresAt != 0 && tkn.AccessExpiresAt > time.Now().Unix() {
			refreshDuration = getRefreshDuration(tkn.AccessExpiresAt)
			a.refreshTicker = time.NewTicker(refreshDuration)
		}

		for {
			select {
			case <-a.refreshTicker.C:
				log.Trace("refreshing token...")
				rCtx, _ := context.WithTimeout(context.Background(), time.Minute)
				err := a.RefreshToken(rCtx)
				if err != nil {
					log.Errorf("error while updating token by refresh token: %v", err)
				}
			case v := <-a.resetTokenRefreshCh:
				log.Tracef("resetting token refresh duration... new unix time: %d", v)
				a.refreshTicker.Reset(getRefreshDuration(v))
			}
		}
	}()
}

// LoginByEmail uses for getting token if you are a client with the biggest access in you company.
// After log in, you can create api keys with limited access for you users using CreateAPIKey method.
// With these api keys your users can log in via SDK using LoginByAPIKey method.
func (a *Account) LoginByEmail(ctx context.Context, req model.LoginByEmailRequest) error {
	log.Trace("LoginByEmail called")

	if req.DeviceID == "" {
		req.DeviceID = uuid.New().String()
	}

	body, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := a.cli.Post(ctx, a.url+model.URLB2BLoginByEmail, bytes.NewBuffer(body), nil)
	if err != nil {
		return errors.Wrap(err, "account couldn't login by email")
	}

	var resp model.LoginByEmailResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return err
	}

	a.processNewToken(resp.LoginBasic.Tokens)

	return a.SetToken(ctx, resp.LoginBasic.Tokens)
}

// LoginByAPIKey allows lo log in into the system for users with limited access.
// See LoginByEmail and CreateAPIKey for more information.
func (a *Account) LoginByAPIKey(ctx context.Context, req model.LoginByAPIKeyRequest) error {
	log.Trace("LoginByAPIKey called")

	body, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := a.cli.Post(ctx, a.url+model.URLB2BLoginByAPIKey, bytes.NewBuffer(body), nil)
	if err != nil {
		return errors.Wrap(err, "account couldn't login by api key")
	}

	var resp model.LoginByAPIKeyResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return err
	}

	a.processNewToken(resp.ApiKeysLogin.Tokens)

	return a.SetToken(ctx, resp.ApiKeysLogin.Tokens)
}

// CreateAPIKey provides creating api keys with limited access to branches, microservices, specific roles and permissions.
func (a *Account) CreateAPIKey(ctx context.Context, req model.CreateAPIKeyRequest) (model.CreateAPIKeyResponse, error) {
	log.Trace("CreateAPIKey called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.CreateAPIKeyResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := a.cli.Post(ctx, a.url+model.URLB2BCreateAPIKey, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.CreateAPIKeyResponse{}, errors.Wrap(err, "account couldn't login")
	}

	var resp model.CreateAPIKeyResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)

	return resp, err
}

func (a *Account) RefreshToken(ctx context.Context) error {
	log.Trace("RefreshToken called")

	// marking this flag that we are refreshing the token
	a.setRefreshingState(true)
	defer a.setRefreshingState(false)

	tkn, err := a.GetToken(ctx)
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

	r, err := a.cli.Post(ctx, a.url+model.URLB2BRefreshToken, bytes.NewBuffer(body), nil)
	if err != nil {
		return errors.Wrap(err, "account couldn't refresh token")
	}

	var resp model.LoginByEmailResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)

	if err != nil {
		return err
	}

	a.processNewToken(resp.LoginBasic.Tokens)

	return a.SetToken(ctx, resp.LoginBasic.Tokens)
}

func getRefreshDuration(v int64) time.Duration {
	expTime := time.Unix(v, 0).Add(-time.Minute * 1) // make trigger little-bit earlier than exp time
	return expTime.UTC().Sub(time.Now())
}

func (a *Account) processNewToken(tkn model.Token) {
	if !a.watchTokenRefreshState {
		a.watchTokenRefresh()
	}

	time.Sleep(time.Millisecond * 100)

	go func() {
		a.resetTokenRefreshCh <- tkn.AccessExpiresAt
	}()
}

func (a *Account) setRefreshingState(state bool) {
	a.Auth.SetTokenRefreshingState(state)
}
