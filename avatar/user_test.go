package avatar

import (
	"github.com/OrbisSystems/orbis-sdk-go/mock"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestClient_GetUserAvatar(t *testing.T) {
	type args struct {
		userID string
		token  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		{
			name:   "testGetUserAvatarPositive",
			fields: getDefaultField(mock.DefaultExpectedResp),
			args: args{
				userID: "testUserID",
				token:  "testToken",
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
			got, err := c.GetUserAvatar(tt.args.userID, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserAvatar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserAvatar() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetUserAvatarByOrbisToken(t *testing.T) {
	type args struct {
		userID     string
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
			name:   "testGetUserAvatarByOrbisTokenPositive",
			fields: getDefaultField(mock.DefaultExpectedResp),
			args: args{
				userID:     "testUserID",
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
			got, err := c.GetUserAvatarByOrbisToken(tt.args.userID, tt.args.orbisToken, tt.args.env)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserAvatarByOrbisToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserAvatarByOrbisToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_RemoveUserAvatar(t *testing.T) {
	type args struct {
		userID string
		token  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		{
			name:   "testRemoveUserAvatarPositive",
			fields: getDefaultField(mock.DefaultExpectedResp),
			args: args{
				userID: "testUserID",
				token:  "testToken",
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
			got, err := c.RemoveUserAvatar(tt.args.userID, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveUserAvatar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveUserAvatar() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_RemoveUserAvatarByOrbisToken(t *testing.T) {
	type args struct {
		userID     string
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
			name:   "testRemoveUserAvatarByOrbisToken",
			fields: getDefaultField(mock.DefaultExpectedResp),
			args: args{
				userID:     "testUserID",
				orbisToken: "testToken",
				env:        "env",
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
			got, err := c.RemoveUserAvatarByOrbisToken(tt.args.userID, tt.args.orbisToken, tt.args.env)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveUserAvatarByOrbisToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveUserAvatarByOrbisToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UploadUserAvatar(t *testing.T) {
	type args struct {
		userID string
		token  string
		avatar os.File
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		{
			name:   "testUploadUserAvatar",
			fields: getDefaultField(mock.DefaultExpectedResp),
			args: args{
				userID: "testUserID",
				token:  "testToken",
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
			got, err := c.UploadUserAvatar(tt.args.userID, tt.args.token, tt.args.avatar)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadUserAvatar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UploadUserAvatar() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UploadUserAvatarByOrbisToken(t *testing.T) {
	type args struct {
		userID     string
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
			name:   "testUploadUserAvatarByOrbisToken",
			fields: getDefaultField(mock.DefaultExpectedResp),
			args: args{
				userID:     "testUserID",
				orbisToken: "testToken",
				env:        "env",
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
			got, err := c.UploadUserAvatarByOrbisToken(tt.args.userID, tt.args.orbisToken, tt.args.env, tt.args.avatar)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadUserAvatarByOrbisToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UploadUserAvatarByOrbisToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
