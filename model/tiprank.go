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
	Num     uint                            `json:"num,default:3"` // The number of analysts to return
	Sort    LatestAnalystRatingsOnStockSort `json:"sort"`          // Sort order of the analyst ratings 		Enums(date, rank)
	Months  uint                            `json:"months"`        // Get analyst ratings from the past X months
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

type LiveFeedRequest struct {
	Num             int             `json:"num"`               // Return the latest X ratings
	BookmarkID      int             `json:"bookmark_id"`       // Get analyst ratings from the past X months
	QuoteAccessType QuoteAccessType `json:"quote_access_type"` // Type of quote to load 		Enums(nbbo, delayed, cboe, bats, realtime, realtime_snapshot)

}

type QuoteAccessType string

const (
	QuoteAccessTypeNBBO             QuoteAccessType = "nbbo"
	QuoteAccessTypeDelayed          QuoteAccessType = "delayed"
	QuoteAccessTypeCBOE             QuoteAccessType = "cboe"
	QuoteAccessTypeBATS             QuoteAccessType = "bats"
	QuoteAccessTypeRealtime         QuoteAccessType = "realtime"
	QuoteAccessTypeRealtimeSnapshot QuoteAccessType = "realtime_snapshot"
)

type LiveFeedResponse struct {
	AnalystName                      string       `json:"analyst_name"`
	AnalystRank                      int          `json:"analyst_rank"`
	NumberOfRankedExperts            int          `json:"number_of_ranked_experts"`
	SuccessRate                      float32      `json:"success_rate"`
	ExcessReturn                     float32      `json:"excess_return"`
	TotalRecommendations             int          `json:"total_recommendations"`
	GoodRecommendations              int          `json:"good_recommendations"`
	NumOfStars                       float32      `json:"num_of_stars"`
	ArticleTitle                     string       `json:"article_title"`
	ArticleSite                      string       `json:"article_site"`
	Action                           string       `json:"action"`
	ExpertUID                        string       `json:"expert_uid"`
	ExpertPictureURL                 string       `json:"expert_picture_url"`
	URL                              string       `json:"url"`
	Recommendation                   string       `json:"recommendation"`
	RecommendationDate               string       `json:"recommendation_date"`
	FirmName                         string       `json:"firm_name"`
	StockTicker                      string       `json:"stock_ticker"`
	CompanyName                      string       `json:"company_name"`
	PriceTarget                      float32      `json:"price_target"`
	OperationID                      int          `json:"operation_id"`
	UniqueOperationID                string       `json:"unique_operation_id"`
	QuoteTitle                       string       `json:"quote_title"`
	TimeAgo                          string       `json:"time_ago"`
	ConvertedPriceTarget             float32      `json:"converted_price_target"`
	ConvertedPriceTargetCurrency     float32      `json:"converted_price_target_currency"`
	ConvertedPriceTargetCurrencyCode string       `json:"converted_price_target_currency_code"`
	Sector                           string       `json:"sector"`
	Timestamp                        string       `json:"timestamp"`
	AddedOnTimestamp                 string       `json:"added_on_timestamp"`
	MarketCap                        float32      `json:"market_cap"`
	Quote                            TiprankQuote `json:"quote"`
}

type TiprankQuote struct {
	AskPx              float64 `json:"ask_px"`
	AskSz              int64   `json:"ask_sz"`
	AvgVol30D          int64   `json:"avg_vol_30d"`
	BidPx              float64 `json:"bid_px"`
	BidSz              int64   `json:"bid_sz"`
	ClosingPx          float64 `json:"closing_px"`
	Column             int64   `json:"column"`
	ComShrsOut         int64   `json:"com_shrs_out"`
	CompDesc           string  `json:"comp_desc"`
	CompName           string  `json:"comp_name"`
	Condition          string  `json:"condition"`
	DivAmt             float64 `json:"div_amt"`
	DivDateEx          string  `json:"div_date_ex"`
	DivYield           float64 `json:"div_yield"`
	High52WPx          float64 `json:"high52w_px"`
	HighPx             float64 `json:"high_px"`
	InfoTime           string  `json:"info_time"`
	IssueMarket        string  `json:"issue_market"`
	IssureMarket       string  `json:"issure_market"`
	LastPx             float64 `json:"last_px"`
	Low52WPx           float64 `json:"low_52w_px"`
	LowPx              float64 `json:"low_px"`
	MktCap             float64 `json:"mkt_cap"`
	OpenPx             float64 `json:"open_px"`
	PayDate            string  `json:"pay_date"`
	PeRatio            float64 `json:"pe_ratio"`
	PostAskPx          float64 `json:"post_ask_px"`
	PostAskSz          int64   `json:"post_ask_sz"`
	PostBidPx          float64 `json:"post_bid_px"`
	PostBidSz          int64   `json:"post_bid_sz"`
	PostLastPx         float64 `json:"post_last_px"`
	PostQuoteTime      string  `json:"post_quote_time"`
	PostTradeTime      string  `json:"post_trade_time"`
	PreAskPx           float64 `json:"pre_ask_px"`
	PreAskSz           int64   `json:"pre_ask_sz"`
	PreBidPx           float64 `json:"pre_bid_px"`
	PreBidSz           int64   `json:"pre_bid_sz"`
	PreLastPx          float64 `json:"pre_last_px"`
	PreQuoteTime       string  `json:"pre_quote_time"`
	PreTradeTime       string  `json:"pre_trade_time"`
	Symbol             string  `json:"symbol"`
	QuoteTime          string  `json:"quote_time"`
	TradePx            float64 `json:"trade_px"`
	TradeRegion        string  `json:"trade_region"`
	TradeSz            int64   `json:"trade_sz"`
	TradeTime          string  `json:"trade_time"`
	Volume             int64   `json:"volume"`
	YestClosePx        float64 `json:"yest_close_px"`
	PriceChange        float64 `json:"price_change"`
	PriceChangePercent float64 `json:"price_change_percent"`
}

type TrendingStocksRequest struct {
	Num             int             `json:"num"`               // "Return the most trending X stocks"
	Filter          FilterType      `json:"filter"`            // "Include ratings of analysts and/or bloggers" 	Enums(analysts, bloggers)
	Period          int             `json:"period"`            // "Trending over the last X days"					Enums(3, 7, 30, 90)
	TrendingType    TrendingType    `json:"trending_type"`     // "Trend Type"									Enums(most-rated, best-rated, worst-rated)
	QuoteAccessType QuoteAccessType `json:"quote_access_type"` // Type of quote to load 		Enums(nbbo, delayed, cboe, bats, realtime, realtime_snapshot)
}

type FilterType string

const (
	FilterAnalysts FilterType = "analysts"
	FilterBloggers FilterType = "bloggers"
)

type TrendingType string

const (
	TrendingTypeMostRated  TrendingType = "most-rated"
	TrendingTypeBestRated  TrendingType = "best-rated"
	TrendingTypeWorstRated TrendingType = "worst-rated"
)

type TrendingStocksResponse struct {
	Ticker             string       `json:"ticker"`
	Popularity         int          `json:"popularity"`
	Sentiment          int          `json:"sentiment"`
	Sector             string       `json:"sector"`
	Hold               float32      `json:"hold"`
	Buy                float32      `json:"buy"`
	Sell               float32      `json:"sell"`
	AveragePriceTarget float32      `json:"average_price_target"`
	Consensus          string       `json:"consensus"`
	LastRatingDate     string       `json:"last_rating_date"`
	MarketCap          int64        `json:"market_cap"`
	Quote              TiprankQuote `json:"quote"`
}

type PortfoliosRequest struct {
	ExpertUID       string          `json:"expert_uid"` //
	Number          int             `json:"number"`
	QuoteAccessType QuoteAccessType `json:"quote_access_type"` // Type of quote to load 		Enums(nbbo, delayed, cboe, bats, realtime, realtime_snapshot)
}

type PortfoliosResponse struct {
	StockTicker  string       `json:"stock_ticker"`
	CompanyName  string       `json:"company_name"`
	Operation    string       `json:"operation"`
	ExcessReturn float64      `json:"excess_return"`
	OpenDate     string       `json:"open_date"`
	Status       string       `json:"status"`
	OpenURL      interface{}  `json:"open_url,omitempty"`
	CloseURL     interface{}  `json:"close_url,omitempty"`
	CloseDate    interface{}  `json:"close_date,omitempty"`
	PriceTarget  float32      `json:"price_target"`
	Sector       string       `json:"sector"`
	Quote        TiprankQuote `json:"quote"`
}

type AnalystProfileResponse struct {
	AnalystName           string             `json:"analyst_name"`
	FirmName              string             `json:"firm_name"`
	ExpertPictureURL      string             `json:"expert_picture_url"`
	ExpertFullPictureURL  string             `json:"expert_full_picture_url"`
	SuccessRate           string             `json:"success_rate"`
	ExcessReturn          string             `json:"excess_return"`
	AnalystRank           int                `json:"analyst_rank"`
	NumberOfRankedExperts int                `json:"number_of_ranked_experts"`
	TotalRecommendations  int                `json:"total_recommendations"`
	GoodRecommendations   int                `json:"good_recommendations"`
	NumOfStars            float32            `json:"num_of_stars"`
	MainSector            string             `json:"main_sector"`
	RatingDistribution    RatingDistribution `json:"rating_distribution"`
}

type RatingDistribution struct {
	BuyPercent  float32  `json:"buy_percent"`
	SellPercent float32  `json:"sell_percent"`
	HoldPercent float32  `json:"hold_percent"`
	Countries   []string `json:"countries"`
}

type SectorConsensusResponse struct {
	Sectors []*SectorConsensus `json:"sectors"`
}

type SectorConsensus struct {
	Sector       string `json:"sector"`
	StrongBuy    int    `json:"strong_buy"`
	ModerateBuy  int    `json:"moderate_buy"`
	Neutral      int    `json:"neutral"`
	ModerateSell int    `json:"moderate_sell"`
	StrongSell   int    `json:"strong_sell"`
	Order        int    `json:"order"`
}

type BestPerformingExpertsResponse struct {
	AnalystName           string                         `json:"analyst_name"`
	FirmName              string                         `json:"firm_name"`
	ExpertUID             string                         `json:"expert_uid"`
	ExpertPictureURL      string                         `json:"expert_picture_url"`
	AnalystRank           int                            `json:"analyst_rank"`
	NumberOfRankedExperts int                            `json:"number_of_ranked_experts"`
	SuccessRate           float32                        `json:"success_rate"`
	ExcessReturn          float32                        `json:"excess_return"`
	TotalRecommendations  int                            `json:"total_recommendations"`
	GoodRecommendations   int                            `json:"good_recommendations"`
	Distribution          *PerformingExpertsDistribution `json:"distribution"`
	NumOfStars            float32                        `json:"num_of_stars"`
	Ratings               []*PerformingExpertsRatings    `json:"ratings"`
	Sector                string                         `json:"sector"`
	FirmAnalystsCount     int                            `json:"firm_analysts_count"`
}

type PerformingExpertsDistribution struct {
	Buy  float32 `json:"buy"`
	Hold float32 `json:"hold"`
	Sell float32 `json:"sell"`
}

type PerformingExpertsRatings struct {
	Ticker      string  `json:"ticker"`
	CompanyName string  `json:"company_name"`
	Rating      string  `json:"rating"`
	PriceTarget float32 `json:"price_target"`
	Date        string  `json:"date"`
}

type AnalystsExpertPictureStoreResponse struct {
	URL string `json:"url"`
}
