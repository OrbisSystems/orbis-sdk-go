package orbis

import (
	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/http"
	orbis_api "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"reflect"
	"testing"
)

func TestClient_CompletePlaid(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.PlaidRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.Complete
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
			got, err := c.CompletePlaid(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompletePlaid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompletePlaid() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetListBankAccountPlaid(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.PlaidRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.ListBankAccount
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
			got, err := c.GetListBankAccountPlaid(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListBankAccountPlaid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListBankAccountPlaid() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetListBankAccountsPlaid(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.PlaidRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.ListBankAccounts
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
			got, err := c.GetListBankAccountsPlaid(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListBankAccountsPlaid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListBankAccountsPlaid() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_InitiatePost(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		newValidation bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.InitiatePost
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
			got, err := c.InitiatePost(tt.args.newValidation)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitiatePost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitiatePost() got = %v, want %v", got, tt.want)
			}
		})
	}
}
