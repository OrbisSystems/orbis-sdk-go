package account

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"

	sdk "github.com/OrbisSystems/orbis-sdk-go/interfaces"
	"github.com/OrbisSystems/orbis-sdk-go/model"
	"github.com/OrbisSystems/orbis-sdk-go/utils"
)

var (
	ErrEmptyRefreshToken = errors.New("empty refresh token from storage")
)

// Account service provides all API for account managing.
type Account struct {
	sdk.Auth

	cli    sdk.HTTPClient
	logger *log.Logger

	loginCallback func(ctx context.Context) error
}

func New(auth sdk.Auth, cli sdk.HTTPClient, logger *log.Logger, disableTokenRefresh bool) *Account {
	a := &Account{
		Auth:   auth,
		cli:    cli,
		logger: logger,
	}

	if !disableTokenRefresh {
		a.watchTokenRefresh()
	}

	return a
}

func (a *Account) watchTokenRefresh() {
	go func() {
		var (
			// pick random offset to avoid 425 response code - Duplicate request when using multiple instances
			// it is still possible to encounter 425 code provided one runs a large number of SDK instances using the same credentials
			refreshTicker = time.NewTicker(5*time.Second + randomOffsetInMilliseconds())
			lastRefresh   = time.Now()
		)

		for range refreshTicker.C {
			// for the case where the time to refresh takes longer than the ticker duration and new refresh starts immediately
			if time.Since(lastRefresh) < time.Second {
				continue
			}

			if err := withRetries(a.RefreshToken); err != nil {
				a.logger.Errorf("failed to update refresh token 3 times: %v", err)

				if a.loginCallback == nil {
					a.logger.Error("unable to complete authentication: no login callback provided")

					lastRefresh = time.Now()
					refreshTicker.Reset(5 * time.Minute)

					continue
				}

				if err = withRetries(a.loginCallback); err != nil {
					a.logger.Errorf("failed to complete authentication: %v", err)

					lastRefresh = time.Now()
					refreshTicker.Reset(5 * time.Minute)

					continue
				}
			}

			lastRefresh = time.Now()
			refreshTicker.Reset(a.getRefreshDurationFromToken())
		}
	}()
}

// SetLoginCallback accepts a function that will be called in case token refresh is failed multiple times
func (a *Account) SetLoginCallback(callback func(ctx context.Context) error) {
	a.loginCallback = callback
}

func (a *Account) getRefreshDurationFromToken() time.Duration {
	var (
		ctx, cancel            = context.WithTimeout(context.Background(), time.Minute)
		defaultRefreshDuration = 5*time.Second + randomOffsetInMilliseconds()
	)

	defer cancel()

	tkn, err := a.GetToken(ctx)
	if err != nil {
		a.logger.Errorf("getRefreshDurationFromToken -> error while getting token from auth: %v", err)

		return defaultRefreshDuration
	}

	if tkn.RefreshToken == "" {
		a.logger.Errorf("getRefreshDurationFromToken -> %v", ErrEmptyRefreshToken)

		return defaultRefreshDuration
	}

	var (
		expTime         = time.Unix(tkn.AccessExpiresAt, 0).Add(-time.Hour * 1) // make trigger little-bit earlier than exp time
		refreshDuration = time.Until(expTime.UTC())
	)
	if refreshDuration.Milliseconds() <= 0 {
		refreshDuration = defaultRefreshDuration
	}

	return refreshDuration
}

// NeedToLogin tells you do you need to call login API or there is still actual token you can use.
// True - need to re-login.
func (a *Account) NeedToLogin(ctx context.Context) (bool, error) {
	tkn, err := a.Auth.GetToken(ctx)
	if err != nil && !errors.Is(err, redis.Nil) {
		return true, err
	}

	if tkn.RefreshToken == "" {
		return true, nil
	}

	return tkn.RefreshExpiresAt <= time.Now().Unix(), nil
}

// LoginByEmail uses for getting token if you are a client with the biggest access in you company.
// After log in, you can create api keys with limited access for you users using CreateAPIKey method.
// With these api keys your users can log in via SDK using LoginByAPIKey method.
func (a *Account) LoginByEmail(ctx context.Context, req model.LoginByEmailRequest) error {
	a.logger.Trace("LoginByEmail called")

	if req.DeviceID == "" {
		req.DeviceID = uuid.New().String()
	}

	body, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := a.cli.Post(ctx, model.URLB2BLoginByEmail, bytes.NewBuffer(body), nil)
	if err != nil {
		return errors.Wrap(err, "account couldn't login by email")
	}

	var resp model.LoginByEmailResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return err
	}

	return a.SetToken(ctx, resp.LoginBasic.Tokens)
}

// LoginByAPIKey allows lo log in into the system for users with limited access.
// See LoginByEmail and CreateAPIKey for more information.
func (a *Account) LoginByAPIKey(ctx context.Context, req model.LoginByAPIKeyRequest) error {
	a.logger.Trace("LoginByAPIKey called")

	body, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := a.cli.Post(ctx, model.URLB2BLoginByAPIKey, bytes.NewBuffer(body), nil)
	if err != nil {
		return errors.Wrap(err, "account couldn't login by api key")
	}

	var resp model.LoginByAPIKeyResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return err
	}

	return a.SetToken(ctx, resp.ApiKeysLogin.Tokens)
}

// CreateAPIKey provides creating api keys with limited access to branches, microservices, specific roles and permissions.
func (a *Account) CreateAPIKey(ctx context.Context, req model.CreateAPIKeyRequest) (model.CreateAPIKeyResponse, error) {
	a.logger.Trace("CreateAPIKey called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.CreateAPIKeyResponse{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := a.cli.Post(ctx, model.URLB2BCreateAPIKey, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.CreateAPIKeyResponse{}, errors.Wrap(err, "account couldn't login")
	}

	var resp model.CreateAPIKeyResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)

	return resp, err
}

func (a *Account) RefreshToken(ctx context.Context) error {
	a.logger.Trace("RefreshToken called")

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

	r, err := a.cli.Post(ctx, model.URLB2BRefreshToken, bytes.NewBuffer(body), nil)
	if err != nil {
		return errors.Wrap(err, "account couldn't refresh token")
	}

	var resp model.LoginByEmailResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)

	if err != nil {
		return err
	}

	return a.SetToken(ctx, resp.LoginBasic.Tokens)
}

func (a *Account) GetUserByID(ctx context.Context, id int) (model.GetB2BUserByIDResponse, error) {
	a.logger.Trace("GetUserByID called")

	r, err := a.cli.Get(ctx, fmt.Sprintf("%s/%d", model.URLB2BGetUserByID, id), nil)
	if err != nil {
		return model.GetB2BUserByIDResponse{}, errors.Wrap(err, "can't get user by id")
	}

	var resp model.GetB2BUserByIDResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)

	return resp, err
}

func withRetries(f func(ctx context.Context) error) error {
	var err error
	for retryNumber := 0; retryNumber < 3; retryNumber++ {
		time.Sleep(exponentialBackoffDuration(retryNumber))

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		err = f(ctx)
		if err == nil {
			return nil
		}
	}

	return err
}

func exponentialBackoffDuration(retryNumber int) time.Duration {
	switch retryNumber {
	case 1:
		return 5*time.Second + randomOffsetInMilliseconds()
	case 2:
		return 10*time.Second + randomOffsetInMilliseconds()
	case 3:
		return 20*time.Second + randomOffsetInMilliseconds()
	default:
		return time.Duration(0)
	}
}

// randomOffsetInMilliseconds() returns random duration in range [100ms:2000ms] with 100ms step
// nolint:gosec // No need to use a strong random generator for a random time offset
func randomOffsetInMilliseconds() time.Duration {
	return time.Duration(((rand.Intn(19) * 100) + 100)) * time.Millisecond
}
