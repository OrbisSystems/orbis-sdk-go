package market

import (
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

func TestWorldMarket_GetContinents(t *testing.T) {
	var (
		expResp = []model.Continent{
			{
				ID:   "1",
				Name: "Adc",
			},
			{
				ID:   "3",
				Name: "Cwe",
			},
		}

		rawResponse = `[{"id": "1", "name": "Adc"}, {"id": "3", "name": "Cwe"}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name          string
		limit, offset int
		hasErr        bool
		fn            func(ctx context.Context, limit, offset int) *WorldMarket
	}{
		{
			name:   "success",
			limit:  1,
			offset: 1,
			hasErr: false,
			fn: func(ctx context.Context, limit, offset int) *WorldMarket {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?limit=%d&offset=%d", model.URLInsightBase+model.URLInsightContinents, limit, offset), nil).Return(httpResponse, nil)

				return &WorldMarket{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			limit:  1,
			offset: 1,
			hasErr: true,
			fn: func(ctx context.Context, limit, offset int) *WorldMarket {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?limit=%d&offset=%d", model.URLInsightBase+model.URLInsightContinents, limit, offset), nil).Return(httpResponse, nil)

				return &WorldMarket{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			limit:  1,
			offset: 1,
			hasErr: true,
			fn: func(ctx context.Context, limit, offset int) *WorldMarket {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?limit=%d&offset=%d", model.URLInsightBase+model.URLInsightContinents, limit, offset), nil).Return(nil, testErr)

				return &WorldMarket{
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
			resp, err := f.GetContinents(ctx, tc.limit, tc.offset)

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

func TestWorldMarket_GetRegions(t *testing.T) {
	var (
		expResp = []model.Region{
			{
				ID:   "1",
				Name: "Adc",
			},
			{
				ID:   "3",
				Name: "Cwe",
			},
		}

		rawResponse = `[{"id": "1", "name": "Adc"}, {"id": "3", "name": "Cwe"}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name          string
		limit, offset int
		hasErr        bool
		fn            func(ctx context.Context, limit, offset int) *WorldMarket
	}{
		{
			name:   "success",
			limit:  1,
			offset: 1,
			hasErr: false,
			fn: func(ctx context.Context, limit, offset int) *WorldMarket {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?limit=%d&offset=%d", model.URLInsightBase+model.URLInsightRegions, limit, offset), nil).Return(httpResponse, nil)

				return &WorldMarket{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			limit:  1,
			offset: 1,
			hasErr: true,
			fn: func(ctx context.Context, limit, offset int) *WorldMarket {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?limit=%d&offset=%d", model.URLInsightBase+model.URLInsightRegions, limit, offset), nil).Return(httpResponse, nil)

				return &WorldMarket{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			limit:  1,
			offset: 1,
			hasErr: true,
			fn: func(ctx context.Context, limit, offset int) *WorldMarket {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?limit=%d&offset=%d", model.URLInsightBase+model.URLInsightRegions, limit, offset), nil).Return(nil, testErr)

				return &WorldMarket{
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
			resp, err := f.GetRegions(ctx, tc.limit, tc.offset)

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

func TestWorldMarket_GetCountryCodes(t *testing.T) {
	var (
		expResp = []model.CountryCode{
			{
				A2Code:        "AA",
				A3Code:        "AAB",
				CountryName:   "Aabcx",
				NumCode:       "123",
				MorningstarID: "332",
			},
		}

		rawResponse = `[{"a2_code": "AA", "a3_code": "AAB", "country_name":"Aabcx", "num_code":"123", "morningstar_id":"332"}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name          string
		limit, offset int
		hasErr        bool
		fn            func(ctx context.Context, limit, offset int) *WorldMarket
	}{
		{
			name:   "success",
			limit:  1,
			offset: 1,
			hasErr: false,
			fn: func(ctx context.Context, limit, offset int) *WorldMarket {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?limit=%d&offset=%d", model.URLInsightBase+model.URLInsightCountryCodes, limit, offset), nil).Return(httpResponse, nil)

				return &WorldMarket{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			limit:  1,
			offset: 1,
			hasErr: true,
			fn: func(ctx context.Context, limit, offset int) *WorldMarket {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?limit=%d&offset=%d", model.URLInsightBase+model.URLInsightCountryCodes, limit, offset), nil).Return(httpResponse, nil)

				return &WorldMarket{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			limit:  1,
			offset: 1,
			hasErr: true,
			fn: func(ctx context.Context, limit, offset int) *WorldMarket {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?limit=%d&offset=%d", model.URLInsightBase+model.URLInsightCountryCodes, limit, offset), nil).Return(nil, testErr)

				return &WorldMarket{
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
			resp, err := f.GetCountryCodes(ctx, tc.limit, tc.offset)

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

func TestWorldMarket_GetGlobalIndexes(t *testing.T) {
	var (
		expResp = []model.GlobalIndexFull{
			{
				Symbol:    "AAPL",
				CountryID: "12",
				RegionID:  "33",
			},
		}

		rawResponse = `[{"symbol":"AAPL","country_id":"12","region_id":"33","continent_id":"","ppp":0,"gdp":0,"country_name":"","region_name":"","continent_name":"","company_name":"","last_price":0,"price_change":0,"price_change_percent":0}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name                 string
		limit, offset        int
		continent, quoteType string
		hasErr               bool
		fn                   func(ctx context.Context, limit, offset int, continent, quoteType string) *WorldMarket
	}{
		{
			name:      "success",
			limit:     1,
			offset:    1,
			continent: "x",
			quoteType: "delayed",
			hasErr:    false,
			fn: func(ctx context.Context, limit, offset int, continent, quoteType string) *WorldMarket {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?limit=%d&offset=%d&continent=%s&quoteType=%s", model.URLInsightBase+model.URLInsightGlobalIndexes, limit, offset, continent, quoteType), nil).Return(httpResponse, nil)

				return &WorldMarket{
					cli: cli,
				}
			},
		},
		{
			name:      "err/unmarshal",
			limit:     1,
			offset:    1,
			continent: "x",
			quoteType: "delayed",
			hasErr:    true,
			fn: func(ctx context.Context, limit, offset int, continent, quoteType string) *WorldMarket {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?limit=%d&offset=%d&continent=%s&quoteType=%s", model.URLInsightBase+model.URLInsightGlobalIndexes, limit, offset, continent, quoteType), nil).Return(httpResponse, nil)

				return &WorldMarket{
					cli: cli,
				}
			},
		},
		{
			name:      "err/cli",
			limit:     1,
			offset:    1,
			continent: "x",
			quoteType: "delayed",
			hasErr:    true,
			fn: func(ctx context.Context, limit, offset int, continent, quoteType string) *WorldMarket {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?limit=%d&offset=%d&continent=%s&quoteType=%s", model.URLInsightBase+model.URLInsightGlobalIndexes, limit, offset, continent, quoteType), nil).Return(nil, testErr)

				return &WorldMarket{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			f := tc.fn(ctx, tc.limit, tc.offset, tc.continent, tc.quoteType)
			resp, err := f.GetGlobalIndexes(ctx, tc.limit, tc.offset, tc.continent, tc.quoteType)

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
