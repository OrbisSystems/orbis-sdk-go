package auth

import "errors"

func (a *Auth) GetToken() (string, error) {
	t, ok, err := a.storage.Get(a.key)
	if err != nil && !ok {
		return "", errors.New("token is empty")
	}

	return string(t), nil
}

func (a *Auth) SetToken(token string) {
	err := a.storage.Save(a.key, []byte(token))
	if err != nil {
		return
	}

}
