package orbis

import (
	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/http"
	orbis_api "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"reflect"
	"testing"
)

func TestClient_GetDefaults(t *testing.T) {
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
			got, err := c.GetDefaults()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDefaults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDefaults() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetTopTenADRS(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		country                *string
		loadQuotes             *bool
		loadEarningRelease     *bool
		loadUpgradesDowngrades *bool
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
			got, err := c.GetTopTenADRS(tt.args.country, *tt.args.loadQuotes, *tt.args.loadEarningRelease, *tt.args.loadUpgradesDowngrades)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTopTenADRS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTopTenADRS() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SearchAdrS(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		country                *string
		loadQuotes             *bool
		loadEarningRelease     *bool
		loadUpgradesDowngrades *bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.AddrSearch
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
			got, err := c.SearchAdrS(tt.args.country, *tt.args.loadQuotes, *tt.args.loadEarningRelease, *tt.args.loadUpgradesDowngrades)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchAdrS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchAdrS() got = %v, want %v", got, tt.want)
			}
		})
	}
}
