package orbis_api

type AdvisoryInfoUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	Country     string `json:"country"`
	Phone       string `json:"phone"`
}

type AvatarUploadRequest struct {
	Contents string `json:"contents"`
	Filename string `json:"filename"`
}

type UploadRequest struct {
	Contents string `json:"contents"`
	Filename string `json:"filename"`
	Tag      string `json:"tag"`
	ID       int64  `json:"id"`
}
