package orbis_api

type ImportInstantRequest struct {
	RemoteDevice         string `json:"remoteDevice"`
	RemoteAddr           string `json:"remoteAddr"`
	Note                 string `json:"note"`
	Account              string `json:"account"`
	BankAccountNumber    string `json:"bankAccountNumber"`
	BankRoutingNumber    string `json:"bankRoutingNumber"`
	BankAccountOwnerName string `json:"bankAccountOwnerName"`
	BankAccountType      string `json:"bankAccountType"`
	Nickname             string `json:"nickname"`
}
