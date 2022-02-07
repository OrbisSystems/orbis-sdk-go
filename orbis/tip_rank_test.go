package orbis

import (
	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/http"
	orbis_api "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"reflect"
	"testing"
)

func TestClient_GetAnalystPortfolio(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		expertUUID string
		count      int64
		loadQuotes *bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.GetAnalystResponse
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
			got, err := c.GetAnalystPortfolio(tt.args.expertUUID, tt.args.count, *tt.args.loadQuotes)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAnalystPortfolio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAnalystPortfolio() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetAnalystProfile(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		expertUUID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.AnalystProfileResponse
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
			got, err := c.GetAnalystProfile(tt.args.expertUUID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAnalystProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAnalystProfile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetBloggerRatings(t *testing.T) {
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
		want    orbis_api.BloggerRatingsResponse
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
			got, err := c.GetBloggerRatings(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBloggerRatings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBloggerRatings() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetConsensusByBlogger(t *testing.T) {
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
		want    orbis_api.ConsensusByBloggerResponse
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
			got, err := c.GetConsensusByBlogger(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConsensusByBlogger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConsensusByBlogger() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetConsensusBySector(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	tests := []struct {
		name    string
		fields  fields
		want    orbis_api.ConsensusBySector
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
			got, err := c.GetConsensusBySector()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConsensusBySector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConsensusBySector() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetHedgeFunds(t *testing.T) {
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
		want    orbis_api.HedgeFundsResponse
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
			got, err := c.GetHedgeFunds(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHedgeFunds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHedgeFunds() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetInsider(t *testing.T) {
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
		want    orbis_api.InsiderResponse
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
			got, err := c.GetInsider(tt.args.symbols)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInsider() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInsider() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetInsiderTransactions(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		symbols []string
		count   int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.InsiderTransactionsResponse
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
			got, err := c.GetInsiderTransactions(tt.args.symbols, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInsiderTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInsiderTransactions() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetLiveFeed(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		count int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.LiveFeedResponse
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
			got, err := c.GetLiveFeed(tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLiveFeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLiveFeed() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetRatingsBySymbol(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		symbol string
		count  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.RatingBySymbolResponse
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
			got, err := c.GetRatingsBySymbol(tt.args.symbol, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRatingsBySymbol() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRatingsBySymbol() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetSimilarStocks(t *testing.T) {
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
			got, err := c.GetSimilarStocks(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSimilarStocks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSimilarStocks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetStockSentiment(t *testing.T) {
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
		want    orbis_api.StockSentimentResponse
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
			got, err := c.GetStockSentiment(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStockSentiment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStockSentiment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetTopAnalyst(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		count int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.GetTopAnalystResponse
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
			got, err := c.GetTopAnalyst(tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTopAnalyst() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTopAnalyst() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetTopStocksTipRanks(t *testing.T) {
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
		want    orbis_api.GetTopStocksResponse
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
			got, err := c.GetTopStocksTipRanks(tt.args.symbols)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTopStocksTipRanks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTopStocksTipRanks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetTrending(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		count  int64
		period int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.TrendingResponse
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
			got, err := c.GetTrending(tt.args.count, tt.args.period)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTrending() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTrending() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Overview(t *testing.T) {
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
		want    orbis_api.OverviewResponse
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
			got, err := c.Overview(tt.args.symbols)
			if (err != nil) != tt.wantErr {
				t.Errorf("Overview() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Overview() got = %v, want %v", got, tt.want)
			}
		})
	}
}
