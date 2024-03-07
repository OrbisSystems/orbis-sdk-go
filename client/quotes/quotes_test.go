package quotes

import (
	"bytes"
	"context"
	"errors"
	"fmt"
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

func TestQuotes_GetQuotesEquityData(t *testing.T) {
	var (
		expResp = []model.QuoteEquityDataResponse{
			{
				AskPx:  12,
				AskSz:  1,
				BidPx:  13,
				BidSz:  2,
				Column: 100,
			},
		}

		rawResponse = `[{"ask_px": 12, "ask_sz": 1, "bid_px": 13, "bid_sz": 2, "column": 100}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name               string
		symbols, quoteType string
		hasErr             bool
		fn                 func(ctx context.Context, symbols, quoteType string) *Quotes
	}{
		{
			name:      "success",
			symbols:   "AAPL",
			quoteType: "delayed",
			hasErr:    false,
			fn: func(ctx context.Context, symbols, quoteType string) *Quotes {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbols=%s&quote_type=%s", model.URLInsightBase+model.URLInsightQuotesEquity, symbols, quoteType), nil).Return(httpResponse, nil)

				return &Quotes{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:      "err/unmarshal",
			symbols:   "AAPL",
			quoteType: "delayed",
			hasErr:    true,
			fn: func(ctx context.Context, symbols, quoteType string) *Quotes {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbols=%s&quote_type=%s", model.URLInsightBase+model.URLInsightQuotesEquity, symbols, quoteType), nil).Return(httpResponse, nil)

				return &Quotes{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:      "err/cli",
			symbols:   "AAPL",
			quoteType: "delayed",
			hasErr:    true,
			fn: func(ctx context.Context, symbols, quoteType string) *Quotes {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbols=%s&quote_type=%s", model.URLInsightBase+model.URLInsightQuotesEquity, symbols, quoteType), nil).Return(nil, testErr)

				return &Quotes{
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
			q := tc.fn(ctx, tc.symbols, tc.quoteType)
			resp, err := q.GetQuotesEquityData(ctx, tc.symbols, tc.quoteType)

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

func TestQuotes_GetQuoteHistory(t *testing.T) {
	var (
		expResp = model.QuoteHistoryResponse{
			Extremum: model.Extremum{
				MinValue: 10,
				MaxValue: 23,
			},
			Quotes: []*model.Quote{
				{
					Symbol: "AAPL",
					Open:   9,
					Close:  10,
					High:   11,
					Low:    8,
					Volume: 56,
				},
				{
					Symbol: "FB",
					Open:   19,
					Close:  23,
					High:   24,
					Low:    17,
					Volume: 98,
				},
			},
			Count: 2,
		}

		rawReq = `{"filter":{"symbol":""},"quote_access_type":"realtime"}`

		rawResponse = `{"extremum": {"min_value": 10, "max_value": 23}, "quotes": [{"symbol": "AAPL", "open": 9, "close": 10, "high": 11, "low": 8, "volume": 56}, {"symbol": "FB", "open": 19, "close": 23, "high": 24, "low": 17, "volume": 98}], "count": 2}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.QuoteHistoryRequest
		hasErr bool
		fn     func(ctx context.Context) *Quotes
	}{
		{
			name: "success",
			input: model.QuoteHistoryRequest{
				QuoteAccessType: model.QuoteAccessTypeRealtime,
			},
			hasErr: false,
			fn: func(ctx context.Context) *Quotes {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightQuoteHistory, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Quotes{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name: "err/unmarshal",
			input: model.QuoteHistoryRequest{
				QuoteAccessType: model.QuoteAccessTypeRealtime,
			},
			hasErr: true,
			fn: func(ctx context.Context) *Quotes {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightQuoteHistory, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Quotes{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name: "err/cli",
			input: model.QuoteHistoryRequest{
				QuoteAccessType: model.QuoteAccessTypeRealtime,
			},
			hasErr: true,
			fn: func(ctx context.Context) *Quotes {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightQuoteHistory, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Quotes{
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
			q := tc.fn(ctx)
			resp, err := q.GetQuoteHistory(ctx, tc.input)

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

func TestQuotes_GetIntradayQuotes(t *testing.T) {
	var (
		expResp = []model.IntradayResponse{
			{
				Price:  12,
				Volume: 100,
				High:   13,
				Low:    11,
				Rsi:    2,
			},
			{
				Price:  13,
				Volume: 34,
				High:   15,
				Low:    12,
				Rsi:    1,
			},
		}

		rawReq = `{"symbol":"AAPL","rsi_span":0,"from":null,"to":null,"raw_volumes":false,"smoothing":null}`

		rawResponse = `[{"price": 12, "volume": 100, "day_high": 13, "day_low": 11, "rsi":2}, {"price": 13, "volume": 34, "day_high": 15, "day_low": 12, "rsi":1}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.IntradayRequest
		hasErr bool
		fn     func(ctx context.Context) *Quotes
	}{
		{
			name: "success",
			input: model.IntradayRequest{
				Symbol: "AAPL",
			},
			hasErr: false,
			fn: func(ctx context.Context) *Quotes {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightIntradayQuotes, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Quotes{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name: "err/unmarshal",
			input: model.IntradayRequest{
				Symbol: "AAPL",
			},
			hasErr: true,
			fn: func(ctx context.Context) *Quotes {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightIntradayQuotes, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Quotes{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name: "err/cli",
			input: model.IntradayRequest{
				Symbol: "AAPL",
			},
			hasErr: true,
			fn: func(ctx context.Context) *Quotes {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightIntradayQuotes, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Quotes{
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
			q := tc.fn(ctx)
			resp, err := q.GetIntradayQuotes(ctx, tc.input)

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
