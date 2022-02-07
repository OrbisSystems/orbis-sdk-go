package orbis_api

type OptionsPreviewReq struct {
	Account AccountOptions `json:"account"`
	Order   OrderOptions   `json:"order"`
}

type AccountOptions struct {
	AccountID     int64  `json:"accountId"`
	AccountNumber string `json:"accountNumber"`
}

type OrderOptions struct {
	Quote       Quote  `json:"quote"`
	Transaction string `json:"transaction"`
	MarketTime  string `json:"marketTime"`
	OrderType   string `json:"orderType"`
	FillType    string `json:"fillType"`
	Validity    string `json:"validity"`
	Quantity    int64  `json:"quantity"`
	LimitPrice  int64  `json:"limitPrice"`
	StopPrice   int64  `json:"stopPrice"`
}

type MultilegStrategyReq struct {
	Account AccountOptions `json:"account"`
	Order   OrderMultiLeg  `json:"order"`
}

type PreviewMultilegReq struct {
	RequiredNbbo bool `json:"requiredNbbo"`
	MultilegStrategyReq
}

type OrderMultiLeg struct {
	SpreadPrice int64  `json:"spreadPrice"`
	Validity    string `json:"validity"`
	OrderType   string `json:"orderType"`
	Legs        []Leg  `json:"legs"`
}

type Leg struct {
	Quantity    int64         `json:"quantity"`
	Quote       QuoteMultileg `json:"quote"`
	Transaction string        `json:"transaction"`
}

type QuoteMultileg struct {
	Symbol string `json:"symbol"`
	Option bool   `json:"option"`
}

type ReplacePreviewReq struct {
	RequireNbbo bool                `json:"requireNbbo"`
	Order       OrderReplacePreview `json:"order"`
}

type OrderReplacePreview struct {
	OrderRef    string              `json:"orderRef"`
	SpreadPrice int64               `json:"spreadPrice"`
	OrderType   string              `json:"orderType"`
	Legs        []LegReplacePreview `json:"legs"`
}

type LegReplacePreview struct {
	Quantity int64 `json:"quantity"`
}
