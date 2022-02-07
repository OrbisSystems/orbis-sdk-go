package orbis

import (
	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/http"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"reflect"
	"testing"
	"time"
)

func TestClient_GetPaymentTypes(t *testing.T) {
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
			got, err := c.GetPaymentTypes()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPaymentTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPaymentTypes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetTypes(t *testing.T) {
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
			got, err := c.GetTypes()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTypes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SearchCorporateActions(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		symbol     *string
		dateFrom   *time.Time
		dateTo     *time.Time
		types      []string
		usePayDate *bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.CorporatesActionsSearch
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
			got, err := c.SearchCorporateActions(tt.args.symbol, tt.args.dateFrom, tt.args.dateTo, tt.args.types, tt.args.usePayDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchCorporateActions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchCorporateActions() got = %v, want %v", got, tt.want)
			}
		})
	}
}
