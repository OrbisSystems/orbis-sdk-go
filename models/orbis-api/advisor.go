package orbis_api

type AdvisorRequest struct {
	Inventory bool `json:"inventory"`
	Order     struct {
		Quote struct {
			Symbol string `json:"symbol"`
		} `json:"quote"`
		Ordertype   string `json:"orderType"`
		Transaction string `json:"transaction"`
		Filltype    string `json:"fillType"`
		Validity    string `json:"validity"`
		Limitprice  int    `json:"limitPrice"`
		Stopprice   int    `json:"stopPrice"`
		Quantity    int    `json:"quantity"`
		Markettime  string `json:"marketTime"`
	} `json:"order"`
}
