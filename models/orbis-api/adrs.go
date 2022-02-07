package orbis_api

type AddrSearch []AddrSearchElement

type AddrSearchElement struct {
	Symbol             string `json:"symbol"`
	Country            string `json:"country"`
	AdrShare           int64  `json:"adrShare"`
	AdrHomeMarketShare int64  `json:"adrHomeMarketShare"`
	HasRecentUpgrade   bool   `json:"hasRecentUpgrade"`
	HasRecentDowngrade bool   `json:"hasRecentDowngrade"`
	Option             bool   `json:"option"`
}
