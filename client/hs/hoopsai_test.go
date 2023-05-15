package hs

import (
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

func TestHoopsAI_DailySummary(t *testing.T) {
	var (
		expResp = map[string]interface{}{
			"test_field": "test_value",
		}

		rawResponse = `{"test_field": "test_value"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		req    model.HSDailySummaryRequest
		hasErr bool
		fn     func(ctx context.Context, req model.HSDailySummaryRequest) *HoopsAI
	}{
		{
			name:   "success",
			req:    model.HSDailySummaryRequest{Asset: "FB", UseCustomerAssets: true, DocGroup: "group"},
			hasErr: false,
			fn: func(ctx context.Context, req model.HSDailySummaryRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s?doc_group=%s&enable_html=true&use_customer_assets=%v", model.URLInsightBase+model.URLInsightHSDailySummary, req.Asset, req.DocGroup, req.UseCustomerAssets), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/building_query",
			req:    model.HSDailySummaryRequest{Asset: "FB", UseCustomerAssets: true},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSDailySummaryRequest) *HoopsAI {
				return &HoopsAI{}
			},
		},
		{
			name:   "err/unmarshal",
			req:    model.HSDailySummaryRequest{Asset: "FB", UseCustomerAssets: true, DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSDailySummaryRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s?doc_group=%s&enable_html=true&use_customer_assets=%v", model.URLInsightBase+model.URLInsightHSDailySummary, req.Asset, req.DocGroup, req.UseCustomerAssets), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			req:    model.HSDailySummaryRequest{Asset: "FB", UseCustomerAssets: true, DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSDailySummaryRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s?doc_group=%s&enable_html=true&use_customer_assets=%v", model.URLInsightBase+model.URLInsightHSDailySummary, req.Asset, req.DocGroup, req.UseCustomerAssets), nil).Return(nil, testErr)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.DailySummary(ctx, tc.req)

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

func TestHoopsAI_WeeklySummary(t *testing.T) {
	var (
		expResp = map[string]interface{}{
			"test_field": "test_value",
		}

		rawResponse = `{"test_field": "test_value"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		req    model.HSWeeklySummaryRequest
		hasErr bool
		fn     func(ctx context.Context, req model.HSWeeklySummaryRequest) *HoopsAI
	}{
		{
			name:   "success",
			req:    model.HSWeeklySummaryRequest{Asset: "FB", UseCustomerAssets: true, DocGroup: "group"},
			hasErr: false,
			fn: func(ctx context.Context, req model.HSWeeklySummaryRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s?doc_group=%s&enable_html=true&use_customer_assets=%v", model.URLInsightBase+model.URLInsightHSWeeklySummary, req.Asset, req.DocGroup, req.UseCustomerAssets), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/building_query",
			req:    model.HSWeeklySummaryRequest{Asset: "FB", UseCustomerAssets: true},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSWeeklySummaryRequest) *HoopsAI {
				return &HoopsAI{}
			},
		},
		{
			name:   "err/unmarshal",
			req:    model.HSWeeklySummaryRequest{Asset: "FB", UseCustomerAssets: true, DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSWeeklySummaryRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s?doc_group=%s&enable_html=true&use_customer_assets=%v", model.URLInsightBase+model.URLInsightHSWeeklySummary, req.Asset, req.DocGroup, req.UseCustomerAssets), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			req:    model.HSWeeklySummaryRequest{Asset: "FB", UseCustomerAssets: true, DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSWeeklySummaryRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s?doc_group=%s&enable_html=true&use_customer_assets=%v", model.URLInsightBase+model.URLInsightHSWeeklySummary, req.Asset, req.DocGroup, req.UseCustomerAssets), nil).Return(nil, testErr)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.WeeklySummary(ctx, tc.req)

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

func TestHoopsAI_Watchlist(t *testing.T) {
	var (
		expResp = map[string]interface{}{
			"test_field": "test_value",
		}

		rawResponse = `{"test_field": "test_value"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		req    model.HSWatchlistRequest
		hasErr bool
		fn     func(ctx context.Context, req model.HSWatchlistRequest) *HoopsAI
	}{
		{
			name:   "success",
			req:    model.HSWatchlistRequest{Asset: "FB", UseCustomerAssets: true, DocGroup: "group"},
			hasErr: false,
			fn: func(ctx context.Context, req model.HSWatchlistRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s?doc_group=%s&enable_html=true&use_customer_assets=%v", model.URLInsightBase+model.URLInsightHSWatchlist, req.Asset, req.DocGroup, req.UseCustomerAssets), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/building_query",
			req:    model.HSWatchlistRequest{Asset: "FB", UseCustomerAssets: true},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSWatchlistRequest) *HoopsAI {
				return &HoopsAI{}
			},
		},
		{
			name:   "err/unmarshal",
			req:    model.HSWatchlistRequest{Asset: "FB", UseCustomerAssets: true, DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSWatchlistRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s?doc_group=%s&enable_html=true&use_customer_assets=%v", model.URLInsightBase+model.URLInsightHSWatchlist, req.Asset, req.DocGroup, req.UseCustomerAssets), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			req:    model.HSWatchlistRequest{Asset: "FB", UseCustomerAssets: true, DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSWatchlistRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s?doc_group=%s&enable_html=true&use_customer_assets=%v", model.URLInsightBase+model.URLInsightHSWatchlist, req.Asset, req.DocGroup, req.UseCustomerAssets), nil).Return(nil, testErr)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.Watchlist(ctx, tc.req)

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

func TestHoopsAI_WatchlistByUserAndName(t *testing.T) {
	var (
		expResp = map[string]interface{}{
			"test_field": "test_value",
		}

		rawResponse = `{"test_field": "test_value"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		req    model.HSWatchlistByUserAndNameRequest
		hasErr bool
		fn     func(ctx context.Context, req model.HSWatchlistByUserAndNameRequest) *HoopsAI
	}{
		{
			name:   "success",
			req:    model.HSWatchlistByUserAndNameRequest{WatchlistName: "FB", UserID: "123", DocGroup: "group"},
			hasErr: false,
			fn: func(ctx context.Context, req model.HSWatchlistByUserAndNameRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s/%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSWatchlist, req.UserID, req.WatchlistName, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/building_query",
			req:    model.HSWatchlistByUserAndNameRequest{WatchlistName: "FB", UserID: "123"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSWatchlistByUserAndNameRequest) *HoopsAI {
				return &HoopsAI{}
			},
		},
		{
			name:   "err/unmarshal",
			req:    model.HSWatchlistByUserAndNameRequest{WatchlistName: "FB", UserID: "123", DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSWatchlistByUserAndNameRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s/%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSWatchlist, req.UserID, req.WatchlistName, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			req:    model.HSWatchlistByUserAndNameRequest{WatchlistName: "FB", UserID: "123", DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSWatchlistByUserAndNameRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s/%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSWatchlist, req.UserID, req.WatchlistName, req.DocGroup), nil).Return(nil, testErr)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.WatchlistByUserAndName(ctx, tc.req)

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

func TestHoopsAI_TopGainers(t *testing.T) {
	var (
		expResp = map[string]interface{}{
			"test_field": "test_value",
		}

		rawResponse = `{"test_field": "test_value"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		req    model.HSMarketResearchRequest
		hasErr bool
		fn     func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI
	}{
		{
			name:   "success",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: false,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSTopGainers, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSTopGainers, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSTopGainers, req.DocGroup), nil).Return(nil, testErr)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.TopGainers(ctx, tc.req)

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

func TestHoopsAI_TopLosers(t *testing.T) {
	var (
		expResp = map[string]interface{}{
			"test_field": "test_value",
		}

		rawResponse = `{"test_field": "test_value"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		req    model.HSMarketResearchRequest
		hasErr bool
		fn     func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI
	}{
		{
			name:   "success",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: false,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSTopLosers, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSTopLosers, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSTopLosers, req.DocGroup), nil).Return(nil, testErr)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.TopLosers(ctx, tc.req)

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

func TestHoopsAI_TopMovers(t *testing.T) {
	var (
		expResp = map[string]interface{}{
			"test_field": "test_value",
		}

		rawResponse = `{"test_field": "test_value"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		req    model.HSMarketResearchRequest
		hasErr bool
		fn     func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI
	}{
		{
			name:   "success",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: false,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSTopMovers, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSTopMovers, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSTopMovers, req.DocGroup), nil).Return(nil, testErr)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.TopMovers(ctx, tc.req)

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

func TestHoopsAI_DownTrend(t *testing.T) {
	var (
		expResp = map[string]interface{}{
			"test_field": "test_value",
		}

		rawResponse = `{"test_field": "test_value"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		req    model.HSMarketResearchRequest
		hasErr bool
		fn     func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI
	}{
		{
			name:   "success",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: false,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSDownTrend, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSDownTrend, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSDownTrend, req.DocGroup), nil).Return(nil, testErr)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.DownTrend(ctx, tc.req)

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

func TestHoopsAI_UpTrend(t *testing.T) {
	var (
		expResp = map[string]interface{}{
			"test_field": "test_value",
		}

		rawResponse = `{"test_field": "test_value"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		req    model.HSMarketResearchRequest
		hasErr bool
		fn     func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI
	}{
		{
			name:   "success",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: false,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSUpTrend, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSUpTrend, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSUpTrend, req.DocGroup), nil).Return(nil, testErr)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.UpTrend(ctx, tc.req)

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

func TestHoopsAI_MarketOverview(t *testing.T) {
	var (
		expResp = map[string]interface{}{
			"test_field": "test_value",
		}

		rawResponse = `{"test_field": "test_value"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		req    model.HSMarketResearchRequest
		hasErr bool
		fn     func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI
	}{
		{
			name:   "success",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: false,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSMarketOverview, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSMarketOverview, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSMarketOverview, req.DocGroup), nil).Return(nil, testErr)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.MarketOverview(ctx, tc.req)

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

func TestHoopsAI_PriceTarget(t *testing.T) {
	var (
		expResp = map[string]interface{}{
			"test_field": "test_value",
		}

		rawResponse = `{"test_field": "test_value"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		req    model.HSMarketResearchRequest
		hasErr bool
		fn     func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI
	}{
		{
			name:   "success",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: false,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSPriceTarget, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSPriceTarget, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSPriceTarget, req.DocGroup), nil).Return(nil, testErr)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.PriceTarget(ctx, tc.req)

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

func TestHoopsAI_UpcomingEarnings(t *testing.T) {
	var (
		expResp = map[string]interface{}{
			"test_field": "test_value",
		}

		rawResponse = `{"test_field": "test_value"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		req    model.HSMarketResearchRequest
		hasErr bool
		fn     func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI
	}{
		{
			name:   "success",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: false,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSUpcomingEarnings, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSUpcomingEarnings, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSUpcomingEarnings, req.DocGroup), nil).Return(nil, testErr)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.UpcomingEarnings(ctx, tc.req)

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

func TestHoopsAI_RecentEarnings(t *testing.T) {
	var (
		expResp = map[string]interface{}{
			"test_field": "test_value",
		}

		rawResponse = `{"test_field": "test_value"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		req    model.HSMarketResearchRequest
		hasErr bool
		fn     func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI
	}{
		{
			name:   "success",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: false,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSRecentEarnings, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSRecentEarnings, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSRecentEarnings, req.DocGroup), nil).Return(nil, testErr)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.RecentEarnings(ctx, tc.req)

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

func TestHoopsAI_RecordHigh(t *testing.T) {
	var (
		expResp = map[string]interface{}{
			"test_field": "test_value",
		}

		rawResponse = `{"test_field": "test_value"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		req    model.HSMarketResearchRequest
		hasErr bool
		fn     func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI
	}{
		{
			name:   "success",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: false,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSRecordHigh, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSRecordHigh, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSRecordHigh, req.DocGroup), nil).Return(nil, testErr)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.RecordHigh(ctx, tc.req)

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

func TestHoopsAI_UnusualHighVolume(t *testing.T) {
	var (
		expResp = map[string]interface{}{
			"test_field": "test_value",
		}

		rawResponse = `{"test_field": "test_value"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		req    model.HSMarketResearchRequest
		hasErr bool
		fn     func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI
	}{
		{
			name:   "success",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: false,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSUnusualHighVolume, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSUnusualHighVolume, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSUnusualHighVolume, req.DocGroup), nil).Return(nil, testErr)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.UnusualHighVolume(ctx, tc.req)

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

func TestHoopsAI_Data(t *testing.T) {
	var (
		expResp = map[string]interface{}{
			"test_field": "test_value",
		}

		rawResponse = `{"test_field": "test_value"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		req    model.HSMarketResearchRequest
		hasErr bool
		fn     func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI
	}{
		{
			name:   "success",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: false,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSAssets, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSAssets, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSAssets, req.DocGroup), nil).Return(nil, testErr)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.Data(ctx, tc.req)

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

func TestHoopsAI_CustomerAssets(t *testing.T) {
	var (
		expResp = map[string]interface{}{
			"test_field": "test_value",
		}

		rawResponse = `{"test_field": "test_value"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		req    model.HSMarketResearchRequest
		hasErr bool
		fn     func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI
	}{
		{
			name:   "success",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: false,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSCustomerAssets, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSCustomerAssets, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			req:    model.HSMarketResearchRequest{DocGroup: "group"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSMarketResearchRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSCustomerAssets, req.DocGroup), nil).Return(nil, testErr)

				return &HoopsAI{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.CustomerAssets(ctx, tc.req)

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
