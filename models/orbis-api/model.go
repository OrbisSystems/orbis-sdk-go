package orbis_api

type ModelPortfolioRequest struct {
	LoadQuotes         *bool
	LoadIndustries     *bool
	LoadUpgrades       *bool
	LoadEarnings       *bool
	LoadCost           *bool
	MarketCapBreakDown *bool
	Combined           *bool
}

type UpdateAdjustmentRequest struct {
	ID        int64   `json:"id"`
	ModelID   int64   `json:"modelId"`
	SourcePct float64 `json:"sourcePct"`
	TargetPct float64 `json:"targetPct"`
	Quote     Quote   `json:"quote"`
}

type CreateModelRequest struct {
	Title                 string      `json:"title"`
	Description           string      `json:"description"`
	Minimum               int64       `json:"minimum"`
	Reserve               int64       `json:"reserve"`
	ReserveType           string      `json:"reserveType"`
	FractionalSharePolicy string      `json:"fractionalSharePolicy"`
	Components            []Component `json:"components"`
}

type Component struct {
	Symbol     string  `json:"symbol"`
	Percentage float64 `json:"percentage"`
	Tag        string  `json:"tag"`
	Note       string  `json:"note"`
}

type UpdateModelRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Minimum     int64  `json:"minimum"`
	Reserve     int64  `json:"reserve"`
	ReserveType string `json:"reserveType"`
}

type UpdateComponentRequest struct {
	Note       string  `json:"note"`
	Tag        string  `json:"tag"`
	Symbol     string  `json:"symbol"`
	Percentage float64 `json:"percentage"`
	Option     bool    `json:"option"`
	Type       string  `json:"type"`
	Model      Model   `json:"model"`
}

type LinkOfMembersRequest struct {
	AccountID int64 `json:"accountId"`
	ModelID   int64 `json:"modelId"`
}

type Model struct {
	ID int64 `json:"id"`
}

type AllocationModelRequest struct {
	Orders  []OrderAllocation `json:"orders"`
	Targets []Target          `json:"targets"`
}

type OrderAllocation struct {
	OurRef   string `json:"ourRef"`
	Quantity int64  `json:"quantity"`
}
