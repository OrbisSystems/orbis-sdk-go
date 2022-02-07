package orbis_api

type DeactivateDeviceBody struct {
	Token string `json:"token"`
}

type DeactivateDeviceResponse struct {
	Success bool `json:"success"`
}
