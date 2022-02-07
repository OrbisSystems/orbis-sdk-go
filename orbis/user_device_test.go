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

func TestClient_DeRegisterUserDevice(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		token string
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
			got, err := c.DeRegisterUserDevice(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeRegisterUserDevice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeRegisterUserDevice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_FirebaseRegister(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.FirebaseRegisterRequest
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
			got, err := c.FirebaseRegister(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirebaseRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirebaseRegister() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_RegisterUserDevice(t *testing.T) {
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
			got, err := c.RegisterUserDevice()
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterUserDevice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterUserDevice() got = %v, want %v", got, tt.want)
			}
		})
	}
}
