package model

type CalculateParamsRequest struct {
	Ticker      string  `json:"ticker"`
	StockPrice  float32 `json:"stock_price"`
	StrikeValue float32 `json:"strike_value"`
	ExpiryDate  Date    `json:"expiry_date"`
	OptionType  string  `json:"option_type"`
}

type Date struct {
	Year  int32 `json:"year"`
	Month int32 `json:"month"`
	Day   int32 `json:"day"`
}

type CalculateParamsResponse struct {
	TickerData TickerData `json:"ticker_data"`
	OptionData OptionData `json:"option_data"`
}

type TickerData struct {
	Ticker             string  `json:"ticker"`
	HistoricVolatility float32 `json:"historic_volatility"`
}

type OptionData struct {
	Ticker           string       `json:"ticker"`
	TheoreticalPrice float32      `json:"theoretical_price"`
	StrikePrice      float32      `json:"strike_price"`
	ExpiryDate       Date         `json:"expiry_date"`
	Values           OptionValues `json:"values"`
	Greeks           OptionGreeks `json:"greeks"`
}

type OptionValues struct {
	IntrinsicValue float32 `json:"intrinsic_value"`
	OptionValue    float32 `json:"option_value"`
	TimeValue      float32 `json:"time_value"`
}

type OptionGreeks struct {
	Delta float32 `json:"delta"`
	Gamma float32 `json:"gamma"`
	Rho   float32 `json:"rho"`
	Theta float32 `json:"theta"`
	Vega  float32 `json:"vega"`
}

type CalculateMatrixParamsRequest struct {
	Ticker        string  `json:"ticker"`
	StockPriceMin float32 `json:"stock_price_min"`
	StockPriceMax float32 `json:"stock_price_max"`
	StrikeValue   float32 `json:"strike_value"`
	ExpiryDate    Date    `json:"expiry_date"`
	OptionType    string  `json:"option_type"`
	MaxIntervals  int32   `json:"max_intervals"`
}

type CalculateMatrixResponse struct {
	TickerData TickerData   `json:"ticker_data"`
	MatrixData []MatrixData `json:"matrix_data"`
}

type MatrixData struct {
	Date       Date         `json:"date"`
	OptionData []OptionData `json:"option_data"`
}
