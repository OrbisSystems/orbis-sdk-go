package orbis_api

type OrderCryptoReq struct {
	Account *Account `json:"account"`
	Order   Order    `json:"order"`
}

type AccountCrypto struct {
	AccountID     int64  `json:"accountId"`
	AccountNumber string `json:"accountNumber"`
}

type OrderCrypto struct {
	OrderType     string `json:"orderType"`
	Quote         Quote  `json:"quote"`
	ExpectedValue string `json:"expectedValue"`
	LimitPx       string `json:"limitPx"`
	Transaction   string `json:"transaction"`
}
