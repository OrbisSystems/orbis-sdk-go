package orbis_api

type TradingLimitManagementRequest struct {
	Account     Account `json:"account"`
	Note        string  `json:"note"`
	Reason      string  `json:"reason"`
	ReferenceNo string  `json:"referenceNo"`
	Adjustment  int64   `json:"adjustment"`
	Currency    string  `json:"currency"`
}
