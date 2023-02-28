package model

import "time"

type CompanyProfile struct {
	Symbol              string `json:"symbol"`
	CompanyName         string `json:"company_name"`
	ShortDescription    string `json:"short_description"`
	LongDescription     string `json:"long_description"`
	Address             string `json:"address"`
	City                string `json:"city"`
	CountryCode         string `json:"country_code"`
	Country             string `json:"country"`
	Phone               string `json:"phone"`
	Fax                 string `json:"fax"`
	Website             string `json:"website"`
	Email               string `json:"email"`
	TotalEmployeeNumber int    `json:"employees_number"`

	Logo             *SymbolLogosResponse `json:"logo,omitempty"`
	SymbolIndustries *SymbolIndustries    `json:"sector,omitempty"`
	EarningRelease   *EarningRelease      `json:"earning_release,omitempty"`

	MarketCap float64 `json:"market_cap"`

	InstitutionalOwnershipPercent float64 `json:"institutional_ownership_percent,omitempty"`
	InsiderOwnershipPercent       float64 `json:"insider_ownership_percent,omitempty"`
	FundOwnershipPercent          float64 `json:"fund_ownership_percent,omitempty"`
}

type SymbolIndustries struct {
	Sector        string `json:"sector"`
	Industry      string `json:"industry"`
	IndustryGroup string `json:"industry_group"`
	SubIndustry   string `json:"sub_industry"`
}

type EarningRelease struct {
	Symbol   string `json:"symbol,omitempty"`
	Date     string `json:"date"`
	Quarter  string `json:"quarter"`
	Status   string `json:"status"`
	Exchange string `json:"exchange"`
}

type GetOwnershipsBySymbolRequest struct {
	Symbols   []string `json:"symbols"`
	OwnerType string   `json:"owner_type"`

	OrderBy *Ordering `json:"order,omitempty"`
	Paging  *Paging   `json:"paging,omitempty"`
}

type GetOwnershipsBySymbolResponse struct {
	Details []OwnershipDetails `json:"details"`
	Count   int                `json:"count"`
}

type OwnershipDetails struct {
	Symbol         string  `json:"symbol"`
	OwnerType      string  `json:"owner_type"`
	Date           string  `json:"date"`
	OwnerCIK       int     `json:"owner_cik"`
	OwnerName      string  `json:"owner_name"`
	NumberOfShares int     `json:"number_of_shares"`
	ShareChange    float64 `json:"share_change"`
}

type GetOwnershipsByIDRequest struct {
	OwnerID string `json:"owner_id"`

	Paging *Paging `json:"paging,omitempty"`
}

type EarningReleasesRequest struct {
	Symbols     []string `json:"symbols"`
	Exchanges   []string `json:"exchanges"`
	DaysFromNow int      `json:"days_from_now"`

	Paging *Paging `json:"paging,omitempty"`
}

type EarningReleasesResponse struct {
	Details []EarningRelease `json:"details"`
	Count   int              `json:"count"`
}

type GetSymbolFundamentalsRequest struct {
	Symbol        string    `json:"symbol"`
	Mode          string    `json:"mode"`
	ReferenceTime time.Time `json:"-"`
}

type GetSymbolFundamentalsResponse map[string]interface{}

type StockScreenerRequest struct {
	Price                *LessAndMore    `json:"price"`
	MarketCapitalization *LessAndMore    `json:"market_cap"`
	PriceToEarningsRatio *LessAndMore    `json:"pe_ratio"`
	DividendPercent      *LessAndMore    `json:"dividend_percent"`
	DividendPayDate      *BeforeAndAfter `json:"dividend_pay_date"`
	DividendExDate       *BeforeAndAfter `json:"dividend_ex_date"`
	Exchange             *IncludeExclude `json:"exchange"`
	Volume               *LessAndMore    `json:"volume"`
	AverageVolume        *LessAndMore    `json:"average_volume"`
	Gap                  *LessAndMore    `json:"gap"`
	StockType            *string         `json:"stock_type"`
	Industry             *string         `json:"industry"`
	QuoteType            *string         `json:"quote_type"`

	Paging *Paging `json:"paging,omitempty"`
}

type BeforeAndAfter struct {
	After  *string `json:"after"`
	Before *string `json:"before"`
}

type IncludeExclude struct {
	Include []string `json:"include"`
	Exclude []string `json:"exclude"`
}

type StockScreenerResponse struct {
	Details []StockScreenerQuote `json:"details"`
	Count   int                  `json:"count"`
}

type StockScreenerQuote struct {
	Symbol               string  `json:"symbol"`
	CompanyName          string  `json:"company_name"`
	LastPrice            float64 `json:"last_price"`
	PriceChange          float64 `json:"price_change"`
	PriceChangePercent   float64 `json:"price_change_percent"`
	Volume               float64 `json:"volume"`
	MarketCapitalization float64 `json:"market_capitalization"`
	DividendPercent      float64 `json:"dividend_percent"`
	DividendExDate       string  `json:"dividend_ex_date"`
	Exchange             string  `json:"exchange"`
	AverageVolume30d     float64 `json:"average_volume_30d"`
}

type StockMarketHeatmapResponse struct {
	MinChange float64                   `json:"min_change"`
	MaxChange float64                   `json:"max_change"`
	Details   []StockMarketHeatmapEntry `json:"details"`
}

type StockMarketHeatmapEntry struct {
	Symbol             string  `json:"symbol"`
	Title              string  `json:"title"`
	PriceChangePercent float64 `json:"price_change_percent"`
}

type GetIndustriesPerformanceRequest struct {
	RankingType string `json:"ranking_type"` // price-weighted, market-weighted
	QuoteType   string `json:"quote_type"`

	QuoteFilters GetIndustriesPerformanceQuoteFilter `json:"quote_filters"`
}

type GetIndustriesPerformanceQuoteFilter struct {
	Limit     int          `json:"limit"`
	Price     *LessAndMore `json:"price"`
	Exchanges string       `json:"exchanges"` // ALL OR US(NYSE/Nasdaq/Amex)
	Order     string       `json:"order"`     // last_price, price_change, volume
}

type GetIndustriesPerformanceResponse struct {
	Sectors    []Sector        `json:"sectors"`
	Industries []IndustryGroup `json:"industries"`

	BestPerformingQuotes  []StockScreenerQuote `json:"best_performing_quotes"`
	WorstPerformingQuotes []StockScreenerQuote `json:"worst_performing_quotes"`
}

type Sector struct {
	SectorCode                   string          `json:"sector_code"`
	SectorName                   string          `json:"sector_name"`
	SectorMarketCap              float64         `json:"sector_market_cap"`
	SectorMarketCap1DayPercent   float64         `json:"sector_market_cap_1day"`
	SectorMarketCap1WeekPercent  float64         `json:"sector_market_cap_1week"`
	SectorMarketCap2WeekPercent  float64         `json:"sector_market_cap_2week"`
	SectorMarketCap1MonthPercent float64         `json:"sector_market_cap_1month"`
	IndustryGroups               []IndustryGroup `json:"industry_groups"`
}

type IndustryGroup struct {
	IndustryGroupCode                   string  `json:"industry_group_code"`
	IndustryGroupName                   string  `json:"industry_group_name"`
	IndustryGroupMarketCap              float64 `json:"industry_group_market_cap"`
	IndustryGroupMarketCap1DayPercent   float64 `json:"industry_group_market_cap_1day"`
	IndustryGroupMarketCap1WeekPercent  float64 `json:"industry_group_market_cap_1week"`
	IndustryGroupMarketCap2WeekPercent  float64 `json:"industry_group_market_cap_2week"`
	IndustryGroupMarketCap1MonthPercent float64 `json:"industry_group_market_cap_1month"`
}

type MomentumRatioGraphRequest struct {
	Symbols   []string `json:"symbols"`
	Benchmark string   `json:"benchmark"`
	Range     string   `json:"range"`
	Period    int      `json:"period"`
	SetEma    bool     `json:"set_ema"`
	Mode      string   `json:"mode"`
}

type MomentumRatioGraphResponse struct {
	AAPL []SymbolMomentum `json:"AAPL"`
	CSCO []SymbolMomentum `json:"CSCO"`
}

type SymbolMomentum struct {
	ResultRatio    float64 `json:"result_ratio"`
	ResultMomentum float64 `json:"result_momentum"`
	Date           string  `json:"date"`
}

type SeasonalityRequest struct {
	Lookahead int      `json:"lookahead"`
	Benchmark string   `json:"benchmark"`
	Symbols   []string `json:"symbols"`
	Years     int      `json:"years"`
}

type SeasonalityResponse map[string]SeasonalityEntry

type SeasonalityEntry struct {
	Score     int               `json:"score"`
	January   TickerSeasonality `json:"january"`
	June      TickerSeasonality `json:"june"`
	May       TickerSeasonality `json:"may"`
	October   TickerSeasonality `json:"october"`
	December  TickerSeasonality `json:"december"`
	March     TickerSeasonality `json:"march"`
	February  TickerSeasonality `json:"february"`
	August    TickerSeasonality `json:"august"`
	July      TickerSeasonality `json:"july"`
	September TickerSeasonality `json:"september"`
	November  TickerSeasonality `json:"november"`
	April     TickerSeasonality `json:"april"`
}

type TickerSeasonality struct {
	Success          float32      `json:"success"`
	Failure          float32      `json:"failure"`
	Average          float32      `json:"average"`
	YearlyData       []YearlyData `json:"yearly_data"`
	TotalYears       int          `json:"total_years"`
	BenchmarkAverage float32      `json:"benchmark_average"`
	Difference       float32      `json:"difference"`
	Score            int          `json:"score"`
}

type YearlyData struct {
	Year             int     `json:"year"`
	Date             int64   `json:"date"`
	Value            float32 `json:"value"`
	BenchmarkAverage float32 `json:"benchmark_average"`
	Difference       float32 `json:"difference"`
}
