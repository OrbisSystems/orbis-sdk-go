package hs

import (
	"bytes"
	"context"
	"encoding/json"
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
	"github.com/OrbisSystems/orbis-sdk-go/utils"
)

var (
	hsTesteExpResp = map[string]interface{}{
		"test_field": "test_value",
	}

	hsTestRawResponse = `{"test_field": "test_value"}`

	hsTestErr = errors.New("process error")
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New(nil, nil))
}

func TestHoopsAI_DailySummary(t *testing.T) {
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

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s?doc_group=%s&enable_html=true&use_customer_assets=%v", model.URLInsightBase+model.URLInsightHSDailySummary, req.Asset, req.DocGroup, req.UseCustomerAssets), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/building_query",
			req:    model.HSDailySummaryRequest{Asset: "FB", UseCustomerAssets: true},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSDailySummaryRequest) *HoopsAI {
				return &HoopsAI{
					logger: logrus.New(),
				}
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
					cli:    cli,
					logger: logrus.New(),
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

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s?doc_group=%s&enable_html=true&use_customer_assets=%v", model.URLInsightBase+model.URLInsightHSDailySummary, req.Asset, req.DocGroup, req.UseCustomerAssets), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.DailySummary(ctx, tc.req)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_WeeklySummary(t *testing.T) {
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

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s?doc_group=%s&enable_html=true&use_customer_assets=%v", model.URLInsightBase+model.URLInsightHSWeeklySummary, req.Asset, req.DocGroup, req.UseCustomerAssets), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/building_query",
			req:    model.HSWeeklySummaryRequest{Asset: "FB", UseCustomerAssets: true},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSWeeklySummaryRequest) *HoopsAI {
				return &HoopsAI{
					logger: logrus.New(),
				}
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
					cli:    cli,
					logger: logrus.New(),
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

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s?doc_group=%s&enable_html=true&use_customer_assets=%v", model.URLInsightBase+model.URLInsightHSWeeklySummary, req.Asset, req.DocGroup, req.UseCustomerAssets), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.WeeklySummary(ctx, tc.req)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_Portfolio(t *testing.T) {
	testCases := []struct {
		name   string
		req    model.HSPortfolioRequest
		hasErr bool
		fn     func(ctx context.Context, req model.HSPortfolioRequest) *HoopsAI
	}{
		{
			name:   "success",
			req:    model.HSPortfolioRequest{},
			hasErr: false,
			fn: func(ctx context.Context, req model.HSPortfolioRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				body, _ := json.Marshal(req)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightHSPortfolio, bytes.NewBuffer(body), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			req:    model.HSPortfolioRequest{},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSPortfolioRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				body, _ := json.Marshal(req)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightHSPortfolio, bytes.NewBuffer(body), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			req:    model.HSPortfolioRequest{},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSPortfolioRequest) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				body, _ := json.Marshal(req)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightHSPortfolio, bytes.NewBuffer(body), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.Portfolio(ctx, tc.req)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_Watchlist(t *testing.T) {
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

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s?doc_group=%s&enable_html=true&use_customer_assets=%v", model.URLInsightBase+model.URLInsightHSWatchlist, req.Asset, req.DocGroup, req.UseCustomerAssets), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/building_query",
			req:    model.HSWatchlistRequest{Asset: "FB", UseCustomerAssets: true},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSWatchlistRequest) *HoopsAI {
				return &HoopsAI{
					logger: logrus.New(),
				}
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
					cli:    cli,
					logger: logrus.New(),
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

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s?doc_group=%s&enable_html=true&use_customer_assets=%v", model.URLInsightBase+model.URLInsightHSWatchlist, req.Asset, req.DocGroup, req.UseCustomerAssets), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.Watchlist(ctx, tc.req)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_WatchlistByUserAndName(t *testing.T) {
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

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s/%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSWatchlist, req.UserID, req.WatchlistName, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/building_query",
			req:    model.HSWatchlistByUserAndNameRequest{WatchlistName: "FB", UserID: "123"},
			hasErr: true,
			fn: func(ctx context.Context, req model.HSWatchlistByUserAndNameRequest) *HoopsAI {
				return &HoopsAI{
					logger: logrus.New(),
				}
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
					cli:    cli,
					logger: logrus.New(),
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

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s/%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSWatchlist, req.UserID, req.WatchlistName, req.DocGroup), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.WatchlistByUserAndName(ctx, tc.req)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_TopGainers(t *testing.T) {
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

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSTopGainers, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
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
					cli:    cli,
					logger: logrus.New(),
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

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSTopGainers, req.DocGroup), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.TopGainers(ctx, tc.req)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_TopLosers(t *testing.T) {
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

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSTopLosers, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
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
					cli:    cli,
					logger: logrus.New(),
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

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSTopLosers, req.DocGroup), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.TopLosers(ctx, tc.req)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_TopMovers(t *testing.T) {
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

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSTopMovers, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
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
					cli:    cli,
					logger: logrus.New(),
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

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSTopMovers, req.DocGroup), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.TopMovers(ctx, tc.req)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_DownTrend(t *testing.T) {
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

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSDownTrend, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
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
					cli:    cli,
					logger: logrus.New(),
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

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSDownTrend, req.DocGroup), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.DownTrend(ctx, tc.req)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_UpTrend(t *testing.T) {
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

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSUpTrend, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
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
					cli:    cli,
					logger: logrus.New(),
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

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSUpTrend, req.DocGroup), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.UpTrend(ctx, tc.req)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_MarketOverview(t *testing.T) {
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

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSMarketOverview, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
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
					cli:    cli,
					logger: logrus.New(),
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

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSMarketOverview, req.DocGroup), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.MarketOverview(ctx, tc.req)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_PriceTarget(t *testing.T) {
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

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSPriceTarget, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
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
					cli:    cli,
					logger: logrus.New(),
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

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSPriceTarget, req.DocGroup), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.PriceTarget(ctx, tc.req)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_UpcomingEarnings(t *testing.T) {
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

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSUpcomingEarnings, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
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
					cli:    cli,
					logger: logrus.New(),
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

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSUpcomingEarnings, req.DocGroup), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.UpcomingEarnings(ctx, tc.req)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_RecentEarnings(t *testing.T) {
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

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSRecentEarnings, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
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
					cli:    cli,
					logger: logrus.New(),
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

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSRecentEarnings, req.DocGroup), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.RecentEarnings(ctx, tc.req)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_RecordHigh(t *testing.T) {
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

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSRecordHigh, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
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
					cli:    cli,
					logger: logrus.New(),
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

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSRecordHigh, req.DocGroup), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.RecordHigh(ctx, tc.req)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_UnusualHighVolume(t *testing.T) {
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

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSUnusualHighVolume, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
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
					cli:    cli,
					logger: logrus.New(),
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

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSUnusualHighVolume, req.DocGroup), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.UnusualHighVolume(ctx, tc.req)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_Data(t *testing.T) {
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

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSAssets, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
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
					cli:    cli,
					logger: logrus.New(),
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

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSAssets, req.DocGroup), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.Data(ctx, tc.req)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_CustomerAssets(t *testing.T) {
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

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSCustomerAssets, req.DocGroup), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
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
					cli:    cli,
					logger: logrus.New(),
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

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?doc_group=%s&enable_html=true", model.URLInsightBase+model.URLInsightHSCustomerAssets, req.DocGroup), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.CustomerAssets(ctx, tc.req)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_GetUsers(t *testing.T) {
	testCases := []struct {
		name   string
		req    string
		hasErr bool
		fn     func(ctx context.Context, req string) *HoopsAI
	}{
		{
			name:   "success",
			req:    "ASD",
			hasErr: false,
			fn: func(ctx context.Context, req string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s", model.URLInsightBase+model.URLInsightHSUsers, req), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			req:    "ASD",
			hasErr: true,
			fn: func(ctx context.Context, req string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s", model.URLInsightBase+model.URLInsightHSUsers, req), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			req:    "ASD",
			hasErr: true,
			fn: func(ctx context.Context, req string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s", model.URLInsightBase+model.URLInsightHSUsers, req), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.req)
			resp, err := hs.GetUsers(ctx, tc.req)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_CreateUser(t *testing.T) {
	testCases := []struct {
		name     string
		customer string
		hasErr   bool
		fn       func(ctx context.Context, customer string) *HoopsAI
	}{
		{
			name:     "success",
			customer: "AAQ",
			hasErr:   false,
			fn: func(ctx context.Context, customer string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, fmt.Sprintf("%s/%s", model.URLInsightBase+model.URLInsightHSUsers, customer), nil, nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:     "err/unmarshal",
			customer: "AAQ",
			hasErr:   true,
			fn: func(ctx context.Context, customer string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, fmt.Sprintf("%s/%s", model.URLInsightBase+model.URLInsightHSUsers, customer), nil, nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:     "err/cli",
			customer: "AAQ",
			hasErr:   true,
			fn: func(ctx context.Context, customer string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Post(ctx, fmt.Sprintf("%s/%s", model.URLInsightBase+model.URLInsightHSUsers, customer), nil, nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.customer)
			resp, err := hs.CreateUser(ctx, tc.customer)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_DeleteUser(t *testing.T) {
	testCases := []struct {
		name             string
		customer, userID string
		hasErr           bool
		fn               func(ctx context.Context, customer, userID string) *HoopsAI
	}{
		{
			name:     "success",
			customer: "AAQ",
			hasErr:   false,
			fn: func(ctx context.Context, customer, userID string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Delete(ctx, fmt.Sprintf("%s/%s/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID), nil, nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:     "err/unmarshal",
			customer: "AAQ",
			hasErr:   true,
			fn: func(ctx context.Context, customer, userID string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Delete(ctx, fmt.Sprintf("%s/%s/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID), nil, nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:     "err/cli",
			customer: "AAQ",
			hasErr:   true,
			fn: func(ctx context.Context, customer, userID string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Delete(ctx, fmt.Sprintf("%s/%s/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID), nil, nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.customer, tc.userID)
			resp, err := hs.DeleteUser(ctx, tc.customer, tc.userID)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_CreateWatchlistByUser(t *testing.T) {
	testCases := []struct {
		name             string
		customer, userID string
		hasErr           bool
		fn               func(ctx context.Context, customer, userID string) *HoopsAI
	}{
		{
			name:     "success",
			customer: "AAQ",
			hasErr:   false,
			fn: func(ctx context.Context, customer, userID string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, fmt.Sprintf("%s/%s/%s/watchlists", model.URLInsightBase+model.URLInsightHSUsers, customer, userID), nil, nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:     "err/unmarshal",
			customer: "AAQ",
			hasErr:   true,
			fn: func(ctx context.Context, customer, userID string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, fmt.Sprintf("%s/%s/%s/watchlists", model.URLInsightBase+model.URLInsightHSUsers, customer, userID), nil, nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:     "err/cli",
			customer: "AAQ",
			hasErr:   true,
			fn: func(ctx context.Context, customer, userID string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Post(ctx, fmt.Sprintf("%s/%s/%s/watchlists", model.URLInsightBase+model.URLInsightHSUsers, customer, userID), nil, nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.customer, tc.userID)
			resp, err := hs.CreateWatchlistByUser(ctx, tc.customer, tc.userID)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_GetWatchlistByUser(t *testing.T) {
	testCases := []struct {
		name                            string
		customer, userID, watchlistName string
		hasErr                          bool
		fn                              func(ctx context.Context, customer, userID, watchlistName string) *HoopsAI
	}{
		{
			name:          "success",
			customer:      "AQ",
			userID:        "QW",
			watchlistName: "DW",
			hasErr:        false,
			fn: func(ctx context.Context, customer, userID, watchlistName string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, watchlistName), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:          "err/unmarshal",
			customer:      "AQ",
			userID:        "QW",
			watchlistName: "DW",
			hasErr:        true,
			fn: func(ctx context.Context, customer, userID, watchlistName string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, watchlistName), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:          "err/cli",
			customer:      "AQ",
			userID:        "QW",
			watchlistName: "DW",
			hasErr:        true,
			fn: func(ctx context.Context, customer, userID, watchlistName string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, watchlistName), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.customer, tc.userID, tc.watchlistName)
			resp, err := hs.GetWatchlistByUser(ctx, tc.customer, tc.userID, tc.watchlistName)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_AddSymbolToWatchlist(t *testing.T) {
	testCases := []struct {
		name                                    string
		customer, userID, watchlistName, symbol string
		hasErr                                  bool
		fn                                      func(ctx context.Context, customer, userID, watchlistName, symbol string) *HoopsAI
	}{
		{
			name:          "success",
			customer:      "AQ",
			userID:        "QW",
			watchlistName: "DW",
			hasErr:        false,
			fn: func(ctx context.Context, customer, userID, watchlistName, symbol string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, watchlistName, symbol), nil, nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:          "err/unmarshal",
			customer:      "AQ",
			userID:        "QW",
			watchlistName: "DW",
			hasErr:        true,
			fn: func(ctx context.Context, customer, userID, watchlistName, symbol string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, watchlistName, symbol), nil, nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:          "err/cli",
			customer:      "AQ",
			userID:        "QW",
			watchlistName: "DW",
			hasErr:        true,
			fn: func(ctx context.Context, customer, userID, watchlistName, symbol string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Post(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, watchlistName, symbol), nil, nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.customer, tc.userID, tc.watchlistName, tc.symbol)
			resp, err := hs.AddSymbolToWatchlist(ctx, tc.customer, tc.userID, tc.watchlistName, tc.symbol)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_DeleteSymbolFromWatchlist(t *testing.T) {
	testCases := []struct {
		name                                    string
		customer, userID, watchlistName, symbol string
		hasErr                                  bool
		fn                                      func(ctx context.Context, customer, userID, watchlistName, symbol string) *HoopsAI
	}{
		{
			name:     "success",
			customer: "AAQ",
			hasErr:   false,
			fn: func(ctx context.Context, customer, userID, watchlistName, symbol string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Delete(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, watchlistName, symbol), nil, nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:     "err/unmarshal",
			customer: "AAQ",
			hasErr:   true,
			fn: func(ctx context.Context, customer, userID, watchlistName, symbol string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Delete(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, watchlistName, symbol), nil, nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:     "err/cli",
			customer: "AAQ",
			hasErr:   true,
			fn: func(ctx context.Context, customer, userID, watchlistName, symbol string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Delete(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, watchlistName, symbol), nil, nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.customer, tc.userID, tc.watchlistName, tc.symbol)
			resp, err := hs.DeleteSymbolFromWatchlist(ctx, tc.customer, tc.userID, tc.watchlistName, tc.symbol)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_AddSymbolsToWatchlist(t *testing.T) {
	testCases := []struct {
		name                            string
		customer, userID, watchlistName string
		symbols                         []string
		hasErr                          bool
		fn                              func(ctx context.Context, customer, userID, watchlistName string, symbols []string) *HoopsAI
	}{
		{
			name:          "success",
			customer:      "AQ",
			userID:        "QW",
			watchlistName: "DW",
			hasErr:        false,
			fn: func(ctx context.Context, customer, userID, watchlistName string, symbols []string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				body, _ := json.Marshal(struct {
					SymbolList string `json:"symbol_list"`
				}{
					SymbolList: utils.ArrayToJoinedString(symbols),
				})

				cli.EXPECT().Post(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, watchlistName), bytes.NewBuffer(body), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:          "err/unmarshal",
			customer:      "AQ",
			userID:        "QW",
			watchlistName: "DW",
			hasErr:        true,
			fn: func(ctx context.Context, customer, userID, watchlistName string, symbols []string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				body, _ := json.Marshal(struct {
					SymbolList string `json:"symbol_list"`
				}{
					SymbolList: utils.ArrayToJoinedString(symbols),
				})

				cli.EXPECT().Post(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, watchlistName), bytes.NewBuffer(body), nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:          "err/cli",
			customer:      "AQ",
			userID:        "QW",
			watchlistName: "DW",
			hasErr:        true,
			fn: func(ctx context.Context, customer, userID, watchlistName string, symbols []string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				body, _ := json.Marshal(struct {
					SymbolList string `json:"symbol_list"`
				}{
					SymbolList: utils.ArrayToJoinedString(symbols),
				})

				cli.EXPECT().Post(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, watchlistName), bytes.NewBuffer(body), nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.customer, tc.userID, tc.watchlistName, tc.symbols)
			resp, err := hs.AddSymbolsToWatchlist(ctx, tc.customer, tc.userID, tc.watchlistName, tc.symbols)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_DeleteWatchlistByName(t *testing.T) {
	testCases := []struct {
		name                            string
		customer, userID, watchlistName string
		hasErr                          bool
		fn                              func(ctx context.Context, customer, userID, watchlistName string) *HoopsAI
	}{
		{
			name:     "success",
			customer: "AAQ",
			hasErr:   false,
			fn: func(ctx context.Context, customer, userID, watchlistName string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Delete(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, watchlistName), nil, nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:     "err/unmarshal",
			customer: "AAQ",
			hasErr:   true,
			fn: func(ctx context.Context, customer, userID, watchlistName string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Delete(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, watchlistName), nil, nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:     "err/cli",
			customer: "AAQ",
			hasErr:   true,
			fn: func(ctx context.Context, customer, userID, watchlistName string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Delete(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, watchlistName), nil, nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.customer, tc.userID, tc.watchlistName)
			resp, err := hs.DeleteWatchlistByName(ctx, tc.customer, tc.userID, tc.watchlistName)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}

func TestHoopsAI_RenameWatchlist(t *testing.T) {
	testCases := []struct {
		name                                                 string
		customer, userID, oldWatchlistName, newWatchlistName string
		hasErr                                               bool
		fn                                                   func(ctx context.Context, customer, userID, oldWatchlistName, newWatchlistName string) *HoopsAI
	}{
		{
			name:     "success",
			customer: "AAQ",
			hasErr:   false,
			fn: func(ctx context.Context, customer, userID, oldWatchlistName, newWatchlistName string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(hsTestRawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Put(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s/rename/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, oldWatchlistName, newWatchlistName), nil, nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:     "err/unmarshal",
			customer: "AAQ",
			hasErr:   true,
			fn: func(ctx context.Context, customer, userID, oldWatchlistName, newWatchlistName string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Put(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s/rename/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, oldWatchlistName, newWatchlistName), nil, nil).Return(httpResponse, nil)

				return &HoopsAI{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:     "err/cli",
			customer: "AAQ",
			hasErr:   true,
			fn: func(ctx context.Context, customer, userID, oldWatchlistName, newWatchlistName string) *HoopsAI {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Put(ctx, fmt.Sprintf("%s/%s/%s/watchlists/%s/rename/%s", model.URLInsightBase+model.URLInsightHSUsers, customer, userID, oldWatchlistName, newWatchlistName), nil, nil).Return(nil, hsTestErr)

				return &HoopsAI{
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
			hs := tc.fn(ctx, tc.customer, tc.userID, tc.oldWatchlistName, tc.newWatchlistName)
			resp, err := hs.RenameWatchlist(ctx, tc.customer, tc.userID, tc.oldWatchlistName, tc.newWatchlistName)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, hsTesteExpResp, resp)
			}
		})
	}
}
