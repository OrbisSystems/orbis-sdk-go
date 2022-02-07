package account

type ListApplicationsReq struct {
	With     []string `json:"with"`
	Search   string   `json:"search"`
	Filters  Filters  `json:"filters"`
	OrderBy  OrderBy  `json:"order_by"`
	CSV      []CSV    `json:"csv"`
	Paginate Paginate `json:"paginate"`
}

type CSV struct {
	Name  string      `json:"name"`
	Alias *AliasUnion `json:"alias"`
}

type AliasClass struct {
	Alias         *string      `json:"alias,omitempty"`
	DateFormatter *string      `json:"date_formatter,omitempty"`
	Concatenate   *Concatenate `json:"concatenate,omitempty"`
}

type Concatenate struct {
	Operator string   `json:"operator"`
	Aliases  []string `json:"aliases"`
}

type ApplicationsFilters struct {
	Status               string   `json:"status"`
	UsOrNot              string   `json:"us_or_not"`
	Firm                 string   `json:"firm"`
	Branch               string   `json:"branch"`
	User                 string   `json:"user"`
	RepCode              string   `json:"rep_code"`
	SubmittedAt          TedAt    `json:"submitted_at"`
	CustodianCompletedAt TedAt    `json:"custodian_completed_at"`
	UpdatedAt            TedAt    `json:"updated_at"`
	CreatedAt            TedAt    `json:"created_at"`
	Answers              []Answer `json:"answers"`
}

type ApplicationsAnswer struct {
	QuestionSystemIdentifier string `json:"question_system_identifier"`
	Answer                   string `json:"answer"`
}

type TedAt struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type ApplicationsOrderBy struct {
	Status               Branch `json:"status"`
	Branch               Branch `json:"branch"`
	RepCode              Branch `json:"rep_code"`
	SubmittedAt          Branch `json:"submitted_at"`
	CustodianCompletedAt Branch `json:"custodian_completed_at"`
	CreatedAt            Branch `json:"created_at"`
	UpdatedAt            Branch `json:"updated_at"`
}

type ApplicationsBranch struct {
	Direction string `json:"direction"`
	Priority  int64  `json:"priority"`
}

type ApplicationsPaginate struct {
	PerPage int64 `json:"per_page"`
	Page    int64 `json:"page"`
}

type AliasUnion struct {
	AliasClass *AliasClass
	String     *string
}
