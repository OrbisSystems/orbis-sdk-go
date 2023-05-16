package model

type GetMarketDatesRequest struct {
	Markets   []string `json:"markets"`
	StartDate Time     `json:"start_date"`
	EndDate   Time     `json:"end_date"`

	Paging *Paging `json:"paging,omitempty"`
}

type GetMarketDatesResponse struct {
	Data  []ExtendedMarketDate `json:"data"`
	Count int                  `json:"count"`
}

type GetMarketHoursResponse struct {
	OpenRightNow   bool                `json:"open_right_now"`
	OpenToday      bool                `json:"open_today"`
	OpenTime       string              `json:"open_time"`
	CloseTime      string              `json:"close_time"`
	TimeZone       string              `json:"time_zone"`
	NextMarketDate *ExtendedMarketDate `json:"next_market_day"`
}

type ExtendedMarketDate struct {
	MarketCode    string `json:"market_code"`
	MarketName    string `json:"market_name"`
	CountryCode   string `json:"country_code"`
	CurrencyCode  string `json:"currency_code"`
	PrimaryMarket bool   `json:"primary_market"`
	OpenDate      string `json:"open_date"`
	OpenTime      string `json:"open_time"`
	CloseTime     string `json:"close_time"`
	Timezone      string `json:"timezone"`
}
