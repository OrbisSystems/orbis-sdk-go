package og

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/OrbisSystems/orbis-sdk-go/interfaces/mock"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New(nil, nil))
}

func TestOptionGreeks_CalculateParams(t *testing.T) {
	var (
		expResp = model.CalculateParamsResponse{
			TickerData: model.TickerData{
				Ticker:             "AAPL",
				HistoricVolatility: 10,
			},
			OptionData: model.OptionData{
				Ticker:           "AAPL",
				TheoreticalPrice: 2,
				StrikePrice:      3,
				ExpiryDate: model.Date{
					Year:  2024,
					Month: 1,
					Day:   2,
				},
				Values: model.OptionValues{
					IntrinsicValue: 2,
					OptionValue:    3,
					TimeValue:      4,
				},
				Greeks: model.OptionGreeks{
					Delta: 2,
					Gamma: 3,
					Rho:   4,
					Theta: 1,
					Vega:  6,
				},
			},
		}

		rawReq = `{"ticker":"AAPL","stock_price":0,"strike_value":0,"expiry_date":{"year":0,"month":0,"day":0},"option_type":""}`

		rawResponse = `{"ticker_data":{"ticker":"AAPL","historic_volatility":10},"option_data":{"ticker":"AAPL","theoretical_price":2,"strike_price":3,"expiry_date":{"year":2024,"month":1,"day":2},"values":{"intrinsic_value":2,"option_value":3,"time_value":4},"greeks":{"delta":2,"gamma":3,"rho":4,"theta":1,"vega":6}}}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.CalculateParamsRequest
		hasErr bool
		fn     func(ctx context.Context) *OptionGreeks
	}{
		{
			name:   "success",
			input:  model.CalculateParamsRequest{Ticker: "AAPL"},
			hasErr: false,
			fn: func(ctx context.Context) *OptionGreeks {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightOptionGreeksCalculateParams, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &OptionGreeks{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  model.CalculateParamsRequest{Ticker: "AAPL"},
			hasErr: true,
			fn: func(ctx context.Context) *OptionGreeks {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightOptionGreeksCalculateParams, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &OptionGreeks{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  model.CalculateParamsRequest{Ticker: "AAPL"},
			hasErr: true,
			fn: func(ctx context.Context) *OptionGreeks {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightOptionGreeksCalculateParams, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &OptionGreeks{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			n := tc.fn(ctx)
			resp, err := n.CalculateParams(ctx, tc.input)

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

func TestOptionGreeks_CalculateMatrix(t *testing.T) {
	var (
		expResp = model.CalculateMatrixResponse{
			TickerData: model.TickerData{
				Ticker:             "AAPL",
				HistoricVolatility: 10,
			},
			MatrixData: []model.MatrixData{
				{
					Date: model.Date{
						Year:  2024,
						Month: 1,
						Day:   2,
					},
					OptionData: []model.OptionData{
						{
							Ticker:           "AAPL",
							TheoreticalPrice: 2,
							StrikePrice:      3,
							ExpiryDate: model.Date{
								Year:  2025,
								Month: 2,
								Day:   4,
							},
							Values: model.OptionValues{
								IntrinsicValue: 2,
								OptionValue:    3,
								TimeValue:      6,
							},
							Greeks: model.OptionGreeks{
								Delta: 2,
								Gamma: 3,
								Rho:   4,
								Theta: 5,
								Vega:  56,
							},
						},
					},
				},
			},
		}

		rawReq = `{"ticker":"AAPL","stock_price_min":0,"stock_price_max":0,"strike_value":0,"expiry_date":{"year":0,"month":0,"day":0},"option_type":"","max_intervals":0}`

		rawResponse = `{"ticker_data":{"ticker":"AAPL","historic_volatility":10},"matrix_data":[{"date":{"year":2024,"month":1,"day":2},"option_data":[{"ticker":"AAPL","theoretical_price":2,"strike_price":3,"expiry_date":{"year":2025,"month":2,"day":4},"values":{"intrinsic_value":2,"option_value":3,"time_value":6},"greeks":{"delta":2,"gamma":3,"rho":4,"theta":5,"vega":56}}]}]}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.CalculateMatrixParamsRequest
		hasErr bool
		fn     func(ctx context.Context) *OptionGreeks
	}{
		{
			name:   "success",
			input:  model.CalculateMatrixParamsRequest{Ticker: "AAPL"},
			hasErr: false,
			fn: func(ctx context.Context) *OptionGreeks {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightOptionGreeksCalculateMatrix, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &OptionGreeks{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  model.CalculateMatrixParamsRequest{Ticker: "AAPL"},
			hasErr: true,
			fn: func(ctx context.Context) *OptionGreeks {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightOptionGreeksCalculateMatrix, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &OptionGreeks{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  model.CalculateMatrixParamsRequest{Ticker: "AAPL"},
			hasErr: true,
			fn: func(ctx context.Context) *OptionGreeks {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightOptionGreeksCalculateMatrix, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &OptionGreeks{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			n := tc.fn(ctx)
			resp, err := n.CalculateMatrix(ctx, tc.input)

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
