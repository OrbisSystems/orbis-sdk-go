package orbis_api

type MutualFundDetails []MutualFundDetail

type Equity []EquityElement

type Search []SearchElement

type Historical []HistoricalElement

type IntradayChartData []IntradayChartDatum

type TopStocks []TopStock

type SimilarStocks []string

type Overview []OverviewElement

type ConsensusBySector []ConsensusBySectorElement

type Trending []TrendingElement

type Defaults []string

type Types []string

type PaymentTypes []string

type Upcoming []UpcomingElement

type Performance []PerformanceElement

type OwnershipDetails []OwnershipDetail

type EarningsCalendar []EarningsCalendarElement

type EarningsCalendarForPortfolio []EarningsCalendarForPortfolioElement

type FundamentalTypes []string

type IndustryHeatMap []IndustryHeatMapElement

type IndustryList []IndustryListElement

type SymbolHeatMap []SymbolHeatMapElement

type UserPortfolioEarnings []UserPortfolioEarning

type UserPreferences []UserPreference

type MemberRealtimeBalances []MemberRealtimeBalance

type ListBankAccounts []ListBankAccount

type TransferStatus []TransferStatusElement

type TransferStatusModel []TransferStatusModelElement

type TransferStatusBranch []TransferStatusBranchElement

type AvailableDocumentTypes []string

type MutualFundDetail struct {
	EstimatedCashAmount     int64                   `json:"estimatedCashAmount"`
	CompanyName             string                  `json:"companyName"`
	Delayed                 bool                    `json:"delayed"`
	CalendarYearHigh        float64                 `json:"calendarYearHigh"`
	DividendYield           float64                 `json:"dividendYield"`
	Put                     bool                    `json:"put"`
	NavPrice                float64                 `json:"navPrice"`
	PriceEarningRatio       int64                   `json:"priceEarningRatio"`
	DayHigh                 int64                   `json:"dayHigh"`
	Nav                     int64                   `json:"nav"`
	SubMarket               int64                   `json:"subMarket"`
	AtTheMoney              bool                    `json:"atTheMoney"`
	DividendRate            float64                 `json:"dividendRate"`
	Level3BuyIRA            int64                   `json:"level3BuyIRA"`
	DayLow                  int64                   `json:"dayLow"`
	AskRegion               RealizedProfitLoss      `json:"askRegion"`
	Nbbo                    bool                    `json:"nbbo"`
	SymbolChange            bool                    `json:"symbolChange"`
	Vwap1                   int64                   `json:"vwap1"`
	DividendExDate          string                  `json:"dividendExDate"`
	MinInitialBuy           int64                   `json:"minInitialBuy"`
	High52Week              float64                 `json:"high52week"`
	Jumbo                   bool                    `json:"jumbo"`
	Mic                     string                  `json:"mic"`
	NotPermissioned         bool                    `json:"notPermissioned"`
	Increment               int64                   `json:"increment"`
	YesterdayVolume         int64                   `json:"yesterdayVolume"`
	AllowNewInvestors       bool                    `json:"allowNewInvestors"`
	IdentType               IdentType               `json:"identType"`
	ChangePercent           int64                   `json:"changePercent"`
	Currency                Currency                `json:"currency"`
	ContractSize            int64                   `json:"contractSize"`
	TradeSize               int64                   `json:"tradeSize"`
	MinIRA                  int64                   `json:"minIRA"`
	YesterdayClose          int64                   `json:"yesterdayClose"`
	Mini                    bool                    `json:"mini"`
	CorpAct                 bool                    `json:"corpAct"`
	BidSize                 int64                   `json:"bidSize"`
	Change                  int64                   `json:"change"`
	UpdateTime              string                  `json:"updateTime"`
	Call                    bool                    `json:"call"`
	Ask                     int64                   `json:"ask"`
	TotalCashAmount         int64                   `json:"totalCashAmount"`
	BidRegion               RealizedProfitLoss      `json:"bidRegion"`
	LastPrice               int64                   `json:"lastPrice"`
	IntradayValue           int64                   `json:"intradayValue"`
	Indicator               int64                   `json:"indicator"`
	Symbol                  string                  `json:"symbol"`
	Country                 Country                 `json:"country"`
	Cusip                   string                  `json:"cusip"`
	Paused                  bool                    `json:"paused"`
	Vwap                    int64                   `json:"vwap"`
	Precision               int64                   `json:"precision"`
	Low52Week               float64                 `json:"low52week"`
	Source                  string                  `json:"source"`
	DividendPaymentDate     string                  `json:"dividendPaymentDate"`
	AskSize                 int64                   `json:"askSize"`
	MutualFund              bool                    `json:"mutualFund"`
	MarketCap               int64                   `json:"marketCap"`
	Halted                  bool                    `json:"halted"`
	MinRIA                  int64                   `json:"minRIA"`
	MinBidSize              int64                   `json:"minBidSize"`
	NavChange               float64                 `json:"navChange"`
	DividendAmount          float64                 `json:"dividendAmount"`
	InTheMoney              bool                    `json:"inTheMoney"`
	Crypto                  bool                    `json:"crypto"`
	Volume                  int64                   `json:"volume"`
	TradeTime               string                  `json:"tradeTime"`
	PutOrCall               PutOrCall               `json:"putOrCall"`
	NavYesterdayPrice       float64                 `json:"navYesterdayPrice"`
	Trailing12MonthsEps     int64                   `json:"trailing12MonthsEps"`
	Exchange                string                  `json:"exchange"`
	TradePrice              int64                   `json:"tradePrice"`
	MinAskSize              int64                   `json:"minAskSize"`
	StrikePrice             int64                   `json:"strikePrice"`
	Option                  bool                    `json:"option"`
	BidTick                 int64                   `json:"bidTick"`
	TradeRegion             RealizedProfitLoss      `json:"tradeRegion"`
	OpenInterest            int64                   `json:"openInterest"`
	MinSubsequentBuy        int64                   `json:"minSubsequentBuy"`
	BidYield                int64                   `json:"bidYield"`
	OpeningPrice            int64                   `json:"openingPrice"`
	Fx                      bool                    `json:"fx"`
	MarketCapClassification MarketCapClassification `json:"marketCapClassification"`
	MinWrap                 int64                   `json:"minWrap"`
	SharesOutstanding       int64                   `json:"sharesOutstanding"`
	Beta                    int64                   `json:"beta"`
	CalendarYearLow         float64                 `json:"calendarYearLow"`
	ImbalanceVolume         int64                   `json:"imbalanceVolume"`
	SymbolMarket            string                  `json:"symbolMarket"`
	MinRetail               int64                   `json:"minRetail"`
	Bond                    bool                    `json:"bond"`
	MarketDescription       string                  `json:"marketDescription"`
	AverageVolume30         int64                   `json:"averageVolume30"`
	AskYield                int64                   `json:"askYield"`
	NotFound                bool                    `json:"notFound"`
	Bid                     int64                   `json:"bid"`
	DividentSoon            bool                    `json:"dividentSoon"`
	Isin                    string                  `json:"isin"`
}

type RealizedProfitLoss struct {
}

// Check Mutual Fund
//
// GET {{url}}/api/quotes/mf/check/{symbol}
//
// Checks if given mutual fund is available to trade and returns a JSON response with status
// of each provided identifier. Supports symbols or cusips separated by commas.
type CheckMutualFund struct {
	Agfqx        bool `json:"AGFQX"`
	The00141M523 bool `json:"00141M523"`
	The92204A793 bool `json:"92204A793"`
}

type EquityElement struct {
	Indicator               int64                   `json:"indicator"`
	Symbol                  string                  `json:"symbol"`
	Country                 Country                 `json:"country"`
	Cusip                   string                  `json:"cusip"`
	Paused                  bool                    `json:"paused"`
	EstimatedCashAmount     int64                   `json:"estimatedCashAmount"`
	CompanyName             string                  `json:"companyName"`
	Vwap                    float64                 `json:"vwap"`
	Precision               int64                   `json:"precision"`
	Low52Week               int64                   `json:"low52week"`
	Delayed                 bool                    `json:"delayed"`
	Source                  Mic                     `json:"source"`
	CalendarYearHigh        float64                 `json:"calendarYearHigh"`
	DividendYield           int64                   `json:"dividendYield"`
	Put                     bool                    `json:"put"`
	CalendarYearHighDate    string                  `json:"calendarYearHighDate"`
	PriceEarningRatio       float64                 `json:"priceEarningRatio"`
	DayHigh                 int64                   `json:"dayHigh"`
	AskSize                 int64                   `json:"askSize"`
	MutualFund              bool                    `json:"mutualFund"`
	NavPrice                int64                   `json:"navPrice"`
	NavChange               int64                   `json:"navChange"`
	NavYesterdayPrice       int64                   `json:"navYesterdayPrice"`
	MarketCap               float64                 `json:"marketCap"`
	Halted                  bool                    `json:"halted"`
	Nav                     int64                   `json:"nav"`
	MinBidSize              int64                   `json:"minBidSize"`
	SubMarket               int64                   `json:"subMarket"`
	DividendAmount          int64                   `json:"dividendAmount"`
	InTheMoney              bool                    `json:"inTheMoney"`
	AtTheMoney              bool                    `json:"atTheMoney"`
	DividendRate            int64                   `json:"dividendRate"`
	CalendarYearLowDate     string                  `json:"calendarYearLowDate"`
	DayLow                  float64                 `json:"dayLow"`
	Sedol                   string                  `json:"sedol"`
	AskRegion               Region                  `json:"askRegion"`
	Volume                  int64                   `json:"volume"`
	TradeTime               string                  `json:"tradeTime"`
	PutOrCall               PutOrCall               `json:"putOrCall"`
	Nbbo                    bool                    `json:"nbbo"`
	SymbolChange            bool                    `json:"symbolChange"`
	Vwap1                   int64                   `json:"vwap1"`
	Trailing12MonthsEps     float64                 `json:"trailing12MonthsEps"`
	Exchange                EquityExchange          `json:"exchange"`
	TradePrice              float64                 `json:"tradePrice"`
	MinAskSize              int64                   `json:"minAskSize"`
	StrikePrice             int64                   `json:"strikePrice"`
	Option                  bool                    `json:"option"`
	High52Week              float64                 `json:"high52week"`
	BidTick                 int64                   `json:"bidTick"`
	TradeRegion             Region                  `json:"tradeRegion"`
	Jumbo                   bool                    `json:"jumbo"`
	OpenInterest            int64                   `json:"openInterest"`
	Mic                     Mic                     `json:"mic"`
	NotPermissioned         bool                    `json:"notPermissioned"`
	Week52HighDate          string                  `json:"week52HighDate"`
	Increment               int64                   `json:"increment"`
	StandardAndPoorsRating  string                  `json:"standardAndPoorsRating"`
	BidYield                int64                   `json:"bidYield"`
	OpeningPrice            int64                   `json:"openingPrice"`
	Week52LowDate           string                  `json:"week52LowDate"`
	YesterdayVolume         int64                   `json:"yesterdayVolume"`
	MarketCapClassification MarketCapClassification `json:"marketCapClassification"`
	IdentType               IdentType               `json:"identType"`
	ChangePercent           float64                 `json:"changePercent"`
	Currency                Currency                `json:"currency"`
	ContractSize            int64                   `json:"contractSize"`
	TradeSize               int64                   `json:"tradeSize"`
	SharesOutstanding       int64                   `json:"sharesOutstanding"`
	Beta                    float64                 `json:"beta"`
	YesterdayClose          float64                 `json:"yesterdayClose"`
	CalendarYearLow         float64                 `json:"calendarYearLow"`
	ImbalanceVolume         int64                   `json:"imbalanceVolume"`
	Mini                    bool                    `json:"mini"`
	CorpAct                 bool                    `json:"corpAct"`
	BidSize                 int64                   `json:"bidSize"`
	SymbolMarket            SymbolMarket            `json:"symbolMarket"`
	Change                  float64                 `json:"change"`
	UpdateTime              string                  `json:"updateTime"`
	Bond                    bool                    `json:"bond"`
	Call                    bool                    `json:"call"`
	MarketDescription       MarketDescription       `json:"marketDescription"`
	AverageVolume30         int64                   `json:"averageVolume30"`
	Ask                     float64                 `json:"ask"`
	AskYield                int64                   `json:"askYield"`
	TotalCashAmount         int64                   `json:"totalCashAmount"`
	NotFound                bool                    `json:"notFound"`
	BidRegion               Region                  `json:"bidRegion"`
	Bid                     float64                 `json:"bid"`
	DividentSoon            bool                    `json:"dividentSoon"`
	Isin                    string                  `json:"isin"`
	LastPrice               float64                 `json:"lastPrice"`
	IntradayValue           int64                   `json:"intradayValue"`
}

type Region struct {
	Code        Mic    `json:"code"`
	Description string `json:"description"`
}

type SearchElement struct {
	Symbol              *string            `json:"symbol,omitempty"`
	CompanyName         *string            `json:"companyName,omitempty"`
	Isin                *string            `json:"isin,omitempty"`
	Cusip               *string            `json:"cusip,omitempty"`
	Country             *Country           `json:"country,omitempty"`
	Mic                 *Mic               `json:"mic,omitempty"`
	Source              *string            `json:"source,omitempty"`
	Exchange            *EquityExchange    `json:"exchange,omitempty"`
	UpdateTime          *string            `json:"updateTime,omitempty"`
	TradeTime           *string            `json:"tradeTime,omitempty"`
	DividendPaymentDate *string            `json:"dividendPaymentDate,omitempty"`
	DividendExDate      *string            `json:"dividendExDate,omitempty"`
	YesterdayClose      *float64           `json:"yesterdayClose,omitempty"`
	OpeningPrice        *float64           `json:"openingPrice,omitempty"`
	DayHigh             *float64           `json:"dayHigh,omitempty"`
	DayLow              *float64           `json:"dayLow,omitempty"`
	Ask                 *float64           `json:"ask,omitempty"`
	Bid                 *float64           `json:"bid,omitempty"`
	LastPrice           *float64           `json:"lastPrice,omitempty"`
	Change              *float64           `json:"change,omitempty"`
	High52Week          *float64           `json:"high52week,omitempty"`
	Low52Week           *float64           `json:"low52week,omitempty"`
	PriceEarningRatio   *float64           `json:"priceEarningRatio,omitempty"`
	ChangePercent       *float64           `json:"changePercent,omitempty"`
	MarketCap           *float64           `json:"marketCap,omitempty"`
	CalendarYearHigh    *float64           `json:"calendarYearHigh,omitempty"`
	CalendarYearLow     *float64           `json:"calendarYearLow,omitempty"`
	DividendYield       *float64           `json:"dividendYield,omitempty"`
	DividendAmount      *float64           `json:"dividendAmount,omitempty"`
	DividendRate        *float64           `json:"dividendRate,omitempty"`
	Beta                *float64           `json:"beta,omitempty"`
	YesterdayVolume     *int64             `json:"yesterdayVolume,omitempty"`
	AskSize             *int64             `json:"askSize,omitempty"`
	BidSize             *int64             `json:"bidSize,omitempty"`
	Volume              *int64             `json:"volume,omitempty"`
	AverageVolume30     *int64             `json:"averageVolume30,omitempty"`
	Indicator           *int64             `json:"indicator,omitempty"`
	BidTick             *int64             `json:"bidTick,omitempty"`
	Precision           *int64             `json:"precision,omitempty"`
	Delayed             *bool              `json:"delayed,omitempty"`
	DividentSoon        *bool              `json:"dividentSoon,omitempty"`
	MarketDescription   *MarketDescription `json:"marketDescription,omitempty"`
	Trailing12MonthsEps *float64           `json:"trailing12MonthsEps,omitempty"`
	SymbolMarket        *SymbolMarket      `json:"symbolMarket,omitempty"`
	Type                *string            `json:"type,omitempty"`
	OldEntry            *Entry             `json:"oldEntry,omitempty"`
	NewEntry            *Entry             `json:"newEntry,omitempty"`
	EFFDate             *string            `json:"effDate,omitempty"`
	ExDate              *string            `json:"exDate,omitempty"`
	RecordDate          *string            `json:"recordDate,omitempty"`
	DeclareDate         *string            `json:"declareDate,omitempty"`
	PayDate             *string            `json:"payDate,omitempty"`
	PaymentTypes        []interface{}      `json:"paymentTypes,omitempty"`
}

type Entry struct {
	Symbol      string `json:"symbol"`
	Exchange    string `json:"exchange"`
	CompanyName string `json:"companyName"`
	Cusip       string `json:"cusip"`
}

type HistoricalElement struct {
	Date    string  `json:"date"`
	Price   float64 `json:"price"`
	Opening float64 `json:"opening"`
	High    float64 `json:"high"`
	Low     float64 `json:"low"`
	Volume  int64   `json:"volume"`
}

type IntradayChartDatum struct {
	Date   string      `json:"date"`
	Price  float64     `json:"price"`
	High   float64     `json:"high"`
	Low    float64     `json:"low"`
	Volume int64       `json:"volume"`
	SMA    interface{} `json:"sma"`
	Ema    interface{} `json:"ema"`
	RSI    interface{} `json:"rsi"`
}

type TopStock struct {
	Quote              TopStockQuote `json:"quote"`
	Option             bool          `json:"option"`
	Symbol             string        `json:"symbol"`
	HasRecentUpgrade   bool          `json:"hasRecentUpgrade"`
	HasRecentDowngrade bool          `json:"hasRecentDowngrade"`
}

type TopStockQuote struct {
	Symbol                  string                  `json:"symbol"`
	CompanyName             string                  `json:"companyName"`
	Country                 *Country                `json:"country,omitempty"`
	Mic                     Mic                     `json:"mic"`
	Currency                Currency                `json:"currency"`
	Source                  *Mic                    `json:"source,omitempty"`
	IdentType               IdentType               `json:"identType"`
	Exchange                EquityExchange          `json:"exchange"`
	BidRegion               RealizedProfitLoss      `json:"bidRegion"`
	AskRegion               RealizedProfitLoss      `json:"askRegion"`
	TradeRegion             RealizedProfitLoss      `json:"tradeRegion"`
	UpdateTime              *string                 `json:"updateTime,omitempty"`
	TradeTime               *string                 `json:"tradeTime,omitempty"`
	Vwap                    float64                 `json:"vwap"`
	Vwap1                   int64                   `json:"vwap1"`
	YesterdayClose          float64                 `json:"yesterdayClose"`
	OpeningPrice            float64                 `json:"openingPrice"`
	StrikePrice             int64                   `json:"strikePrice"`
	DayHigh                 float64                 `json:"dayHigh"`
	DayLow                  float64                 `json:"dayLow"`
	Ask                     float64                 `json:"ask"`
	Bid                     float64                 `json:"bid"`
	AskYield                int64                   `json:"askYield"`
	BidYield                int64                   `json:"bidYield"`
	LastPrice               float64                 `json:"lastPrice"`
	Change                  float64                 `json:"change"`
	OpenInterest            int64                   `json:"openInterest"`
	High52Week              float64                 `json:"high52week"`
	Low52Week               float64                 `json:"low52week"`
	PriceEarningRatio       float64                 `json:"priceEarningRatio"`
	ChangePercent           float64                 `json:"changePercent"`
	MarketCap               float64                 `json:"marketCap"`
	CalendarYearHigh        float64                 `json:"calendarYearHigh"`
	CalendarYearLow         float64                 `json:"calendarYearLow"`
	DividendYield           float64                 `json:"dividendYield"`
	DividendAmount          float64                 `json:"dividendAmount"`
	DividendRate            float64                 `json:"dividendRate"`
	Beta                    float64                 `json:"beta"`
	TotalCashAmount         int64                   `json:"totalCashAmount"`
	EstimatedCashAmount     int64                   `json:"estimatedCashAmount"`
	IntradayValue           int64                   `json:"intradayValue"`
	Nav                     int64                   `json:"nav"`
	TradePrice              float64                 `json:"tradePrice"`
	YesterdayVolume         int64                   `json:"yesterdayVolume"`
	AskSize                 int64                   `json:"askSize"`
	BidSize                 int64                   `json:"bidSize"`
	MinAskSize              int64                   `json:"minAskSize"`
	MinBidSize              int64                   `json:"minBidSize"`
	Volume                  int64                   `json:"volume"`
	ContractSize            int64                   `json:"contractSize"`
	AverageVolume30         int64                   `json:"averageVolume30"`
	ImbalanceVolume         int64                   `json:"imbalanceVolume"`
	SharesOutstanding       int64                   `json:"sharesOutstanding"`
	Increment               int64                   `json:"increment"`
	TradeSize               int64                   `json:"tradeSize"`
	PutOrCall               PutOrCall               `json:"putOrCall"`
	Indicator               int64                   `json:"indicator"`
	SubMarket               int64                   `json:"subMarket"`
	BidTick                 int64                   `json:"bidTick"`
	Precision               int64                   `json:"precision"`
	Delayed                 bool                    `json:"delayed"`
	NotFound                bool                    `json:"notFound"`
	Option                  bool                    `json:"option"`
	NotPermissioned         bool                    `json:"notPermissioned"`
	Halted                  bool                    `json:"halted"`
	Paused                  bool                    `json:"paused"`
	Mini                    bool                    `json:"mini"`
	CorpAct                 bool                    `json:"corpAct"`
	Jumbo                   bool                    `json:"jumbo"`
	MutualFund              bool                    `json:"mutualFund"`
	Bond                    bool                    `json:"bond"`
	Nbbo                    bool                    `json:"nbbo"`
	Crypto                  bool                    `json:"crypto"`
	Fx                      bool                    `json:"fx"`
	SymbolMarket            SymbolMarket            `json:"symbolMarket"`
	Call                    bool                    `json:"call"`
	Put                     bool                    `json:"put"`
	MarketCapClassification MarketCapClassification `json:"marketCapClassification"`
	DividentSoon            bool                    `json:"dividentSoon"`
	SymbolChange            bool                    `json:"symbolChange"`
	MarketDescription       MarketDescription       `json:"marketDescription"`
	AtTheMoney              bool                    `json:"atTheMoney"`
	InTheMoney              bool                    `json:"inTheMoney"`
	Trailing12MonthsEps     float64                 `json:"trailing12MonthsEps"`
	Isin                    *string                 `json:"isin,omitempty"`
	Cusip                   *string                 `json:"cusip,omitempty"`
	Sedol                   *string                 `json:"sedol,omitempty"`
	DividendPaymentDate     *string                 `json:"dividendPaymentDate,omitempty"`
	DividendExDate          *string                 `json:"dividendExDate,omitempty"`
	CalendarYearHighDate    *string                 `json:"calendarYearHighDate,omitempty"`
	CalendarYearLowDate     *string                 `json:"calendarYearLowDate,omitempty"`
	Week52HighDate          *string                 `json:"week52HighDate,omitempty"`
	Week52LowDate           *string                 `json:"week52LowDate,omitempty"`
}

// Analyst Profile
//
// GET {{url}}/api/research/tipranks/analyst/profile/{expertUID}
//
// Returns the analyst profile information.
type AnalystProfileResponse struct {
	AnalystName           string  `json:"analystName"`
	FirmName              string  `json:"firmName"`
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

type OverviewElement struct {
	Ticker            string  `json:"ticker"`
	PriceTarget       float64 `json:"priceTarget"`
	PriceTargetUpside float64 `json:"priceTargetUpside"`
	Buy               int64   `json:"buy"`
	Hold              int64   `json:"hold"`
	Consensus         string  `json:"consensus"`
	Symbol            string  `json:"symbol"`
}

type ConsensusBySectorElement struct {
	Sector       string `json:"sector"`
	StrongBuy    int64  `json:"strongBuy"`
	ModerateBuy  int64  `json:"moderateBuy"`
	Neutral      int64  `json:"neutral"`
	ModerateSell *int64 `json:"moderateSell,omitempty"`
}

type TrendingElement struct {
	Ticker     string        `json:"ticker"`
	Sector     string        `json:"sector"`
	Quote      TrendingQuote `json:"quote"`
	Consensus  string        `json:"consensus"`
	Popularity int64         `json:"popularity"`
	Sentiment  int64         `json:"sentiment"`
	Hold       int64         `json:"hold"`
	Buy        int64         `json:"buy"`
	Symbol     string        `json:"symbol"`
	Sell       *int64        `json:"sell,omitempty"`
}

type TrendingQuote struct {
	Symbol              string            `json:"symbol"`
	CompanyName         string            `json:"companyName"`
	Isin                string            `json:"isin"`
	Cusip               string            `json:"cusip"`
	Country             Country           `json:"country"`
	Mic                 Mic               `json:"mic"`
	Currency            Currency          `json:"currency"`
	Source              string            `json:"source"`
	IdentType           IdentType         `json:"identType"`
	Exchange            EquityExchange    `json:"exchange"`
	UpdateTime          string            `json:"updateTime"`
	TradeTime           string            `json:"tradeTime"`
	YesterdayClose      float64           `json:"yesterdayClose"`
	OpeningPrice        float64           `json:"openingPrice"`
	DayHigh             float64           `json:"dayHigh"`
	DayLow              float64           `json:"dayLow"`
	Ask                 float64           `json:"ask"`
	Bid                 float64           `json:"bid"`
	LastPrice           float64           `json:"lastPrice"`
	Change              float64           `json:"change"`
	High52Week          float64           `json:"high52week"`
	Low52Week           float64           `json:"low52week"`
	ChangePercent       float64           `json:"changePercent"`
	MarketCap           float64           `json:"marketCap"`
	CalendarYearHigh    float64           `json:"calendarYearHigh"`
	CalendarYearLow     float64           `json:"calendarYearLow"`
	Beta                float64           `json:"beta"`
	YesterdayVolume     int64             `json:"yesterdayVolume"`
	AskSize             int64             `json:"askSize"`
	BidSize             int64             `json:"bidSize"`
	Volume              int64             `json:"volume"`
	AverageVolume30     int64             `json:"averageVolume30"`
	Indicator           int64             `json:"indicator"`
	BidTick             int64             `json:"bidTick"`
	Precision           int64             `json:"precision"`
	Delayed             bool              `json:"delayed"`
	MarketDescription   MarketDescription `json:"marketDescription"`
	SymbolMarket        SymbolMarket      `json:"symbolMarket"`
	PriceEarningRatio   *float64          `json:"priceEarningRatio,omitempty"`
	Trailing12MonthsEps *float64          `json:"trailing12MonthsEps,omitempty"`
}

// Init Screener
//
// GET {{url}}/api/research/etf/init
//
// Allows for initial load of some common, mostly non-changing ETF data. Use this endpoint
// to look up available parameter values used in the <b>Screener</b> endpoint.
//
// ### Inclued parameter values are:
// * Sponsors
// * AssetClasses
// * InverseTypes
// * Categories
// * Sectors
// * LeverageTypes
// * SectorCodes
type InitScreener struct {
	Sponsors      map[string]Sponsor `json:"Sponsors"`
	AssetClasses  map[string]string  `json:"AssetClasses"`
	InverseTypes  []string           `json:"InverseTypes"`
	Categories    map[string]string  `json:"Categories"`
	Sectors       []string           `json:"Sectors"`
	LeverageTypes []string           `json:"LeverageTypes"`
	SectorCodes   map[string]string  `json:"SectorCodes"`
}

type Sponsor struct {
	EtfCount int64  `json:"etfCount"`
	Name     string `json:"name"`
}

// Top 10
//
// GET {{url}}/api/research/etf/top10?inverseType=ALL&leverageType=LEVERAGE
//
// Load top 10 ETFs Can filter based on inverseType and leverageType.
//
// ### Inverse Types
// |inverseType|definition|
// |:---|:---|
// |ALL|Any|
// |LONG|Long ETFs|
// |SHORT|Inverse ETFs|
//
// ### Leverage Types
// |leverageType|definition|
// |:---|:---|
// |ALL|Any|
// |LEVERAGE|All leverage ETFs|
// |LEVERAGE_I|Non leverage|
// |LEVERAGE_II|Leverage x2|
// |LEVERAGE_III|Leverage x3|
type Top10 struct {
	Topchangepc    []Bottomchangepc `json:"topchangepc"`
	Topvolume      []Bottomchangepc `json:"topvolume"`
	Topvalue       []Bottomchangepc `json:"topvalue"`
	Bottomchangepc []Bottomchangepc `json:"bottomchangepc"`
}

type Bottomchangepc struct {
	Quote              TopStockQuote `json:"quote"`
	Etf                EtfClass      `json:"etf"`
	Value              int64         `json:"value"`
	Option             bool          `json:"option"`
	HasRecentUpgrade   bool          `json:"hasRecentUpgrade"`
	HasRecentDowngrade bool          `json:"hasRecentDowngrade"`
	Symbol             string        `json:"symbol"`
}

type EtfClass struct {
	Symbol        string        `json:"symbol"`
	SponsorID     int64         `json:"sponsorId"`
	Sponsor       SponsorEnum   `json:"sponsor"`
	SponsorDesc   string        `json:"sponsorDesc"`
	SponsorURL    string        `json:"sponsorUrl"`
	AssetID       int64         `json:"assetId"`
	Asset         Asset         `json:"asset"`
	CategoryID    int64         `json:"categoryId"`
	Category      Category      `json:"category"`
	StructureType StructureType `json:"structureType"`
	StructureID   int64         `json:"structureId"`
	Structure     Structure     `json:"structure"`
	InceptionDate string        `json:"inceptionDate"`
	Leverage      int64         `json:"leverage"`
	ExpenseRatio  float64       `json:"expenseRatio"`
	URL           string        `json:"url"`
	FactSheetURL  string        `json:"factSheetUrl"`
	ProspectusURL string        `json:"prospectusUrl"`
	OtherUrls     string        `json:"otherUrls"`
	CountryPct    int64         `json:"countryPct"`
	ContinentPct  int64         `json:"continentPct"`
	SectorCode    int64         `json:"sectorCode"`
	SectorPct     int64         `json:"sectorPct"`
	PriceRank     int64         `json:"priceRank"`
	PriceChange   int64         `json:"priceChange"`
	PriceChange1W int64         `json:"priceChange1w"`
	PriceChange2W int64         `json:"priceChange2w"`
	PriceChange1M int64         `json:"priceChange1m"`
	Exposure      int64         `json:"exposure"`
	Inverse       bool          `json:"inverse"`
	Active        bool          `json:"active"`
	Option        bool          `json:"option"`
}

// Details
//
// GET {{url}}/api/research/etf/details?symbol=VXF&loadQuotes=false
//
// Load various details on the selected ETF symbol:
//
// * Maturities
// * Sectors
// * Countries
// * Holdings
// * MarketCaps
// * Continents
// * AssetAllocations
type Details struct {
	Maturities       []AssetAllocation `json:"Maturities"`
	Sectors          []Continent       `json:"Sectors"`
	Countries        []AssetAllocation `json:"Countries"`
	Holdings         []Holding         `json:"Holdings"`
	MarketCaps       []AssetAllocation `json:"MarketCaps"`
	Continents       []Continent       `json:"Continents"`
	AssetAllocations []AssetAllocation `json:"AssetAllocations"`
}

type AssetAllocation struct {
	Breakdown  string  `json:"breakdown"`
	Percentage float64 `json:"percentage"`
}

type Continent struct {
	Code       string  `json:"code"`
	Breakdown  string  `json:"breakdown"`
	Percentage float64 `json:"percentage"`
}

type Holding struct {
	Symbol      string  `json:"symbol"`
	CompanyName string  `json:"companyName"`
	Rank        int64   `json:"rank"`
	Percentage  float64 `json:"percentage"`
	Option      bool    `json:"option"`
}

// Breakdown by Sector
//
// GET {{url}}/api/research/etf/breakdown?duration=1d&displaySize=5&exposure=90
//
// Load top ETFs by sector.
type BreakdownBySector struct {
	Energy                []ConsumerDiscretionary `json:"ENERGY"`
	Materials             []ConsumerDiscretionary `json:"MATERIALS"`
	Industrials           []ConsumerDiscretionary `json:"INDUSTRIALS"`
	ConsumerDiscretionary []ConsumerDiscretionary `json:"CONSUMER_DISCRETIONARY"`
	ConsumerStaples       []ConsumerDiscretionary `json:"CONSUMER_STAPLES"`
	HealthCare            []ConsumerDiscretionary `json:"HEALTH_CARE"`
	Financials            []ConsumerDiscretionary `json:"FINANCIALS"`
	RealEstate            []ConsumerDiscretionary `json:"REAL_ESTATE"`
	InformationTechnology []ConsumerDiscretionary `json:"INFORMATION_TECHNOLOGY"`
	Telecommunication     []interface{}           `json:"TELECOMMUNICATION"`
	Utilities             []ConsumerDiscretionary `json:"UTILITIES"`
}

type ConsumerDiscretionary struct {
	Symbol                string        `json:"symbol"`
	Option                bool          `json:"option"`
	LevelCode             int64         `json:"levelCode"`
	Rank                  int64         `json:"rank"`
	MarketCap             int64         `json:"marketCap"`
	CurrentValue          float64       `json:"currentValue"`
	OneWeekValue          float64       `json:"oneWeekValue"`
	TwoWeeksValue         float64       `json:"twoWeeksValue"`
	OneMonthValue         float64       `json:"oneMonthValue"`
	Quote                 TopStockQuote `json:"quote"`
	Duration              Duration      `json:"duration"`
	BreakdownPct          float64       `json:"breakdownPct"`
	EtfID                 int64         `json:"etfId"`
	RankingPriceChangePct float64       `json:"rankingPriceChangePct"`
}

// Alpha Tracker
//
// GET
// {{url}}/api/quotes/rst?benchmark=SPY&symbols=AAPL,GOOG&range=2y&period=180&setEMA=true
//
// Returns an array of data points used to create a momentum ratio graph.
type AlphaTracker struct {
	Goog []GOOGElement `json:"GOOG"`
	Aapl []GOOGElement `json:"AAPL"`
}

type GOOGElement struct {
	ResultRatio    float64 `json:"resultRatio"`
	ResultMomentum float64 `json:"resultMomentum"`
	Date           string  `json:"date"`
}

// Seasonality
//
// GET {{url}}/api/quotes/seasonality?symbols=AAPL,IBM,GOOG,TSLA&benchmark=SPY&years=5
//
// ## Seasonality
type Seasonality struct {
	Goog map[string]AAPLValue `json:"GOOG"`
	Aapl map[string]AAPLValue `json:"AAPL"`
	Tsla map[string]AAPLValue `json:"TSLA"`
	IBM  map[string]AAPLValue `json:"IBM"`
}

type AAPLValue struct {
	Success          *float64      `json:"success,omitempty"`
	Failure          *float64      `json:"failure,omitempty"`
	Average          float64       `json:"average"`
	YearlyData       []YearlyDatum `json:"yearlyData"`
	TotalYears       int64         `json:"totalYears"`
	BenchmarkAverage float64       `json:"benchmarkAverage"`
	Difference       float64       `json:"difference"`
}

type YearlyDatum struct {
	Year             int64    `json:"year"`
	Date             int64    `json:"date"`
	Value            *float64 `json:"value,omitempty"`
	BenchmarkAverage float64  `json:"benchmarkAverage"`
	Difference       float64  `json:"difference"`
}

// Weekly Seasonality
//
// GET
// {{url}}/api/quotes/seasonality/weekly?symbols=AAPL,IBM,GOOG,TSLA&benchmark=SPY&years=5
//
// ## Seasonality weekly data
type WeeklySeasonality struct {
	Goog map[string]AAPLValue `json:"GOOG"`
	Aapl map[string]AAPLValue `json:"AAPL"`
	Tsla map[string]AAPLValue `json:"TSLA"`
	IBM  map[string]AAPLValue `json:"IBM"`
}

type UpcomingElement struct {
	Symbol                          string        `json:"symbol"`
	CompanyName                     string        `json:"companyName"`
	ProspectusURL                   string        `json:"prospectusUrl"`
	FileDate                        string        `json:"fileDate"`
	IPODate                         string        `json:"ipoDate"`
	IPODateApprox                   bool          `json:"ipoDateApprox"`
	OfferPrice                      float64       `json:"offerPrice"`
	OfferPriceMax                   float64       `json:"offerPriceMax"`
	OfferPriceMin                   float64       `json:"offerPriceMin"`
	OfferShares                     float64       `json:"offerShares"`
	FirstDayOpen                    int64         `json:"firstDayOpen"`
	FirstDayClose                   int64         `json:"firstDayClose"`
	IPOReturn                       int64         `json:"ipoReturn"`
	Status                          string        `json:"status"`
	Underwriters                    []Underwriter `json:"underwriters"`
	Ceo                             string        `json:"ceo"`
	TotalExpenses                   float64       `json:"totalExpenses"`
	SharesOverAlloted               int64         `json:"sharesOverAlloted"`
	SharesOutstanding               int64         `json:"sharesOutstanding"`
	LockupPeriod                    int64         `json:"lockupPeriod"`
	OfferAmount                     int64         `json:"offerAmount"`
	IndustryValue                   int64         `json:"industryValue"`
	IndustryPercentage              int64         `json:"industryPercentage"`
	FirstDayReturn                  int64         `json:"firstDayReturn"`
	TotalReturn                     int64         `json:"totalReturn"`
	OfferSize                       float64       `json:"offerSize"`
	QuoteYesterdayCloseWithoutSplit int64         `json:"quoteYesterdayCloseWithoutSplit"`
	LeadUnderwriter                 Underwriter   `json:"leadUnderwriter"`
	FirstTradingDay                 bool          `json:"firstTradingDay"`
	Cik                             int64         `json:"cik"`
	Value                           int64         `json:"value"`
	Option                          bool          `json:"option"`
	GroupValue                      int64         `json:"groupValue"`
	SectorValue                     int64         `json:"sectorValue"`
	IndValue                        int64         `json:"indValue"`
	SubIndustryValue                int64         `json:"subIndustryValue"`
	StateOfInc                      *string       `json:"stateOfInc,omitempty"`
}

type Underwriter struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

type PerformanceElement struct {
	Symbol                          string  `json:"symbol"`
	CompanyName                     string  `json:"companyName"`
	ProspectusURL                   string  `json:"prospectusUrl"`
	FileDate                        string  `json:"fileDate"`
	IPODate                         string  `json:"ipoDate"`
	IPODateApprox                   bool    `json:"ipoDateApprox"`
	OfferPrice                      float64 `json:"offerPrice"`
	OfferPriceMax                   int64   `json:"offerPriceMax"`
	OfferPriceMin                   int64   `json:"offerPriceMin"`
	OfferShares                     float64 `json:"offerShares"`
	FirstDayOpen                    float64 `json:"firstDayOpen"`
	FirstDayClose                   float64 `json:"firstDayClose"`
	IPOReturn                       float64 `json:"ipoReturn"`
	Status                          Status  `json:"status"`
	Ceo                             string  `json:"ceo"`
	StateOfInc                      *string `json:"stateOfInc,omitempty"`
	TotalExpenses                   float64 `json:"totalExpenses"`
	SharesOverAlloted               int64   `json:"sharesOverAlloted"`
	SharesOutstanding               int64   `json:"sharesOutstanding"`
	LockupPeriod                    int64   `json:"lockupPeriod"`
	LockupExpiration                string  `json:"lockupExpiration"`
	QuietPeriodExpiration           string  `json:"quietPeriodExpiration"`
	OfferAmount                     int64   `json:"offerAmount"`
	IndustryValue                   int64   `json:"industryValue"`
	IndustryPercentage              int64   `json:"industryPercentage"`
	FirstDayReturn                  float64 `json:"firstDayReturn"`
	TotalReturn                     int64   `json:"totalReturn"`
	OfferSize                       float64 `json:"offerSize"`
	QuoteYesterdayCloseWithoutSplit int64   `json:"quoteYesterdayCloseWithoutSplit"`
	FirstTradingDay                 bool    `json:"firstTradingDay"`
	Cik                             int64   `json:"cik"`
	Value                           int64   `json:"value"`
	Option                          bool    `json:"option"`
	GroupValue                      int64   `json:"groupValue"`
	SectorValue                     int64   `json:"sectorValue"`
	IndValue                        int64   `json:"indValue"`
	SubIndustryValue                int64   `json:"subIndustryValue"`
}

// Option Ratios
//
// GET {{url}}/api/research/option/ratios?symbols=AAPL,BABA
//
// Analyze option ratios by symbols to find positive and negative bias towards specific
// equities.
type OptionRatios struct {
	Baba []BABAElement `json:"BABA"`
	Aapl []BABAElement `json:"AAPL"`
}

type BABAElement struct {
	UnderlyingSymbol     string  `json:"underlyingSymbol"`
	Expiration           string  `json:"expiration"`
	CallVolume           int64   `json:"callVolume"`
	PutVolume            int64   `json:"putVolume"`
	TotalVolume          int64   `json:"totalVolume"`
	CallOi               int64   `json:"callOi"`
	PutOi                int64   `json:"putOi"`
	CallVolAvg30Day      float64 `json:"callVolAvg30Day"`
	CallOiAvg30Day       float64 `json:"callOiAvg30Day"`
	PutVolAvg30Day       float64 `json:"putVolAvg30Day"`
	PutOiAvg30Day        float64 `json:"putOiAvg30Day"`
	CallVolCallOiRatio   float64 `json:"callVolCallOiRatio"`
	PutVolPutOiRatio     float64 `json:"putVolPutOiRatio"`
	CallOiPutOiRatio     float64 `json:"callOiPutOiRatio"`
	PutOiCallOiRatio     float64 `json:"putOiCallOiRatio"`
	CallVolPutVolRatio   float64 `json:"callVolPutVolRatio"`
	PutVolCallVolRatio   float64 `json:"putVolCallVolRatio"`
	CallVol30DayAvgRatio float64 `json:"callVol30DayAvgRatio"`
	PutVol30DayAvgRatio  float64 `json:"putVol30DayAvgRatio"`
	CallOi30DayAvgRatio  float64 `json:"callOi30DayAvgRatio"`
	PutOi30DayAvgRatio   float64 `json:"putOi30DayAvgRatio"`
	Updated              Updated `json:"updated"`
	Symbol               string  `json:"symbol"`
	Option               bool    `json:"option"`
}

// Option Radar
//
// GET {{url}}/api/research/option/radar?filter=
//
// Analyze option ratios across the whole option market to find positive and negative bias
// towards specific equities.
//
// |filter|Description|
// |-------|--------|
// |vol_oi|Volume: Open Interest ratios|
// |vol|Open Interest ratios|
// |oi|Volume ratios|
type OptionRadar struct {
	CallVolPutVolRatios []BABAElement `json:"callVolPutVolRatios"`
	PutOiCallOiRatios   []BABAElement `json:"putOiCallOiRatios"`
	PutVolPutOiRatios   []BABAElement `json:"putVolPutOiRatios"`
	PutVolCallVolRatios []BABAElement `json:"putVolCallVolRatios"`
	CallOiPutOiRatios   []BABAElement `json:"callOiPutOiRatios"`
	CallVolCallOiRatios []BABAElement `json:"callVolCallOiRatios"`
}

type OwnershipDetail struct {
	Details              []Detail `json:"details"`
	TotalAvailable       int64    `json:"totalAvailable"`
	InstitutionalPercent *float64 `json:"institutionalPercent,omitempty"`
	FundPercent          *float64 `json:"fundPercent,omitempty"`
	OtherPercent         *float64 `json:"otherPercent,omitempty"`
}

type Detail struct {
	Symbol             string  `json:"symbol"`
	OwnerID            string  `json:"ownerId"`
	OwnerType          *string `json:"ownerType,omitempty"`
	Date               string  `json:"date"`
	OwnerCIK           int64   `json:"ownerCIK"`
	OwnerName          string  `json:"ownerName"`
	NumberOfShares     *int64  `json:"numberOfShares,omitempty"`
	ShareChangePercent int64   `json:"shareChangePercent"`
}

// Owner Details
//
// GET {{url}}/api/research/owner/details?ids=PS000000YH,PS000004TI
type OwnerDetails struct {
	Ps000004TI Ps00000 `json:"PS000004TI"`
	Ps000000Yh Ps00000 `json:"PS000000YH"`
}

type Ps00000 struct {
	OwnerName string   `json:"ownerName"`
	Details   []Detail `json:"details"`
}

type EarningsCalendarElement struct {
	InPortfolio        bool   `json:"inPortfolio"`
	InWatchlist        bool   `json:"inWatchlist"`
	Estimate           int64  `json:"estimate"`
	Eps                int64  `json:"eps"`
	Symbol             string `json:"symbol"`
	Time               Time   `json:"time"`
	CreatedOn          string `json:"createdOn"`
	Date               string `json:"date"`
	IndustryValue      int64  `json:"industryValue"`
	IndustryPercentage int64  `json:"industryPercentage"`
	Value              int64  `json:"value"`
	Option             bool   `json:"option"`
	GroupValue         int64  `json:"groupValue"`
	IndValue           int64  `json:"indValue"`
	SubIndustryValue   int64  `json:"subIndustryValue"`
	SectorValue        int64  `json:"sectorValue"`
}

type EarningsCalendarForPortfolioElement struct {
	InPortfolio        bool          `json:"inPortfolio"`
	InWatchlist        bool          `json:"inWatchlist"`
	Estimate           int64         `json:"estimate"`
	Eps                int64         `json:"eps"`
	Symbol             string        `json:"symbol"`
	Time               Time          `json:"time"`
	Quote              TopStockQuote `json:"quote"`
	CreatedOn          string        `json:"createdOn"`
	Date               string        `json:"date"`
	IndustryValue      int64         `json:"industryValue"`
	IndustryPercentage int64         `json:"industryPercentage"`
	Value              int64         `json:"value"`
	Option             bool          `json:"option"`
	GroupValue         int64         `json:"groupValue"`
	IndValue           int64         `json:"indValue"`
	SubIndustryValue   int64         `json:"subIndustryValue"`
	SectorValue        int64         `json:"sectorValue"`
}

// Screener
//
// POST {{url}}/api/research/screener
//
// * Field values: `MarketCap`, `LastPrice`, `Beta`, `DividendAmount`, `DividendRate`,
// `DividendYield`, `DividendPayDate`, `DividendExDate`, `Exchange`, `Gap`, `Adr`, `Etf`,
// `Industry`, `Volume`, `Volume30D`, `PriceEarningRatio`
// * Comparators: `Equals`, `LessThan`, `GreaterThan`, `NotEquals`
type Screener struct {
	ResultCount int64          `json:"resultCount"`
	TotalCount  int64          `json:"totalCount"`
	Items       []ScreenerItem `json:"items"`
}

type ScreenerItem struct {
	Symbol string       `json:"symbol"`
	Values PurpleValues `json:"values"`
}

type PurpleValues struct {
	Adr       bool    `json:"Adr"`
	MarketCap float64 `json:"MarketCap"`
}

// Screener V2
//
// POST {{url}}/api/research/screener
//
// * Request Start Index, Page Size (limit) is now passed along with the body
// * Starts at index 0.
// * If start exceeds the length of the list, an empty list is returned
// * Field values: `MarketCap`, `LastPrice`, `Beta`, `DividendAmount`, `DividendRate`,
// `DividendYield`, `DividendPayDate`, `DividendExDate`, `Exchange`, `Gap`, `Adr`, `Etf`,
// `Industry`, `Volume`, `Volume30D`, `PriceEarningRatio`
// * Comparators: `Equals`, `LessThan`, `GreaterThan`, `NotEquals`
type ScreenerV2 struct {
	ResultCount int64            `json:"resultCount"`
	Start       int64            `json:"start"`
	TotalCount  int64            `json:"totalCount"`
	Items       []ScreenerV2Item `json:"items"`
}

type ScreenerV2Item struct {
	Symbol string       `json:"symbol"`
	Values FluffyValues `json:"values"`
	Option bool         `json:"option"`
}

type FluffyValues struct {
	Industry string `json:"Industry"`
}

type IndustryHeatMapElement struct {
	Industry                Industry `json:"industry"`
	PriceRank               int64    `json:"priceRank"`
	PriceChange             float64  `json:"priceChange"`
	PriceChangeOneWeek      float64  `json:"priceChangeOneWeek"`
	PriceChangeTwoWeeks     float64  `json:"priceChangeTwoWeeks"`
	PriceChangeOneMonth     float64  `json:"priceChangeOneMonth"`
	MarketCapRank           int64    `json:"marketCapRank"`
	MarketCap               float64  `json:"marketCap"`
	MarketCapChange         float64  `json:"marketCapChange"`
	MarketCapChangeOneWeek  float64  `json:"marketCapChangeOneWeek"`
	MarketCapChangeTwoWeeks float64  `json:"marketCapChangeTwoWeeks"`
	MarketCapChangeOneMonth float64  `json:"marketCapChangeOneMonth"`
}

type Industry struct {
	IndustryType IndustryType  `json:"industryType"`
	LevelCode    int64         `json:"levelCode"`
	IndustryName string        `json:"industryName"`
	Symbols      []interface{} `json:"symbols"`
	Size         int64         `json:"size"`
}

type IndustryListElement struct {
	IndustryName string `json:"industryName"`
	LevelCode    int64  `json:"levelCode"`
	IndustryType string `json:"industryType"`
}

type SymbolHeatMapElement struct {
	Indicator              int64             `json:"indicator"`
	Symbol                 string            `json:"symbol"`
	Country                Country           `json:"country"`
	Cusip                  string            `json:"cusip"`
	High52Week             float64           `json:"high52week"`
	Mic                    Mic               `json:"mic"`
	CompanyName            string            `json:"companyName"`
	Precision              int64             `json:"precision"`
	Week52HighDate         string            `json:"week52HighDate"`
	Low52Week              float64           `json:"low52week"`
	Source                 Source            `json:"source"`
	OpeningPrice           float64           `json:"openingPrice"`
	CalendarYearHigh       float64           `json:"calendarYearHigh"`
	CalendarYearHighDate   string            `json:"calendarYearHighDate"`
	Week52LowDate          string            `json:"week52LowDate"`
	YesterdayVolume        *int64            `json:"yesterdayVolume,omitempty"`
	IdentType              IdentType         `json:"identType"`
	ChangePercent          float64           `json:"changePercent"`
	Currency               Currency          `json:"currency"`
	DayHigh                float64           `json:"dayHigh"`
	Beta                   *float64          `json:"beta,omitempty"`
	YesterdayClose         float64           `json:"yesterdayClose"`
	CalendarYearLow        float64           `json:"calendarYearLow"`
	MarketCap              *float64          `json:"marketCap,omitempty"`
	SymbolMarket           SymbolMarket      `json:"symbolMarket"`
	Change                 float64           `json:"change"`
	UpdateTime             string            `json:"updateTime"`
	CalendarYearLowDate    string            `json:"calendarYearLowDate"`
	DayLow                 float64           `json:"dayLow"`
	Sedol                  string            `json:"sedol"`
	Volume                 int64             `json:"volume"`
	MarketDescription      MarketDescription `json:"marketDescription"`
	TradeTime              string            `json:"tradeTime"`
	AverageVolume30        *int64            `json:"averageVolume30,omitempty"`
	Vwap1                  float64           `json:"vwap1"`
	Exchange               EquityExchange    `json:"exchange"`
	Isin                   string            `json:"isin"`
	LastPrice              float64           `json:"lastPrice"`
	StandardAndPoorsRating *string           `json:"standardAndPoorsRating,omitempty"`
	DividendPaymentDate    *string           `json:"dividendPaymentDate,omitempty"`
	DividendAmount         *float64          `json:"dividendAmount,omitempty"`
	DividendExDate         *string           `json:"dividendExDate,omitempty"`
	PriceEarningRatio      *float64          `json:"priceEarningRatio,omitempty"`
	Trailing12MonthsEps    *float64          `json:"trailing12MonthsEps,omitempty"`
	DividendYield          *float64          `json:"dividendYield,omitempty"`
	DividendRate           *float64          `json:"dividendRate,omitempty"`
	AskSize                *int64            `json:"askSize,omitempty"`
	BidSize                *int64            `json:"bidSize,omitempty"`
	Ask                    *float64          `json:"ask,omitempty"`
	Bid                    *float64          `json:"bid,omitempty"`
	DividentSoon           *bool             `json:"dividentSoon,omitempty"`
}

// Historical Data
//
// GET {{url}}/api/research/historicalPrices?symbol&dates
type HistoricalData struct {
	The5Y The1_M `json:"5y"`
	The1W The1_M `json:"1w"`
	The2Y The1_M `json:"2y"`
	The1M The1_M `json:"1m"`
}

type The1_M struct {
	Date  string  `json:"date"`
	Price float64 `json:"price"`
}

// Account trading limit
//
// GET {{url}}/api/user/balance?account&accountId&currency
//
// Return buying power (aka trading limit) on individual account. Not all fields may be
// available on all systems:
//
// ```javascript
// {
// "availableCash": 229851.66, // key field, indicates account's buying power
// "marginCalls": {},
// "netFundsAvailable": 115925.34,
// "sodTotalTradeBalance": 71769.69,
// "sodDayBP": 551826.36,
// "pendingDebitInterest": 0,
// "sodEquity": 88280.79,
// "sodMoneyMarketFunds": 0,
// "buyingPower": 0,
// "sodOptionsShort": 0,
// "fullyPaidUnsettledBalance": 0,
// "sodOvernightBP": 231850.67,
// "sodCashOnAccount": 71769.69,
// "sodSma": 115925.34,
// "sodEquityShort": 0,
// "grossFundsAvailable": 115925.34,
// "currency": "USD",
// "sodTotalEquity": 160050.48,
// "sodMarginBalance": 71769.69,
// "cash": 137956.59,
// "sodMaintenanceExcess": 137956.59,
// "sodEquityMarketValue": 0,
// "cashOnAccount": 70768.18,
// "sellToOpenCash": 0,
// "pendingTransfer": 0,
// "recentDeposit": 0,
// "sodCash": 0,
// "sodShortBalance": 0,
// "pendingDebitDividend": 0,
// "sodTaxWitholding": 0,
// "sodDebitBalance": 0,
// "sodAvailable": true,
// "sodOptionsLong": 0,
// "sodCashEquity": 0,
// "sodCashBalance": 0,
// "sodMarginExcessEquity": 137956.59
// }
// ```
type AccountTradingLimit struct {
	Currency      Currency `json:"currency"`
	BuyingPower   float64  `json:"buyingPower"`
	AvailableCash float64  `json:"availableCash"`
}

// /user/watchlist
//
// GET {{url}}/api/user/watchlist?loadQuotes=false
type UserWatchlist struct {
	Entries []EntryElement `json:"Entries"`
	Titles  []Title        `json:"Titles"`
}

type EntryElement struct {
	Symbol      string `json:"symbol"`
	WatchlistID int64  `json:"watchlistId"`
}

type Title struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type UserPortfolioEarning struct {
	Symbol    string    `json:"symbol"`
	Time      Time      `json:"time"`
	CreatedOn CreatedOn `json:"createdOn"`
	Date      string    `json:"date"`
}

type UserPreference struct {
	Symbol    string    `json:"symbol"`
	Time      Time      `json:"time"`
	CreatedOn CreatedOn `json:"createdOn"`
	Date      string    `json:"date"`
}

type MemberRealtimeBalance struct {
	AccountID           int64   `json:"accountId"`
	ModelID             *int64  `json:"modelId,omitempty"`
	Balance             float64 `json:"balance"`
	TotalDividendAmount int64   `json:"totalDividendAmount"`
	TotalDividendYield  int64   `json:"totalDividendYield"`
}

// Portfolio
//
// GET
//
// {{url}}/api/v2/advisory/model/portfolios/{modelId}?loadQuotes=false&combined=false&loadIndustries=false&loadUpgrades=false&loadEarnings=false&loadCost=false&marketCapWeights=false
//
// If `modelId` specified, returns all positions contained within requested model. Each
// entry will indicate account it belongs to by `accountId` field. Entries can be groupped
// together by having `combined=true` parameter. Grouping will loose `accountId` and cost
// will be averaged out.
//
// If `modelId` equals `all`, returns all positions across all models
type Portfolio struct {
	Homogeneous        bool               `json:"homogeneous"`
	RealizedProfitLoss RealizedProfitLoss `json:"realizedProfitLoss"`
	Entries            []EntryClass       `json:"entries"`
	Total              int64              `json:"total"`
	TotalGain          int64              `json:"totalGain"`
	ContainsEquities   bool               `json:"containsEquities"`
	TotalGainPct       int64              `json:"totalGainPct"`
	ContainsShort      bool               `json:"containsShort"`
	Currency           Currency           `json:"currency"`
	ContainsOptions    bool               `json:"containsOptions"`
	TotalDaysValue     int64              `json:"totalDaysValue"`
	TotalCost          float64            `json:"totalCost"`
}

type EntryClass struct {
	Quote                  EntryQuote             `json:"quote"`
	Currency               Currency               `json:"currency"`
	Symbol                 Symbol                 `json:"symbol"`
	Price                  float64                `json:"price"`
	Quantity               int64                  `json:"quantity"`
	Type                   int64                  `json:"type"`
	LiquidationQuantity    int64                  `json:"liquidationQuantity"`
	LiquidationTransaction LiquidationTransaction `json:"liquidationTransaction"`
	Long                   bool                   `json:"long"`
	AccountType            int64                  `json:"accountType"`
	CostValue              float64                `json:"costValue"`
	LongCash               bool                   `json:"longCash"`
	CostOrders             []RealizedProfitLoss   `json:"costOrders"`
	AccountID              *int64                 `json:"accountId,omitempty"`
}

type EntryQuote struct {
	Symbol    Symbol         `json:"symbol"`
	Currency  Currency       `json:"currency"`
	IdentType IdentType      `json:"identType"`
	Exchange  PurpleExchange `json:"exchange"`
}

// /group/portfolios/{groupId}
//
// GET
//
// {{url}}/api/v2/advisory/group/portfolios/{groupId}?loadQuotes=false&combined=false&loadIndustries=false&loadUpgrades=false&loadEarnings=false&loadCost=false&marketCapWeights=false
type GroupPortfoliosGroupID struct {
	Entries            []EntryClass       `json:"entries"`
	Total              int64              `json:"total"`
	TotalGain          int64              `json:"totalGain"`
	ContainsEquities   bool               `json:"containsEquities"`
	TotalGainPct       int64              `json:"totalGainPct"`
	ContainsShort      bool               `json:"containsShort"`
	Currency           Currency           `json:"currency"`
	ContainsOptions    bool               `json:"containsOptions"`
	TotalDaysValue     int64              `json:"totalDaysValue"`
	TotalCost          float64            `json:"totalCost"`
	Homogeneous        bool               `json:"homogeneous"`
	RealizedProfitLoss RealizedProfitLoss `json:"realizedProfitLoss"`
}

// Create
//
// POST {{url}}/api/transfer/relationship/create
//
// Creates new ACH relationship. Micro-deposits will be issued to the bank account which
// should then be verified via the `approve` endpoint. If you wish to use the IAV method
// instead please view the Instant Account Verification endpoints.
type Create struct {
	EntryID              int64    `json:"entryId"`
	BankAccount          string   `json:"bankAccount"`
	BankAccountType      string   `json:"bankAccountType"`
	CreationTime         string   `json:"creationTime"`
	BankRoutingNumber    string   `json:"bankRoutingNumber"`
	Approval             Approval `json:"approval"`
	Nickname             string   `json:"nickname"`
	ApprovalMethod       string   `json:"approvalMethod"`
	ID                   string   `json:"id"`
	Account              string   `json:"account"`
	BankAccountOwnerName string   `json:"bankAccountOwnerName"`
	Status               string   `json:"status"`
}

type Approval struct {
	ApprovedBy   ApprovedBy `json:"approvedBy"`
	ApprovalTime string     `json:"approvalTime"`
}

type ApprovedBy struct {
	UserEntity string `json:"userEntity"`
	UserClass  string `json:"userClass"`
	UserName   string `json:"userName"`
}

// Approve
//
// POST {{url}}/api/transfer/relationship/approve
//
// Approve a pending relationship with micro deposits.
type Approve struct {
	BankAccount          string `json:"bankAccount"`
	BankAccountType      string `json:"bankAccountType"`
	CreationTime         string `json:"creationTime"`
	Nickname             string `json:"nickname"`
	ApprovalMethod       string `json:"approvalMethod"`
	EntryID              int64  `json:"entryId"`
	Account              string `json:"account"`
	BankAccountOwnerName string `json:"bankAccountOwnerName"`
	Status               string `json:"status"`
}

// Cancel
//
// POST {{url}}/api/transfer/relationship/cancel/{entryId}
//
// Cancel the relationship.
type Cancel struct {
	BankAccount          string       `json:"bankAccount"`
	Cancellation         Cancellation `json:"cancellation"`
	BankAccountType      string       `json:"bankAccountType"`
	CreationTime         string       `json:"creationTime"`
	BankRoutingNumber    string       `json:"bankRoutingNumber"`
	Approval             Approval     `json:"approval"`
	Nickname             string       `json:"nickname"`
	ApprovalMethod       string       `json:"approvalMethod"`
	EntryID              int64        `json:"entryId"`
	Account              string       `json:"account"`
	BankAccountOwnerName string       `json:"bankAccountOwnerName"`
	Status               string       `json:"status"`
}

type Cancellation struct {
	Reason           string `json:"reason"`
	CancellationTime string `json:"cancellationTime"`
	Comment          string `json:"comment"`
}

// Rename
//
// POST {{url}}/api/transfer/relationship/rename/{entryId}
//
// Change the nickname of a relationship.
type Rename struct {
	BankAccount          string `json:"bankAccount"`
	BankAccountType      string `json:"bankAccountType"`
	CreationTime         string `json:"creationTime"`
	BankRoutingNumber    string `json:"bankRoutingNumber"`
	Nickname             string `json:"nickname"`
	ApprovalMethod       string `json:"approvalMethod"`
	EntryID              int64  `json:"entryId"`
	Account              string `json:"account"`
	BankAccountOwnerName string `json:"bankAccountOwnerName"`
	Status               string `json:"status"`
}

// Initiate Post
//
// POST {{url}}/api/iav/init
//
// Initiates IAV process for supported 3rd Party Services for account aggregation. Orbis
// supports multiple IAV providers.
//
//
// Plaid
// Obtain the public key required to activate IAV on the client side.
//
// Please see the documentation on client-side link configuration:
// https://plaid.com/docs/quickstart/#cs-link-config
//
// Yodlee
// Returns all the necessary data to get started with account aggregation
//
// To receive the UserJWTToken, create a post request for this endpoint and include
// 'subject' in the body.
//
// Please see the documentation on client-side link configuration:
// https://developer.yodlee.com/tools
type InitiatePost struct {
	PublicKey          *string `json:"publicKey,omitempty"`
	Type               string  `json:"type"`
	JwtUserToken       *string `json:"jwtUserToken,omitempty"`
	APIURL             *string `json:"apiUrl,omitempty"`
	JwtAdminToken      *string `json:"jwtAdminToken,omitempty"`
	AccountCreationURL *string `json:"accountCreationUrl,omitempty"`
}

// Complete
//
// POST {{url}}/api/iav/complete
//
// For Plaid, Send the public token and the account id received on the client side to
// complete the relationship.
// For Yodlee, the Yodlee Unique Account Id ('externalIdentifier') is necessary and the
// publicToken is not needed.
type Complete struct {
	BankAccount          string   `json:"bankAccount"`
	BankAccountType      string   `json:"bankAccountType"`
	CreationTime         string   `json:"creationTime"`
	BankRoutingNumber    string   `json:"bankRoutingNumber"`
	Approval             Approval `json:"approval"`
	Nickname             string   `json:"nickname"`
	ApprovalMethod       string   `json:"approvalMethod"`
	ID                   string   `json:"id"`
	Account              string   `json:"account"`
	BankAccountOwnerName string   `json:"bankAccountOwnerName"`
	Status               string   `json:"status"`
}

type ListBankAccount struct {
	BankAccountID string `json:"bankAccountId"`
	Routing       string `json:"routing"`
	Type          string `json:"type"`
	Subtype       string `json:"subtype"`
	Name          string `json:"name"`
	Mask          string `json:"mask"`
	OfficialName  string `json:"officialName"`
}

// Link Bank Account
//
// POST {{url}}/api/iav/link
//
// Establish the relationship by sending the bankAccountId from the bank list and the
// trading account. Once the relationship is complete, money transfers can be done
// immediately. For Yodlee, send the Yodlee user identifier ("externalIdentifier")
// addtionally to link the account.
type LinkBankAccount struct {
	BankAccount             string    `json:"bankAccount"`
	BankAccountType         string    `json:"bankAccountType"`
	CreationTime            *string   `json:"creationTime,omitempty"`
	BankRoutingNumber       string    `json:"bankRoutingNumber"`
	Approval                *Approval `json:"approval,omitempty"`
	Nickname                string    `json:"nickname"`
	ApprovalMethod          string    `json:"approvalMethod"`
	ID                      string    `json:"id"`
	Account                 *string   `json:"account,omitempty"`
	BankAccountOwnerName    *string   `json:"bankAccountOwnerName,omitempty"`
	Status                  string    `json:"status"`
	RelationshipType        *string   `json:"relationshipType,omitempty"`
	MaskedBankAccount       *string   `json:"maskedBankAccount,omitempty"`
	Source                  *string   `json:"source,omitempty"`
	MaskedBankRoutingNumber *string   `json:"maskedBankRoutingNumber,omitempty"`
	EntryID                 *int64    `json:"entryId,omitempty"`
}

// Initiate (Deprecated)
//
// GET {{url}}/api/iav/init
//
// Initiates IAV process for supported 3rd Party Services for account aggregation. Orbis
// supports multiple IAV providers.
//
//
// Plaid
// Obtain the public key required to activate IAV on the client side.
//
// Please see the documentation on client-side link configuration:
// https://plaid.com/docs/quickstart/#cs-link-config
//
// Yodlee
// Returns all the necessary data to get started with account aggregation
//
// To receive the UserJWTToken, create a post request for this endpoint and include
// 'subject' in the body.
//
// Please see the documentation on client-side link configuration:
// https://developer.yodlee.com/tools
type InitiateDeprecated struct {
	APIURL             *string `json:"apiUrl,omitempty"`
	Type               *string `json:"type,omitempty"`
	JwtAdminToken      *string `json:"jwtAdminToken,omitempty"`
	AccountCreationURL *string `json:"accountCreationUrl,omitempty"`
	PublicKey          *string `json:"publicKey,omitempty"`
}

// Deposit
//
// POST {{url}}/api/transfer/ach/deposit
//
// Initiate an ach deposit into the specified account.
// If `instantAmount` is specified then `controlSource`, `controlTrxId`, `controlResolution`
// are required.
type Deposit struct {
	TransferTime string `json:"transferTime"`
	TransferID   string `json:"transferId"`
	Direction    string `json:"direction"`
	State        string `json:"state"`
	Amount       int64  `json:"amount"`
}

// Withdraw
//
// POST {{url}}/api/transfer/ach/withdraw
//
// Initiate an ach withdrawal to the specified bank.
type Withdraw struct {
	TransferTime string `json:"transferTime"`
	TransferID   string `json:"transferId"`
	Direction    string `json:"direction"`
	State        string `json:"state"`
	Amount       int64  `json:"amount"`
}

type TransferStatusElement struct {
	TransferID                  string       `json:"transferId"`
	ExternalTransferID          string       `json:"externalTransferId"`
	Mechanism                   string       `json:"mechanism"`
	Direction                   string       `json:"direction"`
	State                       string       `json:"state"`
	EstimatedFundsAvailableDate string       `json:"estimatedFundsAvailableDate"`
	CreatedOn                   string       `json:"createdOn"`
	UpdatedOn                   string       `json:"updatedOn"`
	Amount                      int64        `json:"amount"`
	Relationship                Relationship `json:"relationship"`
}

type Relationship struct {
	EntryID     int64  `json:"entryId"`
	ID          string `json:"id"`
	Nickname    string `json:"nickname"`
	AccountType string `json:"accountType"`
}

type TransferStatusModelElement struct {
	TransferID                  string             `json:"transferId"`
	ExternalTransferID          string             `json:"externalTransferId"`
	Mechanism                   string             `json:"mechanism"`
	Direction                   string             `json:"direction"`
	State                       string             `json:"state"`
	EstimatedFundsAvailableDate string             `json:"estimatedFundsAvailableDate"`
	Amount                      int64              `json:"amount"`
	Relationship                RealizedProfitLoss `json:"relationship"`
}

type TransferStatusBranchElement struct {
	AccountID                   int64              `json:"accountId"`
	TransferID                  string             `json:"transferId"`
	ExternalTransferID          string             `json:"externalTransferId"`
	Mechanism                   string             `json:"mechanism"`
	Direction                   string             `json:"direction"`
	State                       string             `json:"state"`
	EstimatedFundsAvailableDate string             `json:"estimatedFundsAvailableDate"`
	Amount                      int64              `json:"amount"`
	Relationship                RealizedProfitLoss `json:"relationship"`
}

// Amount Available
//
// GET {{url}}/api/transfer/available/{account}
//
// Look up amount available for withdrawal.
type AmountAvailable struct {
	Total                 int64 `json:"total"`
	UnAdjustedTotal       int64 `json:"unAdjustedTotal"`
	StartDayCashAvailable int64 `json:"startDayCashAvailable"`
}

// Confirm
//
// GET {{url}}/api/documents/confirm?account=5VA0XXXX&date=03/03/2018
//
// Load a confirm by date.
type Confirm struct {
	URL string `json:"url"`
}

// List of Confirms
//
// GET
//
// {{url}}/api/documents/confirm/list?account=5VA0XXXX&beginDate=01/01/2017&endDate=03/01/2018
//
// Load a list of confirms through a specified date-range.
type ListOfConfirms struct {
	Confirm []ConfirmElement `json:"confirm"`
}

type ConfirmElement struct {
	ProcessDate string `json:"processDate"`
	ConfirmURL  string `json:"confirmUrl"`
}

// Document list
//
// GET
//
// {{url}}/api/documents/list?account=5VA0XXXX&documentType=All&fromMonth=1/1/2021&toMonth=12/1/2021
//
// Used to load a variety of diffrerent document types ranging from account statements to
// tax forms. All forms are returned by default for the specified range. You can load a
// specific document type by passing the `documentType` param. List of available document
// types can be obtained via document type listing endpoint.
type DocumentList struct {
	DocumentList []DocumentListElement `json:"documentList"`
}

type DocumentListElement struct {
	DocID         string `json:"docID"`
	AccountNumber string `json:"accountNumber"`
	DocumentDate  string `json:"documentDate"`
	AccountName   string `json:"accountName"`
	RepID         string `json:"repId"`
	DocType       int64  `json:"docType"`
	URL           string `json:"url"`
	Insert1       string `json:"insert1"`
	Insert2       string `json:"insert2"`
	Key           string `json:"key"`
}

// Available document types and codes
//
// GET {{api}}/api/documents/typesAndCodes
//
// *DEPRECATED*
type AvailableDocumentTypesAndCodes struct {
	All              int64 `json:"All"`
	Form5498         int64 `json:"Form5498"`
	Sdira            int64 `json:"SDIRA"`
	Form1099         int64 `json:"Form1099"`
	Form1099R        int64 `json:"Form1099R"`
	Form1042S        int64 `json:"Form1042S"`
	Form5498ESA      int64 `json:"Form5498_ESA"`
	AccountStatement int64 `json:"AccountStatement"`
	Fmv              int64 `json:"FMV"`
}

type Country string

const (
	Usa Country = "USA"
)

type Currency string

const (
	Usd Currency = "USD"
)

type IdentType string

const (
	ExchangeSymbol IdentType = "ExchangeSymbol"
)

type MarketCapClassification string

const (
	Large MarketCapClassification = "Large"
	Mega  MarketCapClassification = "Mega"
	Micro MarketCapClassification = "Micro"
	Mid   MarketCapClassification = "Mid"
	Nano  MarketCapClassification = "Nano"
	Small MarketCapClassification = "Small"
)

type PutOrCall string

const (
	Empty PutOrCall = "\u0000"
)

type Mic string

const (
	Arcx Mic = "ARCX"
	Bats Mic = "BATS"
	Ootc Mic = "OOTC"
	Xngs Mic = "XNGS"
	Xnms Mic = "XNMS"
	Xnys Mic = "XNYS"
)

type EquityExchange string

const (
	ExchangeAMEX EquityExchange = "AMEX"
	ExchangeNYSE EquityExchange = "NYSE"
	Nasdaq       EquityExchange = "NASDAQ"
	OTCi         EquityExchange = "OTCi"
)

type MarketDescription string

const (
	MarketDescriptionAMEX   MarketDescription = "AMEX"
	MarketDescriptionNYSE   MarketDescription = "NYSE"
	MarketDescriptionNasdaq MarketDescription = "Nasdaq"
	Otc                     MarketDescription = "OTC"
)

type SymbolMarket string

const (
	A SymbolMarket = "A"
	N SymbolMarket = "N"
	Q SymbolMarket = "Q"
	V SymbolMarket = "V"
)

type Asset string

const (
	Alternative Asset = "Alternative"
)

type Category string

const (
	TradingInverseCommodities   Category = "Trading--Inverse Commodities"
	TradingInverseEquity        Category = "Trading--Inverse Equity"
	TradingLeveragedCommodities Category = "Trading--Leveraged Commodities"
	TradingLeveragedEquity      Category = "Trading--Leveraged Equity"
	Volatility                  Category = "Volatility"
)

type SponsorEnum string

const (
	CreditSuisseAG  SponsorEnum = "Credit Suisse AG"
	DirexionFunds   SponsorEnum = "Direxion Funds"
	ProShares       SponsorEnum = "ProShares"
	UBSGroupAG      SponsorEnum = "UBS Group AG"
	USCFInvestments SponsorEnum = "USCF Investments"
)

type Structure string

const (
	GrantorTrust                   Structure = "Grantor Trust"
	OpenEndedInvestmentCompany     Structure = "Open Ended Investment Company"
	UncollateralizedDebtInstrument Structure = "Uncollateralized Debt Instrument"
)

type StructureType string

const (
	Etf StructureType = "ETF"
)

type Duration string

const (
	The1D Duration = "1d"
)

type TitleEnum string

const (
	CoManager       TitleEnum = "Co Manager"
	JointBookrunner TitleEnum = "Joint Bookrunner"
)

type Status string

const (
	Priced Status = "PRICED"
)

type Updated string

const (
	The07122019000000Edt Updated = "07/12/2019 00:00:00 EDT"
)

type ShareChangePercent string

const (
	NaN ShareChangePercent = "NaN"
)

type Time string

const (
	AfterClose Time = "AFTER_CLOSE"
	During     Time = "DURING"
	Other      Time = "OTHER"
	PriorOpen  Time = "PRIOR_OPEN"
)

type IndustryType string

const (
	Group IndustryType = "GROUP"
)

type Source string

const (
	Cboe1    Source = "CBOE1"
	Realtime Source = "REALTIME"
)

type CreatedOn string

const (
	The11082017061144Est CreatedOn = "11/08/2017 06:11:44 EST"
	The11092017061006Est CreatedOn = "11/09/2017 06:10:06 EST"
	The11092017061007Est CreatedOn = "11/09/2017 06:10:07 EST"
	The11092017061008Est CreatedOn = "11/09/2017 06:10:08 EST"
	The11092017061009Est CreatedOn = "11/09/2017 06:10:09 EST"
	The11092017061010Est CreatedOn = "11/09/2017 06:10:10 EST"
	The11092017061011Est CreatedOn = "11/09/2017 06:10:11 EST"
	The11092017061012Est CreatedOn = "11/09/2017 06:10:12 EST"
	The11092017061013Est CreatedOn = "11/09/2017 06:10:13 EST"
	The11092017061014Est CreatedOn = "11/09/2017 06:10:14 EST"
	The11092017061015Est CreatedOn = "11/09/2017 06:10:15 EST"
	The11092017061016Est CreatedOn = "11/09/2017 06:10:16 EST"
	The11092017061017Est CreatedOn = "11/09/2017 06:10:17 EST"
	The11092017061018Est CreatedOn = "11/09/2017 06:10:18 EST"
	The11092017061019Est CreatedOn = "11/09/2017 06:10:19 EST"
	The11092017061150Est CreatedOn = "11/09/2017 06:11:50 EST"
)

type LiquidationTransaction string

const (
	Sell LiquidationTransaction = "SELL"
)

type PurpleExchange string

const (
	Undefined PurpleExchange = "UNDEFINED"
)

type Symbol string

const (
	Exg Symbol = "EXG"
	Voo Symbol = "VOO"
)
