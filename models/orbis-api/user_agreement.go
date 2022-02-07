package orbis_api

type AttestStatusRequest struct {
	Professional bool `json:"professional"`
}

type SignAgreementRequest struct {
	Code    string           `json:"code"`
	Entries []EntryAgreement `json:"entries"`
}

type EntryAgreement struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type SignCryptoAgreementRequest struct {
	Signature string  `json:"signature"`
	Account   Account `json:"account"`
}
