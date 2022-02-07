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

func TestClient_AddEntryUserWatchlist(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.EntryUserWatchlistRequest
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
			got, err := c.AddEntryUserWatchlist(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddEntryUserWatchlist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddEntryUserWatchlist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CreateUserWatchlist(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.CreateWatchlistRequest
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
			got, err := c.CreateUserWatchlist(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUserWatchlist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUserWatchlist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_DeleteUserWatchlist(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.DeleteUserWatchlistRequest
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
			got, err := c.DeleteUserWatchlist(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteUserWatchlist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUserWatchlist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_FirebaseUserRegister(t *testing.T) {
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
			got, err := c.FirebaseUserRegister(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirebaseUserRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirebaseUserRegister() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetAccountBalance(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		accountNumber string
		currency      string
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
			got, err := c.GetAccountBalance(tt.args.accountNumber, tt.args.currency)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountBalance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetAccountBalanceHistory(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.AccountHistoryRequest
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
			got, err := c.GetAccountBalanceHistory(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountBalanceHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountBalanceHistory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetAccountHistory(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.AccountHistoryRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.AccountHistory
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
			got, err := c.GetAccountHistory(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountHistory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetAccountInformation(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		accountNumber string
		accountID     int64
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
			got, err := c.GetAccountInformation(&tt.args.accountNumber, &tt.args.accountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountInformation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountInformation() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetAccountPortfolio(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.AccountPortfolioRequest
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
			got, err := c.GetAccountPortfolio(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountPortfolio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountPortfolio() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetAccountPortfolioHistory(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.AccountHistoryRequest
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
			got, err := c.GetAccountPortfolioHistory(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountPortfolioHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountPortfolioHistory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetAccountTradingLimit(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		accountNumber string
		currency      string
		accountID     int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.AccountTradingLimit
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
			got, err := c.GetAccountTradingLimit(tt.args.accountNumber, tt.args.currency, tt.args.accountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountTradingLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountTradingLimit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetFullAccountPortfolio(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.FullAccountPortfolioRequest
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
			got, err := c.GetFullAccountPortfolio(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFullAccountPortfolio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFullAccountPortfolio() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetMappedAccountPortfolioHistory(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.AccountHistoryRequest
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
			got, err := c.GetMappedAccountPortfolioHistory(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMappedAccountPortfolioHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMappedAccountPortfolioHistory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetPortfolioEarnings(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		account orbis_api.UserAccount
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.UserPortfolioEarnings
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
			got, err := c.GetPortfolioEarnings(tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPortfolioEarnings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPortfolioEarnings() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetPosition(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		symbol        *string
		accountNumber *string
		accountID     int64
		option        bool
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
			got, err := c.GetPosition(tt.args.symbol, tt.args.accountNumber, tt.args.accountID, tt.args.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPosition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPosition() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetUserInformation(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		loadRtb       bool
		accountName   *string
		accountNumber *string
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
			got, err := c.GetUserInformation(tt.args.loadRtb, tt.args.accountName, tt.args.accountNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInformation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserInformation() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetUserPreferences(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	tests := []struct {
		name    string
		fields  fields
		want    orbis_api.UserPreferences
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
			got, err := c.GetUserPreferences()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserPreferences() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserPreferences() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetUserWatchlist(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		loadQuotes bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    orbis_api.UserWatchlist
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
			got, err := c.GetUserWatchlist(tt.args.loadQuotes)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserWatchlist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserWatchlist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_PositionSearch(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		symbols    []string
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
			got, err := c.PositionSearch(tt.args.symbols, tt.args.loadQuotes)
			if (err != nil) != tt.wantErr {
				t.Errorf("PositionSearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PositionSearch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_RemoveEntryUserWatchlist(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		symbol string
		req    orbis_api.EntryUserWatchlistRequest
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
			got, err := c.RemoveEntryUserWatchlist(tt.args.symbol, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveEntryUserWatchlist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveEntryUserWatchlist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_RenameUserWatchlist(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req orbis_api.RenameUserWatchlistRequest
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
			got, err := c.RenameUserWatchlist(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RenameUserWatchlist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RenameUserWatchlist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UserPreferences(t *testing.T) {
	type fields struct {
		API    auth.API
		client http.Client
		config config.OrbisConfig
	}
	type args struct {
		req    map[string]string
		action string
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
			got, err := c.UserPreferences(tt.args.action, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserPreferences() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserPreferences() got = %v, want %v", got, tt.want)
			}
		})
	}
}
