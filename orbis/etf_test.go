package orbis

import (
	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/http"
	orbis_api "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"reflect"
	"testing"
)

func TestClient_GetBreakDownBySector(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		duration    *string
		displaySize *int64
		exposure    *int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.BreakdownBySector
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
			got, err := c.GetBreakDownBySector(tt.args.duration, tt.args.displaySize, tt.args.exposure)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBreakDownBySector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBreakDownBySector() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetDetails(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		symbol     *string
		loadQuotes *bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.Details
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
			got, err := c.GetDetails(tt.args.symbol, tt.args.loadQuotes)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDetails() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetInitScreener(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	tests := []struct {
		name    string
		fields  fields
		want    orbis_api.InitScreener
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
			got, err := c.GetInitScreener()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInitScreener() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInitScreener() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetScreener(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.ScreenerRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.ScrenerResponse
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
			got, err := c.GetScreener(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetScreener() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetScreener() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetTopTenETF(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		inverseType  *string
		leverageType *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.Top10
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
			got, err := c.GetTopTenETF(tt.args.inverseType, tt.args.leverageType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTopTen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTopTen() got = %v, want %v", got, tt.want)
			}
		})
	}
}
