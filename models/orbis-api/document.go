package orbis_api

import "time"

type ConfirmDocumentResponse struct {
	URL string `json:"url"`
}

type ServiceDocsRequest struct {
	AccountNumber *string
	DocType       *string
	BeginMonth    *time.Time
	BeginYear     *time.Time
	EndMonth      *time.Time
	EndYear       *time.Time
}

type ServiceDocsResponse struct {
	PensonStatus PensonStatus          `json:"pensonStatus"`
	DocumentList []DocumentListElement `json:"documentList"`
}

type PensonStatus struct {
	StatusDescription string `json:"statusDescription"`
	StatusDateTime    string `json:"statusDateTime"`
}
