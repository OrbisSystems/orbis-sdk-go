package model

import (
	"strings"
	"time"
)

type QuoteEquityDataResponse struct {
	AskPx         float64 `json:"ask_px"`
	AskSz         int     `json:"ask_sz"`
	AvgVol30D     int     `json:"avg_vol30d"`
	BidPx         float64 `json:"bid_px"`
	BidSz         int     `json:"bid_sz"`
	ClosingPx     float64 `json:"closing_px"`
	Column        int     `json:"column"`
	ComShrsOut    int     `json:"com_shrs_out"`
	CompDesc      string  `json:"comp_desc"`
	CompName      string  `json:"comp_name"`
	Condition     int     `json:"condition"`
	DivAmt        float64 `json:"div_amt"`
	DivDateEx     string  `json:"div_date_ex"`
	DivYield      float64 `json:"div_yield"`
	High52WPx     float64 `json:"high52w_px"`
	HighPx        float64 `json:"high_px"`
	InfoTime      string  `json:"info_time"`
	IssueMarket   string  `json:"issue_market"`
	IssureMarket  string  `json:"issure_market"`
	LastPx        float64 `json:"last_px"`
	Low52WPx      float64 `json:"low52w_px"`
	LowPx         float64 `json:"low_px"`
	MktCap        int     `json:"mkt_cap"`
	OpenPx        float64 `json:"open_px"`
	PayDate       string  `json:"pay_date"`
	PeRatio       float64 `json:"pe_ratio"`
	PostAskPx     float64 `json:"post_ask_px"`
	PostAskSz     int     `json:"post_ask_sz"`
	PostBidPx     float64 `json:"post_bid_px"`
	PostBidSz     int     `json:"post_bid_sz"`
	PostLastPx    float64 `json:"post_last_px"`
	PostQuoteTime string  `json:"post_quote_time"`
	PostTradeTime string  `json:"post_trade_time"`
	PreAskPx      float64 `json:"pre_ask_px"`
	PreAskSz      int     `json:"pre_ask_sz"`
	PreBidPx      float64 `json:"pre_bid_px"`
	PreBidSz      int     `json:"pre_bid_sz"`
	PreLastPx     float64 `json:"pre_last_px"`
	PreQuoteTime  string  `json:"pre_quote_time"`
	PreTradeTime  string  `json:"pre_trade_time"`
	QuoteTime     string  `json:"quote_time"`
	Symbol        string  `json:"symbol"`
	TradePx       float64 `json:"trade_px"`
	TradeRegion   string  `json:"trade_region"`
	TradeSz       int     `json:"trade_sz"`
	TradeTime     string  `json:"trade_time"`
	Volume        int     `json:"volume"`
	YestClosePx   float64 `json:"yest_close_px"`
}

type QuoteHistoryRequest struct {
	Filter          QuoteHistoryRequestFilter `json:"filter"` // Use ONLY one of the next filters: TimeRange OR Period
	Order           *Ordering                 `json:"order,omitempty"`
	Paging          *Paging                   `json:"paging,omitempty"`
	QuoteAccessType QuoteAccessType           `json:"quote_access_type"` // Type of quote to load 		Enums(nbbo, delayed, cboe, bats, realtime, realtime_snapshot)	(REQUIRED)
}

type QuoteHistoryRequestFilter struct {
	Symbol    string           `json:"symbol"`
	TimeRange *TimeRangeFilter `json:"time_range,omitempty"`
	Period    *PeriodFilter    `json:"period,omitempty"`
}

type TimeRangeFilter struct {
	StartTimeUnixTimestamp int64 `json:"start_time"` // Use unix timestamp in seconds
	EndTimeUnixTimestamp   int64 `json:"end_time"`   // Use unix timestamp in seconds
}

type PeriodFilter struct {
	Type  PeriodType `json:"type" enums:"day,week,month,year"`
	Value int        `json:"value"`
}

type PeriodType string

const (
	PeriodDay   PeriodType = "day"
	PeriodWeek  PeriodType = "week"
	PeriodMonth PeriodType = "month"
	PeriodYear  PeriodType = "year"
)

type QuoteHistoryResponse struct {
	Extremum Extremum `json:"extremum"`
	Quotes   []*Quote `json:"quotes"`
	Count    int      `json:"count"`
}

type Extremum struct {
	MinValue     float64   `json:"min_value"`
	MinValueDate time.Time `json:"min_value_date"`
	MaxValue     float64   `json:"max_value"`
	MaxValueDate time.Time `json:"max_value_date"`
}

type Quote struct {
	Symbol             string    `json:"symbol"`
	Date               time.Time `json:"date"`
	Open               float64   `json:"open"`
	Close              float64   `json:"close"`
	High               float64   `json:"high"`
	Low                float64   `json:"low"`
	Volume             int64     `json:"volume"`
	PriceChange        float64   `json:"price_change"`
	PriceChangePercent float64   `json:"price_change_percent"`
}

type IntradayRequest struct {
	Symbol     string  `json:"symbol"`      // (REQUIRE)
	RsiSpan    float32 `json:"rsi_span"`    // (REQUIRE)
	From       *Time   `json:"from"`        // Format: YYYY-MM-DD	(OPTIONAL)
	To         *Time   `json:"to"`          // Format: YYYY-MM-DD	(OPTIONAL)
	RawVolumes bool    `json:"raw_volumes"` // (OPTIONAL)
	Smoothing  *int    `json:"smoothing"`   // (OPTIONAL)
}

type Time struct {
	time.Time
}

func (c *Time) UnmarshalJSON(b []byte) error {
	t, err := time.Parse("2006-01-02", strings.Trim(string(b), `"`))
	if err != nil {
		return err
	}
	*c = Time{t}
	return nil
}

type IntradayResponse struct {
	Date               time.Time `json:"date"`
	Price              float32   `json:"price"`
	Volume             int64     `json:"volume,omitempty"`
	High               float32   `json:"day_high"`
	Low                float32   `json:"day_low"`
	Rsi                float32   `json:"rsi"`
	PriceChange        float64   `json:"price_change"`
	PriceChangePercent float64   `json:"price_change_percent"`
	Sma                float32   `json:"sma"`
	Ema                float32   `json:"ema"`
}
