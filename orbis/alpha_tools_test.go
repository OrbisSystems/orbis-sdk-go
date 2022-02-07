package orbis

import (
	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/http"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"reflect"
	"testing"
)

func TestClient_GetAlphaTracker(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		benchmark *string
		rangE     *string
		symbols   []string
		period    *int64
		setEMA    *bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.AlphaTracker
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
			got, err := c.GetAlphaTracker(tt.args.benchmark, tt.args.rangE, tt.args.symbols, tt.args.period, tt.args.setEMA)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAlphaTracker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAlphaTracker() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetSeasonality(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		benchmark *string
		symbols   []string
		years     int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.AlphaTracker
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
			got, err := c.GetSeasonality(tt.args.benchmark, tt.args.symbols, tt.args.years)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSeasonality() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSeasonality() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetWeaklySeasonality(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		benchmark *string
		symbols   []string
		years     int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.AlphaTracker
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
			got, err := c.GetWeaklySeasonality(tt.args.benchmark, tt.args.symbols, tt.args.years)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWeaklySeasonality() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWeaklySeasonality() got = %v, want %v", got, tt.want)
			}
		})
	}
}
