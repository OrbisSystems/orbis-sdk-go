package orbis

import (
	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/http"
	orbis_api "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"reflect"
	"testing"
)

func TestClient_DepositACH(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.ACHDepositRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.Deposit
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
			got, err := c.DepositACH(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DepositACH() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DepositACH() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_DepositIRA(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.DepositIRARequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.Deposit
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
			got, err := c.DepositIRA(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DepositIRA() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DepositIRA() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Withdraw(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.WithdrawRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.Withdraw
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
			got, err := c.Withdraw(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Withdraw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Withdraw() got = %v, want %v", got, tt.want)
			}
		})
	}
}
