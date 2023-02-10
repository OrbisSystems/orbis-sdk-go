package logos

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/pkg/errors"

	sdk "github.com/OrbisSystems/orbis-sdk-go"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

// Logos service provides API for getting different information about symbol's logo etc.
type Logos struct {
	sdk.Auth

	cfg       config.Config
	cli       sdk.HTTPClient
	validator sdk.Validator
}

func New(cfg config.Config, auth sdk.Auth, cli sdk.HTTPClient, validator sdk.Validator) *Logos {
	return &Logos{
		Auth:      auth,
		cfg:       cfg,
		cli:       cli,
		validator: validator,
	}
}

// TODO add validator

// SymbolLogos returns logos info for symbol.
func (l *Logos) SymbolLogos(ctx context.Context, symbol string) (model.SymbolLogosResponse, error) {
	r, err := l.cli.Get(ctx, fmt.Sprintf("%s?symbol=%s", l.cfg.AuthHost+model.URLInsightBase+model.URLInsightLogosSymbolLogos, symbol), nil)
	if err != nil {
		return model.SymbolLogosResponse{}, errors.Wrap(err, "couldn't get symbol logos")
	}

	var resp model.SymbolLogosResponse
	err = l.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.SymbolLogosResponse{}, err
	}

	return resp, err
}

// SocialSymbolLogos returns urls of social network logos for symbol.
func (l *Logos) SocialSymbolLogos(ctx context.Context, symbol string) (model.SymbolSocialsResponse, error) {
	r, err := l.cli.Get(ctx, fmt.Sprintf("%s?symbol=%s", l.cfg.AuthHost+model.URLInsightBase+model.URLInsightLogosSocialSymbolLogos, symbol), nil)
	if err != nil {
		return model.SymbolSocialsResponse{}, errors.Wrap(err, "couldn't get social symbol logos")
	}

	var resp model.SymbolSocialsResponse
	err = l.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.SymbolSocialsResponse{}, err
	}

	return resp, err
}

// DirectSymbolLogo returns symbol logo as file.
func (l *Logos) DirectSymbolLogo(ctx context.Context, symbol string) (io.ReadCloser, error) {
	r, err := l.cli.Get(ctx, fmt.Sprintf("%s?symbol=%s", l.cfg.AuthHost+model.URLInsightBase+model.URLInsightLogosDirectSymbolLogos, symbol), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get direct symbol logos")
	}

	return l.cli.GetBodyAndCheckOK(r)
}

// CryptoSymbolLogo returns crypto logos info for symbol
func (l *Logos) CryptoSymbolLogo(ctx context.Context, symbol string) (model.SymbolLogosResponse, error) {
	r, err := l.cli.Get(ctx, fmt.Sprintf("%s?symbol=%s", l.cfg.AuthHost+model.URLInsightBase+model.URLInsightLogosCryptoSymbolLogo, symbol), nil)
	if err != nil {
		return model.SymbolLogosResponse{}, errors.Wrap(err, "couldn't get crypto symbol logos")
	}

	var resp model.SymbolLogosResponse
	err = l.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.SymbolLogosResponse{}, err
	}

	return resp, err
}

// DirectCryptoSymbolLogo returns crypto symbol logo as file.
func (l *Logos) DirectCryptoSymbolLogo(ctx context.Context, symbol string) (io.ReadCloser, error) {
	r, err := l.cli.Get(ctx, fmt.Sprintf("%s?symbol=%s", l.cfg.AuthHost+model.URLInsightBase+model.URLInsightLogosDirectCryptoSymbolLogos, symbol), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get direct crypto symbol logos")
	}

	return l.cli.GetBodyAndCheckOK(r)
}

// MultiSymbolLogos returns logos details for many symbols.
func (l *Logos) MultiSymbolLogos(ctx context.Context, req model.MultipleSymbolLogosRequest) ([]model.SymbolLogosResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := l.cli.Post(ctx, l.cfg.AuthHost+model.URLInsightBase+model.URLInsightLogosMultiSymbolLogos, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get multi symbol logos")
	}

	var resp []model.SymbolLogosResponse
	err = l.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// ConvertedSymbolLogo returns converted logo as file.
func (l *Logos) ConvertedSymbolLogo(ctx context.Context, req model.SymbolLogoConvertedRequest) (io.ReadCloser, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := l.cli.Post(ctx, l.cfg.AuthHost+model.URLInsightBase+model.URLInsightLogosConvertedSymbolLogos, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get converted symbol logo")
	}

	return l.cli.GetBodyAndCheckOK(r)
}

// MultipleCryptoSymbolLogo returns logos details for many crypto symbols.
func (l *Logos) MultipleCryptoSymbolLogo(ctx context.Context, req model.MultipleCryptoLogosRequest) ([]model.SymbolLogosResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := l.cli.Post(ctx, l.cfg.AuthHost+model.URLInsightBase+model.URLInsightLogosMultipleCryptoSymbolLogo, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get multi crypto symbol logo")
	}

	var resp []model.SymbolLogosResponse
	err = l.cli.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// ConvertedCryptoSymbolLogo returns converted crypto logo as file.
func (l *Logos) ConvertedCryptoSymbolLogo(ctx context.Context, req model.SymbolLogoConvertedRequest) (io.ReadCloser, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := l.cli.Post(ctx, l.cfg.AuthHost+model.URLInsightBase+model.URLInsightLogosConvertedCryptoSymbolLogo, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get converted crypto symbol logo")
	}

	return l.cli.GetBodyAndCheckOK(r)
}