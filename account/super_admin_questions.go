package account

import (
	"github.com/OrbisSystems/orbis-sdk-go/models/account"
)

const (
	apiUrlSuperAdminQuestionsList            = "/super-admin/questions/list"
	apiUrlSuperAdminGetWhereHasOptions       = "/super-admin/questions/get-where-has-options"
	apiUrlSuperAdminQuestionsCreate          = "/super-admin/questions/create"
	apiUrlSuperAdminQuestionsUpdate          = "/super-admin/questions/update"
	apiUrlSuperAdminQuestionsCopy            = "/super-admin/questions/copy"
	apiUrlSuperAdminQuestionsDelete          = "/super-admin/questions/delete"
	apiUrlSuperAdminQuestionsUpdatePositions = "/super-admin/questions/update-positions"
)

func (c *Client) GetListQuestions(req account.GetListSuperAdminQuestions) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminQuestionsList)
}

func (c *Client) GetWhereHasOptionsQuestion(applicationTypeID int64, with []string) (interface{}, error) {
	body := map[string]interface{}{
		"application_type_id": applicationTypeID,
		"with":                with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminGetWhereHasOptions)
}

func (c *Client) CreateQuestion(req account.CreateSuperAdminQuestion) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminQuestionsCreate)
}

func (c *Client) UpdateQuestion(req account.UpdateSuperAdminQuestionReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminQuestionsUpdate)
}

func (c *Client) CopyQuestion(req account.CopySuperAdminQuestion) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminQuestionsCopy)
}

func (c *Client) DeleteQuestion(ID int64) (interface{}, error) {
	body := map[string]int64{
		"id": ID,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlSuperAdminQuestionsDelete)
}

func (c *Client) UpdateQuestionPosition(req account.UpdateSuperAdminQuestionPosition) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlSuperAdminQuestionsUpdatePositions)
}
