package orbis_api

type UserNotesRequest struct {
	UserID  int64  `json:"userId"`
	Content string `json:"content"`
}

type AccountNotesRequest struct {
	AccountID int64  `json:"accountId"`
	Content   string `json:"content"`
}
