package orbis_api

type ManageGroupRequest struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	AccountIDS  []int64 `json:"accountIds"`
}

type PortfolioGroupRequest struct {
	LoadQuotes       *bool
	LoadIndustries   *bool
	LoadUpgrades     *bool
	LoadEarnings     *bool
	LoadCost         *bool
	MarketCapWeights *bool
	Combined         *bool
}
