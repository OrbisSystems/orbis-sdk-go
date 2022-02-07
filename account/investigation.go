package account

import "os"

const (
	apiUrlUploadDocument = "/applications/custodian/investigations/upload-document"
	apiUrlAppeal         = "/applications/custodian/investigations/appeal"
)

//TODO add uploading file
func (c *Client) UploadDocument(applicationID, name string, data os.File) (interface{}, error) {
	panic("implement me")
}

func (c *Client) Appeal(investigationID int64, action, notes string) (interface{}, error) {
	body := map[string]interface{}{
		"investigation_id": investigationID,
		"action":           action,
		"notes":            notes,
	}
	return c.client.MarshallAndSendPost(body, c.config.ApiUrl+apiUrlUserEnableBranch)
}
