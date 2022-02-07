package orbis_api

type ACHDepositRequest struct {
	Account           string `json:"account"`
	EntryID           int64  `json:"entryId"`
	Amount            int64  `json:"amount"`
	InstantAmount     int64  `json:"instantAmount"`
	RemoteAddr        string `json:"remoteAddr"`
	RemoteDevice      string `json:"remoteDevice"`
	ControlSource     string `json:"controlSource"`
	ControlTrxID      string `json:"controlTrxId"`
	ControlResolution string `json:"controlResolution"`
}

type DepositIRARequest struct {
	Account          string `json:"account"`
	EntryID          int64  `json:"entryId"`
	Amount           int64  `json:"amount"`
	ContributionType string `json:"contributionType"`
	ContributionYear string `json:"contributionYear"`
}

type WithdrawRequest struct {
	Account      string `json:"account"`
	EntryID      int64  `json:"entryId"`
	Amount       int64  `json:"amount"`
	RemoteAddr   string `json:"remoteAddr"`
	RemoteDevice string `json:"remoteDevice"`
}
