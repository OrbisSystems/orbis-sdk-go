package model

import "github.com/google/uuid"

type GetFundDetailsResponse struct {
	Symbol       string  `json:"symbol"`
	ExpenseRatio float64 `json:"expense_ratio"`
	Structure    string  `json:"structure"`
	Leverage     float64 `json:"leverage"`
	Inverse      bool    `json:"inverse"`
	Asset        string  `json:"asset"`
	Category     string  `json:"category"`

	City        string `json:"city"`
	Address     string `json:"address"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	Phone       string `json:"phone"`
	Fax         string `json:"fax"`
	Website     string `json:"website"`
	Email       string `json:"email"`

	FundType string `json:"fund_type"`

	Maturities       []MaturityBreakdown        `json:"maturities"`
	Sectors          []SectorBreakdown          `json:"sectors"`
	Countries        []CountryBreakdown         `json:"countries"`
	Holdings         []HoldingBreakdown         `json:"holdings"`
	MarketCaps       []MarketCapBreakdown       `json:"market_caps"`
	Continents       []ContinentBreakdown       `json:"continents"`
	AssetAllocations []AssetAllocationBreakdown `json:"asset_allocations"`
}

type MaturityBreakdown struct {
	Breakdown  string  `json:"breakdown"`
	Percentage float64 `json:"percentage"`
}

type SectorBreakdown struct {
	Code       string  `json:"code"`
	Breakdown  string  `json:"breakdown"`
	Percentage float64 `json:"percentage"`
}

type CountryBreakdown struct {
	Breakdown  string  `json:"breakdown"`
	Percentage float64 `json:"percentage"`
}

type HoldingBreakdown struct {
	Symbol      string  `json:"symbol"`
	CompanyName string  `json:"company_name"`
	Rank        int     `json:"rank"`
	Percentage  float64 `json:"percentage"`
}

type MarketCapBreakdown struct {
	Breakdown  string  `json:"breakdown"`
	Percentage float64 `json:"percentage"`
}

type ContinentBreakdown struct {
	Code       string  `json:"code"`
	Breakdown  string  `json:"breakdown"`
	Percentage float64 `json:"percentage"`
}

type AssetAllocationBreakdown struct {
	Breakdown  string  `json:"breakdown"`
	Percentage float64 `json:"percentage"`
}

type GetFundScreenerFiltersResponse struct {
	Assets     []FundFilterAsset    `json:"assets"`
	Categories []FundFilterCategory `json:"categories"`
	Sponsors   []FundFilterSponsor  `json:"sponsors"`
	Inverse    []string             `json:"inverse"`
	Leverage   []string             `json:"leverage"`
	Sectors    []FundSector         `json:"sectors"`
}

type FundSector struct {
	SectorID   string `json:"sector_id"`
	SectorName string `json:"sector_name"`
}

type FundFilterAsset struct {
	AssetID   uuid.UUID `json:"asset_id"`
	AssetName string    `json:"asset_name"`
}

type FundFilterCategory struct {
	CategoryID   uuid.UUID `json:"category_id"`
	CategoryName string    `json:"category_name"`
}

type FundFilterSponsor struct {
	SponsorID       uuid.UUID `json:"sponsor_id"`
	SponsorName     string    `json:"sponsor_name"`
	NumberOfEntries int       `json:"number_of_entries"`
}

type FundScreenerRequest struct {
	Assets     []uuid.UUID `json:"assets"`
	Categories []uuid.UUID `json:"categories"`
	Sponsors   []uuid.UUID `json:"sponsors"`
	Inverse    string      `json:"inverse"`
	Leverage   string      `json:"leverage"`
	FundType   string      `json:"fund_type"`
	QuoteType  string      `json:"quote_type"`

	Paging *Paging `json:"paging,omitempty"`
}

type FundScreenerResponse struct {
	Details []FundScreenerQuote `json:"details"`
	Count   int                 `json:"count"`
}

type FundScreenerQuote struct {
	Leverage     float64 `json:"leverage"`
	Inverse      bool    `json:"inverse"`
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	Category     string  `json:"category"`
	ExpenseRatio float64 `json:"expense_ratio"`
	LastPrice    float64 `json:"last_price"`
	PriceChange  float64 `json:"price_change"`
	FundType     string  `json:"fund_type"`
}

type FundSectorScreenerRequest struct {
	SectorCode      string       `json:"sector_code"`
	Exposure        *LessAndMore `json:"exposure"`
	DividendPercent *LessAndMore `json:"dividend_percent"`
	Sponsors        []uuid.UUID  `json:"sponsors"`
	Inverse         string       `json:"inverse"`
	Leverage        string       `json:"leverage"`
	Holdings        []string     `json:"holdings"`

	QuoteType string `json:"quote_type"`

	Ordering *Ordering `json:"ordering,omitempty"`
	Paging   *Paging   `json:"paging,omitempty"`
}

type LessAndMore struct {
	MoreThan *float64 `json:"more_than"`
	LessThan *float64 `json:"less_than"`
}

type GetTopFundsRequest struct {
	Inverse      string `json:"inverse"`
	Leverage     string `json:"leverage"`
	FundType     string `json:"fund_type"`
	QuoteType    string `json:"quote_type"`
	ReverseOrder bool   `json:"reverse_order"`

	Paging *Paging `json:"paging,omitempty"`
}

type GetFundsForHoldingRequest struct {
	Symbol          string          `json:"symbol"`
	Leverage        string          `json:"leverage"`
	QuoteAccessType QuoteAccessType `json:"quote_access_type"`

	Ordering *Ordering `json:"order,omitempty"`
	Paging   *Paging   `json:"paging"`
}

type GetFundsForHoldingResponse struct {
	Data  []FundForHoldingWithQuote `json:"data"`
	Count int                       `json:"count"`
}

type FundForHoldingWithQuote struct {
	Leverage       float64 `json:"leverage"`
	Inverse        bool    `json:"inverse"`
	Symbol         string  `json:"symbol"`
	Name           string  `json:"name"`
	HoldingPercent float64 `json:"holding_percent"`
	LastPrice      float64 `json:"last_price"`
	PriceChange    float64 `json:"price_change"`
	Volume         float64 `json:"volume"`
	Sponsor        string  `json:"sponsor"`
	MarketCap      float64 `json:"market_cap"`
	FundType       string  `json:"fund_type"`
}
