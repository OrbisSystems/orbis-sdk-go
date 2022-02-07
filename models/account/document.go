package account

type UpdateDocumentReq struct {
	With          []string `json:"with"`
	ApplicationID int64    `json:"application_id"`
	DocumentID    int64    `json:"document_id"`
	Status        string   `json:"status"`
	Reason        string   `json:"reason"`
}
