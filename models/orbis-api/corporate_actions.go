package orbis_api

type CorporatesActionsSearch []CorporatesActionsSearchElement

type CorporatesActionsSearchElement struct {
	Type           string                `json:"type"`
	OldEntry       CorporateActionsEntry `json:"oldEntry"`
	NewEntry       CorporateActionsEntry `json:"newEntry"`
	EFFDate        string                `json:"effDate"`
	ExDate         string                `json:"exDate"`
	RecordDate     string                `json:"recordDate"`
	DeclareDate    string                `json:"declareDate"`
	PayDate        string                `json:"payDate"`
	DividendAmount float64               `json:"dividendAmount"`
	PaymentTypes   []interface{}         `json:"paymentTypes,omitempty"`
}

type CorporateActionsEntry struct {
	Symbol      string `json:"symbol"`
	Exchange    string `json:"exchange"`
	CompanyName string `json:"companyName"`
	Cusip       string `json:"cusip"`
}
