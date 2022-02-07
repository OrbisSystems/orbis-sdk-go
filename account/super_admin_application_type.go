package account

import (
	"github.com/OrbisSystems/orbis-sdk-go/models/account"
)

const (
	apiUrlSuperAdminSaveNewVersion        = "/super-admin/application-types/versions/save"
	apiUrlSuperAdminRevertToVersion       = "/super-admin/application-types/versions/revert"
	apiUrlSuperAdminGetApplicationTypes   = "/super-admin/application-types/list"
	apiUrlSuperAdminGetApplicationType    = "/super-admin/application-types/get"
	apiUrlSuperAdminCreateApplication     = "/super-admin/application-types/create"
	apiUrlSuperAdminUpdateApplicationType = "/super-admin/application-types/update"
	apiUrlSuperAdminCopyApplicationType   = "/super-admin/application-types/copy"
	apiUrlSuperAdminDeleteApplicationType = "/super-admin/application-types/delete"
	apiUrlSuperAdminPushApplicationType   = "/super-admin/application-types/push"
	apiUrlSuperAdminApplyChanges          = "/super-admin/application-types/sync"
)

func (c *Client) GetApplicationTypeList(req account.ApplicationTypeListReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminGetApplicationTypes)
}

func (c *Client) GetApplicationTypeByID(ID int64, with []string) (interface{}, error) {
	body := map[string]interface{}{
		"with": with,
		"id":   ID,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminGetApplicationType)
}

func (c *Client) CreateApplicationType(req account.CreateApplicationTypeReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminCreateApplication)
}

func (c *Client) UpdateApplicationType(req account.UpdateApplicationTypeReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminUpdateApplicationType)
}

func (c *Client) CopyApplicationType(req account.CopyApplicationTypeReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminCopyApplicationType)
}

func (c *Client) DeleteApplicationType(ID int64) (interface{}, error) {
	body := map[string]int64{
		"id": ID,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminDeleteApplicationType)
}

func (c *Client) PushApplicationToEnv(req account.PushToEnvReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminPushApplicationType)
}

func (c *Client) ApplyChanges(req account.ApplicationTypeApplyChanges) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminApplyChanges)
}

func (c *Client) SaveNew(req account.VersionReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminSaveNewVersion)
}

func (c *Client) RevertToVersion(req account.RevertVersionReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminRevertToVersion)
}
