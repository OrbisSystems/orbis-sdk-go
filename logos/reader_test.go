package logos

import (
	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/http"
	"github.com/OrbisSystems/orbis-sdk-go/mock"
	"github.com/OrbisSystems/orbis-sdk-go/storage"
	"io"
	http2 "net/http"
	"reflect"
	"testing"
)

type fields struct {
	API    auth.API
	client http.Client
	config config.OrbisConfig
}

func getDefaultField(resp mock.ExpectedResp) fields {
	return fields{
		API: auth.NewAuth(storage.NewInMemory(), http.DefaultAuthKey),
		client: mock.NewMock(func() (*http2.Response, error) {
			return &http2.Response{
				StatusCode: 200,
				Body:       mock.GetReaderCloser(resp),
			}, nil
		}, http.DefaultClient),
		config: config.OrbisConfig{},
	}
}

func TestClient_GetAllLogos(t *testing.T) {
	type args struct {
		page int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		{
			name:    "GetAllLogos",
			fields:  getDefaultField(mock.DefaultExpectedResp),
			args:    args{page: 10},
			want:    mock.GetReaderCloser(mock.DefaultExpectedResp),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.GetAllLogos(tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllLogos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllLogos() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetSymbolLogo(t *testing.T) {
	type args struct {
		symbol string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		{
			name:    "GetSymbolLogo",
			fields:  getDefaultField(mock.DefaultExpectedResp),
			args:    args{symbol: "symbol"},
			want:    mock.GetReaderCloser(mock.DefaultExpectedResp),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.GetSymbolLogo(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSymbolLogo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSymbolLogo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetSymbolLogoConverted(t *testing.T) {
	type args struct {
		symbol     string
		conversion string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		{
			name:    "GetSymbolLogoConverted",
			fields:  getDefaultField(mock.DefaultExpectedResp),
			args:    args{symbol: "symbol", conversion: "png"},
			want:    mock.GetReaderCloser(mock.DefaultExpectedResp),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.GetSymbolLogoConverted(tt.args.symbol, tt.args.conversion)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSymbolLogoConverted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSymbolLogoConverted() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetSymbolLogos(t *testing.T) {
	type args struct {
		symbol string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		{
			name:    "GetSymbolLogos",
			fields:  getDefaultField(mock.DefaultExpectedResp),
			args:    args{symbol: "symbol"},
			want:    mock.GetReaderCloser(mock.DefaultExpectedResp),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.GetSymbolLogos(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSymbolLogos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSymbolLogos() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetSymbolsLogos(t *testing.T) {
	type args struct {
		symbols []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		{
			name:    "GetSymbolsLogos",
			fields:  getDefaultField(mock.DefaultExpectedResp),
			args:    args{symbols: []string{"symbol"}},
			want:    mock.GetReaderCloser(mock.DefaultExpectedResp),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.GetSymbolsLogos(tt.args.symbols)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSymbolsLogos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSymbolsLogos() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetSymbolsSocials(t *testing.T) {
	type args struct {
		symbol string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		{
			name:    "GetSymbolsSocials",
			fields:  getDefaultField(mock.DefaultExpectedResp),
			args:    args{symbol: "symbol"},
			want:    mock.GetReaderCloser(mock.DefaultExpectedResp),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.GetSymbolsSocials(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSymbolsSocials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSymbolsSocials() got = %v, want %v", got, tt.want)
			}
		})
	}
}
