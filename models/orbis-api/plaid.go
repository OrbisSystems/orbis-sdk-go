package orbis_api

type PlaidRequest struct {
	Account       string  `json:"account"`
	PublicToken   string  `json:"publicToken"`
	BankAccountID *string `json:"bankAccountId,omitempty"`
	NewValidation string  `json:"newValidation"`
}
