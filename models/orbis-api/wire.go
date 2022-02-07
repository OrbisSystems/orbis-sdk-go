package orbis_api

type WireRequest struct {
	Account          string        `json:"account"`
	DisbursementType string        `json:"disbursementType"`
	Amount           int64         `json:"amount"`
	RemoteAddr       string        `json:"remoteAddr"`
	RemoteDevice     string        `json:"remoteDevice"`
	International    bool          `json:"international"`
	EntryID          int64         `json:"entryId"`
	RecipientBank    RecipientBank `json:"recipientBank"`
	Beneficiary      Beneficiary   `json:"beneficiary"`
	IraDetails       IraDetails    `json:"iraDetails"`
}

type Beneficiary struct {
	Account    string      `json:"account"`
	ThirdParty bool        `json:"thirdParty"`
	Name       string      `json:"name"`
	Address    AddressWire `json:"address"`
}

type AddressWire struct {
	StreetAddress []string `json:"streetAddress"`
	City          string   `json:"city"`
	StateProvince string   `json:"stateProvince"`
	Country       string   `json:"country"`
	PostalCode    string   `json:"postalCode"`
}

type IraDetails struct {
	ContributionYear int64  `json:"contributionYear"`
	ContributionType string `json:"contributionType"`
}

type RecipientBank struct {
	Identifier            Identifier `json:"identifier"`
	Country               string     `json:"country"`
	City                  string     `json:"city"`
	PostalCode            string     `json:"postalCode"`
	Name                  string     `json:"name"`
	StateProvince         string     `json:"stateProvince"`
	AdditionalInformation string     `json:"additionalInformation"`
}

type Identifier struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
