package orbis_api

type ScheduleAllocationRequest struct {
	Adjustment *Adjustment `json:"adjustment"`
	Allocation Allocation  `json:"allocation"`
}

type Adjustment struct {
	ID int64 `json:"id"`
}

type Allocation struct {
	Transaction string   `json:"transaction"`
	Quote       Quote    `json:"quote"`
	Targets     []Target `json:"targets"`
}

type Target struct {
	Quantity float64           `json:"quantity"`
	Account  AccountAllocation `json:"account"`
}

type AccountAllocation struct {
	AccountNumber string `json:"accountNumber"`
}
