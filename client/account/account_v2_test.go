package account

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/OrbisSystems/orbis-sdk-go/interfaces/mock"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

func TestAccount_GetB2BUsersV2(t *testing.T) {
	var (
		expResp = model.GetB2BUsersV2Response{}

		req = model.GetB2BUsersV2Request{
			Filters: model.GetB2BUsersV2Filters{
				EmailVerified: true,
				Ids:           []int{1},
			},
			Pagination: model.GetB2BUsersV2Pagination{
				Page: 0,
				Size: 100,
			},
		}

		rawResp = `{"current_page":1,"pages_count":0,"status":1,"total_records":0}`
		rawReq  = `{"include_chiefadmin":false,"filters":{"email_verified":true,"ids":[1],"neither_verified":false,"phone_verified":false},"pagination":{"page":0,"size":100}}`

		testErr = errors.New("process error")
	)

	expResp.Status = 1
	expResp.CurrentPage = 1
	expResp.PagesCount = 0

	testCases := []struct {
		name   string
		hasErr bool
		input  model.GetB2BUsersV2Request
		fn     func(ctx context.Context, req model.GetB2BUsersV2Request) *Account
	}{
		{
			name:   "success",
			hasErr: false,
			input:  req,
			fn: func(ctx context.Context, req model.GetB2BUsersV2Request) *Account {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResp))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLB2BGetUsersV2, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Account{
					cli: cli,

					refreshTicker: time.NewTicker(time.Hour * 100),
				}
			},
		},
		{
			name:   "unmarshal",
			hasErr: true,
			input:  req,
			fn: func(ctx context.Context, req model.GetB2BUsersV2Request) *Account {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("test"))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}
				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLB2BGetUsersV2, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Account{
					cli: cli,

					refreshTicker: time.NewTicker(time.Hour * 100),
				}
			},
		},
		{
			name:   "cli",
			hasErr: true,
			input:  req,
			fn: func(ctx context.Context, req model.GetB2BUsersV2Request) *Account {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResp))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}
				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLB2BGetUsersV2, bytes.NewBuffer(bb), nil).Return(httpResponse, testErr)

				return &Account{
					cli: cli,

					refreshTicker: time.NewTicker(time.Hour * 100),
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			a := tc.fn(ctx, tc.input)
			resp, err := a.GetB2BUsersV2(ctx, tc.input)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, expResp, resp)
			}
		})
	}
}
