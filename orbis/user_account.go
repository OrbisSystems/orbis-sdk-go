package orbis

import (
	"fmt"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
	"net/http"
	"strconv"
)

var (
	getUsersList    = "/v2/advisory/users/list"
	getUserAccounts = "/v2/advisory/users/accounts"
	getAccountStats = "/v2/advisory/users/accounts/stats"
	getAccountNotes = "/v2/advisory/users/account/notes"
	getNotes        = "/v2/advisory/users/notes"
	userNotes       = "/v2/advisory/users/notes/%s"
	accountNotes    = "/v2/advisory/users/notes/%s"
)

func (c *Client) GetUsersList() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+getUsersList, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetUserAccounts(ID int64, loadRtb bool) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getUserAccounts, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("id", strconv.FormatInt(ID, 10))
	q.Add("loadRtb", strconv.FormatBool(loadRtb))
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetAccountStats() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+getAccountStats, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetAccountNotes(accountID int64) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getAccountNotes, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("aid", strconv.FormatInt(accountID, 10))
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetNotes(userID int64) (io.ReadCloser, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getNotes, nil)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	q.Add("userId", strconv.FormatInt(userID, 10))
	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) UserNotes(action string, req models.UserNotesRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+fmt.Sprintf(userNotes, action))
}

func (c *Client) AccountNotes(action string, req models.AccountNotesRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, c.config.ApiUrl+fmt.Sprintf(accountNotes, action))
}
