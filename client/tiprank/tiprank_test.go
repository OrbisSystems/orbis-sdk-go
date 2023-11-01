package tiprank

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

func TestTipRank_AnalystConsensus(t *testing.T) {
	var (
		expResp = []model.AnalystConsensusResponse{
			{
				LowPriceTarget:          1,
				HighPriceTarget:         2,
				PriceTarget:             3,
				Buy:                     4,
				Sell:                    5,
				Hold:                    6,
				Consensus:               "Ac",
				PriceTargetUpside:       3,
				Ticker:                  "AAPL",
				CompanyName:             "Apple",
				PriceTargetCurrencyCode: "CA",
				TotalAnalysts:           123,
			},
		}

		rawReq = `{"tickers":["AAPL","FB"]}`

		rawResponse = `[{"low_price_target":1,"high_price_target":2,"price_target":3,"buy":4,"sell":5,"hold":6,"consensus":"Ac","price_target_upside":3,"ticker":"AAPL","company_name":"Apple","price_target_currency_code":"CA","total_analysts":123}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.AnalystConsensusRequest
		hasErr bool
		fn     func(ctx context.Context) *TipRank
	}{
		{
			name: "success",
			input: model.AnalystConsensusRequest{
				Tickers: []string{"AAPL", "FB"},
			},
			hasErr: false,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightTipRankAnalystConsensus, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name: "err/unmarshal",
			input: model.AnalystConsensusRequest{
				Tickers: []string{"AAPL", "FB"},
			},
			hasErr: true,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightTipRankAnalystConsensus, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name: "err/cli",
			input: model.AnalystConsensusRequest{
				Tickers: []string{"AAPL", "FB"},
			},
			hasErr: true,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightTipRankAnalystConsensus, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &TipRank{
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
			resp, err := r.AnalystConsensus(ctx, tc.input)

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

func TestTipRank_LatestAnalystRatingsOnStock(t *testing.T) {
	var (
		expResp = []model.LatestAnalystRatingsOnStockResponse{
			{
				AnalystName:              "Awd",
				FirmName:                 "OOO",
				AnalystRank:              12,
				NumberOfRankedExperts:    32,
				SuccessRate:              4,
				ExcessReturn:             5,
				StockGoodRecommendations: 3,
			},
		}

		rawReq = `{"tickers":["AAPL","FB"],"num":0,"sort":"","months":0}`

		rawResponse = `[{"analyst_name":"Awd","firm_name":"OOO","recommendation":"","recommendation_date":"","expert_uid":"","url":"","expert_picture_url":"","expert_full_picture_url":"","analyst_rank":12,"number_of_ranked_experts":32,"success_rate":4,"excess_return":5,"total_recommendations":0,"good_recommendations":0,"num_of_stars":0,"stock_success_rate":0,"stock_avg_return":0,"article_title":"","article_site":"","article_quote":"","price_target":0,"price_target_currency":0,"price_target_currency_code":"","converted_price_target":0,"converted_price_target_currency":0,"converted_price_target_currency_code":"","ticker":"","timestamp":"","analyst_action":"","stock_total_recommendations":0,"stock_good_recommendations":3}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.LatestAnalystRatingsOnStockRequest
		hasErr bool
		fn     func(ctx context.Context) *TipRank
	}{
		{
			name: "success",
			input: model.LatestAnalystRatingsOnStockRequest{
				Tickers: []string{"AAPL", "FB"},
			},
			hasErr: false,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightTipRankAnalystMulti, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name: "err/unmarshal",
			input: model.LatestAnalystRatingsOnStockRequest{
				Tickers: []string{"AAPL", "FB"},
			},
			hasErr: true,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightTipRankAnalystMulti, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name: "err/cli",
			input: model.LatestAnalystRatingsOnStockRequest{
				Tickers: []string{"AAPL", "FB"},
			},
			hasErr: true,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightTipRankAnalystMulti, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &TipRank{
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
			resp, err := r.LatestAnalystRatingsOnStock(ctx, tc.input)

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

func TestTipRank_LiveFeed(t *testing.T) {
	var (
		expResp = []model.LiveFeedResponse{
			{
				AnalystName:           "Awd",
				FirmName:              "OOO",
				AnalystRank:           12,
				NumberOfRankedExperts: 32,
				SuccessRate:           4,
				ExcessReturn:          5,
			},
		}

		rawReq = `{"num":12,"bookmark_id":0,"quote_access_type":""}`

		rawResponse = `[{"analyst_name":"Awd","firm_name":"OOO","recommendation":"","recommendation_date":"","expert_uid":"","url":"","expert_picture_url":"","expert_full_picture_url":"","analyst_rank":12,"number_of_ranked_experts":32,"success_rate":4,"excess_return":5,"total_recommendations":0,"good_recommendations":0,"num_of_stars":0,"stock_success_rate":0,"stock_avg_return":0,"article_title":"","article_site":"","article_quote":"","price_target":0,"price_target_currency":0,"price_target_currency_code":"","converted_price_target":0,"converted_price_target_currency":0,"converted_price_target_currency_code":"","ticker":"","timestamp":"","analyst_action":"","stock_total_recommendations":0}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.LiveFeedRequest
		hasErr bool
		fn     func(ctx context.Context) *TipRank
	}{
		{
			name: "success",
			input: model.LiveFeedRequest{
				Num: 12,
			},
			hasErr: false,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightTipRankLiveFeed, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name: "err/unmarshal",
			input: model.LiveFeedRequest{
				Num: 12,
			},
			hasErr: true,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightTipRankLiveFeed, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name: "err/cli",
			input: model.LiveFeedRequest{
				Num: 12,
			},
			hasErr: true,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightTipRankLiveFeed, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &TipRank{
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
			resp, err := r.LiveFeed(ctx, tc.input)

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

func TestTipRank_TrendingStocks(t *testing.T) {
	var (
		expResp = []model.TrendingStocksResponse{
			{
				Ticker: "AAPL",
				Hold:   123,
				Buy:    145,
				Sell:   169,
			},
		}

		rawReq = `{"num":12,"filter":"","period":0,"trending_type":"","quote_access_type":""}`

		rawResponse = `[{"ticker":"AAPL","hold":123,"buy":145,"sell":169}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.TrendingStocksRequest
		hasErr bool
		fn     func(ctx context.Context) *TipRank
	}{
		{
			name: "success",
			input: model.TrendingStocksRequest{
				Num: 12,
			},
			hasErr: false,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightTipRankTrendingStocks, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name: "err/unmarshal",
			input: model.TrendingStocksRequest{
				Num: 12,
			},
			hasErr: true,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightTipRankTrendingStocks, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name: "err/cli",
			input: model.TrendingStocksRequest{
				Num: 12,
			},
			hasErr: true,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightTipRankTrendingStocks, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &TipRank{
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
			resp, err := r.TrendingStocks(ctx, tc.input)

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

func TestTipRank_AnalystPortfolios(t *testing.T) {
	var (
		expResp = []model.PortfoliosResponse{
			{
				StockTicker: "AAPL",
				CompanyName: "Apple",
				PriceTarget: 180,
			},
		}

		rawReq = `{"expert_uid":"","number":12,"quote_access_type":""}`

		rawResponse = `[{"stock_ticker":"AAPL","company_name":"Apple","price_target":180}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.PortfoliosRequest
		hasErr bool
		fn     func(ctx context.Context) *TipRank
	}{
		{
			name: "success",
			input: model.PortfoliosRequest{
				Number: 12,
			},
			hasErr: false,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightTipRankAnalystPortfolio, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name: "err/unmarshal",
			input: model.PortfoliosRequest{
				Number: 12,
			},
			hasErr: true,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightTipRankAnalystPortfolio, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name: "err/cli",
			input: model.PortfoliosRequest{
				Number: 12,
			},
			hasErr: true,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightTipRankAnalystPortfolio, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &TipRank{
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
			resp, err := r.AnalystPortfolios(ctx, tc.input)

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

func TestTipRank_AnalystProfile(t *testing.T) {
	var (
		expResp = model.AnalystProfileResponse{
			AnalystName:           "Wdew",
			FirmName:              "Yjjs",
			AnalystRank:           1,
			NumberOfRankedExperts: 2,
			TotalRecommendations:  3,
			GoodRecommendations:   4,
		}

		rawResponse = `{"analyst_name":"Wdew","firm_name":"Yjjs","expert_picture_url":"","expert_full_picture_url":"","success_rate":"","excess_return":"","analyst_rank":1,"number_of_ranked_experts":2,"total_recommendations":3,"good_recommendations":4,"num_of_stars":0,"main_sector":""}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		id     string
		hasErr bool
		fn     func(ctx context.Context, id string) *TipRank
	}{
		{
			name:   "success",
			id:     "123",
			hasErr: false,
			fn: func(ctx context.Context, id string) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?id=%s", model.URLInsightBase+model.URLInsightTipRankAnalystProfile, id), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			id:     "123",
			hasErr: true,
			fn: func(ctx context.Context, id string) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?id=%s", model.URLInsightBase+model.URLInsightTipRankAnalystProfile, id), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			id:     "123",
			hasErr: true,
			fn: func(ctx context.Context, id string) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?id=%s", model.URLInsightBase+model.URLInsightTipRankAnalystProfile, id), nil).Return(nil, testErr)

				return &TipRank{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			r := tc.fn(ctx, tc.id)
			resp, err := r.AnalystProfile(ctx, tc.id)

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

func TestTipRank_SectorConsensus(t *testing.T) {
	var (
		expResp = model.SectorConsensusResponse{
			Sectors: []*model.SectorConsensus{
				{
					Sector:       "Qaz",
					StrongBuy:    1,
					ModerateBuy:  2,
					Neutral:      3,
					ModerateSell: 4,
					StrongSell:   3,
					Order:        2,
				},
				{
					Sector:       "Arrs",
					StrongBuy:    3,
					ModerateBuy:  1,
					Neutral:      6,
					ModerateSell: 3,
					StrongSell:   9,
					Order:        2,
				},
			},
		}

		rawResponse = `{"sectors":[{"sector":"Qaz","strong_buy":1,"moderate_buy":2,"neutral":3,"moderate_sell":4,"strong_sell":3,"order":2},{"sector":"Arrs","strong_buy":3,"moderate_buy":1,"neutral":6,"moderate_sell":3,"strong_sell":9,"order":2}]}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		hasErr bool
		fn     func(ctx context.Context) *TipRank
	}{
		{
			name:   "success",
			hasErr: false,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, model.URLInsightBase+model.URLInsightTipRankSectorConsensus, nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			hasErr: true,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, model.URLInsightBase+model.URLInsightTipRankSectorConsensus, nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			hasErr: true,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, model.URLInsightBase+model.URLInsightTipRankSectorConsensus, nil).Return(nil, testErr)

				return &TipRank{
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
			resp, err := r.SectorConsensus(ctx)

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

func TestTipRank_BestPerformingExperts(t *testing.T) {
	var (
		expResp = []model.BestPerformingExpertsResponse{
			{
				AnalystName:           "Oll",
				FirmName:              "Pdd AAC",
				NumberOfRankedExperts: 2,
				SuccessRate:           3,
				ExcessReturn:          2,
				TotalRecommendations:  12,
			},
		}

		rawResponse = `[{"analyst_name":"Oll","firm_name":"Pdd AAC","number_of_ranked_experts":2,"success_rate":3,"excess_return":2,"total_recommendations":12}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		num    int
		hasErr bool
		fn     func(ctx context.Context, num int) *TipRank
	}{
		{
			name:   "success",
			num:    123,
			hasErr: false,
			fn: func(ctx context.Context, num int) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?num=%d", model.URLInsightBase+model.URLInsightTipRankBestPerformingExperts, num), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			num:    123,
			hasErr: true,
			fn: func(ctx context.Context, num int) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?num=%d", model.URLInsightBase+model.URLInsightTipRankBestPerformingExperts, num), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			num:    123,
			hasErr: true,
			fn: func(ctx context.Context, num int) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?num=%d", model.URLInsightBase+model.URLInsightTipRankBestPerformingExperts, num), nil).Return(nil, testErr)

				return &TipRank{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			r := tc.fn(ctx, tc.num)
			resp, err := r.BestPerformingExperts(ctx, tc.num)

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

func TestTipRank_StocksSimilarStocks(t *testing.T) {
	var (
		expResp = []string{"1", "2", "3"}

		rawResponse = `["1", "2", "3"]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		ticker string
		hasErr bool
		fn     func(ctx context.Context, ticker string) *TipRank
	}{
		{
			name:   "success",
			ticker: "APPL",
			hasErr: false,
			fn: func(ctx context.Context, ticker string) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?ticker=%s", model.URLInsightBase+model.URLInsightTipRankStocksSimilarStocks, ticker), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			ticker: "APPL",
			hasErr: true,
			fn: func(ctx context.Context, ticker string) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?ticker=%s", model.URLInsightBase+model.URLInsightTipRankStocksSimilarStocks, ticker), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			ticker: "APPL",
			hasErr: true,
			fn: func(ctx context.Context, ticker string) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?ticker=%s", model.URLInsightBase+model.URLInsightTipRankStocksSimilarStocks, ticker), nil).Return(nil, testErr)

				return &TipRank{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			r := tc.fn(ctx, tc.ticker)
			resp, err := r.StocksSimilarStocks(ctx, tc.ticker)

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

func TestTipRank_AnalystsExpertPictureStore(t *testing.T) {
	var (
		expResp = model.AnalystsExpertPictureStoreResponse{
			URL: "https://localhost.com",
		}

		rawResponse = `{"url": "https://localhost.com"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		hasErr bool
		fn     func(ctx context.Context) *TipRank
	}{
		{
			name:   "success",
			hasErr: false,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, model.URLInsightBase+model.URLInsightTipRankAnalystsExpertPictureStore, nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			hasErr: true,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, model.URLInsightBase+model.URLInsightTipRankAnalystsExpertPictureStore, nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			hasErr: true,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, model.URLInsightBase+model.URLInsightTipRankAnalystsExpertPictureStore, nil).Return(nil, testErr)

				return &TipRank{
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
			resp, err := r.AnalystsExpertPictureStore(ctx)

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

func TestTipRank_SupportedTickers(t *testing.T) {
	var (
		expResp = model.SupportedTickersResponse{
			Tickers: []string{"AAPL", "FB"},
		}

		rawResponse = `{"tickers": ["AAPL", "FB"]}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		hasErr bool
		fn     func(ctx context.Context) *TipRank
	}{
		{
			name:   "success",
			hasErr: false,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, model.URLInsightBase+model.URLInsightTipRankSupportedTickers, nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			hasErr: true,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, model.URLInsightBase+model.URLInsightTipRankSupportedTickers, nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			hasErr: true,
			fn: func(ctx context.Context) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, model.URLInsightBase+model.URLInsightTipRankSupportedTickers, nil).Return(nil, testErr)

				return &TipRank{
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
			resp, err := r.SupportedTickers(ctx)

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

func TestTipRank_GeneralStockUpdates(t *testing.T) {
	var (
		expResp = model.GeneralStockUpdatesResponse{
			Stocks: []string{"AAPL", "FB"},
		}

		rawResponse = `{"stocks": ["AAPL", "FB"]}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name             string
		hasErr           bool
		utcTime, details string
		fn               func(ctx context.Context, utcTime, details string) *TipRank
	}{
		{
			name:    "success",
			hasErr:  false,
			utcTime: "2021-01-01",
			details: "full",
			fn: func(ctx context.Context, utcTime, details string) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?utc_time=%s&details=%s", model.URLInsightBase+model.URLInsightTipRankGeneralStockUpdates, utcTime, details), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:    "err/unmarshal",
			hasErr:  true,
			utcTime: "2021-01-01",
			details: "full",
			fn: func(ctx context.Context, utcTime, details string) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?utc_time=%s&details=%s", model.URLInsightBase+model.URLInsightTipRankGeneralStockUpdates, utcTime, details), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:    "err/cli",
			hasErr:  true,
			utcTime: "2021-01-01",
			details: "full",
			fn: func(ctx context.Context, utcTime, details string) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?utc_time=%s&details=%s", model.URLInsightBase+model.URLInsightTipRankGeneralStockUpdates, utcTime, details), nil).Return(nil, testErr)

				return &TipRank{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			r := tc.fn(ctx, tc.utcTime, tc.details)
			resp, err := r.GeneralStockUpdates(ctx, tc.utcTime, tc.details)

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

func TestTipRank_InsidersOverview(t *testing.T) {
	var (
		expResp = model.InsidersOverviewResponse{
			Name:  "Del",
			Stars: 12,
		}

		rawResponse = `{"Name": "Del", "stars": 12}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name      string
		hasErr    bool
		expertUID string
		fn        func(ctx context.Context, expertUID string) *TipRank
	}{
		{
			name:      "success",
			hasErr:    false,
			expertUID: "1234567890",
			fn: func(ctx context.Context, expertUID string) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?expert_uid=%s", model.URLInsightBase+model.URLInsightTipRankInsidersOverview, expertUID), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:      "err/unmarshal",
			hasErr:    true,
			expertUID: "1234567890",
			fn: func(ctx context.Context, expertUID string) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?expert_uid=%s", model.URLInsightBase+model.URLInsightTipRankInsidersOverview, expertUID), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:      "err/cli",
			hasErr:    true,
			expertUID: "1234567890",
			fn: func(ctx context.Context, expertUID string) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?expert_uid=%s", model.URLInsightBase+model.URLInsightTipRankInsidersOverview, expertUID), nil).Return(nil, testErr)

				return &TipRank{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			r := tc.fn(ctx, tc.expertUID)
			resp, err := r.InsidersOverview(ctx, tc.expertUID)

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

func TestTipRank_InsidersBestPerformingExperts(t *testing.T) {
	var (
		expResp = []model.InsidersBestPerformingExpertsResponse{
			{
				Name:                   "Asdf",
				Rank:                   12,
				NumberOfRankedInsiders: 2,
				ExcessReturn:           4,
				NumOfStars:             45,
			},
		}

		rawResponse = `[{"name":"Asdf","rank":12,"number_of_ranked_insiders":2,"excess_return":4,"num_of_stars":45}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		hasErr bool
		num    int
		fn     func(ctx context.Context, num int) *TipRank
	}{
		{
			name:   "success",
			hasErr: false,
			num:    112,
			fn: func(ctx context.Context, num int) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?num=%d", model.URLInsightBase+model.URLInsightTipRankInsidersBestPerformingExperts, num), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			hasErr: true,
			num:    112,
			fn: func(ctx context.Context, num int) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?num=%d", model.URLInsightBase+model.URLInsightTipRankInsidersBestPerformingExperts, num), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			hasErr: true,
			num:    112,
			fn: func(ctx context.Context, num int) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?num=%d", model.URLInsightBase+model.URLInsightTipRankInsidersBestPerformingExperts, num), nil).Return(nil, testErr)

				return &TipRank{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			r := tc.fn(ctx, tc.num)
			resp, err := r.InsidersBestPerformingExperts(ctx, tc.num)

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

func TestTipRank_InsidersLiveFeed(t *testing.T) {
	var (
		expResp = []model.InsidersLiveFeedResponse{
			{
				Name:          "Asdf",
				Rank:          12,
				NumOfStars:    45,
				TransactionID: 123,
			},
		}

		rawResponse = `[{"name":"Asdf","rank":12,"num_of_stars":45,"transaction_id":123}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		hasErr bool
		num    int
		sort   string
		fn     func(ctx context.Context, num int, sort string) *TipRank
	}{
		{
			name:   "success",
			hasErr: false,
			num:    112,
			sort:   "desc",
			fn: func(ctx context.Context, num int, sort string) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?num=%d&sort=%s", model.URLInsightBase+model.URLInsightTipRankInsidersLiveFeed, num, sort), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			hasErr: true,
			num:    112,
			sort:   "desc",
			fn: func(ctx context.Context, num int, sort string) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?num=%d&sort=%s", model.URLInsightBase+model.URLInsightTipRankInsidersLiveFeed, num, sort), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			hasErr: true,
			num:    112,
			sort:   "desc",
			fn: func(ctx context.Context, num int, sort string) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?num=%d&sort=%s", model.URLInsightBase+model.URLInsightTipRankInsidersLiveFeed, num, sort), nil).Return(nil, testErr)

				return &TipRank{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			r := tc.fn(ctx, tc.num, tc.sort)
			resp, err := r.InsidersLiveFeed(ctx, tc.num, tc.sort)

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

func TestTipRank_HedgeFundsBestPerformingExperts(t *testing.T) {
	var (
		expResp = []model.HedgeFundsBestPerformingExpertsResponse{
			{
				Name:          "Asdf",
				Rank:          12,
				NumOfStars:    45,
				HedgeFundName: "Aopl",
			},
		}

		rawResponse = `[{"name":"Asdf","rank":12,"num_of_stars":45,"hedge_fund_name":"Aopl"}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		hasErr bool
		num    int
		fn     func(ctx context.Context, num int) *TipRank
	}{
		{
			name:   "success",
			hasErr: false,
			num:    112,
			fn: func(ctx context.Context, num int) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?num=%d", model.URLInsightBase+model.URLInsightTipRankHedgefundsBestPerformingExperts, num), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:   "err/unmarshal",
			hasErr: true,
			num:    112,
			fn: func(ctx context.Context, num int) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?num=%d", model.URLInsightBase+model.URLInsightTipRankHedgefundsBestPerformingExperts, num), nil).Return(httpResponse, nil)

				return &TipRank{
					cli: cli,
				}
			},
		},
		{
			name:   "err/cli",
			hasErr: true,
			num:    112,
			fn: func(ctx context.Context, num int) *TipRank {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?num=%d", model.URLInsightBase+model.URLInsightTipRankHedgefundsBestPerformingExperts, num), nil).Return(nil, testErr)

				return &TipRank{
					cli: cli,
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			r := tc.fn(ctx, tc.num)
			resp, err := r.HedgeFundsBestPerformingExperts(ctx, tc.num)

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
