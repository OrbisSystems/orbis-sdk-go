package account

type UpdateServicesReq struct {
	Data []Datum `json:"data"`
}

type Datum struct {
	ID       int64    `json:"id"`
	Services []string `json:"services"`
}
