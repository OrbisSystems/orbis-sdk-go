package orbis_api

type JournalRequest struct {
	DeliverFrom  UserAccount `json:"deliverFrom"`
	DeliverTo    UserAccount `json:"deliverTo"`
	RemoteAddr   string      `json:"remoteAddr"`
	RemoteDevice string      `json:"remoteDevice"`
	Amount       int64       `json:"amount"`
	DeliveryDate string      `json:"deliveryDate"`
	Reason       string      `json:"reason"`
	Purpose      string      `json:"purpose"`
	Symbol       string      `json:"symbol"`
	Loa          Loa         `json:"loa"`
}

type Loa struct {
	RemoteAddr   string `json:"remoteAddr"`
	RemoteDevice string `json:"remoteDevice"`
	FullName     string `json:"fullName"`
	Timestamp    string `json:"timestamp"`
	Signature    string `json:"signature"`
}

type JournalSearchRequest struct {
	TransferID string      `json:"transferId"`
	Status     string      `json:"status"`
	Account    UserAccount `json:"account"`
}
