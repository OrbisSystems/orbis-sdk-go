package account

import "github.com/OrbisSystems/orbis-sdk-go/models/account"

const (
	apiUrlUsersList        = "/users/list"
	apiUrlUserGet          = "/users/list"
	apiUrlUserCreate       = "/users/list"
	apiUrlUserUpdate       = "/users/list"
	apiUrlUserEnableBranch = "/users/list"
)

func (c *Client) GetUserList(userReq account.GetUsers) (interface{}, error) {
	return c.client.MarshallAndSendPost(userReq, c.config.ApiUrl+apiUrlUsersList)
}
func (c *Client) GetUser(ID int64, with []string) (interface{}, error) {
	body := map[string]interface{}{
		"user_id": ID,
		"with":    with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlUserGet)
}
func (c *Client) CreateUser(createUser account.CreateUser) (interface{}, error) {
	return c.client.MarshallAndSendPost(createUser, c.config.ApiUrl+apiUrlUserCreate)
}
func (c *Client) UpdateUser(createUser account.CreateUser) (interface{}, error) {
	return c.client.MarshallAndSendPost(createUser, c.config.ApiUrl+apiUrlUserUpdate)
}
func (c *Client) EnableBranchToUser(userID string, branchID string) (interface{}, error) {
	body := map[string]string{
		"user_id":   userID,
		"branch_id": branchID,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlUserEnableBranch)
}
