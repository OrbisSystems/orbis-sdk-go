package orbis_api

import "time"

type AccountHistory []AccountHistoryRequestElement

type AccountPortfolioRequest struct {
	LoadQuotes     *bool
	AccountID      *int64
	AccountNumber  *string
	LoadEarnings   *bool
	LoadUpgrades   *bool
	LoadIndustries *bool
}
type FullAccountPortfolioRequest struct {
	AccountPortfolioRequest
	MarketCapBreakdown *bool
}

type AccountHistoryRequest struct {
	UserAccount
	Page, Count      int64
	DateFrom, DateTo *time.Time
}

type AccountHistoryRequestElement struct {
	Symbol                   *string  `json:"symbol,omitempty"`
	ExchangeDate             *string  `json:"exchangeDate,omitempty"`
	Quantity                 *int64   `json:"quantity,omitempty"`
	CashUpdateEntry          bool     `json:"cashUpdateEntry"`
	StockCredit              bool     `json:"stockCredit"`
	Change                   bool     `json:"change"`
	DividendPayment          bool     `json:"dividendPayment"`
	Description              *string  `json:"description,omitempty"`
	CashActivity             bool     `json:"cashActivity"`
	Type                     string   `json:"type"`
	StockActivity            bool     `json:"stockActivity"`
	DivRate                  *float64 `json:"divRate,omitempty"`
	CouponRate               *int64   `json:"couponRate,omitempty"`
	Split                    bool     `json:"split"`
	TradeEntry               bool     `json:"tradeEntry"`
	Currency                 string   `json:"currency"`
	DivInt                   *int64   `json:"divInt,omitempty"`
	WithholdAmt              *int64   `json:"withholdAmt,omitempty"`
	PayDate                  *string  `json:"payDate,omitempty"`
	Status                   *string  `json:"status,omitempty"`
	Option                   bool     `json:"option"`
	Addition                 bool     `json:"addition"`
	IntEffectiveDate         *string  `json:"intEffectiveDate,omitempty"`
	UserEntryDate            *string  `json:"userEntryDate,omitempty"`
	ActivityIndicator        *string  `json:"activityIndicator,omitempty"`
	CusipOrSymbol            *string  `json:"cusipOrSymbol,omitempty"`
	OriginalQty              *int64   `json:"originalQty,omitempty"`
	ExecPrice                *float64 `json:"execPrice,omitempty"`
	Amount                   *float64 `json:"amount,omitempty"`
	CashType                 *string  `json:"cashType,omitempty"`
	AcatsNo                  *string  `json:"acatsNo,omitempty"`
	EntryTypeCode            *string  `json:"entryTypeCode,omitempty"`
	CashSubType              *string  `json:"cashSubType,omitempty"`
	WithholdTaxIndicator     *string  `json:"withholdTaxIndicator,omitempty"`
	EffectiveDate            *string  `json:"effectiveDate,omitempty"`
	WithholdTaxTypeCode      *string  `json:"withholdTaxTypeCode,omitempty"`
	CertificateShortDesc     *string  `json:"certificateShortDesc,omitempty"`
	SettleDate               *string  `json:"settleDate,omitempty"`
	SecuritySubType          *string  `json:"securitySubType,omitempty"`
	Details                  *string  `json:"details,omitempty"`
	StockReceivedOrDelivered *string  `json:"stockReceivedOrDelivered,omitempty"`
	SequenceEntryDate        *string  `json:"sequenceEntryDate,omitempty"`
	TradeDate                *string  `json:"tradeDate,omitempty"`
	ReinvestmentAmt          *int64   `json:"reinvestmentAmt,omitempty"`
	SecurityType             *string  `json:"securityType,omitempty"`
	ProcessDate              *string  `json:"processDate,omitempty"`
}

type RenameUserWatchlistRequest struct {
	ListID string `json:"ListId"`
	Title  string `json:"Title"`
}
type DeleteUserWatchlistRequest struct {
	ListID string `json:"ListId"`
}
type EntryUserWatchlistRequest struct {
	ListID string `json:"ListId"`
	Symbol string `json:"Symbol"`
}
