package orbis_api

type IndexComponents []IndexComponent

type CompanyResearch struct {
	Address                         Address           `json:"address"`
	StockSummary                    StockSummary      `json:"stockSummary"`
	EarningEstimates                EarningEstimates  `json:"earningEstimates"`
	BusinessSummary                 string            `json:"businessSummary"`
	LongSummary                     string            `json:"longSummary"`
	Phone                           string            `json:"phone"`
	Fax                             string            `json:"fax"`
	URL                             string            `json:"url"`
	Email                           string            `json:"email"`
	EmployeeDescription             string            `json:"employeeDescription"`
	EmployeeAsOf                    string            `json:"employeeAsOf"`
	EmployeeCount                   int64             `json:"employeeCount"`
	ShareHolderCount                int64             `json:"shareHolderCount"`
	ShareholderAsOf                 string            `json:"shareholderAsOf"`
	AnnualMeeting                   string            `json:"annualMeeting"`
	IncorpCountry                   string            `json:"incorpCountry"`
	IncorpState                     string            `json:"incorpState"`
	IncorpDate                      string            `json:"incorpDate"`
	IssueDescription                string            `json:"issueDescription"`
	IssueReference                  string            `json:"issueReference"`
	IssuePrimaryExchange            string            `json:"issuePrimaryExchange"`
	IssuePrimaryTicker              string            `json:"issuePrimaryTicker"`
	IssueIADCurrency                string            `json:"issueIADCurrency"`
	OutStandingDPFC                 string            `json:"outStandingDPFC"`
	OutStandingAsOf                 string            `json:"outStandingAsOf"`
	OutStandingShares               int64             `json:"outStandingShares"`
	ShareOutstandingAmount          int64             `json:"shareOutstandingAmount"`
	SharesOutstandingDate           string            `json:"sharesOutstandingDate"`
	FileDate                        string            `json:"fileDate"`
	ReportDate                      string            `json:"reportDate"`
	ShareOutstandingPrimaryExchange string            `json:"shareOutstandingPrimaryExchange"`
	ShareOutstandingPrimaryTicker   string            `json:"shareOutstandingPrimaryTicker"`
	FidelitySector                  string            `json:"fidelitySector"`
	FidelityIndustry                string            `json:"fidelityIndustry"`
	GlobalSector                    string            `json:"globalSector"`
	GlobalGroup                     string            `json:"globalGroup"`
	GlobalIndustry                  string            `json:"globalIndustry"`
	GlobalSubIndustry               string            `json:"globalSubIndustry"`
	GlobalDescription               string            `json:"globalDescription"`
	GlobalSectorCode                int64             `json:"globalSectorCode"`
	GlobalGroupCode                 int64             `json:"globalGroupCode"`
	GlobalIndustryCode              int64             `json:"globalIndustryCode"`
	GlobalSubIndustryCode           int64             `json:"globalSubIndustryCode"`
	InsiderOwnPct                   float64           `json:"insiderOwnPct"`
	InstitutionalOwnPC              float64           `json:"institutionalOwnPc"`
	SectorRanking                   Ranking           `json:"sectorRanking"`
	IndustryRanking                 Ranking           `json:"industryRanking"`
	IndustryMarketCap               float64           `json:"industryMarketCap"`
	IndustryMaxRank                 int64             `json:"industryMaxRank"`
	BusinessSummaryNew              BusinessShort     `json:"businessSummary_New"`
	CapitalHistory                  BusinessShort     `json:"capitalHistory"`
	DividendsPaid                   BusinessShort     `json:"dividendsPaid"`
	PriceRange                      BusinessShort     `json:"priceRange"`
	BusinessShort                   BusinessShort     `json:"businessShort"`
	StockSplits                     BusinessShort     `json:"stockSplits"`
	Summaries                       map[string]string `json:"summaries"`
}

type Address struct {
	Streets []string `json:"streets"`
	City    string   `json:"city"`
	State   string   `json:"state"`
	Zipcode string   `json:"zipcode"`
	Country string   `json:"country"`
}

type BusinessShort struct {
	Date  string `json:"date"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

type EarningEstimates struct {
	Estimates      map[string]Quarter1 `json:"estimates"`
	LongTermGrowth LongTermGrowth      `json:"longTermGrowth"`
	Num            int64               `json:"num"`
	Quarter1       Quarter1            `json:"quarter1"`
	Quarter2       Quarter1            `json:"quarter2"`
	Year1          Quarter1            `json:"year1"`
	Year2          Quarter1            `json:"year2"`
}

type Quarter1 struct {
	PreviousYearActual *float64 `json:"previousYearActual,omitempty"`
	Count              int64    `json:"count"`
	Change             float64  `json:"change"`
	High               float64  `json:"high"`
	Low                float64  `json:"low"`
	Mean               float64  `json:"mean"`
	End                string   `json:"end"`
	Actual             *float64 `json:"actual,omitempty"`
	Estimate           *float64 `json:"estimate,omitempty"`
}

type LongTermGrowth struct {
	Count     int64   `json:"count"`
	High      float64 `json:"high"`
	Low       int64   `json:"low"`
	Estimates float64 `json:"estimates"`
}

type Ranking struct {
	Industry            Industry `json:"industry"`
	PriceRank           int64    `json:"priceRank"`
	PriceChange         float64  `json:"priceChange"`
	PriceChangeOneWeek  float64  `json:"priceChangeOneWeek"`
	PriceChangeTwoWeeks float64  `json:"priceChangeTwoWeeks"`
	PriceChangeOneMonth float64  `json:"priceChangeOneMonth"`
	MarketCap           float64  `json:"marketCap"`
	MarketCapChange     float64  `json:"marketCapChange"`
}

type StockSummary struct {
	BusinessShort   BusinessShort `json:"BUSINESS_SHORT"`
	DividendsPaid   BusinessShort `json:"DIVIDENDS_PAID"`
	BusinessLong    BusinessShort `json:"BUSINESS_LONG"`
	BusinessSummary BusinessShort `json:"BUSINESS_SUMMARY"`
	PriceRange      BusinessShort `json:"PRICE_RANGE"`
	StockSplits     BusinessShort `json:"STOCK_SPLITS"`
	CapitalHistory  BusinessShort `json:"CAPITAL_HISTORY"`
}

type IndexComponent struct {
	Symbol                  string  `json:"symbol"`
	CompanyName             string  `json:"companyName"`
	Isin                    string  `json:"isin"`
	Cusip                   string  `json:"cusip"`
	Sedol                   string  `json:"sedol"`
	Country                 string  `json:"country"`
	Mic                     string  `json:"mic"`
	Currency                string  `json:"currency"`
	Source                  string  `json:"source"`
	IdentType               string  `json:"identType"`
	Exchange                string  `json:"exchange"`
	BidRegion               Region  `json:"bidRegion"`
	AskRegion               Region  `json:"askRegion"`
	TradeRegion             Region  `json:"tradeRegion"`
	UpdateTime              string  `json:"updateTime"`
	TradeTime               string  `json:"tradeTime"`
	DividendPaymentDate     string  `json:"dividendPaymentDate"`
	DividendExDate          string  `json:"dividendExDate"`
	CalendarYearHighDate    string  `json:"calendarYearHighDate"`
	CalendarYearLowDate     string  `json:"calendarYearLowDate"`
	Week52HighDate          string  `json:"week52HighDate"`
	Week52LowDate           string  `json:"week52LowDate"`
	Vwap                    float64 `json:"vwap"`
	Vwap1                   int64   `json:"vwap1"`
	YesterdayClose          float64 `json:"yesterdayClose"`
	OpeningPrice            float64 `json:"openingPrice"`
	StrikePrice             int64   `json:"strikePrice"`
	DayHigh                 float64 `json:"dayHigh"`
	DayLow                  float64 `json:"dayLow"`
	Ask                     float64 `json:"ask"`
	Bid                     float64 `json:"bid"`
	AskYield                int64   `json:"askYield"`
	BidYield                int64   `json:"bidYield"`
	LastPrice               float64 `json:"lastPrice"`
	Change                  float64 `json:"change"`
	OpenInterest            int64   `json:"openInterest"`
	High52Week              float64 `json:"high52week"`
	Low52Week               float64 `json:"low52week"`
	PriceEarningRatio       float64 `json:"priceEarningRatio"`
	ChangePercent           float64 `json:"changePercent"`
	MarketCap               float64 `json:"marketCap"`
	CalendarYearHigh        float64 `json:"calendarYearHigh"`
	CalendarYearLow         float64 `json:"calendarYearLow"`
	DividendYield           float64 `json:"dividendYield"`
	DividendAmount          float64 `json:"dividendAmount"`
	DividendRate            float64 `json:"dividendRate"`
	Beta                    float64 `json:"beta"`
	TotalCashAmount         int64   `json:"totalCashAmount"`
	EstimatedCashAmount     int64   `json:"estimatedCashAmount"`
	IntradayValue           int64   `json:"intradayValue"`
	Nav                     int64   `json:"nav"`
	TradePrice              float64 `json:"tradePrice"`
	YesterdayVolume         int64   `json:"yesterdayVolume"`
	AskSize                 int64   `json:"askSize"`
	BidSize                 int64   `json:"bidSize"`
	MinAskSize              int64   `json:"minAskSize"`
	MinBidSize              int64   `json:"minBidSize"`
	Volume                  int64   `json:"volume"`
	ContractSize            int64   `json:"contractSize"`
	AverageVolume30         int64   `json:"averageVolume30"`
	ImbalanceVolume         int64   `json:"imbalanceVolume"`
	SharesOutstanding       int64   `json:"sharesOutstanding"`
	Increment               int64   `json:"increment"`
	TradeSize               int64   `json:"tradeSize"`
	PutOrCall               string  `json:"putOrCall"`
	Indicator               int64   `json:"indicator"`
	SubMarket               int64   `json:"subMarket"`
	BidTick                 int64   `json:"bidTick"`
	Precision               int64   `json:"precision"`
	Delayed                 bool    `json:"delayed"`
	NotFound                bool    `json:"notFound"`
	Option                  bool    `json:"option"`
	NotPermissioned         bool    `json:"notPermissioned"`
	Halted                  bool    `json:"halted"`
	Paused                  bool    `json:"paused"`
	Mini                    bool    `json:"mini"`
	CorpAct                 bool    `json:"corpAct"`
	Jumbo                   bool    `json:"jumbo"`
	MutualFund              bool    `json:"mutualFund"`
	Bond                    bool    `json:"bond"`
	Nbbo                    bool    `json:"nbbo"`
	Crypto                  bool    `json:"crypto"`
	Fx                      bool    `json:"fx"`
	SymbolMarket            string  `json:"symbolMarket"`
	Call                    bool    `json:"call"`
	MarketCapClassification string  `json:"marketCapClassification"`
	DividentSoon            bool    `json:"dividentSoon"`
	SymbolChange            bool    `json:"symbolChange"`
	MarketDescription       string  `json:"marketDescription"`
	AtTheMoney              bool    `json:"atTheMoney"`
	InTheMoney              bool    `json:"inTheMoney"`
	Trailing12MonthsEps     float64 `json:"trailing12MonthsEps"`
	Put                     bool    `json:"put"`
}

type Fundamentals struct {
	FiscalYearEnd int64     `json:"FiscalYearEnd"`
	CurrencyID    string    `json:"CurrencyId"`
	LastUpdated   string    `json:"LastUpdated"`
	V             int64     `json:"V"`
	The2016Q3     The2016Q3 `json:"2016Q3"`
}

type The2016Q3 struct {
	Date          string   `json:"Date"`
	The12M        struct{} `json:"12M"`
	Amortization  int64    `json:"Amortization"`
	CostOfRevenue int64    `json:"CostOfRevenue"`
}

type ScreenerResearch []ScreenerElement

type ScreenerElement struct {
	Field      string     `json:"field"`
	Comparator string     `json:"comparator"`
	Value      *Value     `json:"value"`
	DateRange  *DateRange `json:"dateRange,omitempty"`
}

type DateRange struct {
	LowerBound string `json:"lowerBound"`
	UpperBound string `json:"upperBound"`
}

type Value struct {
	Bool    *bool
	Integer *int64
}
