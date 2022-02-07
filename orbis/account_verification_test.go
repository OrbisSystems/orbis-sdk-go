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

func TestClient_CompleteAccount(t *testing.T) {
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
			got, err := c.CompleteAccount(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompleteAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompleteAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetListBankAccount(t *testing.T) {
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
			got, err := c.GetListBankAccount(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListBankAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListBankAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetListBankAccounts(t *testing.T) {
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
			got, err := c.GetListBankAccounts(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListBankAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListBankAccounts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Import(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.ImportInstantRequest
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
			got, err := c.Import(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Import() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Import() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_InitiateAccountVerification(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		subject string
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
			got, err := c.InitiateAccountVerification(tt.args.subject)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitiateAccountVerification() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitiateAccountVerification() got = %v, want %v", got, tt.want)
			}
		})
	}
}
