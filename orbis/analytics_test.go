package orbis

import (
	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/http"
	"io"
	"reflect"
	"testing"
)

func TestClient_GetModelDriftReport(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		modelID   string
		threshold string
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
			got, err := c.GetModelDriftReport(tt.args.modelID, tt.args.threshold)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetModelDriftReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetModelDriftReport() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetModelPerformance(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		modelID          string
		performanceRange int64
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
			got, err := c.GetModelPerformance(tt.args.modelID, tt.args.performanceRange)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetModelPerformance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetModelPerformance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetOverexposureAnalysis(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		overexposureType *string
		scope            *string
		ID               *string
		threshold        *string
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
			got, err := c.GetOverexposureAnalysis(*tt.args.overexposureType, *tt.args.scope, *tt.args.ID, *tt.args.threshold)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOverexposureAnalysis() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOverexposureAnalysis() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetRealtimeBalanceBreakDown(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		modelID    string
		bucketSize int64
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
			got, err := c.GetRealtimeBalanceBreakDown(tt.args.modelID, tt.args.bucketSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRealtimeBalanceBreakDown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRealtimeBalanceBreakDown() got = %v, want %v", got, tt.want)
			}
		})
	}
}
