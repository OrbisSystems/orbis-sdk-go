package orbis_api

type PortfolioRequest struct {
	LoadQuotes         *bool
	LoadIndustries     *bool
	LoadUpgrades       *bool
	LoadEarnings       *bool
	LoadRtb            *bool
	MarketCapBreakDown *bool
	Combined           *bool
}

type OptionExpirationsDetailsRequest struct {
	Maturity   string `json:"maturity"`
	LoadQuotes string `json:"loadQuotes"`
}
