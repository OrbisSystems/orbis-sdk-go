package orbis

import (
	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/http"
	orbis_api "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
	"reflect"
	"testing"
)

func TestClient_GetCompanyResearch(t *testing.T) {
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
		want    orbis_api.CompanyResearch
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
			got, err := c.GetCompanyResearch(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCompanyResearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCompanyResearch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetEarningCalendar(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		excludeQuotes *bool
		loadQuotes    *bool
		days          *int64
		symbols       []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.EarningsCalendar
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
			got, err := c.GetEarningCalendar(tt.args.excludeQuotes, tt.args.loadQuotes, tt.args.days, tt.args.symbols)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEarningCalendar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEarningCalendar() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetEarningCalendarForPortfolio(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		excludeQuotes *bool
		loadQuotes    *bool
		days          *int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.EarningsCalendarForPortfolio
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
			got, err := c.GetEarningCalendarForPortfolio(tt.args.excludeQuotes, tt.args.loadQuotes, tt.args.days)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEarningCalendarForPortfolio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEarningCalendarForPortfolio() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetFundamentalTypes(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
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
			got, err := c.GetFundamentalTypes()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFundamentalTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFundamentalTypes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetFundamentals(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		typeFund string
		symbol   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.Fundamentals
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
			got, err := c.GetFundamentals(tt.args.typeFund, tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFundamentals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFundamentals() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetHistoricalData(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		symbol string
		dates  []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.HistoricalData
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
			got, err := c.GetHistoricalData(tt.args.symbol, tt.args.dates)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHistoricalData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHistoricalData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetIndexComponents(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		index      string
		loadQuotes *bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.IndexComponents
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
			got, err := c.GetIndexComponents(tt.args.index, tt.args.loadQuotes)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIndexComponents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIndexComponents() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetIndexOverview(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	tests := []struct {
		name    string
		fields  fields
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
			got, err := c.GetIndexOverview()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIndexOverview() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIndexOverview() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetIndustryHeatMap(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		industryType string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.IndustryHeatMap
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
			got, err := c.GetIndustryHeatMap(tt.args.industryType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIndustryHeatMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIndustryHeatMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetIndustryList(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		industrySector string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.IndustryList
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
			got, err := c.GetIndustryList(tt.args.industrySector)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIndustryList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIndustryList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetOptionRadar(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		filter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.OptionRadar
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
			got, err := c.GetOptionRadar(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOptionRadar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOptionRadar() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetOptionRatios(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
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
			got, err := c.GetOptionRatios(tt.args.symbols)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOptionRatios() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOptionRatios() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetOwnerDetails(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		ids []string
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
			got, err := c.GetOwnerDetails(tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOwnerDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOwnerDetails() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetOwnershipDetails(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		symbols []string
		filter  string
		sort    string
		count   int64
		page    int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.OwnershipDetails
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
			got, err := c.GetOwnershipDetails(tt.args.symbols, tt.args.filter, tt.args.sort, tt.args.count, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOwnershipDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOwnershipDetails() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetStockUpgrades(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
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
			got, err := c.GetStockUpgrades(tt.args.symbols)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStockUpgrades() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStockUpgrades() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetSymbolHeatMap(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		amountLimit    *float64
		numSymbols     *int64
		industryCode   *int64
		exchangeFilter string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.SymbolHeatMap
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
			got, err := c.GetSymbolHeatMap(tt.args.amountLimit, tt.args.numSymbols, tt.args.industryCode, tt.args.exchangeFilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSymbolHeatMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSymbolHeatMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetTopQuotesByCategory(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		category string
		market   string
		order    string
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
			got, err := c.GetTopQuotesByCategory(tt.args.category, tt.args.market, tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTopQuotesByCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTopQuotesByCategory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Screener(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.ScreenerResearch
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.Screener
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
			got, err := c.Screener(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Screener() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Screener() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ScreenerV2(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.ScreenerResearch
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.Screener
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
			got, err := c.ScreenerV2(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScreenerV2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScreenerV2() got = %v, want %v", got, tt.want)
			}
		})
	}
}
