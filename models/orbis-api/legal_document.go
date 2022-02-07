package orbis_api

type UpdatePolicyRequest struct {
	Account      AccountLegalDocument `json:"account"`
	Policy       string               `json:"policy"`
	RemoteDevice string               `json:"remoteDevice"`
	RemoteAddr   string               `json:"remoteAddr"`
	Esigned      bool                 `json:"esigned"`
}

type AccountLegalDocument struct {
	AccountNumber string `json:"accountNumber"`
	AccountID     int64  `json:"accountId"`
}
