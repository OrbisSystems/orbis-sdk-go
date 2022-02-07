package orbis_api

type NotificationLinkRequest struct {
	Name     string `json:"name"`
	QueueURL string `json:"queueUrl"`
}

type NotificationUnLinkRequest struct {
	Name string `json:"name"`
}
