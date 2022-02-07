package account

import (
	"github.com/OrbisSystems/orbis-sdk-go/models/account"
)

const (
	apiUrlSuperAdminGetListSections       = "/super-admin/sections/list"
	apiUrlSuperAdminCreateSection         = "/super-admin/sections/create"
	apiUrlSuperAdminUpdateSection         = "/super-admin/sections/update"
	apiUrlSuperAdminCopySection           = "/super-admin/sections/update"
	apiUrlSuperAdminDeleteSection         = "/super-admin/sections/delete"
	apiUrlSuperAdminUpdatePositionSection = "/super-admin/sections/update-positions"
)

func (c *Client) GetListSections(with []string) (interface{}, error) {
	body := map[string][]string{
		"with": with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminGetListSections)
}

func (c *Client) CreateSection(req account.CreateSuperAdminSectionReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminCreateSection)
}

func (c *Client) UpdateSection(req account.UpdateSuperAdminSectionReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminUpdateSection)
}

func (c *Client) CopySection(req account.CopySuperAdminSectionReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminCopySection)
}

func (c *Client) DeleteSection(ID int64) (interface{}, error) {
	body := map[string]int64{
		"id": ID,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminDeleteSection)
}

func (c *Client) UpdateSectionPosition(req account.UpdateSuperAdminSectionPositionReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminUpdatePositionSection)
}
