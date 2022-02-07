package orbis_api

type LiveFeedResponse []LiveFeedResponseElement
type GetAnalystResponse []GetAnalyticResponseElement
type GetTopAnalystResponse []GetTopAnalystElement
type RatingBySymbolResponse []RatingBySymbolResponseElement
type InsiderTransactionsResponse []InsiderTransactionsResponseElement
type OverviewResponse []OverviewResponseElement
type GetTopStocksResponse []GetTopStocksResponseElement
type TrendingResponse []TrendingResponseElement

type LiveFeedResponseElement struct {
	AnalystName           string        `json:"analystName"`
	FirmName              string        `json:"firmName"`
	StockTicker           string        `json:"stockTicker"`
	CompanyName           string        `json:"companyName"`
	Recommendation        string        `json:"recommendation"`
	ExpertUID             string        `json:"expertUID"`
	ExpertPictureURL      string        `json:"expertPictureUrl"`
	ExpertPictureFullURL  string        `json:"expertPictureFullUrl"`
	URL                   string        `json:"url"`
	RecommendationDate    string        `json:"recommendationDate"`
	Quote                 QuoteLiveFeed `json:"quote"`
	SuccessRate           float64       `json:"successRate"`
	ExcessReturn          float64       `json:"excessReturn"`
	NumOfStars            float64       `json:"numOfStars"`
	PriceTarget           int64         `json:"priceTarget"`
	NumberOfRankedExperts int64         `json:"numberOfRankedExperts"`
	TotalRecommendations  int64         `json:"totalRecommendations"`
	GoodRecommendations   int64         `json:"goodRecommendations"`
	AnalystRank           int64         `json:"analystRank"`
	Symbol                string        `json:"symbol"`
	PhotoAvailable        bool          `json:"photoAvailable"`
	NumOfStarsWhole       int64         `json:"numOfStarsWhole"`
}

type QuoteLiveFeed struct {
	Symbol              string  `json:"symbol"`
	CompanyName         string  `json:"companyName"`
	Isin                string  `json:"isin"`
	Cusip               string  `json:"cusip"`
	Country             string  `json:"country"`
	Mic                 string  `json:"mic"`
	Currency            string  `json:"currency"`
	Source              string  `json:"source"`
	IdentType           string  `json:"identType"`
	Exchange            string  `json:"exchange"`
	UpdateTime          string  `json:"updateTime"`
	TradeTime           string  `json:"tradeTime"`
	YesterdayClose      float64 `json:"yesterdayClose"`
	OpeningPrice        int64   `json:"openingPrice"`
	DayHigh             int64   `json:"dayHigh"`
	DayLow              float64 `json:"dayLow"`
	Ask                 float64 `json:"ask"`
	LastPrice           float64 `json:"lastPrice"`
	Change              float64 `json:"change"`
	High52Week          float64 `json:"high52week"`
	Low52Week           float64 `json:"low52week"`
	PriceEarningRatio   float64 `json:"priceEarningRatio"`
	ChangePercent       float64 `json:"changePercent"`
	MarketCap           float64 `json:"marketCap"`
	CalendarYearHigh    float64 `json:"calendarYearHigh"`
	CalendarYearLow     float64 `json:"calendarYearLow"`
	YesterdayVolume     int64   `json:"yesterdayVolume"`
	AskSize             int64   `json:"askSize"`
	Volume              int64   `json:"volume"`
	AverageVolume30     int64   `json:"averageVolume30"`
	Indicator           int64   `json:"indicator"`
	BidTick             int64   `json:"bidTick"`
	Precision           int64   `json:"precision"`
	Delayed             bool    `json:"delayed"`
	SymbolMarket        string  `json:"symbolMarket"`
	MarketDescription   string  `json:"marketDescription"`
	Trailing12MonthsEps float64 `json:"trailing12MonthsEps"`
}

type GetAnalyticResponseElement struct {
	StockTicker  string        `json:"stockTicker"`
	Quote        QuoteLiveFeed `json:"quote"`
	Status       string        `json:"status"`
	OpenDate     string        `json:"openDate"`
	Operation    string        `json:"operation"`
	ExcessReturn float64       `json:"excessReturn"`
	PriceTarget  int64         `json:"priceTarget"`
	Symbol       string        `json:"symbol"`
}

type GetTopAnalystElement struct {
	AnalystName           string  `json:"analystName"`
	FirmName              string  `json:"firmName"`
	ExpertUID             string  `json:"expertUID"`
	ExpertPictureURL      string  `json:"expertPictureUrl"`
	SuccessRate           float64 `json:"successRate"`
	ExcessReturn          float64 `json:"excessReturn"`
	NumOfStars            int64   `json:"numOfStars"`
	NumberOfRankedExperts int64   `json:"numberOfRankedExperts"`
	TotalRecommendations  int64   `json:"totalRecommendations"`
	GoodRecommendations   int64   `json:"goodRecommendations"`
	AnalystRank           int64   `json:"analystRank"`
	PhotoAvailable        bool    `json:"photoAvailable"`
	NumOfStarsWhole       int64   `json:"numOfStarsWhole"`
}

type RatingBySymbolResponseElement struct {
	AnalystName           string  `json:"analystName"`
	FirmName              string  `json:"firmName"`
	Recommendation        string  `json:"recommendation"`
	ExpertUID             string  `json:"expertUID"`
	URL                   string  `json:"url"`
	RecommendationDate    string  `json:"recommendationDate"`
	SuccessRate           float64 `json:"successRate"`
	ExcessReturn          float64 `json:"excessReturn"`
	NumOfStars            float64 `json:"numOfStars"`
	PriceTarget           int64   `json:"priceTarget"`
	NumberOfRankedExperts int64   `json:"numberOfRankedExperts"`
	TotalRecommendations  int64   `json:"totalRecommendations"`
	GoodRecommendations   int64   `json:"goodRecommendations"`
	AnalystRank           int64   `json:"analystRank"`
	NumOfStarsWhole       int64   `json:"numOfStarsWhole"`
}

type InsiderResponse struct {
	Aapl Aapl `json:"AAPL"`
}

type Aapl struct {
	YearlyInsiderTransactions []YearlyInsiderTransaction `json:"yearlyInsiderTransactions"`
	ConfidenceSignal          ConfidenceSignal           `json:"confidenceSignal"`
	Transactions              []Transaction              `json:"transactions"`
	DiscretionaryTransactions int64                      `json:"discretionaryTransactions"`
	UninformativeTransactions int64                      `json:"uninformativeTransactions"`
}

type ConfidenceSignal struct {
	Score       string  `json:"score"`
	SectorScore float64 `json:"sectorScore"`
	StockScore  float64 `json:"stockScore"`
}

type Transaction struct {
	InsiderName string  `json:"insiderName"`
	ExpertUid   string  `json:"expertUid"`
	Position    string  `json:"position"`
	Stars       string  `json:"stars"`
	Transaction string  `json:"transaction"`
	Amount      float64 `json:"amount"`
	Date        string  `json:"date"`
}

type YearlyInsiderTransaction struct {
	Month      int64  `json:"month"`
	Year       int64  `json:"year"`
	SellCount  *int64 `json:"sellCount,omitempty"`
	SellAmount *int64 `json:"sellAmount,omitempty"`
	BuyCount   *int64 `json:"buyCount,omitempty"`
	BuyAmount  *int64 `json:"buyAmount,omitempty"`
}

type InsiderTransactionsResponseElement struct {
	Symbol      string  `json:"symbol"`
	InsiderName string  `json:"insiderName"`
	ExpertUid   string  `json:"expertUid"`
	Position    string  `json:"position"`
	Stars       string  `json:"stars"`
	Transaction string  `json:"transaction"`
	Amount      float64 `json:"amount"`
	Date        string  `json:"date"`
}

type OverviewResponseElement struct {
	Ticker            string  `json:"ticker"`
	PriceTarget       float64 `json:"priceTarget"`
	PriceTargetUpside float64 `json:"priceTargetUpside"`
	Buy               int64   `json:"buy"`
	Hold              int64   `json:"hold"`
	Consensus         string  `json:"consensus"`
	Symbol            string  `json:"symbol"`
}

type GetTopStocksResponseElement struct {
	Ticker         string         `json:"ticker"`
	CompanyName    string         `json:"companyName"`
	Quote          QuoteTopStocks `json:"quote"`
	Consensus      string         `json:"consensus"`
	PriceTarget    float64        `json:"priceTarget"`
	LastRatingDate string         `json:"lastRatingDate"`
	Symbol         string         `json:"symbol"`
}

type QuoteTopStocks struct {
	Symbol              string  `json:"symbol"`
	CompanyName         string  `json:"companyName"`
	Isin                string  `json:"isin"`
	Cusip               string  `json:"cusip"`
	Country             string  `json:"country"`
	Mic                 string  `json:"mic"`
	Currency            string  `json:"currency"`
	Source              string  `json:"source"`
	IdentType           string  `json:"identType"`
	Exchange            string  `json:"exchange"`
	UpdateTime          string  `json:"updateTime"`
	TradeTime           string  `json:"tradeTime"`
	DividendPaymentDate string  `json:"dividendPaymentDate"`
	DividendExDate      string  `json:"dividendExDate"`
	YesterdayClose      int64   `json:"yesterdayClose"`
	OpeningPrice        float64 `json:"openingPrice"`
	DayHigh             float64 `json:"dayHigh"`
	DayLow              float64 `json:"dayLow"`
	Ask                 float64 `json:"ask"`
	Bid                 int64   `json:"bid"`
	LastPrice           float64 `json:"lastPrice"`
	Change              float64 `json:"change"`
	High52Week          float64 `json:"high52week"`
	Low52Week           float64 `json:"low52week"`
	ChangePercent       float64 `json:"changePercent"`
	MarketCap           float64 `json:"marketCap"`
	CalendarYearHigh    float64 `json:"calendarYearHigh"`
	CalendarYearLow     float64 `json:"calendarYearLow"`
	DividendYield       float64 `json:"dividendYield"`
	DividendAmount      float64 `json:"dividendAmount"`
	DividendRate        int64   `json:"dividendRate"`
	YesterdayVolume     int64   `json:"yesterdayVolume"`
	AskSize             int64   `json:"askSize"`
	BidSize             int64   `json:"bidSize"`
	Volume              int64   `json:"volume"`
	AverageVolume30     int64   `json:"averageVolume30"`
	Indicator           int64   `json:"indicator"`
	BidTick             int64   `json:"bidTick"`
	Precision           int64   `json:"precision"`
	Delayed             bool    `json:"delayed"`
	MarketDescription   string  `json:"marketDescription"`
	SymbolMarket        string  `json:"symbolMarket"`
}

type ConsensusByBloggerResponse struct {
	BearishRatio      float64             `json:"bearishRatio"`
	BlogsDistribution []BlogsDistribution `json:"blogsDistribution"`
	BullishRatio      float64             `json:"bullishRatio"`
	Consensus         string              `json:"consensus"`
	NumOfBloggers     int64               `json:"numOfBloggers"`
	SectorBullRatio   float64             `json:"sectorBullRatio"`
	SectorName        string              `json:"sectorName"`
}

type BlogsDistribution struct {
	Percentage float64 `json:"percentage"`
	Site       string  `json:"site"`
}

type TrendingResponseElement struct {
	Ticker     string        `json:"ticker"`
	Sector     string        `json:"sector"`
	Quote      QuoteTrending `json:"quote"`
	Consensus  string        `json:"consensus"`
	Popularity int64         `json:"popularity"`
	Sentiment  int64         `json:"sentiment"`
	Hold       int64         `json:"hold"`
	Buy        int64         `json:"buy"`
	Symbol     string        `json:"symbol"`
	Sell       *int64        `json:"sell,omitempty"`
}

type QuoteTrending struct {
	Symbol              string   `json:"symbol"`
	CompanyName         string   `json:"companyName"`
	Isin                string   `json:"isin"`
	Cusip               string   `json:"cusip"`
	Country             string   `json:"country"`
	Mic                 string   `json:"mic"`
	Currency            string   `json:"currency"`
	Source              string   `json:"source"`
	IdentType           string   `json:"identType"`
	Exchange            string   `json:"exchange"`
	UpdateTime          string   `json:"updateTime"`
	TradeTime           string   `json:"tradeTime"`
	YesterdayClose      float64  `json:"yesterdayClose"`
	OpeningPrice        float64  `json:"openingPrice"`
	DayHigh             float64  `json:"dayHigh"`
	DayLow              float64  `json:"dayLow"`
	Ask                 float64  `json:"ask"`
	Bid                 float64  `json:"bid"`
	LastPrice           float64  `json:"lastPrice"`
	Change              float64  `json:"change"`
	High52Week          float64  `json:"high52week"`
	Low52Week           float64  `json:"low52week"`
	ChangePercent       float64  `json:"changePercent"`
	MarketCap           float64  `json:"marketCap"`
	CalendarYearHigh    float64  `json:"calendarYearHigh"`
	CalendarYearLow     float64  `json:"calendarYearLow"`
	Beta                float64  `json:"beta"`
	YesterdayVolume     int64    `json:"yesterdayVolume"`
	AskSize             int64    `json:"askSize"`
	BidSize             int64    `json:"bidSize"`
	Volume              int64    `json:"volume"`
	AverageVolume30     int64    `json:"averageVolume30"`
	Indicator           int64    `json:"indicator"`
	BidTick             int64    `json:"bidTick"`
	Precision           int64    `json:"precision"`
	Delayed             bool     `json:"delayed"`
	MarketDescription   string   `json:"marketDescription"`
	SymbolMarket        string   `json:"symbolMarket"`
	PriceEarningRatio   *float64 `json:"priceEarningRatio,omitempty"`
	Trailing12MonthsEps *float64 `json:"trailing12MonthsEps,omitempty"`
}

type HedgeFundsResponse struct {
	HoldingData           HoldingData            `json:"holdingData"`
	InstitutionalHoldings []InstitutionalHolding `json:"institutionalHoldings"`
	LastQSharesTraded     int64                  `json:"lastQSharesTraded"`
	SignalData            SignalData             `json:"signalData"`
}

type HoldingData struct {
	Holdings        []HoldingHedgeFund `json:"holdings"`
	TotalHedgeFunds int64              `json:"totalHedgeFunds"`
}

type HoldingHedgeFund struct {
	Date       string `json:"date"`
	SharesHeld int64  `json:"sharesHeld"`
}

type InstitutionalHolding struct {
	Action                   string `json:"action"`
	ExpertPictureURL         string `json:"expertPictureUrl"`
	ExpertPictureFullURL     string `json:"expertPictureFullUrl"`
	ExpertUID                string `json:"expertUID"`
	HedgeFundRank            int64  `json:"hedgeFundRank"`
	HoldingChange            int64  `json:"holdingChange"`
	InstitutionName          string `json:"institutionName"`
	ManagerName              string `json:"managerName"`
	NumberOfRankedHedgeFunds int64  `json:"numberOfRankedHedgeFunds"`
}

type SignalData struct {
	BasedOnNumHedgeFunds int64   `json:"basedOnNumHedgeFunds"`
	Rating               string  `json:"rating"`
	Sentiment            float64 `json:"sentiment"`
}

type StockSentimentResponse struct {
	Sentiment       string    `json:"sentiment"`
	StockSentiment  Sentiment `json:"stockSentiment"`
	SectorSentiment Sentiment `json:"sectorSentiment"`
	WordCloud       []string  `json:"wordCloud"`
	Buzz            Buzz      `json:"buzz"`
}

type Buzz struct {
	WeeklyAverage float64 `json:"weeklyAverage"`
	ThisWeek      int64   `json:"thisWeek"`
}

type Sentiment struct {
	Positive float64 `json:"positive"`
	Neutral  float64 `json:"neutral"`
	Negative float64 `json:"negative"`
}

type BloggerRatingsResponse []BloggerRatingsResponseElement

type BloggerRatingsResponseElement struct {
	BloggerName                            string  `json:"bloggerName"`
	StockSentimentResponseExpertPictureURL string  `json:"expertPictureUrl"`
	ExpertPictureFullURL                   string  `json:"expertPictureFullUrl"`
	ExpertUID                              string  `json:"expertUID"`
	FirmName                               string  `json:"firmName"`
	NumOfStars                             float64 `json:"numOfStars"`
	Recommendation                         string  `json:"recommendation"`
	RecommendationDate                     string  `json:"recommendationDate"`
	URL                                    string  `json:"url"`
	ExpertPictureURL                       string  `json:"expertPictureURL"`
}
