package model

type AnalystConsensusRequest struct {
	Tickers         []string `json:"tickers"` // Comma-separated tickers
	IncludeAnalysts bool     `json:"include_analysts,omitempty"`
}

type AnalystConsensusResponse struct {
	LowPriceTarget          float32 `json:"low_price_target"`
	HighPriceTarget         float32 `json:"high_price_target"`
	PriceTarget             float32 `json:"price_target"`
	Buy                     float32 `json:"buy"`
	Sell                    float32 `json:"sell"`
	Hold                    float32 `json:"hold"`
	Consensus               string  `json:"consensus"`
	PriceTargetUpside       float32 `json:"price_target_upside"`
	Ticker                  string  `json:"ticker"`
	CompanyName             string  `json:"company_name"`
	PriceTargetCurrencyCode string  `json:"price_target_currency_code"`
	TotalAnalysts           float32 `json:"total_analysts"`
}

type LatestAnalystRatingsOnStockRequest struct {
	Tickers []string                        `json:"tickers"`       // Comma-separated tickers
	Num     uint                            `json:"num,default:3"` // The number of analysts to return (OPTIONAL)
	Sort    LatestAnalystRatingsOnStockSort `json:"sort"`          // Sort order of the analyst ratings 		Enums(date, rank)	(OPTIONAL)
	Months  uint                            `json:"months"`        // Get analyst ratings from the past X months	(OPTIONAL)
}

type LatestAnalystRatingsOnStockSort string

const (
	LatestAnalystRatingsOnStockSortDate LatestAnalystRatingsOnStockSort = "date"
	LatestAnalystRatingsOnStockSortRank LatestAnalystRatingsOnStockSort = "rank"
)

type LatestAnalystRatingsOnStockResponse struct {
	AnalystName                      string  `json:"analyst_name"`
	FirmName                         string  `json:"firm_name"`
	Recommendation                   string  `json:"recommendation"`
	RecommendationDate               string  `json:"recommendation_date"`
	ExpertUID                        string  `json:"expert_uid"`
	URL                              string  `json:"url"`
	ExpertPictureURL                 string  `json:"expert_picture_url"`
	ExpertFullPictureURL             string  `json:"expert_full_picture_url"`
	AnalystRank                      int     `json:"analyst_rank"`
	NumberOfRankedExperts            int     `json:"number_of_ranked_experts"`
	SuccessRate                      float32 `json:"success_rate"`
	ExcessReturn                     float32 `json:"excess_return"`
	TotalRecommendations             int     `json:"total_recommendations"`
	GoodRecommendations              int     `json:"good_recommendations"`
	NumOfStars                       float32 `json:"num_of_stars"`
	StockSuccessRate                 float32 `json:"stock_success_rate"`
	StockAvgReturn                   float32 `json:"stock_avg_return"`
	ArticleTitle                     string  `json:"article_title"`
	ArticleSite                      string  `json:"article_site"`
	ArticleQuote                     string  `json:"article_quote"`
	PriceTarget                      float32 `json:"price_target"`
	PriceTargetCurrency              float32 `json:"price_target_currency"`
	PriceTargetCurrencyCode          string  `json:"price_target_currency_code"`
	ConvertedPriceTarget             float32 `json:"converted_price_target"`
	ConvertedPriceTargetCurrency     float32 `json:"converted_price_target_currency"`
	ConvertedPriceTargetCurrencyCode string  `json:"converted_price_target_currency_code"`
	Ticker                           string  `json:"ticker"`
	Timestamp                        string  `json:"timestamp"`
	AnalystAction                    string  `json:"analyst_action"`
	StockTotalRecommendations        float32 `json:"stock_total_recommendations"`
	StockGoodRecommendations         float32 `json:"stock_good_recommendations"`
}
