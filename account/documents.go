package account

import "github.com/OrbisSystems/orbis-sdk-go/models/account"

const (
	apiUrlDocumentsGenerateQRCode = "/applications/generate-qr-code"
	apiUrlDocumentsGetQuestions   = "/applications/get-document-questions"
	apiUrlDocumentsGetAnswers     = "/applications/get-document-answers" // ask questions ?
	apiUrlDocumentsUpdateStatus   = "/documents/update-status"           // ask questions ?
	apiUrlDocumentsGet            = "/documents/get"
)

func (c *Client) GenerateQRCode(applicationID int64) (interface{}, error) {
	body := map[string]int64{
		"application_id": applicationID,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlDocumentsGenerateQRCode)
}

func (c *Client) GetDocumentQuestions(applicationID int64, with []string) (interface{}, error) {
	body := map[string]interface{}{
		"application_id": applicationID,
		"with":           with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlDocumentsGetQuestions)
}

func (c *Client) GetDocumentAnswers(applicationID int64, with []string) (interface{}, error) {
	body := map[string]interface{}{
		"application_id": applicationID,
		"with":           with,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlDocumentsGetAnswers)
}

func (c *Client) UpdateDocumentStatus(req account.UpdateDocumentReq) (interface{}, error) {
	return c.client.MarshallAndSendPost(req, c.config.ApiUrl+apiUrlDocumentsUpdateStatus)
}

func (c *Client) GetDocument(applicationID, documentID int64) (interface{}, error) {
	body := map[string]int64{
		"application_id": applicationID,
		"document_id":    documentID,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlDocumentsGet)
}
