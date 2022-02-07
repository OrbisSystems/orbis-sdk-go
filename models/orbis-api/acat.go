package orbis_api

type InitiateACATScheduleRequest []InitiateACATRequestElement

type SearchRequest struct {
	Status       *string
	Account      *string
	AccountID    *int64
	RequestID    *string
	BatchID      *string
	LoadAccounts *bool
}

type CancelACATRequest struct {
	Account      string `json:"account"`
	AccountID    int64  `json:"accountId"`
	BatchID      string `json:"batchId"`
	RequestID    string `json:"requestId"`
	Comment      string `json:"comment"`
	Status       string `json:"status"`
	RemoteIP     string `json:"remoteIp"`
	RemoteDevice string `json:"remoteDevice"`
}

type InitiateACATRequest struct {
	Account           UserAccount `json:"account"`
	BatchID           string      `json:"batchId"`
	Comment           string      `json:"comment"`
	Status            string      `json:"status"`
	RemoteIP          string      `json:"remoteIp"`
	RemoteDevice      string      `json:"remoteDevice"`
	ContraNumber      string      `json:"contraNumber"`
	ContraAccount     string      `json:"contraAccount"`
	ContraTitle       string      `json:"contraTitle"`
	ContraAccountType string      `json:"contraAccountType"`
	PrimaryTaxID      string      `json:"primaryTaxId"`
	SecondaryTaxID    string      `json:"secondaryTaxId"`
	ParticipantNumber string      `json:"participantNumber"`
	Type              string      `json:"type"`
	Assets            []AssetACAT `json:"assets"`
}

type AssetACAT struct {
	Symbol     string    `json:"symbol"`
	Quote      QuoteACAT `json:"quote"`
	Quantity   int64     `json:"quantity"`
	OptionFlag int64     `json:"optionFlag"`
	IsShort    int64     `json:"isShort"`
}

type QuoteACAT struct {
	Symbol string `json:"symbol"`
	Cusip  string `json:"cusip"`
}

type InitiateACATRequestElement struct {
	Account           UserAccount `json:"account"`
	BatchID           string      `json:"batchId"`
	Comment           string      `json:"comment"`
	RemoteIP          string      `json:"remoteIp"`
	RemoteDevice      string      `json:"remoteDevice"`
	ContraNumber      string      `json:"contraNumber"`
	ContraAccount     string      `json:"contraAccount"`
	ContraTitle       string      `json:"contraTitle"`
	ContraAccountType string      `json:"contraAccountType"`
	PrimaryTaxID      string      `json:"primaryTaxId"`
	SecondaryTaxID    string      `json:"secondaryTaxId"`
	ParticipantNumber string      `json:"participantNumber"`
	Type              string      `json:"type"`
	Assets            []AssetACAT `json:"assets"`
}
