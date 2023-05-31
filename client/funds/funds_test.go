package funds

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

	"github.com/OrbisSystems/orbis-sdk-go/interfaces/mock"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New(nil))
}

func TestFunds_GetFundDetails(t *testing.T) {
	var (
		expResp = model.GetFundDetailsResponse{
			Symbol:   "AAPL",
			Leverage: 1.2,
			City:     "London",
			Fax:      "aaa.asdas.de",
			Country:  "UK",
		}

		rawResponse = `{"symbol":"AAPL","leverage":1.2,"city":"London","fax":"aaa.asdas.de","country":"UK"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  string
		hasErr bool
		fn     func(ctx context.Context, symbol string) *Funds
	}{
		{
			name:   "success",
			input:  "AAPL",
			hasErr: false,
			fn: func(ctx context.Context, symbol string) *Funds {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightFundsDetails, symbol), nil).Return(httpResponse, nil)

				return &Funds{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  "AAPL",
			hasErr: true,
			fn: func(ctx context.Context, symbol string) *Funds {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightFundsDetails, symbol), nil).Return(httpResponse, nil)

				return &Funds{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			input:  "AAPL",
			hasErr: true,
			fn: func(ctx context.Context, symbol string) *Funds {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightFundsDetails, symbol), nil).Return(nil, testErr)

				return &Funds{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			f := tc.fn(ctx, tc.input)
			resp, err := f.GetFundDetails(ctx, tc.input)

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

func TestFunds_GetFundScreenerFilters(t *testing.T) {
	var (
		expResp = model.GetFundScreenerFiltersResponse{
			Assets: []model.FundFilterAsset{
				{
					AssetName: "AAC",
				},
			},
			Inverse:  []string{"Q", "C"},
			Leverage: []string{"L1"},
		}

		rawResponse = `{"assets":[{"asset_name": "AAC"}], "inverse": ["Q", "C"], "leverage": ["L1"]}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		hasErr bool
		fn     func(ctx context.Context) *Funds
	}{
		{
			name:   "success",
			hasErr: false,
			fn: func(ctx context.Context) *Funds {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, model.URLInsightBase+model.URLInsightFundsScreenerFilters, nil).Return(httpResponse, nil)

				return &Funds{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			hasErr: true,
			fn: func(ctx context.Context) *Funds {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, model.URLInsightBase+model.URLInsightFundsScreenerFilters, nil).Return(httpResponse, nil)

				return &Funds{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			hasErr: true,
			fn: func(ctx context.Context) *Funds {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, model.URLInsightBase+model.URLInsightFundsScreenerFilters, nil).Return(nil, testErr)

				return &Funds{
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
			resp, err := f.GetFundScreenerFilters(ctx)

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

func TestFunds_ScreenFunds(t *testing.T) {
	var (
		expResp = model.FundScreenerResponse{
			Details: []model.FundScreenerQuote{
				{
					Leverage: 1.12,
					Symbol:   "AAPL",
					Name:     "Apple",
					Category: "Tech",
				},
			},
			Count: 1,
		}
		rawReq      = `{"assets":null,"categories":null,"sponsors":null,"inverse":"","leverage":"1.12","fund_type":"","quote_type":""}`
		rawResponse = `{"details":[{"leverage": 1.12, "symbol": "AAPL", "name": "Apple", "category": "Tech"}], "count": 1}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		hasErr bool
		input  model.FundScreenerRequest
		fn     func(ctx context.Context) *Funds
	}{
		{
			name:   "success",
			hasErr: false,
			input: model.FundScreenerRequest{
				Leverage: "1.12",
			},
			fn: func(ctx context.Context) *Funds {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightFundsScreener, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Funds{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			hasErr: true,
			input: model.FundScreenerRequest{
				Leverage: "1.12",
			},
			fn: func(ctx context.Context) *Funds {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				r := io.NopCloser(strings.NewReader("x"))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightFundsScreener, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Funds{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			hasErr: true,
			input: model.FundScreenerRequest{
				Leverage: "1.12",
			},
			fn: func(ctx context.Context) *Funds {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightFundsScreener, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Funds{
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
			resp, err := f.ScreenFunds(ctx, tc.input)

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

func TestFunds_ScreenSectorFunds(t *testing.T) {
	var (
		expResp = model.FundScreenerResponse{
			Details: []model.FundScreenerQuote{
				{
					Leverage: 1.12,
					Symbol:   "AAPL",
					Name:     "Apple",
					Category: "Tech",
				},
			},
			Count: 1,
		}
		rawReq      = `{"sector_code":"","exposure":null,"dividend_percent":null,"sponsors":null,"inverse":"","leverage":"1.12","holdings":null,"quote_type":""}`
		rawResponse = `{"details":[{"leverage": 1.12, "symbol": "AAPL", "name": "Apple", "category": "Tech"}], "count": 1}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		hasErr bool
		input  model.FundSectorScreenerRequest
		fn     func(ctx context.Context) *Funds
	}{
		{
			name:   "success",
			hasErr: false,
			input: model.FundSectorScreenerRequest{
				Leverage: "1.12",
			},
			fn: func(ctx context.Context) *Funds {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightFundsSectorScreener, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Funds{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			hasErr: true,
			input: model.FundSectorScreenerRequest{
				Leverage: "1.12",
			},
			fn: func(ctx context.Context) *Funds {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				r := io.NopCloser(strings.NewReader("x"))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightFundsSectorScreener, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Funds{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			hasErr: true,
			input: model.FundSectorScreenerRequest{
				Leverage: "1.12",
			},
			fn: func(ctx context.Context) *Funds {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightFundsSectorScreener, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Funds{
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
			resp, err := f.ScreenSectorFunds(ctx, tc.input)

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

func TestFunds_GetTopFunds(t *testing.T) {
	var (
		expResp     = []string{"A", "B"}
		rawReq      = `{"inverse":"","leverage":"1.12","fund_type":"","quote_type":"","reverse_order":false}`
		rawResponse = `["A", "B"]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		hasErr bool
		input  model.GetTopFundsRequest
		fn     func(ctx context.Context) *Funds
	}{
		{
			name:   "success",
			hasErr: false,
			input: model.GetTopFundsRequest{
				Leverage: "1.12",
			},
			fn: func(ctx context.Context) *Funds {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightFundsTop, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Funds{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			hasErr: true,
			input: model.GetTopFundsRequest{
				Leverage: "1.12",
			},
			fn: func(ctx context.Context) *Funds {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				r := io.NopCloser(strings.NewReader("x"))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightFundsTop, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Funds{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			hasErr: true,
			input: model.GetTopFundsRequest{
				Leverage: "1.12",
			},
			fn: func(ctx context.Context) *Funds {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightFundsTop, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Funds{
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
			resp, err := f.GetTopFunds(ctx, tc.input)

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
