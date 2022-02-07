package account

import (
	"github.com/OrbisSystems/orbis-sdk-go/models/account"
)

const (
	apiUrlSuperAdminGetPagesList       = "/super-admin/pages/list"
	apiUrlSuperAdminCreatePage         = "/super-admin/pages/create"
	apiUrlSuperAdminUpdatePage         = "/super-admin/pages/update"
	apiUrlSuperAdminCopyPage           = "/super-admin/pages/copy"
	apiUrlSuperAdminDeletePage         = "/super-admin/pages/delete"
	apiUrlSuperAdminUpdatePagePosition = "/super-admin/pages/update-positions"
)

func (c *Client) GetListPages(with []string) (interface{}, error) {
	body := map[string][]string{
		"with": with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminGetPagesList)
}

func (c *Client) CreatePage(req account.CreateApplicationTypePageReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminCreatePage)
}

func (c *Client) UpdatePage(req account.UpdateApplicationTypePageReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminUpdatePage)
}

func (c *Client) CopyPage(req account.CopyPageSuperAdmin) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminCopyPage)
}

func (c *Client) DeletePage(ID int64) (interface{}, error) {
	body := map[string]interface{}{
		"id": ID,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminDeletePage)
}

func (c *Client) UpdatePagePosition(req account.UpdatePositionsSuperAdminReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminUpdatePagePosition)
}
