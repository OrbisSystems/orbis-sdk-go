package orbis

import (
	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/http"
	orbis_api "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"reflect"
	"testing"
)

func TestClient_GetPerformance(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		loadQuotes *bool
		time       *string
		sort       *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.PerformanceIPOS
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
			got, err := c.GetPerformance(tt.args.loadQuotes, tt.args.time, tt.args.sort)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPerformance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPerformance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetRecent(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		loadQuotes *bool
		time       *string
		sort       *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.GetIPOS
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
			got, err := c.GetRecent(tt.args.loadQuotes, tt.args.time, tt.args.sort)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRecent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRecent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetUpcoming(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	tests := []struct {
		name    string
		fields  fields
		want    orbis_api.GetIPOS
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
			got, err := c.GetUpcoming()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUpcoming() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUpcoming() got = %v, want %v", got, tt.want)
			}
		})
	}
}
