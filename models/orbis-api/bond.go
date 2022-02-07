package orbis_api

type BondPlacementReq struct {
	Quote       QuoteBond `json:"quote"`
	Quantity    int64     `json:"quantity"`
	Transaction string    `json:"transaction"`
	Price       int64     `json:"price"`
}

type QuoteBond struct {
	Cusip   string `json:"cusip"`
	QuoteID string `json:"quoteId"`
}

type QuoteBondPlacementReq struct {
	Quote       Quote  `json:"quote"`
	Quantity    int64  `json:"quantity"`
	Transaction string `json:"transaction"`
}

type QuoteBondPlacement struct {
	Cusip string `json:"cusip"`
}

type BondQuoteAcceptanceReq struct {
	OrderRef string `json:"orderRef"`
}
