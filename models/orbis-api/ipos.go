package orbis_api

type GetIPOS []GetRecentIPO
type PerformanceIPOS []PerformanceIPO

type GetRecentIPO struct {
	Symbol                          string        `json:"symbol"`
	CompanyName                     string        `json:"companyName"`
	ProspectusURL                   string        `json:"prospectusUrl"`
	FileDate                        string        `json:"fileDate"`
	IPODate                         string        `json:"ipoDate"`
	IPODateApprox                   bool          `json:"ipoDateApprox"`
	OfferPrice                      float64       `json:"offerPrice"`
	OfferPriceMax                   int64         `json:"offerPriceMax"`
	OfferPriceMin                   int64         `json:"offerPriceMin"`
	OfferShares                     float64       `json:"offerShares"`
	FirstDayOpen                    float64       `json:"firstDayOpen"`
	FirstDayClose                   float64       `json:"firstDayClose"`
	IPOReturn                       float64       `json:"ipoReturn"`
	Status                          string        `json:"status"`
	Underwriters                    []Underwriter `json:"underwriters"`
	Ceo                             string        `json:"ceo"`
	TotalExpenses                   int64         `json:"totalExpenses"`
	SharesOverAlloted               int64         `json:"sharesOverAlloted"`
	SharesOutstanding               int64         `json:"sharesOutstanding"`
	LockupPeriod                    int64         `json:"lockupPeriod"`
	LockupExpiration                string        `json:"lockupExpiration"`
	QuietPeriodExpiration           string        `json:"quietPeriodExpiration"`
	OfferAmount                     int64         `json:"offerAmount"`
	IndustryValue                   int64         `json:"industryValue"`
	IndustryPercentage              int64         `json:"industryPercentage"`
	FirstDayReturn                  float64       `json:"firstDayReturn"`
	TotalReturn                     int64         `json:"totalReturn"`
	OfferSize                       float64       `json:"offerSize"`
	QuoteYesterdayCloseWithoutSplit int64         `json:"quoteYesterdayCloseWithoutSplit"`
	LeadUnderwriter                 Underwriter   `json:"leadUnderwriter"`
	FirstTradingDay                 bool          `json:"firstTradingDay"`
	Cik                             int64         `json:"cik"`
	Value                           int64         `json:"value"`
	Option                          bool          `json:"option"`
	GroupValue                      int64         `json:"groupValue"`
	SectorValue                     int64         `json:"sectorValue"`
	IndValue                        int64         `json:"indValue"`
	SubIndustryValue                int64         `json:"subIndustryValue"`
}

type PerformanceIPO struct {
	Symbol                          string  `json:"symbol"`
	CompanyName                     string  `json:"companyName"`
	ProspectusURL                   string  `json:"prospectusUrl"`
	FileDate                        string  `json:"fileDate"`
	IPODate                         string  `json:"ipoDate"`
	IPODateApprox                   bool    `json:"ipoDateApprox"`
	OfferPrice                      float64 `json:"offerPrice"`
	OfferPriceMax                   int64   `json:"offerPriceMax"`
	OfferPriceMin                   int64   `json:"offerPriceMin"`
	OfferShares                     float64 `json:"offerShares"`
	FirstDayOpen                    float64 `json:"firstDayOpen"`
	FirstDayClose                   float64 `json:"firstDayClose"`
	IPOReturn                       float64 `json:"ipoReturn"`
	Status                          string  `json:"status"`
	Ceo                             string  `json:"ceo"`
	StateOfInc                      *string `json:"stateOfInc,omitempty"`
	TotalExpenses                   float64 `json:"totalExpenses"`
	SharesOverAlloted               int64   `json:"sharesOverAlloted"`
	SharesOutstanding               int64   `json:"sharesOutstanding"`
	LockupPeriod                    int64   `json:"lockupPeriod"`
	LockupExpiration                string  `json:"lockupExpiration"`
	QuietPeriodExpiration           string  `json:"quietPeriodExpiration"`
	OfferAmount                     int64   `json:"offerAmount"`
	IndustryValue                   int64   `json:"industryValue"`
	IndustryPercentage              int64   `json:"industryPercentage"`
	FirstDayReturn                  float64 `json:"firstDayReturn"`
	TotalReturn                     int64   `json:"totalReturn"`
	OfferSize                       float64 `json:"offerSize"`
	QuoteYesterdayCloseWithoutSplit int64   `json:"quoteYesterdayCloseWithoutSplit"`
	FirstTradingDay                 bool    `json:"firstTradingDay"`
	Cik                             int64   `json:"cik"`
	Value                           int64   `json:"value"`
	Option                          bool    `json:"option"`
	GroupValue                      int64   `json:"groupValue"`
	SectorValue                     int64   `json:"sectorValue"`
	IndValue                        int64   `json:"indValue"`
	SubIndustryValue                int64   `json:"subIndustryValue"`
}
