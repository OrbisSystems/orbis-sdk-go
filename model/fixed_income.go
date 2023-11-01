package model

import "time"

type FixedIncome struct {
	Isin                    *string    `json:"isin"`
	Cusip                   *string    `json:"cusip"`
	Issuer                  *string    `json:"issuer"`
	Sector                  *string    `json:"sector"`
	Description             *string    `json:"description"`
	DescriptionShort        *string    `json:"description_short"`
	Type                    *string    `json:"type"`
	IssueSize               *int64     `json:"issue_size"`
	IssuePrice              *float32   `json:"issue_price"`
	IssueDate               *time.Time `json:"issue_date"`
	MaturityDate            *time.Time `json:"maturity_date"`
	SpRating                *string    `json:"sp_rating"`
	Duration                *float32   `json:"duration"`
	Coupon                  *float32   `json:"coupon"`
	CouponFrequency         *string    `json:"coupon_frequency"`
	CouponType              *string    `json:"coupon_type"`
	NextCouponDate          *time.Time `json:"next_coupon_date"`
	Callable                *bool      `json:"callable"`
	CallType                *string    `json:"call_type"`
	NextCallDate            *time.Time `json:"next_call_date"`
	NextCallPrice           *float32   `json:"next_call_price"`
	Convertible             *bool      `json:"convertible"`
	ParValue                *int64     `json:"par_value"`
	Insured                 *bool      `json:"insured"`
	AccruedInterest         *float32   `json:"accrued_interest"`
	Status                  *string    `json:"status"`
	MunicipalState          *string    `json:"municipal_state"`
	MunicipalType           *string    `json:"municipal_type"`
	MunicipalTaxableFederal *bool      `json:"municipal_taxable_federal"`
	Purpose                 *string    `json:"purpose"`
	BankQualified           *bool      `json:"bank_qualified"`
	TreasurySubtype         *string    `json:"treasury_subtype"`
	Ticker                  *string    `json:"ticker"`
	UpdatedAt               *time.Time `json:"updated_at"`
}

type GetFixedIncomeEntriesRequest struct {
	Statuses                []string        `json:"statuses"`
	Sectors                 []string        `json:"sectors"`
	Types                   []string        `json:"types"`
	SPRatings               []string        `json:"sp_ratings"`
	Coupon                  *LessAndMore    `json:"coupon"`
	CouponFrequencies       []string        `json:"coupon_frequencies"`
	IssueDate               *BeforeAndAfter `json:"issue_date"`
	MaturityDate            *BeforeAndAfter `json:"maturity"`
	Callable                *bool           `json:"callable"`
	Insured                 *bool           `json:"insured"`
	Convertible             *bool           `json:"convertible"`
	MunicipalStates         []string        `json:"municipal_states"`
	MunicipalTypes          []string        `json:"municipal_types"`
	MunicipalTaxableFederal *bool           `json:"municipal_tax"`
	Purposes                []string        `json:"purposes"`
	Tickers                 []string        `json:"tickers"`
	BankQualified           *bool           `json:"bank_qualified"`

	Paging *Paging `json:"paging"`
}

type GetFixedIncomeEntriesResponse struct {
	Data  []FixedIncome `json:"data"`
	Count int           `json:"count"`
}

type GetFixedIncomeHistoricalRequest struct {
	InstrumentID string `json:"instrument_id"`
	Frequency    string `json:"frequency"`
	Start        string `json:"start"`
	End          string `json:"end"`
}

type GetFixedIncomeHistoricalResponse struct {
	Success   bool                           `json:"success"`
	Timestamp time.Time                      `json:"timestamp"`
	Status    int                            `json:"status"`
	Message   string                         `json:"message"`
	Data      []GetFixedIncomeHistoricalData `json:"data"`
}

type GetFixedIncomeHistoricalData struct {
	Timestamp    time.Time `json:"timestamp"`
	Price        float32   `json:"price"`
	YieldToWorst float32   `json:"yield_to_worst"`
}
