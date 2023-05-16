package model

type IPOResponse struct {
	IPOList      []IPO `json:"ipo_list"`
	TotalRecords int   `json:"total_records"`
}

type IPO struct {
	Symbol                string                 `json:"symbol"`
	CompanyName           string                 `json:"company_name"`
	ProspectusURL         string                 `json:"prospectus_url"`
	FileDate              string                 `json:"file_date"`
	IpoDate               string                 `json:"ipo_date"`
	IpoDateApprox         bool                   `json:"ipo_date_approx"`
	OfferPrice            float32                `json:"offer_price"`
	OfferPriceMax         float32                `json:"offer_price_max"`
	OfferPriceMin         float32                `json:"offer_price_min"`
	OfferShares           float32                `json:"offer_shares"`
	FirstDayOpen          float32                `json:"first_day_open"`
	FirstDayClose         float32                `json:"first_day_close"`
	IpoReturn             float32                `json:"ipo_return"`
	Status                IpoStatus              `json:"status"`
	Underwriters          []UnderwriterResponse  `json:"underwriters"`
	Ceo                   string                 `json:"ceo"`
	TotalExpenses         float32                `json:"total_expenses"`
	SharesOverAlloted     int64                  `json:"shares_over_alloted"`
	SharesOutstanding     int64                  `json:"shares_outstanding"`
	LockupPeriod          int                    `json:"lockup_period"`
	LockupExpiration      string                 `json:"lockup_expiration"`
	QuietPeriodExpiration string                 `json:"quiet_period_expiration"`
	OfferAmount           int64                  `json:"offer_amount"`
	Cik                   int                    `json:"cik"`
	Quote                 map[string]interface{} `json:"quote,omitempty"`
}

type UnderwriterResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

type IpoStatus string

const (
	PricedIpoStatus    IpoStatus = "Priced"
	FiledIpoStatus     IpoStatus = "Filed"
	WithdrawnIpoStatus IpoStatus = "Withdrawn"
	PostponedIpoStatus IpoStatus = "Postponed"
)

type RecentIPORequest struct {
	LoadQuotes bool      `json:"load_quotes"`
	Time       string    `json:"time"`
	Ordering   *Ordering `json:"order"`
	Paging     Paging    `json:"paging"`
}
