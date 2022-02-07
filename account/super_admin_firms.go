package account

const (
	apiUrlSuperAdminGetListFirms = "/super-admin/firms/list"
	apiUrlSuperAdminGetFirm      = "/super-admin/firms/get"
	apiUrlSuperAdminCreateFirm   = "/super-admin/firms/create"
	apiUrlSuperAdminUpdateFirm   = "/super-admin/firms/update"
	apiUrlSuperAdminDeleteFirm   = "/super-admin/firms/delete"
)

func (c *Client) GetListFirms(with []string) (interface{}, error) {
	body := map[string][]string{
		"with": with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminGetListFirms)
}

func (c *Client) GetFirm(ID int64, with []string) (interface{}, error) {
	body := map[string]interface{}{
		"with": with,
		"id":   ID,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminGetFirm)
}

func (c *Client) CreateFirm(name string) (interface{}, error) {
	body := map[string]string{
		"name": name,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminCreateFirm)
}

func (c *Client) UpdateFirm(ID int64, name string) (interface{}, error) {
	body := map[string]interface{}{
		"id":   ID,
		"name": name,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminUpdateFirm)
}

func (c *Client) DeleteFirm(ID int64) (interface{}, error) {
	body := map[string]int64{
		"id": ID,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminDeleteFirm)
}
