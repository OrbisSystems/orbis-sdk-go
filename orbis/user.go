package orbis

import (
	"bytes"
	"encoding/json"
	"fmt"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var (
	getUserInformation               = "/user/info"
	getAccountInformation            = "/user/account"
	getAccountPortfolio              = "/user/portfolio"
	getFullAccountPortfolio          = "/user/portfolio/full"
	getAccountTradingLimit           = "/user/balance"
	getAccountBalance                = "/user/rtb"
	getAccountBalanceHistory         = "/user/rtb/history"
	getAccountPortfolioHistory       = "/user/position/history"
	getMappedAccountPortfolioHistory = "/user/position/history/mapped"
	getAccountHistory                = "/user/account/history"
	getPosition                      = "/user/position"
	positionSearch                   = "/user/position/search"
	getUserWatchlist                 = "/user/watchlist"
	getPortfolioEarnings             = "/user/portfolio/earnings"
	getUserPreferences               = "/user/preferences"
	userPreferences                  = "/user/preferences/%s"
	createUserWatchlist              = "/user/watchlist/create"
	renameUserWatchlist              = "/user/watchlist/rename"
	deleteUserWatchlist              = "/user/watchlist/delete"
	addEntryUserWatchlist            = "/user/watchlist/addEntry"
	removeEntryUserWatchlist         = "/user/watchlist/removeEntry"
	firebaseUserRegister             = "/user/droid/register"
)

func (c *Client) GetUserInformation(loadRtb bool, accountName, accountNumber *string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getUserInformation, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("loadRtb", strconv.FormatBool(loadRtb))
	if accountName != nil {
		q.Add("accountName", *accountName)
	}
	if accountNumber != nil {
		q.Add("accountNumber", *accountNumber)
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetAccountInformation(accountNumber *string, accountID *int64) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getAccountInformation, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	if accountID != nil {
		q.Add("accountId", strconv.FormatInt(*accountID, 10))
	}
	if accountNumber != nil {
		q.Add("accountNumber", *accountNumber)
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetAccountPortfolio(req models.AccountPortfolioRequest) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getAccountPortfolio, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	if req.LoadEarnings != nil {
		q.Add("loadEarnings", strconv.FormatBool(*req.LoadEarnings))
	}
	if req.LoadIndustries != nil {
		q.Add("loadIndustries", strconv.FormatBool(*req.LoadIndustries))
	}
	if req.LoadUpgrades != nil {
		q.Add("loadUpgrades", strconv.FormatBool(*req.LoadUpgrades))
	}
	if req.LoadQuotes != nil {
		q.Add("loadQuotes", strconv.FormatBool(*req.LoadQuotes))
	}
	if req.AccountID != nil {
		q.Add("accountId", strconv.FormatInt(*req.AccountID, 10))
	}
	if req.AccountNumber != nil {
		q.Add("accountNumber", *req.AccountNumber)
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetFullAccountPortfolio(req models.FullAccountPortfolioRequest) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getFullAccountPortfolio, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	if req.LoadEarnings != nil {
		q.Add("loadEarnings", strconv.FormatBool(*req.LoadEarnings))
	}
	if req.LoadIndustries != nil {
		q.Add("loadIndustries", strconv.FormatBool(*req.LoadIndustries))
	}
	if req.LoadUpgrades != nil {
		q.Add("loadUpgrades", strconv.FormatBool(*req.LoadUpgrades))
	}
	if req.LoadQuotes != nil {
		q.Add("loadQuotes", strconv.FormatBool(*req.LoadQuotes))
	}
	if req.AccountID != nil {
		q.Add("accountId", strconv.FormatInt(*req.AccountID, 10))
	}
	if req.AccountNumber != nil {
		q.Add("accountNumber", *req.AccountNumber)
	}
	if req.MarketCapBreakdown != nil {
		q.Add("marketCapBreakdown", strconv.FormatBool(*req.MarketCapBreakdown))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetAccountTradingLimit(accountNumber, currency string, accountID int64) (models.AccountTradingLimit, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getAccountTradingLimit, nil)
	if err != nil {
		return models.AccountTradingLimit{}, err
	}
	q := httpReq.URL.Query()
	q.Add("accountID", strconv.FormatInt(accountID, 10))
	q.Add("accountNumber", accountNumber)
	q.Add("currency", currency)

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return models.AccountTradingLimit{}, err
	}

	var resp models.AccountTradingLimit
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.AccountTradingLimit{}, err
	}

	return resp, nil
}

func (c *Client) GetAccountBalance(accountNumber, currency string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getAccountTradingLimit, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("accountNumber", accountNumber)
	q.Add("currency", currency)

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetAccountBalanceHistory(req models.AccountHistoryRequest) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getAccountTradingLimit, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()

	q.Add("count", strconv.FormatInt(req.Count, 10))
	q.Add("page", strconv.FormatInt(req.Page, 10))
	q.Add("accountId", strconv.FormatInt(req.AccountID, 10))
	q.Add("account", req.AccountNumber)
	if req.DateFrom != nil {
		q.Add("dateFrom", req.DateFrom.Format("01/02/2006"))
	}
	if req.DateTo != nil {
		q.Add("dateTo", req.DateFrom.Format("01/02/2006"))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetAccountPortfolioHistory(req models.AccountHistoryRequest) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getAccountPortfolioHistory, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()

	q.Add("count", strconv.FormatInt(req.Count, 10))
	q.Add("page", strconv.FormatInt(req.Page, 10))
	q.Add("accountId", strconv.FormatInt(req.AccountID, 10))
	q.Add("account", req.AccountNumber)
	if req.DateFrom != nil {
		q.Add("dateFrom", req.DateFrom.Format("01/02/2006"))
	}
	if req.DateTo != nil {
		q.Add("dateTo", req.DateFrom.Format("01/02/2006"))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetMappedAccountPortfolioHistory(req models.AccountHistoryRequest) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getMappedAccountPortfolioHistory, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()

	q.Add("count", strconv.FormatInt(req.Count, 10))
	q.Add("page", strconv.FormatInt(req.Page, 10))
	q.Add("accountId", strconv.FormatInt(req.AccountID, 10))
	q.Add("account", req.AccountNumber)
	if req.DateFrom != nil {
		q.Add("dateFrom", req.DateFrom.Format("01/02/2006"))
	}
	if req.DateTo != nil {
		q.Add("dateTo", req.DateFrom.Format("01/02/2006"))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetAccountHistory(req models.AccountHistoryRequest) (models.AccountHistory, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getAccountPortfolioHistory, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()

	q.Add("count", strconv.FormatInt(req.Count, 10))
	q.Add("page", strconv.FormatInt(req.Page, 10))
	q.Add("accountId", strconv.FormatInt(req.AccountID, 10))
	q.Add("account", req.AccountNumber)
	if req.DateFrom != nil {
		q.Add("dateFrom", req.DateFrom.Format("01/02/2006"))
	}
	if req.DateTo != nil {
		q.Add("dateTo", req.DateFrom.Format("01/02/2006"))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp models.AccountHistory
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetPosition(symbol, accountNumber *string, accountID int64, option bool) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getAccountPortfolioHistory, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()

	q.Add("accountId", strconv.FormatInt(accountID, 10))
	q.Add("option", strconv.FormatBool(option))
	if accountNumber != nil {
		q.Add("account", *accountNumber)
	}
	if symbol != nil {
		q.Add("symbol", *symbol)
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) PositionSearch(symbols []string, loadQuotes bool) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+positionSearch, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()

	if len(symbols) > 0 {
		q.Add("symbols", strings.Join(symbols, ","))
	}
	q.Add("loadQuotes", strconv.FormatBool(loadQuotes))

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetUserWatchlist(loadQuotes bool) (models.UserWatchlist, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+positionSearch, nil)
	if err != nil {
		return models.UserWatchlist{}, err
	}
	q := httpReq.URL.Query()

	q.Add("loadQuotes", strconv.FormatBool(loadQuotes))

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return models.UserWatchlist{}, err
	}

	var resp models.UserWatchlist
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.UserWatchlist{}, err
	}

	return resp, nil
}

func (c *Client) GetPortfolioEarnings(account models.UserAccount) (models.UserPortfolioEarnings, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getPortfolioEarnings, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("account", account.AccountNumber)
	q.Add("accountId", strconv.FormatInt(account.AccountID, 10))

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp models.UserPortfolioEarnings
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.UserPortfolioEarnings{}, err
	}

	return resp, nil
}

func (c *Client) GetUserPreferences() (models.UserPreferences, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getUserPreferences, nil)
	if err != nil {
		return nil, err
	}
	r, err := c.client.Get(httpReq.URL.Query().Encode(), nil)
	if err != nil {
		return nil, err
	}

	var resp models.UserPreferences
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.UserPreferences{}, err
	}

	return resp, nil
}

func (c *Client) UserPreferences(action string, req map[string]string) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+fmt.Sprintf(userPreferences, action), nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()

	b, err := json.Marshal(req)
	r, err := c.client.Post(q.Encode(), bytes.NewBuffer(b), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) CreateUserWatchlist(req models.CreateWatchlistRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+createUserWatchlist)
}

func (c *Client) RenameUserWatchlist(req models.RenameUserWatchlistRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+renameUserWatchlist)
}

func (c *Client) DeleteUserWatchlist(req models.DeleteUserWatchlistRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+deleteUserWatchlist)
}

func (c *Client) AddEntryUserWatchlist(req models.EntryUserWatchlistRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+addEntryUserWatchlist)
}

func (c *Client) RemoveEntryUserWatchlist(symbol string, req models.EntryUserWatchlistRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+removeEntryUserWatchlist)
}

func (c *Client) FirebaseUserRegister(token string) (io.ReadCloser, error) {
	req := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+firebaseUserRegister)
}
