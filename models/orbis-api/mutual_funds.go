package orbis_api

type MutualFundsRequest struct {
	Account AccountMutualFund `json:"account"`
	Order   OrderMutualFund   `json:"order"`
}

type AccountMutualFund struct {
	AccountNumber string `json:"accountNumber"`
}

type OrderMutualFund struct {
	Value        int64   `json:"value"`
	Quantity     float64 `json:"quantity"`
	Quote        Quote   `json:"quote"`
	Transaction  string  `json:"transaction"`
	CapitalGains string  `json:"capitalGains"`
	Dividends    string  `json:"dividends"`
}
