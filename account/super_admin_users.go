package account

const (
	apiUrlSuperAdminAddFirmToUser                  = "/super-admin/users/firm/add"
	apiUrlSuperAdminRemoveFirmToUser               = "/super-admin/users/firm/remove"
	apiUrlSuperAdminSetDefaultComplianceFirmToUser = "/super-admin/users/firm/set-default-compliance"
)

func (c *Client) AddToFirm(userID, firmID int64, with []string) (interface{}, error) {
	return sendAndCheck(userID, firmID, with, apiUrlSuperAdminAddFirmToUser, c)
}

func (c *Client) RemoveFromFirm(userID, firmID int64, with []string) (interface{}, error) {
	return sendAndCheck(userID, firmID, with, apiUrlSuperAdminRemoveFirmToUser, c)
}

func (c *Client) SetDefault(userID, firmID int64, with []string) (interface{}, error) {
	return sendAndCheck(userID, firmID, with, apiUrlSuperAdminSetDefaultComplianceFirmToUser, c)
}

func sendAndCheck(userID int64, firmID int64, with []string, url string, c *Client) (interface{}, error) {
	body := map[string]interface{}{
		"user_id": userID,
		"firm_id": firmID,
		"with":    with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+url)
}
