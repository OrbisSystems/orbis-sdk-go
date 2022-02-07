package orbis_api

import "time"

type HistoricalRequest struct {
	Symbol    string     `json:"symbol"`
	Start     *time.Time `json:"start"`
	End       *time.Time `json:"end"`
	Range     *string    `json:"range"`
	SetWeekly bool       `json:"setWeekly"`
	RsiSpan   float64    `json:"rsiSpan"`
}

type GetShortability struct {
	Shortable      bool    `json:"shortable"`
	Located        bool    `json:"located"`
	LocateRequired bool    `json:"locateRequired"`
	Rate           float64 `json:"rate"`
	Shares         int64   `json:"shares"`
}
