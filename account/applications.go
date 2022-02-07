package account

import (
	"github.com/OrbisSystems/orbis-sdk-go/models/account"
	"os"
)

const (
	apiUrlGetAvailableAPIStatuses           = "/applications/get-statuses-list"
	apiUrlGetApplicationList                = "/applications/list"
	apiUrlGetApplication                    = "/applications/get"
	apiUrlGetApplicationType                = "/applications/get-application-types"
	apiUrlGetApplicationForm                = "/applications/get-form"
	apiUrlCreateApplication                 = "/applications/create"
	apiUrlApplicationAnswer                 = "/applications/answer"
	apiUrlApplicationAnswerMultiple         = "/answer-multiple"
	apiUrlApplicationDeleteAnswer           = "/applications/delete-answer"
	apiUrlApplicationSubmit                 = "applications/submit"
	apiUrlApplicationEnableReSubmit         = "/applications/enable-resubmit"
	apiUrlApplicationBackgroundCheck        = "/applications/background-check"
	apiUrlApplicationComplianceStatusUpdate = "/applications/background-check"
	apiUrlApplicationDelete                 = "/applications/delete"
	apiUrlApplicationAddNote                = "/applications/add-note"
	apiUrlApplicationGetHistory             = "/applications/get-history"
	apiUrlApplicationAccountExist           = "/applications/account-exists"
	apiUrlApplicationVerifyDocuments        = "/applications/verify-documents"
)

func (c *Client) GetAvailableAppStatuses() (interface{}, error) {
	return c.client.MarshallAndSendPost(nil, c.config.ApiUrl+apiUrlGetAvailableAPIStatuses)
}

func (c *Client) GetApplicationsList(req account.ListApplicationsReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlGetApplicationList)
}

func (c *Client) GetApplication(applicationID int64, with []string) (interface{}, error) {
	body := map[string]interface{}{
		"application_id": applicationID,
		"with":           with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlGetApplication)
}

func (c *Client) GetApplicationTypes(with []string) (interface{}, error) {
	body := map[string][]string{
		"with": with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlGetApplicationType)
}

func (c *Client) GetApplicationForm(applicationTypeID int64, with []string) (interface{}, error) {
	body := map[string]interface{}{
		"application_type_id": applicationTypeID,
		"with":                with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlGetApplicationForm)
}

func (c *Client) CreateApplication(userID int64, with []string) (interface{}, error) {
	body := map[string]interface{}{
		"user_id": userID,
		"with":    with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlCreateApplication)
}

//todo implement later
func (c *Client) Answer(applicationID, questionID string, data os.File) (interface{}, error) {
	panic("implement me")
}

func (c *Client) AnswerMultiply(req account.AnswerMultiple) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlApplicationAnswerMultiple)
}

//TODO define two different GET APPLICATIONS func

func (c *Client) DeleteAnswer(applicationID, answerID int64) (interface{}, error) {
	body := map[string]int64{
		"application_id": applicationID,
		"answer_id":      answerID,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlApplicationDeleteAnswer)
}
func (c *Client) Submit(applicationID int64) (interface{}, error) {
	body := map[string]int64{
		"application_id": applicationID,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlApplicationSubmit)
}
func (c *Client) EnableResubmit(applicationID int64) (interface{}, error) {
	body := map[string]int64{
		"application_id": applicationID,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlApplicationEnableReSubmit)
}
func (c *Client) BackGroundCheck(applicationID int64, with []string) (interface{}, error) {
	body := map[string]interface{}{
		"application_id": applicationID,
		"with":           with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlApplicationBackgroundCheck)
}
func (c *Client) ComplianceStatusUpdate(applicationID int64, status string, with []string) (interface{}, error) {
	body := map[string]interface{}{
		"application_id": applicationID,
		"status":         status,
		"with":           with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlApplicationComplianceStatusUpdate)
}
func (c *Client) DeleteApplication(applicationID int64) (interface{}, error) {
	body := map[string]int64{
		"application_id": applicationID,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlApplicationDelete)
}
func (c *Client) NoteApplication(applicationID int64, note string, with []string) (interface{}, error) {
	body := map[string]interface{}{
		"application_id": applicationID,
		"note":           note,
		"with":           with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlApplicationAddNote)
}

func (c *Client) GetHistoryOfApplication(applicationHistoryID int64, with []string) (interface{}, error) {
	body := map[string]interface{}{
		"application_history_id": applicationHistoryID,
		"with":                   with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlApplicationGetHistory)
}

func (c *Client) AccountExist(IDNumber string) (interface{}, error) {
	body := map[string]string{
		"id_number": IDNumber,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlApplicationAccountExist)
}

func (c *Client) ApplicationVerifyDocument(applicationID int64, applicantType string, with []string) (interface{}, error) {
	body := map[string]interface{}{
		"application_id":  applicationID,
		"applicationType": applicantType,
		"with":            with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlApplicationVerifyDocuments)
}
