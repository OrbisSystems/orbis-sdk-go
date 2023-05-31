package model

type Continent struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Region struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CountryCode struct {
	A2Code        string `json:"a2_code"`
	A3Code        string `json:"a3_code"`
	CountryName   string `json:"country_name"`
	NumCode       string `json:"num_code"`
	MorningstarID string `json:"morningstar_id"`
}

type GlobalIndexFull struct {
	Symbol        string  `json:"symbol"`
	CountryID     string  `json:"country_id"`
	RegionID      string  `json:"region_id"`
	ContinentID   string  `json:"continent_id"`
	Ppp           float32 `json:"ppp"`
	Gdp           float32 `json:"gdp"`
	CountryName   string  `json:"country_name"`
	RegionName    string  `json:"region_name"`
	ContinentName string  `json:"continent_name"`

	CompanyName        string  `json:"company_name"`
	LastPrice          float64 `json:"last_price"`
	PriceChange        float64 `json:"price_change"`
	PriceChangePercent float64 `json:"price_change_percent"`
}
