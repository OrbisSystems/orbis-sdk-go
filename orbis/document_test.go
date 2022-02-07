package orbis

import (
	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/http"
	orbis_api "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
	"reflect"
	"testing"
	"time"
)

func TestClient_ConfirmDocument(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		account string
		date    *time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.ConfirmDocumentResponse
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
			got, err := c.ConfirmDocument(tt.args.account, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfirmDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfirmDocument() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ConfirmListDocuments(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		account   string
		beginDate *time.Time
		endDate   *time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.ListOfConfirms
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
			got, err := c.ConfirmListDocuments(tt.args.account, tt.args.beginDate, tt.args.endDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfirmListDocuments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfirmListDocuments() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Download(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		account *string
		key     *string
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
			got, err := c.Download(*tt.args.account, *tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Download() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetAvailableDocumentTypes(t *testing.T) {
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
			got, err := c.GetAvailableDocumentTypes()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAvailableDocumentTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAvailableDocumentTypes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetAvailableDocumentTypesAndCodes(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	tests := []struct {
		name    string
		fields  fields
		want    orbis_api.AvailableDocumentTypesAndCodes
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
			got, err := c.GetAvailableDocumentTypesAndCodes()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAvailableDocumentTypesAndCodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAvailableDocumentTypesAndCodes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetServiceDocs(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.ServiceDocsRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.ServiceDocsResponse
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
			got, err := c.GetServiceDocs(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetServiceDocs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetServiceDocs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
