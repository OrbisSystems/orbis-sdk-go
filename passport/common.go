package passport

import (
	"github.com/OrbisSystems/orbis-sdk-go/models/passport"
	"github.com/pkg/errors"
)

var (
	errAPIKeyIsEmpty    = errors.New("api key must be not empty")
	errAPISecretIsEmpty = errors.New("api secret must be not empty")
)

func checkSecret(secret passport.Secret) error {
	if secret.APISecret == "" {
		return errAPISecretIsEmpty
	}
	if secret.APIKey == "" {
		return errAPIKeyIsEmpty
	}

	return nil
}
