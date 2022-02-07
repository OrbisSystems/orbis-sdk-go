package orbis_api

type CancelReplaceReq struct {
	Order OrderCancelReplace `json:"order"`
}

type CancellationReq struct {
	Order struct {
		OrderRef string `json:"orderRef"`
	} `json:"order"`
}

type OrderCancelReplace struct {
	OrderRef   string  `json:"orderRef"`
	MarketCode string  `json:"marketCode"`
	StopPrice  float64 `json:"stopPrice"`
	LimitPrice float64 `json:"limitPrice"`
	Quantity   int64   `json:"quantity"`
}
