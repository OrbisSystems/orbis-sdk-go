package orbis_api

type BankCodesViaAddressRequest struct {
	Country       string   `json:"country"`
	City          string   `json:"city"`
	StreetAddress []string `json:"streetAddress"`
	PostalCode    string   `json:"postalCode"`
	StateProvince string   `json:"stateProvince"`
}

type CreateTransferRequest struct {
	Account              string `json:"account"`
	BankRoutingNumber    string `json:"bankRoutingNumber"`
	BankAccount          string `json:"bankAccount"`
	BankAccountOwnerName string `json:"bankAccountOwnerName"`
	BankAccountType      string `json:"bankAccountType"`
	Nickname             string `json:"nickname"`
}

type RelationshipInformationRequest struct {
	Nickname      string        `json:"nickname"`
	Note          string        `json:"note"`
	International bool          `json:"international"`
	UserAccount   UserAccount   `json:"userAccount"`
	Beneficiary   Beneficiary   `json:"beneficiary"`
	RecipientBank RecipientBank `json:"recipientBank"`
}

type ApproveTransferRequest struct {
	Account string `json:"account"`
	ID      string `json:"id"`
	Amount1 string `json:"amount1"`
	Amount2 string `json:"amount2"`
}

type CancelRelationshipTransferRequest struct {
	Account string `json:"account"`
	Reason  string `json:"reason"`
}
