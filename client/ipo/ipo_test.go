package ipo

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/OrbisSystems/orbis-sdk-go/interfaces/mock"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New(nil))
}

func TestIPO_GetUpcomingIPOs(t *testing.T) {
	var (
		expResp = model.IPOResponse{
			IPOList: []model.IPO{
				{
					Symbol:      "AAPL",
					CompanyName: "Apple",
				},
			},
			TotalRecords: 100,
		}

		rawResponse = `{"ipo_list":[{"symbol":"AAPL","company_name":"Apple","prospectus_url":"","file_date":"","ipo_date":"","ipo_date_approx":false,"offer_price":0,"offer_price_max":0,"offer_price_min":0,"offer_shares":0,"first_day_open":0,"first_day_close":0,"ipo_return":0,"status":"","underwriters":null,"ceo":"","total_expenses":0,"shares_over_alloted":0,"shares_outstanding":0,"lockup_period":0,"lockup_expiration":"","quiet_period_expiration":"","offer_amount":0,"cik":0}],"total_records":100}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name          string
		limit, offset int
		hasErr        bool
		fn            func(ctx context.Context, limit, offset int) *IPO
	}{
		{
			name:   "success",
			limit:  1,
			offset: 1,
			hasErr: false,
			fn: func(ctx context.Context, limit, offset int) *IPO {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?limit=%d&offset=%d", model.URLInsightBase+model.URLInsightIPOsUpcoming, limit, offset), nil).Return(httpResponse, nil)

				return &IPO{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			limit:  1,
			offset: 1,
			hasErr: true,
			fn: func(ctx context.Context, limit, offset int) *IPO {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?limit=%d&offset=%d", model.URLInsightBase+model.URLInsightIPOsUpcoming, limit, offset), nil).Return(httpResponse, nil)

				return &IPO{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			limit:  1,
			offset: 1,
			hasErr: true,
			fn: func(ctx context.Context, limit, offset int) *IPO {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?limit=%d&offset=%d", model.URLInsightBase+model.URLInsightIPOsUpcoming, limit, offset), nil).Return(nil, testErr)

				return &IPO{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			f := tc.fn(ctx, tc.limit, tc.offset)
			resp, err := f.GetUpcomingIPOs(ctx, tc.limit, tc.offset)

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

func TestIPO_GetRecentIPOs(t *testing.T) {
	var (
		expResp = model.IPOResponse{
			IPOList: []model.IPO{
				{
					Symbol:      "AAPL",
					CompanyName: "Apple",
				},
			},
			TotalRecords: 100,
		}

		rawReq = `{"load_quotes":false,"time":"10:10","order":null,"paging":{}}`

		rawResponse = `{"ipo_list":[{"symbol":"AAPL","company_name":"Apple","prospectus_url":"","file_date":"","ipo_date":"","ipo_date_approx":false,"offer_price":0,"offer_price_max":0,"offer_price_min":0,"offer_shares":0,"first_day_open":0,"first_day_close":0,"ipo_return":0,"status":"","underwriters":null,"ceo":"","total_expenses":0,"shares_over_alloted":0,"shares_outstanding":0,"lockup_period":0,"lockup_expiration":"","quiet_period_expiration":"","offer_amount":0,"cik":0}],"total_records":100}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.RecentIPORequest
		hasErr bool
		fn     func(ctx context.Context) *IPO
	}{
		{
			name: "success",
			input: model.RecentIPORequest{
				Time: "10:10",
			},
			hasErr: false,
			fn: func(ctx context.Context) *IPO {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightIPOsRecent, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &IPO{
					cli: cli,
				}
			},
		},
		{
			name: "err/unmarshal",
			input: model.RecentIPORequest{
				Time: "10:10",
			},
			hasErr: true,
			fn: func(ctx context.Context) *IPO {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightIPOsRecent, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &IPO{
					cli: cli,
				}
			},
		},
		{
			name: "err/cli",
			input: model.RecentIPORequest{
				Time: "10:10",
			},
			hasErr: true,
			fn: func(ctx context.Context) *IPO {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightIPOsRecent, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &IPO{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			f := tc.fn(ctx)
			resp, err := f.GetRecentIPOs(ctx, tc.input)

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
