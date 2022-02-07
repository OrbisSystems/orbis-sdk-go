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

func TestClient_Allocations(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.AllocationModelRequest
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
			got, err := c.Allocations(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Allocations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Allocations() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CreateModel(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.CreateModelRequest
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
			got, err := c.CreateModel(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateModel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_DeleteModel(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		ID int64
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
			got, err := c.DeleteModel(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteModel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetAdjustmentPreview(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		adjustmentID string
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
			got, err := c.GetAdjustmentPreview(tt.args.adjustmentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAdjustmentPreview() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAdjustmentPreview() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetAdjustments(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		modelID    *string
		status     *string
		loadQuotes bool
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
			got, err := c.GetAdjustments(*tt.args.modelID, *tt.args.status, tt.args.loadQuotes)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAdjustments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAdjustments() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetAllModelBuyingPower(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		fCurrency      *string
		sCurrencyQuery *string
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
			got, err := c.GetAllModelBuyingPower(*tt.args.fCurrency, *tt.args.sCurrencyQuery)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllModelBuyingPower() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllModelBuyingPower() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetBuyingPower(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		currency *string
		modelID  *string
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
			got, err := c.GetBuyingPower(*tt.args.currency, *tt.args.modelID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBuyingPower() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBuyingPower() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetComponentTags(t *testing.T) {
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
			got, err := c.GetComponentTags()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetComponentTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetComponentTags() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetMemberAccountStatistic(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		modelID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &Client{
				API:    tt.fields.API,
				client: tt.fields.client,
				config: tt.fields.config,
			}
		})
	}
}

func TestClient_GetMemberAccounts(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		modelID        string
		loadRtb        bool
		loadRtbHistory bool
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
			got, err := c.GetMemberAccounts(tt.args.modelID, tt.args.loadRtb, tt.args.loadRtbHistory)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMemberAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMemberAccounts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetMemberRealtimeBalances(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		modelID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.MemberRealtimeBalances
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
			got, err := c.GetMemberRealtimeBalances(tt.args.modelID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMemberRealtimeBalances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMemberRealtimeBalances() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetModelListing(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		loadQuotes   bool
		loadAccounts bool
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
			got, err := c.GetModelListing(tt.args.loadQuotes, tt.args.loadAccounts)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetModelListing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetModelListing() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetPortfolio(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		modelID string
		req     orbis_api.ModelPortfolioRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.Portfolio
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
			got, err := c.GetPortfolio(tt.args.modelID, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPortfolio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPortfolio() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetPositionLookup(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		modelID *string
		symbol  *string
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
			got, err := c.GetPositionLookup(*tt.args.modelID, *tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPositionLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPositionLookup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetPositionLookupOptions(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		modelID *string
		symbol  *string
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
			got, err := c.GetPositionLookupOptions(*tt.args.modelID, *tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPositionLookupOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPositionLookupOptions() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetRealtimeBalance(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		modelID string
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
			got, err := c.GetRealtimeBalance(tt.args.modelID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRealtimeBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRealtimeBalance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetRealtimeBalanceHistory(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		modelID  string
		dateFrom *time.Time
		dateTo   *time.Time
		page     int64
		count    int64
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
			got, err := c.GetRealtimeBalanceHistory(tt.args.modelID, tt.args.dateFrom, tt.args.dateTo, tt.args.page, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRealtimeBalanceHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRealtimeBalanceHistory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_LinkOfMembersAccounts(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.LinkOfMembersRequest
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
			got, err := c.LinkOfMembersAccounts(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkOfMembersAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkOfMembersAccounts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_OrphanedAccounts(t *testing.T) {
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
			got, err := c.OrphanedAccounts()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrphanedAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrphanedAccounts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UnlinkMemberAccount(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.LinkOfMembersRequest
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
			got, err := c.UnlinkMemberAccount(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnlinkMemberAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnlinkMemberAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UpdateAdjustment(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		action string
		req    orbis_api.UpdateAdjustmentRequest
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
			got, err := c.UpdateAdjustment(tt.args.action, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateAdjustment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateAdjustment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UpdateComponent(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.UpdateComponentRequest
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
			got, err := c.UpdateComponent(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateComponent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateComponent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UpdateModel(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.UpdateModelRequest
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
			got, err := c.UpdateModel(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateModel() got = %v, want %v", got, tt.want)
			}
		})
	}
}
