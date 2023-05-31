package research

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

func TestResearch_GetCompanyProfile(t *testing.T) {
	var (
		expResp = model.CompanyProfile{
			Symbol:              "AAPL",
			CompanyName:         "Apple",
			Country:             "USA",
			Website:             "https://www.apple.com",
			TotalEmployeeNumber: 1000,
		}

		rawResponse = `{"symbol": "AAPL", "company_name": "Apple", "country": "USA", "website": "https://www.apple.com", "employees_number": 1000}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		symbol string
		hasErr bool
		fn     func(ctx context.Context, symbol string) *Research
	}{
		{
			name:   "success",
			symbol: "AAPL",
			hasErr: false,
			fn: func(ctx context.Context, symbol string) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightCompanyProfile, symbol), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			symbol: "AAPL",
			hasErr: true,
			fn: func(ctx context.Context, symbol string) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightCompanyProfile, symbol), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			symbol: "AAPL",
			hasErr: true,
			fn: func(ctx context.Context, symbol string) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightCompanyProfile, symbol), nil).Return(nil, testErr)

				return &Research{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			r := tc.fn(ctx, tc.symbol)
			resp, err := r.GetCompanyProfile(ctx, tc.symbol)

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

func TestResearch_GetCombinedProfile(t *testing.T) {
	var (
		expResp = model.CompanyProfile{
			Symbol:              "AAPL",
			CompanyName:         "Apple",
			Country:             "USA",
			Website:             "https://www.apple.com",
			TotalEmployeeNumber: 1000,
		}

		rawResponse = `{"symbol": "AAPL", "company_name": "Apple", "country": "USA", "website": "https://www.apple.com", "employees_number": 1000}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		symbol string
		hasErr bool
		fn     func(ctx context.Context, symbol string) *Research
	}{
		{
			name:   "success",
			symbol: "AAPL",
			hasErr: false,
			fn: func(ctx context.Context, symbol string) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightQuoteProfile, symbol), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			symbol: "AAPL",
			hasErr: true,
			fn: func(ctx context.Context, symbol string) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightQuoteProfile, symbol), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			symbol: "AAPL",
			hasErr: true,
			fn: func(ctx context.Context, symbol string) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightQuoteProfile, symbol), nil).Return(nil, testErr)

				return &Research{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			r := tc.fn(ctx, tc.symbol)
			resp, err := r.GetCombinedProfile(ctx, tc.symbol)

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

func TestResearch_GetOwnershipsBySymbol(t *testing.T) {
	var (
		expResp = model.GetOwnershipsBySymbolResponse{
			Details: []model.OwnershipDetails{
				{
					Symbol:         "AAPL",
					OwnerType:      "QW",
					OwnerCIK:       12,
					NumberOfShares: 23,
				},
				{
					Symbol:         "FB",
					OwnerType:      "OK",
					OwnerCIK:       32,
					NumberOfShares: 123,
				},
			},
			Count: 2,
		}

		rawReq = `{"symbols":["AAPL","FB"],"owner_type":""}`

		rawResponse = `{"details": [{"symbol": "AAPL", "owner_type": "QW", "owner_cik": 12, "number_of_shares": 23}, {"symbol": "FB", "owner_type": "OK", "owner_cik": 32, "number_of_shares": 123}], "count": 2}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.GetOwnershipsBySymbolRequest
		hasErr bool
		fn     func(ctx context.Context) *Research
	}{
		{
			name: "success",
			input: model.GetOwnershipsBySymbolRequest{
				Symbols: []string{"AAPL", "FB"},
			},
			hasErr: false,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightSymbolOwnerships, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name: "err/unmarshal",
			input: model.GetOwnershipsBySymbolRequest{
				Symbols: []string{"AAPL", "FB"},
			},
			hasErr: true,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightSymbolOwnerships, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name: "err/cli",
			input: model.GetOwnershipsBySymbolRequest{
				Symbols: []string{"AAPL", "FB"},
			},
			hasErr: true,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightSymbolOwnerships, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Research{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			r := tc.fn(ctx)
			resp, err := r.GetOwnershipsBySymbol(ctx, tc.input)

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

func TestResearch_GetOwnershipsByID(t *testing.T) {
	var (
		expResp = model.GetOwnershipsBySymbolResponse{
			Details: []model.OwnershipDetails{
				{
					Symbol:         "AAPL",
					OwnerType:      "QW",
					OwnerCIK:       12,
					NumberOfShares: 23,
				},
				{
					Symbol:         "FB",
					OwnerType:      "OK",
					OwnerCIK:       32,
					NumberOfShares: 123,
				},
			},
			Count: 2,
		}

		rawReq = `{"owner_id":"12"}`

		rawResponse = `{"details": [{"symbol": "AAPL", "owner_type": "QW", "owner_cik": 12, "number_of_shares": 23}, {"symbol": "FB", "owner_type": "OK", "owner_cik": 32, "number_of_shares": 123}], "count": 2}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.GetOwnershipsByIDRequest
		hasErr bool
		fn     func(ctx context.Context) *Research
	}{
		{
			name: "success",
			input: model.GetOwnershipsByIDRequest{
				OwnerID: "12",
			},
			hasErr: false,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightOwnerships, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name: "err/unmarshal",
			input: model.GetOwnershipsByIDRequest{
				OwnerID: "12",
			},
			hasErr: true,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightOwnerships, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name: "err/cli",
			input: model.GetOwnershipsByIDRequest{
				OwnerID: "12",
			},
			hasErr: true,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightOwnerships, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Research{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			r := tc.fn(ctx)
			resp, err := r.GetOwnershipsByID(ctx, tc.input)

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

func TestResearch_GetEarningReleases(t *testing.T) {
	var (
		expResp = model.EarningReleasesResponse{
			Details: []model.EarningRelease{
				{
					Symbol:   "AAPL",
					Quarter:  "3",
					Status:   "1",
					Exchange: "2",
				},
				{
					Symbol:   "FB",
					Quarter:  "2",
					Status:   "3",
					Exchange: "1",
				},
			},
			Count: 2,
		}

		rawReq = `{"symbols":["AAPL","FB"],"exchanges":null,"days_from_now":0}`

		rawResponse = `{"details": [{"symbol": "AAPL", "quarter": "3", "status": "1", "exchange": "2"}, {"symbol": "FB", "quarter": "2",  "status": "3", "exchange": "1"}], "count": 2}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.EarningReleasesRequest
		hasErr bool
		fn     func(ctx context.Context) *Research
	}{
		{
			name: "success",
			input: model.EarningReleasesRequest{
				Symbols: []string{"AAPL", "FB"},
			},
			hasErr: false,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightEarningsCalendar, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name: "err/unmarshal",
			input: model.EarningReleasesRequest{
				Symbols: []string{"AAPL", "FB"},
			},
			hasErr: true,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightEarningsCalendar, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name: "err/cli",
			input: model.EarningReleasesRequest{
				Symbols: []string{"AAPL", "FB"},
			},
			hasErr: true,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightEarningsCalendar, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Research{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			r := tc.fn(ctx)
			resp, err := r.GetEarningReleases(ctx, tc.input)

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

func TestResearch_GetSymbolFundamentals(t *testing.T) {
	var (
		expResp = model.GetSymbolFundamentalsResponse{
			"qwe": "yes",
		}

		rawReq = `{"symbol":"AAPL","mode":""}`

		rawResponse = `{"qwe": "yes"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.GetSymbolFundamentalsRequest
		hasErr bool
		fn     func(ctx context.Context) *Research
	}{
		{
			name: "success",
			input: model.GetSymbolFundamentalsRequest{
				Symbol: "AAPL",
			},
			hasErr: false,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightFundamentals, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name: "err/unmarshal",
			input: model.GetSymbolFundamentalsRequest{
				Symbol: "AAPL",
			},
			hasErr: true,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightFundamentals, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name: "err/cli",
			input: model.GetSymbolFundamentalsRequest{
				Symbol: "AAPL",
			},
			hasErr: true,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightFundamentals, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Research{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			r := tc.fn(ctx)
			resp, err := r.GetSymbolFundamentals(ctx, tc.input)

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

func TestResearch_Screener(t *testing.T) {
	var (
		expResp = model.StockScreenerResponse{
			Details: []model.StockScreenerQuote{
				{
					Symbol:      "AAPL",
					CompanyName: "Apple",
					LastPrice:   134,
					PriceChange: 2,
				},
			},
			Count: 1,
		}

		rawReq = `{"price":null,"market_cap":null,"pe_ratio":null,"dividend_percent":null,"dividend_pay_date":null,"dividend_ex_date":null,"exchange":null,"volume":null,"average_volume":null,"gap":null,"stock_type":null,"industry":null,"quote_type":null,"paging":{"limit":1}}`

		rawResponse = `{"details": [{"symbol": "AAPL", "company_name": "Apple", "last_price": 134, "price_change": 2}], "count": 1}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.StockScreenerRequest
		hasErr bool
		fn     func(ctx context.Context) *Research
	}{
		{
			name: "success",
			input: model.StockScreenerRequest{
				Paging: &model.Paging{Limit: 1},
			},
			hasErr: false,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightStockScreener, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name: "err/unmarshal",
			input: model.StockScreenerRequest{
				Paging: &model.Paging{Limit: 1},
			},
			hasErr: true,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightStockScreener, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name: "err/cli",
			input: model.StockScreenerRequest{
				Paging: &model.Paging{Limit: 1},
			},
			hasErr: true,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightStockScreener, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Research{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			r := tc.fn(ctx)
			resp, err := r.Screener(ctx, tc.input)

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

func TestResearch_StockMarketHeatmap(t *testing.T) {
	var (
		expResp = model.StockMarketHeatmapResponse{
			MinChange: 12,
			MaxChange: 14,
			Details: []model.StockMarketHeatmapEntry{
				{
					Symbol:             "AAPL",
					Title:              "Apple",
					PriceChangePercent: 1,
				},
				{
					Symbol:             "FB",
					Title:              "Facebook",
					PriceChangePercent: 2,
				},
			},
		}

		rawResponse = `{"min_change": 12, "max_change": 14, "details": [{"symbol": "AAPL", "title": "Apple", "price_change_percent": 1}, {"symbol": "FB", "title": "Facebook", "price_change_percent": 2}]}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name                   string
		heatmapName, quoteType string
		hasErr                 bool
		fn                     func(ctx context.Context, heatmapName, quoteType string) *Research
	}{
		{
			name:        "success",
			heatmapName: "qwerty",
			quoteType:   "delayed",
			hasErr:      false,
			fn: func(ctx context.Context, heatmapName, quoteType string) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?heatmapName=%s&quoteType=%s", model.URLInsightBase+model.URLInsightHeatmaps, heatmapName, quoteType), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name:        "err/unmarshal",
			heatmapName: "qwerty",
			quoteType:   "delayed",
			hasErr:      true,
			fn: func(ctx context.Context, heatmapName, quoteType string) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?heatmapName=%s&quoteType=%s", model.URLInsightBase+model.URLInsightHeatmaps, heatmapName, quoteType), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name:        "err/cli",
			heatmapName: "qwerty",
			quoteType:   "delayed",
			hasErr:      true,
			fn: func(ctx context.Context, heatmapName, quoteType string) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?heatmapName=%s&quoteType=%s", model.URLInsightBase+model.URLInsightHeatmaps, heatmapName, quoteType), nil).Return(nil, testErr)

				return &Research{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			r := tc.fn(ctx, tc.heatmapName, tc.quoteType)
			resp, err := r.StockMarketHeatmap(ctx, tc.heatmapName, tc.quoteType)

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

func TestResearch_GetIndustriesPerformance(t *testing.T) {
	var (
		expResp = model.GetIndustriesPerformanceResponse{
			Sectors: []model.Sector{
				{
					SectorCode:                   "C1",
					SectorName:                   "Code",
					SectorMarketCap:              1,
					SectorMarketCap1DayPercent:   2,
					SectorMarketCap1WeekPercent:  3,
					SectorMarketCap2WeekPercent:  4,
					SectorMarketCap1MonthPercent: 5,
				},
			},
		}

		rawReq = `{"ranking_type":"","quote_type":"","quote_filters":{"limit":1,"price":null,"exchanges":"","order":""}}`

		rawResponse = `{"sectors": [{"sector_code":"C1", "sector_name":"Code", "sector_market_cap":1,"sector_market_cap_1day":2, "sector_market_cap_1week":3, "sector_market_cap_2week":4, "sector_market_cap_1month":5}]}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.GetIndustriesPerformanceRequest
		hasErr bool
		fn     func(ctx context.Context) *Research
	}{
		{
			name: "success",
			input: model.GetIndustriesPerformanceRequest{
				QuoteFilters: model.GetIndustriesPerformanceQuoteFilter{
					Limit: 1,
				},
			},
			hasErr: false,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightIndustriesPerformance, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name: "err/unmarshal",
			input: model.GetIndustriesPerformanceRequest{
				QuoteFilters: model.GetIndustriesPerformanceQuoteFilter{
					Limit: 1,
				},
			},
			hasErr: true,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightIndustriesPerformance, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name: "err/cli",
			input: model.GetIndustriesPerformanceRequest{
				QuoteFilters: model.GetIndustriesPerformanceQuoteFilter{
					Limit: 1,
				},
			},
			hasErr: true,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightIndustriesPerformance, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Research{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			r := tc.fn(ctx)
			resp, err := r.GetIndustriesPerformance(ctx, tc.input)

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

func TestResearch_GetMomentumRatioGraph(t *testing.T) {
	var (
		expResp = model.MomentumRatioGraphResponse{
			"AAPL": []model.SymbolMomentum{
				{
					ResultRatio:    1,
					ResultMomentum: 2,
					Date:           "2023-01-01",
				},
			},
			"FB": []model.SymbolMomentum{
				{
					ResultRatio:    2,
					ResultMomentum: 3,
					Date:           "2023-02-02",
				},
			},
		}

		rawReq = `{"symbols":["AAPL","FB"],"benchmark":"","range":"","period":0,"set_ema":false,"mode":""}`

		rawResponse = `{"AAPL": [{"result_ratio": 1, "result_momentum": 2, "date": "2023-01-01"}], "FB": [{"result_ratio": 2, "result_momentum": 3, "date": "2023-02-02"}]}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.MomentumRatioGraphRequest
		hasErr bool
		fn     func(ctx context.Context) *Research
	}{
		{
			name: "success",
			input: model.MomentumRatioGraphRequest{
				Symbols: []string{"AAPL", "FB"},
			},
			hasErr: false,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightGetMomentumRatioGraph, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name: "err/unmarshal",
			input: model.MomentumRatioGraphRequest{
				Symbols: []string{"AAPL", "FB"},
			},
			hasErr: true,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightGetMomentumRatioGraph, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name: "err/cli",
			input: model.MomentumRatioGraphRequest{
				Symbols: []string{"AAPL", "FB"},
			},
			hasErr: true,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightGetMomentumRatioGraph, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Research{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			r := tc.fn(ctx)
			resp, err := r.GetMomentumRatioGraph(ctx, tc.input)

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

func TestResearch_GetSeasonality(t *testing.T) {
	var (
		expResp = model.SeasonalityResponse{
			"AAPL": model.SeasonalityEntry{
				Score: 13,
			},
			"FB": model.SeasonalityEntry{
				Score: 34,
			},
		}

		rawReq = `{"lookahead":0,"benchmark":"","symbols":["AAPL","FB"],"years":0}`

		rawResponse = `{"AAPL": {"score": 13}, "FB": {"score": 34}}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.SeasonalityRequest
		hasErr bool
		fn     func(ctx context.Context) *Research
	}{
		{
			name: "success",
			input: model.SeasonalityRequest{
				Symbols: []string{"AAPL", "FB"},
			},
			hasErr: false,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightSeasonality, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name: "err/unmarshal",
			input: model.SeasonalityRequest{
				Symbols: []string{"AAPL", "FB"},
			},
			hasErr: true,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightSeasonality, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Research{
					cli: cli,
				}
			},
		},
		{
			name: "err/cli",
			input: model.SeasonalityRequest{
				Symbols: []string{"AAPL", "FB"},
			},
			hasErr: true,
			fn: func(ctx context.Context) *Research {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightSeasonality, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Research{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			r := tc.fn(ctx)
			resp, err := r.GetSeasonality(ctx, tc.input)

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
