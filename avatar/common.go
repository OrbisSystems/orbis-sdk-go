package avatar

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	emptyIDError         = errors.New("ID is empty")
	emptyTokenError      = errors.New("token is empty")
	emptyOrbisTokenError = errors.New("orbis token is empty")
	emptyEnvError        = errors.New("env is empty")
)

func getQueryParamsToken(ID, token string) (string, error) {
	if ID == "" {
		return "", emptyIDError
	}

	if token == "" {
		return "", emptyTokenError
	}

	return fmt.Sprintf("token=%s", token), nil
}

func getQueryParamsOrbisToken(ID, orbisToken, env string) (string, error) {
	if ID == "" {
		return "", emptyIDError
	}

	if orbisToken == "" {
		return "", emptyOrbisTokenError
	}

	if env == "" {
		return "", emptyEnvError
	}

	return fmt.Sprintf("orbisTokem=%s&env=%s", orbisToken, env), nil
}
