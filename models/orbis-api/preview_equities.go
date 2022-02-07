package orbis_api

type EquitiesRequest struct {
	RequireNbbo bool    `json:"requireNbbo"`
	Account     Account `json:"account"`
	Order       Order   `json:"order"`
}

type Account struct {
	AccountID     *int64 `json:"accountId"`
	AccountNumber string `json:"accountNumber,omitempty"`
}

type Order struct {
	Quote         Quote   `json:"quote"`
	Transaction   string  `json:"transaction"`
	MarketTime    string  `json:"marketTime"`
	OrderType     string  `json:"orderType"`
	FillType      string  `json:"fillType"`
	Validity      string  `json:"validity"`
	Quantity      int64   `json:"quantity"`
	LimitPrice    int64   `json:"limitPrice"`
	StopPrice     int64   `json:"stopPrice"`
	ExpectedValue float64 `json:"expectedValue"`
}

type Quote struct {
	Symbol string `json:"symbol"`
}
