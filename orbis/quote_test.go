package orbis

import (
	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/http"
	orbis_api "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
	"reflect"
	"testing"
	"time"
)

func TestClient_GetBonds(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		cusip []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.GetBonds(tt.args.cusip)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBonds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBonds() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetEquity(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		symbol string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.EquityElement
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.GetEquity(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEquity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEquity() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetEquityTicks(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		symbol    *string
		startTime *time.Time
		endTime   *time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.GetEquityTicks(tt.args.symbol, tt.args.startTime, tt.args.endTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEquityTicks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEquityTicks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetHistorical(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.HistoricalRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.Historical
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.GetHistorical(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHistorical() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHistorical() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetIntradayChartData(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		symbol     string
		rsiSpa     *float64
		from       *int64
		to         *int64
		rowVolumes *bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.IntradayChartData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.GetIntradayChartData(&tt.args.symbol, tt.args.rsiSpa, tt.args.from, tt.args.to, tt.args.rowVolumes)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIntradayChartData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIntradayChartData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetOptionChain(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		symbol string
		date   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.GetOptionChain(tt.args.symbol, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOptionChain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOptionChain() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetOptionChainDates(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		symbol  string
		symbols []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.GetOptionChainDates(tt.args.symbol, tt.args.symbols)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOptionChainDates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOptionChainDates() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetOptionChainGreek(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		symbol  string
		symbols []string
		date    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.GetOptionChainGreek(tt.args.symbol, tt.args.date, tt.args.symbols)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOptionChainGreek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOptionChainGreek() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetOptions(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		symbol string
		greeks bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.GetOptions([]string{tt.args.symbol}, tt.args.greeks)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOptions() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetOptionsGreeks(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		symbol string
		greeks bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.GetOptionsGreeks([]string{tt.args.symbol}, tt.args.greeks)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOptionsGreeks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOptionsGreeks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetOptionsPricingType(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		underlying *string
		symbol     *string
		expectedPx *float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.GetOptionsPricingType(tt.args.underlying, tt.args.symbol, tt.args.expectedPx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOptionsPricingType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOptionsPricingType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetShortability(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		symbol string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.GetShortability
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.GetShortability(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetShortability() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetShortability() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetTopStocksQuote(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		category *string
		market   *string
		order    *string
		count    int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.TopStocks
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.GetTopStocksQuote(tt.args.category, tt.args.market, tt.args.order, &tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTopStocksQuote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTopStocksQuote() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SearchQuote(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		criteria   string
		limit      int
		loadQuotes bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.TopStockQuote
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.SearchQuote(tt.args.criteria, tt.args.limit, tt.args.loadQuotes)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchQuote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchQuote() got = %v, want %v", got, tt.want)
			}
		})
	}
}
