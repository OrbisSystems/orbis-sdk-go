package account

type WatchmenCompleteReq struct {
	ApplicationID string `json:"application_id"`
	Username      string `json:"username"`
	Password      string `json:"password"`
}

type PushALEsReq struct {
	Data []ALEsDatum `json:"data"`
}

type ALEsDatum struct {
	ID        int64       `json:"id"`
	Payload   PayloadALEs `json:"payload"`
	Timestamp string      `json:"timestamp"`
}

type PayloadALEs struct {
	Status        string `json:"status"`
	Account       string `json:"account"`
	ApplicationID string `json:"application_id"`
	ReferenceID   string `json:"reference_id"`
}

type PushSketchesReq struct {
	Data []PushSketchesDatum `json:"data"`
}

type PushSketchesDatum struct {
	ID        int64           `json:"id"`
	Payload   PayloadSketches `json:"payload"`
	Timestamp string          `json:"timestamp"`
}

type PayloadSketches struct {
	Source              string              `json:"source"`
	ApplicationID       string              `json:"application_id"`
	Account             string              `json:"account"`
	AutoAppealed        bool                `json:"auto_appealed"`
	CipID               string              `json:"cip_id"`
	ReferenceID         string              `json:"reference_id"`
	Status              string              `json:"status"`
	SketchInvestigation SketchInvestigation `json:"sketch_investigation"`
}

type SketchInvestigation struct {
	ID            string        `json:"id"`
	Status        string        `json:"status"`
	Archived      bool          `json:"archived"`
	History       []History     `json:"history"`
	Result        Result        `json:"result"`
	CipCategories []CipCategory `json:"cip_categories"`
}

type CipCategory struct {
	Name          string `json:"name"`
	CategoryState string `json:"category_state"`
}

type History struct {
	User        string `json:"user"`
	Timestamp   string `json:"timestamp"`
	StateChange string `json:"state_change"`
}

type Result struct {
	Evaluation    Evaluation    `json:"evaluation"`
	EquifaxResult EquifaxResult `json:"equifax_result"`
	OfacResult    OfacResult    `json:"ofac_result"`
}

type EquifaxResult struct {
	State   string  `json:"state"`
	Reasons Reasons `json:"reasons"`
	Results Results `json:"results"`
}

type Reasons struct {
	IdentityLocatedOnTertiaryDataSourceGoodMatch string `json:"IDENTITY_LOCATED_ON_TERTIARY_DATA_SOURCE_GOOD_MATCH"`
}

type Results struct {
	Accept Reasons `json:"accept"`
}

type Evaluation struct {
	EvaluatedState string `json:"evaluated_state"`
}

type OfacResult struct {
	OfacMatch string `json:"ofac_match"`
}
