package avatar

import (
	"github.com/OrbisSystems/orbis-sdk-go/mock"

	"io"
	"reflect"
	"testing"
)

func TestClient_AvatarAuth(t *testing.T) {
	type args struct {
		env        string
		orbisToken string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		{
			name:   "testAvatarAuthPositive",
			fields: getDefaultField(mock.DefaultExpectedResp),
			args: args{
				env:        "testEnv",
				orbisToken: "testToken",
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
			got, err := c.AvatarAuth(tt.args.env, tt.args.orbisToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("AvatarAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AvatarAuth() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_AvatarRenewToken(t *testing.T) {
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
		{
			name:   "testAvatarRenewToken",
			fields: getDefaultField(mock.DefaultExpectedResp),
			args: args{
				token: "testToken",
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
			got, err := c.AvatarRenewToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("AvatarRenewToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AvatarRenewToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_AvatarTokenInvalidate(t *testing.T) {
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
		{
			name:   "testAvatarAuthPositive",
			fields: getDefaultField(mock.DefaultExpectedResp),
			args: args{
				token: "testToken",
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
			got, err := c.AvatarTokenInvalidate(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("AvatarTokenInvalidate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AvatarTokenInvalidate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
