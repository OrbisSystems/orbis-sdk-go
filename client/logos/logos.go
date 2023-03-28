package logos

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	sdk "github.com/OrbisSystems/orbis-sdk-go/interfaces"
	"github.com/OrbisSystems/orbis-sdk-go/model"
	"github.com/OrbisSystems/orbis-sdk-go/utils"
)

// Logos service provides API for getting different information about symbol's logo etc.
type Logos struct {
	cli sdk.HTTPClient
}

func New(cli sdk.HTTPClient) *Logos {
	return &Logos{
		cli: cli,
	}
}

// SymbolLogos returns logos info for symbol.
func (l *Logos) SymbolLogos(ctx context.Context, symbol string) (model.SymbolLogosResponse, error) {
	log.Trace("SymbolLogos called")

	r, err := l.cli.Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightLogosSymbolLogos, symbol), nil)
	if err != nil {
		return model.SymbolLogosResponse{}, errors.Wrap(err, "couldn't get symbol logos")
	}

	var resp model.SymbolLogosResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.SymbolLogosResponse{}, err
	}

	return resp, err
}

// SocialSymbolLogos returns urls of social network logos for symbol.
func (l *Logos) SocialSymbolLogos(ctx context.Context, symbol string) (model.SymbolSocialsResponse, error) {
	log.Trace("SocialSymbolLogos called")

	r, err := l.cli.Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightLogosSocialSymbolLogos, symbol), nil)
	if err != nil {
		return model.SymbolSocialsResponse{}, errors.Wrap(err, "couldn't get social symbol logos")
	}

	var resp model.SymbolSocialsResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.SymbolSocialsResponse{}, err
	}

	return resp, err
}

// DirectSymbolLogo returns symbol logo as file.
func (l *Logos) DirectSymbolLogo(ctx context.Context, symbol string) (io.ReadCloser, error) {
	log.Trace("DirectSymbolLogo called")

	r, err := l.cli.Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightLogosDirectSymbolLogos, symbol), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get direct symbol logos")
	}

	return utils.GetBodyAndCheckOK(r)
}

// CryptoSymbolLogo returns crypto logos info for symbol
func (l *Logos) CryptoSymbolLogo(ctx context.Context, symbol string) (model.SymbolLogosResponse, error) {
	log.Trace("CryptoSymbolLogo called")

	r, err := l.cli.Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightLogosCryptoSymbolLogo, symbol), nil)
	if err != nil {
		return model.SymbolLogosResponse{}, errors.Wrap(err, "couldn't get crypto symbol logos")
	}

	var resp model.SymbolLogosResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return model.SymbolLogosResponse{}, err
	}

	return resp, err
}

// DirectCryptoSymbolLogo returns crypto symbol logo as file.
func (l *Logos) DirectCryptoSymbolLogo(ctx context.Context, symbol string) (io.ReadCloser, error) {
	log.Trace("DirectCryptoSymbolLogo called")

	r, err := l.cli.Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightLogosDirectCryptoSymbolLogos, symbol), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get direct crypto symbol logos")
	}

	return utils.GetBodyAndCheckOK(r)
}

// MultiSymbolLogos returns logos details for many symbols.
func (l *Logos) MultiSymbolLogos(ctx context.Context, req model.MultipleSymbolLogosRequest) ([]model.SymbolLogosResponse, error) {
	log.Trace("MultiSymbolLogos called")

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := l.cli.Post(ctx, model.URLInsightBase+model.URLInsightLogosMultiSymbolLogos, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get multi symbol logos")
	}

	var resp []model.SymbolLogosResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// ConvertedSymbolLogo returns converted logo as file.
func (l *Logos) ConvertedSymbolLogo(ctx context.Context, req model.SymbolLogoConvertedRequest) (io.ReadCloser, error) {
	log.Trace("ConvertedSymbolLogo called")

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := l.cli.Post(ctx, model.URLInsightBase+model.URLInsightLogosConvertedSymbolLogos, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get converted symbol logo")
	}

	return utils.GetBodyAndCheckOK(r)
}

// MultipleCryptoSymbolLogo returns logos details for many crypto symbols.
func (l *Logos) MultipleCryptoSymbolLogo(ctx context.Context, req model.MultipleCryptoLogosRequest) ([]model.SymbolLogosResponse, error) {
	log.Trace("MultipleCryptoSymbolLogo called")

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := l.cli.Post(ctx, model.URLInsightBase+model.URLInsightLogosMultipleCryptoSymbolLogo, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get multi crypto symbol logo")
	}

	var resp []model.SymbolLogosResponse
	err = utils.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// ConvertedCryptoSymbolLogo returns converted crypto logo as file.
func (l *Logos) ConvertedCryptoSymbolLogo(ctx context.Context, req model.SymbolLogoConvertedRequest) (io.ReadCloser, error) {
	log.Trace("ConvertedCryptoSymbolLogo called")

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := l.cli.Post(ctx, model.URLInsightBase+model.URLInsightLogosConvertedCryptoSymbolLogo, bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get converted crypto symbol logo")
	}

	return utils.GetBodyAndCheckOK(r)
}
