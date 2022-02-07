package account

import (
	"github.com/OrbisSystems/orbis-sdk-go/models/account"
)

const (
	apiUrlSuperAdminGetListBranches = "/super-admin/branches/list"
	apiUrlSuperAdminGetBranch       = "/super-admin/branches/get"
	apiUrlSuperAdminCreateBranch    = "/super-admin/branches/create"
	apiUrlSuperAdminUpdateBranch    = "/super-admin/branches/update"
	apiUrlSuperAdminDeleteBranch    = "/super-admin/branches/delete"
)

func (c *Client) GetSuperAdminListBranches(firmID int64, with []string) (interface{}, error) {
	body := map[string]interface{}{
		"firm_id": firmID,
		"with":    with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminGetListBranches)
}

func (c *Client) GetSuperAdminBranch(ID int64) (interface{}, error) {
	body := map[string]int64{
		"id": ID,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminGetBranch)
}

func (c *Client) CreateSuperAdminBranch(req account.CreateBranchSuperAdminReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminCreateBranch)
}

func (c *Client) UpdateSuperAdminBranch(req account.UpdateBranchSuperAdminReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminUpdateBranch)
}

func (c *Client) DeleteSuperAdminBranch(ID int64) (interface{}, error) {
	body := map[string]int64{
		"id": ID,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminDeleteBranch)
}
