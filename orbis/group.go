package orbis

import (
	"fmt"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"io"
	"net/http"
	"strconv"
)

var (
	manageGroup       = "/v1/groups/manage/%s"
	getGroupsList     = "/v1/groups/list"
	getGroupAccount   = "/v2/advisory/group/accounts/stats/%s"
	getGroupPortfolio = "/v2/advisory/group/portfolios/%s"
	getGroupRtb       = "/v2/advisory/group/rtb/%s"
	getGroupRtbs      = "/v2/advisory/group/rtbs/%s"
)

func (c *Client) ManageGroup(action string, req models.ManageGroupRequest) (io.ReadCloser, error) {
	return c.client.DoPostAndReturnBody(req, fmt.Sprintf(manageGroup, action))
}

func (c *Client) GetGroupsList() (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+getGroupsList, nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetGroupAccount(groupsID string) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(getGroupAccount, groupsID), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetGroupPortfolio(req models.PortfolioGroupRequest) (models.GroupPortfoliosGroupID, error) {
	httpReq, err := http.NewRequest(http.MethodGet, c.config.ApiUrl+getGroupPortfolio, nil)
	if err != nil {
		return models.GroupPortfoliosGroupID{}, err
	}
	q := httpReq.URL.Query()
	if req.LoadQuotes != nil {
		q.Add("loadQuotes", strconv.FormatBool(*req.LoadQuotes))
	}
	if req.Combined != nil {
		q.Add("combined", strconv.FormatBool(*req.Combined))
	}
	if req.LoadIndustries != nil {
		q.Add("loadIndustries", strconv.FormatBool(*req.LoadIndustries))
	}
	if req.LoadUpgrades != nil {
		q.Add("loadUpgrades", strconv.FormatBool(*req.LoadUpgrades))
	}
	if req.LoadEarnings != nil {
		q.Add("loadEarnings", strconv.FormatBool(*req.LoadEarnings))
	}
	if req.LoadCost != nil {
		q.Add("loadCost", strconv.FormatBool(*req.LoadCost))
	}
	if req.MarketCapWeights != nil {
		q.Add("marketCapWeights", strconv.FormatBool(*req.MarketCapWeights))
	}

	r, err := c.client.Get(q.Encode(), nil)
	if err != nil {
		return models.GroupPortfoliosGroupID{}, err
	}

	var resp models.GroupPortfoliosGroupID
	err = c.client.UnmarshalAndCheckOk(&resp, r)
	if err != nil {
		return models.GroupPortfoliosGroupID{}, err
	}

	return resp, nil
}

func (c *Client) GetGroupRtb(groupID string) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(getGroupRtb, groupID), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}

func (c *Client) GetGroupRtbs(groupID string) (io.ReadCloser, error) {
	r, err := c.client.Get(c.config.ApiUrl+fmt.Sprintf(getGroupRtbs, groupID), nil)
	if err != nil {
		return nil, err
	}

	return c.client.GetBodyAndCheckOK(r)
}
