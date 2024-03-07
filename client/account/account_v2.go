package account

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/OrbisSystems/orbis-sdk-go/model"
	"github.com/OrbisSystems/orbis-sdk-go/utils"
	"github.com/pkg/errors"
)

func (a *Account) GetB2BUsersV2(ctx context.Context, req model.GetB2BUsersV2Request) (model.GetB2BUsersV2Response, error) {
	a.logger.Trace("GetB2BUsersV2 called")

	body, err := json.Marshal(req)
	if err != nil {
		return model.GetB2BUsersV2Response{}, errors.Wrap(err, "couldn't marshal input parameters")
	}

	r, err := a.cli.Post(ctx, model.URLB2BGetUsersV2, bytes.NewBuffer(body), nil)
	if err != nil {
		return model.GetB2BUsersV2Response{}, errors.Wrap(err, "can't get b2b users")
	}

	var resp model.GetB2BUsersV2Response
	err = utils.UnmarshalAndCheckOk(&resp, r)

	return resp, err
}
