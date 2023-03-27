package dates

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/OrbisSystems/orbis-sdk-go/interface/mock"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New("", nil))
}

func TestMarketDates_GetMarketDatesHistory(t *testing.T) {
	var (
		expResp = model.GetMarketDatesResponse{
			Data: []model.ExtendedMarketDate{
				{
					MarketCode:    "1",
					MarketName:    "asd",
					CountryCode:   "ce1",
					CurrencyCode:  "usd",
					PrimaryMarket: true,
				},
			},
			Count: 10,
		}
		rawReq      = `{"markets":null,"start_date":"0001-01-01T00:00:00Z","end_date":"0001-01-01T00:00:00Z","paging":{"limit":10}}`
		rawResponse = `{"data":[{"market_code":"1","market_name":"asd","country_code":"ce1","currency_code":"usd","primary_market":true,"open_date":"","open_time":"","close_time":"","timezone":""}],"count":10}`
		url         = "http://localhost"
		testErr     = errors.New("process error")
	)
	testCases := []struct {
		name   string
		input  model.GetMarketDatesRequest
		hasErr bool
		fn     func(ctx context.Context, req model.GetMarketDatesRequest) *MarketDates
	}{
		{
			name: "success",
			input: model.GetMarketDatesRequest{
				Paging: &model.Paging{
					Limit: 10,
				},
			},
			hasErr: false,
			fn: func(ctx context.Context, req model.GetMarketDatesRequest) *MarketDates {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				r := io.NopCloser(strings.NewReader(rawResponse))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, url+model.URLInsightBase+model.URLInsightMarketDatesHistory, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &MarketDates{
					url: url,
					cli: cli,
				}
			},
		},
		{
			name: "err/unmarshal",
			input: model.GetMarketDatesRequest{
				Paging: &model.Paging{
					Limit: 10,
				},
			},
			hasErr: true,
			fn: func(ctx context.Context, req model.GetMarketDatesRequest) *MarketDates {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				r := io.NopCloser(strings.NewReader("x"))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, url+model.URLInsightBase+model.URLInsightMarketDatesHistory, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &MarketDates{
					url: url,
					cli: cli,
				}
			},
		},
		{
			name: "err/cli",
			input: model.GetMarketDatesRequest{
				Paging: &model.Paging{
					Limit: 10,
				},
			},
			hasErr: true,
			fn: func(ctx context.Context, req model.GetMarketDatesRequest) *MarketDates {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, url+model.URLInsightBase+model.URLInsightMarketDatesHistory, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &MarketDates{
					url: url,
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			a := tc.fn(ctx, tc.input)
			resp, err := a.GetMarketDatesHistory(ctx, tc.input)

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

func TestMarketDates_GetTodayMarketHours(t *testing.T) {
	var (
		expResp = model.GetMarketHoursResponse{
			OpenRightNow: true,
			OpenToday:    true,
			OpenTime:     "10:00:00",
			CloseTime:    "20:00:00",
			TimeZone:     "UTC",
		}
		rawResponse = `{"open_right_now":true,"open_today":true,"open_time":"10:00:00","close_time":"20:00:00","time_zone":"UTC"}`
		url         = "http://localhost"
		testErr     = errors.New("process error")
	)
	testCases := []struct {
		name   string
		input  string
		hasErr bool
		fn     func(ctx context.Context, market string) *MarketDates
	}{
		{
			name:   "success",
			input:  "market",
			hasErr: false,
			fn: func(ctx context.Context, market string) *MarketDates {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?market=%s", url+model.URLInsightBase+model.URLInsightMarketDatesToday, market), nil).Return(httpResponse, nil)

				return &MarketDates{
					url: url,
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  "market",
			hasErr: true,
			fn: func(ctx context.Context, market string) *MarketDates {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?market=%s", url+model.URLInsightBase+model.URLInsightMarketDatesToday, market), nil).Return(httpResponse, nil)

				return &MarketDates{
					url: url,
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			input:  "market",
			hasErr: true,
			fn: func(ctx context.Context, market string) *MarketDates {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?market=%s", url+model.URLInsightBase+model.URLInsightMarketDatesToday, market), nil).Return(nil, testErr)

				return &MarketDates{
					url: url,
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			a := tc.fn(ctx, tc.input)
			resp, err := a.GetTodayMarketHours(ctx, tc.input)

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
