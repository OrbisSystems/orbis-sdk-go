package orbis_api

import "time"

type TransferStatusRequest struct {
	Direction *string    `json:"direction"`
	Mechanism *string    `json:"mechanism"`
	DateFrom  *time.Time `json:"dateFrom"`
	DateTo    *time.Time `json:"dateTo"`
	Range     *string    `json:"range"`
	Start     *string    `json:"start"`
	PerPage   *int64     `json:"perPage"`
}

type CancelTransferRequest struct {
	Account      string `json:"account"`
	TransferID   string `json:"transferId"`
	RemoteAddr   string `json:"remoteAddr"`
	RemoteDevice string `json:"remoteDevice"`
}

type FundTransferRequest struct {
	UserAccount      UserAccount `json:"userAccount"`
	EntryID          int64       `json:"entryId"`
	Amount           int64       `json:"amount"`
	Currency         string      `json:"currency"`
	AwaitRedirectURL bool        `json:"awaitRedirectUrl"`
	AwaitTime        int64       `json:"awaitTime"`
	CallbackURL      string      `json:"callbackUrl"`
	RemoteAddr       string      `json:"remoteAddr"`
	RemoteDevice     string      `json:"remoteDevice"`
}

type UserAccount struct {
	AccountNumber string `json:"accountNumber"`
	AccountID     int64  `json:"accountId"`
}

type TransferRequest struct {
	DeliverFrom  UserAccount `json:"deliverFrom"`
	DeliverTo    UserAccount `json:"deliverTo"`
	Quote        Quote       `json:"quote"`
	RemoteAddr   string      `json:"remoteAddr"`
	RemoteDevice string      `json:"remoteDevice"`
	Quantity     int64       `json:"quantity"`
	Reason       string      `json:"reason"`
	Purpose      string      `json:"purpose"`
}
