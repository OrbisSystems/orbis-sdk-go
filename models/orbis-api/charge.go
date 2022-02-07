package orbis_api

import "time"

type SearchChargesRequest struct {
	UserAccount
	DateFrom  *time.Time
	DateTo    *time.Time
	ChargeRef *string
}

type ApplyChargeRequest struct {
	AccountID int64   `json:"accountId"`
	Date      string  `json:"date"`
	Amount    float64 `json:"amount"`
}
