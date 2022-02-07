package avatar

import (
	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/http"
	"github.com/OrbisSystems/orbis-sdk-go/mock"
	"github.com/OrbisSystems/orbis-sdk-go/storage"
	"io"
	http2 "net/http"
	"os"
	"reflect"
	"testing"
)

type fields struct {
	API    auth.API
	client http.Client
	config config.OrbisConfig
}

func getDefaultField(expectedRes mock.ExpectedResp) fields {
	return struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}{API: auth.NewAuth(storage.NewInMemory(), http.DefaultAuthKey),
		client: mock.NewMock(func() (*http2.Response, error) {
			return &http2.Response{
				StatusCode: 200,
				Body:       mock.GetReaderCloser(expectedRes),
			}, nil
		}, http.DefaultClient), config: config.OrbisConfig{ApiUrl: "test"}}
}

func TestClient_GetAccountAvatar(t *testing.T) {
	type args struct {
		accountID string
		token     string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		{name: "testPositive", fields: getDefaultField(mock.DefaultExpectedResp),
			args: args{
				accountID: "testID",
				token:     "testToken",
			},
			want:    mock.GetReaderCloser(mock.DefaultExpectedResp),
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
			got, err := c.GetAccountAvatar(tt.args.accountID, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountAvatar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountAvatar() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetAccountAvatarByOrbisToken(t *testing.T) {
	type args struct {
		accountID  string
		orbisToken string
		env        string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		{
			name:   "testGetAccountAvatarByOrbisTokenPositive",
			fields: getDefaultField(mock.DefaultExpectedResp),
			args: args{
				accountID:  "testAccountID",
				orbisToken: "testToken",
				env:        "testEnv",
			},
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
			got, err := c.GetAccountAvatarByOrbisToken(tt.args.accountID, tt.args.orbisToken, tt.args.env)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountAvatarByOrbisToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountAvatarByOrbisToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_RemoveAccountAvatar(t *testing.T) {
	type args struct {
		accountID string
		token     string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		{
			name:   "testRemoveAccountAvatarPositive",
			fields: getDefaultField(mock.DefaultExpectedResp),
			args: args{
				accountID: "testAccountID",
				token:     "testToken",
			},
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
			got, err := c.RemoveAccountAvatar(tt.args.accountID, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveAccountAvatar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAccountAvatar() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_RemoveAccountAvatarByOrbisToken(t *testing.T) {
	type args struct {
		accountID  string
		orbisToken string
		env        string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		{
			name:   "testRemoveAccountAvatarByOrbisTokenPositivi",
			fields: getDefaultField(mock.DefaultExpectedResp),
			args: args{
				accountID:  "testAccountID",
				orbisToken: "testToken",
				env:        "testEnv",
			},
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
			got, err := c.RemoveAccountAvatarByOrbisToken(tt.args.accountID, tt.args.orbisToken, tt.args.env)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveAccountAvatarByOrbisToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAccountAvatarByOrbisToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UploadAccountAvatar(t *testing.T) {
	type args struct {
		accountID string
		token     string
		avatar    os.File
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		{
			name:   "testUploadAccountAvatarPositive",
			fields: getDefaultField(mock.DefaultExpectedResp),
			args: args{
				accountID: "testAccountID",
				token:     "testToken",
			},
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
			got, err := c.UploadAccountAvatar(tt.args.accountID, tt.args.token, tt.args.avatar)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadAccountAvatar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UploadAccountAvatar() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UploadAccountAvatarByOrbisToken(t *testing.T) {
	type args struct {
		accountID  string
		orbisToken string
		env        string
		avatar     os.File
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		{
			name:   "testUploadAccountAvatarByOrbisTokenPositivi",
			fields: getDefaultField(mock.DefaultExpectedResp),
			args: args{
				accountID:  "testAccountID",
				orbisToken: "testToken",
				env:        "testEnv",
			},
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
			got, err := c.UploadAccountAvatarByOrbisToken(tt.args.accountID, tt.args.orbisToken, tt.args.env, tt.args.avatar)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadAccountAvatarByOrbisToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UploadAccountAvatarByOrbisToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
