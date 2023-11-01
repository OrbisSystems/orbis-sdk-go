package fi

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

func toPointer[T any](val T) *T {
	return &val
}

func TestNew(t *testing.T) {
	assert.NotNil(t, New(nil))
}

func TestFunds_GetFixedIncomeEntryByID(t *testing.T) {
	var (
		testID = "US52107QAJ40"

		expResp = model.FixedIncome{
			Isin:             toPointer("US52107QAJ40"),
			Cusip:            toPointer("52107QAJ4"),
			Issuer:           toPointer("LAZARD GROUP LLC"),
			Sector:           toPointer("FINANCIAL SERVICES"),
			DescriptionShort: toPointer("LAZ 4.5% 09/19/2028"),
			Type:             toPointer("corporate"),
			IssueSize:        toPointer(int64(500000000)),
			SpRating:         toPointer("BBB+"),
			CouponFrequency:  toPointer(("semi_annual")),
			CouponType:       toPointer(("fixed")),
			Callable:         toPointer(true),
			CallType:         toPointer("ordinary"),
			Convertible:      toPointer(false),
			ParValue:         toPointer(int64(1000)),
			Insured:          toPointer(false),
			Status:           toPointer("outstanding"),
			Purpose:          toPointer("General Purpose"),
			Ticker:           toPointer("LAZ"),
		}

		rawResponse = `{"isin":"US52107QAJ40","cusip":"52107QAJ4","issuer":"LAZARD GROUP LLC","sector":"FINANCIAL SERVICES","description_short":"LAZ 4.5% 09/19/2028","type":"corporate","issue_size":500000000,"sp_rating":"BBB+","coupon_frequency":"semi_annual","coupon_type":"fixed","callable":true,"call_type":"ordinary","convertible":false,"par_value":1000,"insured":false,"status":"outstanding","purpose":"General Purpose","ticker":"LAZ"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  string
		hasErr bool
		fn     func(ctx context.Context, id string) *FixedIncome
	}{
		{
			name:   "success",
			input:  testID,
			hasErr: false,
			fn: func(ctx context.Context, id string) *FixedIncome {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s%s", model.URLInsightBase+model.URLInsightFixedIncome, id), nil).Return(httpResponse, nil)

				return &FixedIncome{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  testID,
			hasErr: true,
			fn: func(ctx context.Context, id string) *FixedIncome {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s%s", model.URLInsightBase+model.URLInsightFixedIncome, id), nil).Return(httpResponse, nil)

				return &FixedIncome{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			input:  testID,
			hasErr: true,
			fn: func(ctx context.Context, id string) *FixedIncome {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s%s", model.URLInsightBase+model.URLInsightFixedIncome, id), nil).Return(nil, testErr)

				return &FixedIncome{
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
			resp, err := f.GetFixedIncomeEntryByID(ctx, tc.input)

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

func TestFunds_GetFixedIncomeEntries(t *testing.T) {
	fixedIncome := model.FixedIncome{
		Isin:             toPointer("US52107QAJ40"),
		Cusip:            toPointer("52107QAJ4"),
		Issuer:           toPointer("LAZARD GROUP LLC"),
		Sector:           toPointer("FINANCIAL SERVICES"),
		DescriptionShort: toPointer("LAZ 4.5% 09/19/2028"),
		Type:             toPointer("corporate"),
		IssueSize:        toPointer(int64(500000000)),
		SpRating:         toPointer("BBB+"),
		CouponFrequency:  toPointer(("semi_annual")),
		CouponType:       toPointer(("fixed")),
		Callable:         toPointer(true),
		CallType:         toPointer("ordinary"),
		Convertible:      toPointer(false),
		ParValue:         toPointer(int64(1000)),
		Insured:          toPointer(false),
		Status:           toPointer("outstanding"),
		Purpose:          toPointer("General Purpose"),
		Ticker:           toPointer("LAZ"),
	}
	var (
		expResp = model.GetFixedIncomeEntriesResponse{
			Data: []model.FixedIncome{
				fixedIncome,
			},
			Count: 1,
		}
		rawReq      = `{"statuses":null,"sectors":null,"types":null,"sp_ratings":null,"coupon":null,"coupon_frequencies":null,"issue_date":null,"maturity":null,"callable":null,"insured":null,"convertible":null,"municipal_states":null,"municipal_types":null,"municipal_tax":null,"purposes":null,"tickers":null,"bank_qualified":null,"paging":{"limit":1}}`
		rawResponse = `{"count":1,"data":[{"isin":"US52107QAJ40","cusip":"52107QAJ4","issuer":"LAZARD GROUP LLC","sector":"FINANCIAL SERVICES","description_short":"LAZ 4.5% 09/19/2028","type":"corporate","issue_size":500000000,"sp_rating":"BBB+","coupon_frequency":"semi_annual","coupon_type":"fixed","callable":true,"call_type":"ordinary","convertible":false,"par_value":1000,"insured":false,"status":"outstanding","purpose":"General Purpose","ticker":"LAZ"}]}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		hasErr bool
		input  model.GetFixedIncomeEntriesRequest
		fn     func(ctx context.Context) *FixedIncome
	}{
		{
			name:   "success",
			hasErr: false,
			input: model.GetFixedIncomeEntriesRequest{
				Paging: &model.Paging{Limit: 1},
			},
			fn: func(ctx context.Context) *FixedIncome {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightFixedIncome, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &FixedIncome{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			hasErr: true,
			input: model.GetFixedIncomeEntriesRequest{
				Paging: &model.Paging{Limit: 1},
			},
			fn: func(ctx context.Context) *FixedIncome {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				r := io.NopCloser(strings.NewReader("x"))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightFixedIncome, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &FixedIncome{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			hasErr: true,
			input: model.GetFixedIncomeEntriesRequest{
				Paging: &model.Paging{Limit: 1},
			},
			fn: func(ctx context.Context) *FixedIncome {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightFixedIncome, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &FixedIncome{
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
			resp, err := f.GetFixedIncomeEntries(ctx, tc.input)

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

func TestFunds_GetFixedIncomeHistorical(t *testing.T) {
	var (
		expResp = model.GetFixedIncomeHistoricalResponse{
			Success: true,
			Status:  1,
			Message: "Asd",
			Data: []model.GetFixedIncomeHistoricalData{
				{
					Price:        12,
					YieldToWorst: 13,
				},
			},
		}
		rawReq      = `{"instrument_id":"US52107QAJ40","frequency":"","start":"","end":""}`
		rawResponse = `{"success":true,"timestamp":"0001-01-01T00:00:00Z","status":1,"message":"Asd","data":[{"timestamp":"0001-01-01T00:00:00Z","price":12,"yield_to_worst":13}]}
`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		hasErr bool
		input  model.GetFixedIncomeHistoricalRequest
		fn     func(ctx context.Context) *FixedIncome
	}{
		{
			name:   "success",
			hasErr: false,
			input: model.GetFixedIncomeHistoricalRequest{
				InstrumentID: "US52107QAJ40",
			},
			fn: func(ctx context.Context) *FixedIncome {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightFixedIncomeHistorical, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &FixedIncome{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			hasErr: true,
			input: model.GetFixedIncomeHistoricalRequest{
				InstrumentID: "US52107QAJ40",
			},
			fn: func(ctx context.Context) *FixedIncome {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				r := io.NopCloser(strings.NewReader("x"))

				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightFixedIncomeHistorical, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &FixedIncome{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			hasErr: true,
			input: model.GetFixedIncomeHistoricalRequest{
				InstrumentID: "US52107QAJ40",
			},
			fn: func(ctx context.Context) *FixedIncome {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightFixedIncomeHistorical, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &FixedIncome{
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
			resp, err := f.GetFixedIncomeHistorical(ctx, tc.input)

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
