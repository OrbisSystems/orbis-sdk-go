package account

type Name string

const (
	English Name = "English"
)

type VersionReq struct {
	With        []string `json:"with"`
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
}

type RevertVersionReq struct {
	With    []string `json:"with"`
	ID      int64    `json:"id"`
	Version int64    `json:"version"`
}

type ApplicationTypeListReq struct {
	With   []string `json:"with"`
	FirmID int64    `json:"firm_id"`
}

type CreateApplicationTypeReq struct {
	BranchID     int64         `json:"branch_id"`
	FirmID       int64         `json:"firm_id"`
	Translations []Translation `json:"translations"`
}

type UpdateApplicationTypeReq struct {
	ID           int64         `json:"id"`
	BranchID     int64         `json:"branch_id"`
	Translations []Translation `json:"translations"`
}

type CopyApplicationTypeReq struct {
	With                []string      `json:"with"`
	ID                  int64         `json:"id"`
	ToApplicationTypeID int64         `json:"to_application_type_id"`
	Translations        []Translation `json:"translations"`
}

type PushToEnvReq struct {
	ID     int64  `json:"id"`
	PushTo string `json:"push_to"`
}

type ApplicationTypeApplyChanges struct {
	ApplicationType ApplicationType `json:"application_type"`
}

type ApplicationType struct {
	ID           int64              `json:"id"`
	Version      int64              `json:"version"`
	CreatedAt    string             `json:"created_at"`
	UpdatedAt    string             `json:"updated_at"`
	Translations []Translation      `json:"translations"`
	Pages        []PageApplyChanges `json:"pages"`
	Versions     []Version          `json:"versions"`
}

type PageApplyChanges struct {
	ID                int64                 `json:"id"`
	ApplicationTypeID int64                 `json:"application_type_id"`
	Position          int64                 `json:"position"`
	CreatedAt         string                `json:"created_at"`
	UpdatedAt         string                `json:"updated_at"`
	Translations      []Translation         `json:"translations"`
	Sections          []SectionApplyChanges `json:"sections"`
}

type SectionApplyChanges struct {
	ID           int64         `json:"id"`
	PageID       int64         `json:"page_id"`
	Position     int64         `json:"position"`
	CreatedAt    string        `json:"created_at"`
	UpdatedAt    string        `json:"updated_at"`
	Translations []Translation `json:"translations"`
	Questions    []Question    `json:"questions"`
}

type Question struct {
	ID                         int64                `json:"id"`
	SectionID                  int64                `json:"section_id"`
	TypeID                     int64                `json:"type_id"`
	RuleSetID                  int64                `json:"rule_set_id"`
	CustodianExchangeFormKeyID *int64               `json:"custodian_exchange_form_key_id"`
	ApexFormKeyID              *int64               `json:"apex_form_key_id"`
	DocumentTypeID             *int64               `json:"document_type_id"`
	PrefillQuestionID          interface{}          `json:"prefill_question_id"`
	IsOptional                 int64                `json:"is_optional"`
	HasMultipleAnswers         int64                `json:"has_multiple_answers"`
	IsEncrypted                int64                `json:"is_encrypted"`
	SystemIdentifier           string               `json:"system_identifier"`
	Position                   int64                `json:"position"`
	CreatedAt                  string               `json:"created_at"`
	UpdatedAt                  string               `json:"updated_at"`
	Translations               []Translation        `json:"translations"`
	Description                *Description         `json:"description"`
	ParentQuestions            []interface{}        `json:"parent_questions"`
	Options                    []OptionApplyChanges `json:"options"`
}

type Description struct {
	ID           int64         `json:"id"`
	QuestionID   int64         `json:"question_id"`
	CreatedAt    string        `json:"created_at"`
	UpdatedAt    string        `json:"updated_at"`
	Translations []interface{} `json:"translations"`
}

type OptionApplyChanges struct {
	ID               int64                     `json:"id"`
	QuestionID       int64                     `json:"question_id"`
	SystemIdentifier string                    `json:"system_identifier"`
	Position         int64                     `json:"position"`
	CreatedAt        string                    `json:"created_at"`
	UpdatedAt        string                    `json:"updated_at"`
	Translations     []TranslationApplyChanges `json:"translations"`
	ParentOptions    []interface{}             `json:"parent_options"`
}

type TranslationApplyChanges struct {
	ID               int64      `json:"id"`
	TranslatableID   int64      `json:"translatable_id"`
	TranslatableType string     `json:"translatable_type"`
	LocaleID         int64      `json:"locale_id"`
	Data             string     `json:"data"`
	CreatedAt        string     `json:"created_at"`
	UpdatedAt        string     `json:"updated_at"`
	Language         Locale     `json:"language"`
	Locale           LocaleInfo `json:"locale"`
}

type LocaleInfo struct {
	ID        int64       `json:"id"`
	Locale    Locale      `json:"locale"`
	Name      Name        `json:"name"`
	CreatedAt interface{} `json:"created_at"`
	UpdatedAt interface{} `json:"updated_at"`
}

type Version struct {
	ID                int64  `json:"id"`
	ApplicationTypeID int64  `json:"application_type_id"`
	Version           int64  `json:"version"`
	Form              string `json:"form"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type CreateBranchSuperAdminReq struct {
	With                     []string  `json:"with"`
	FirmID                   int64     `json:"firm_id"`
	Name                     string    `json:"name"`
	Code                     string    `json:"code"`
	EncryptionTrigger        string    `json:"encryption_trigger"`
	RetailLoginEnabled       bool      `json:"retail_login_enabled"`
	RequirePhoneVerification bool      `json:"require_phone_verification"`
	ApplicationTypeID        int64     `json:"application_type_id"`
	Services                 []Service `json:"services"`
}

type Service struct {
	ID int64 `json:"id"`
}

type UpdateBranchSuperAdminReq struct {
	With                     []string           `json:"with"`
	ID                       int64              `json:"id"`
	FirmID                   int64              `json:"firm_id"`
	Name                     string             `json:"name"`
	Code                     string             `json:"code"`
	EncryptionTrigger        string             `json:"encryption_trigger"`
	RetailLoginEnabled       bool               `json:"retail_login_enabled"`
	RequirePhoneVerification bool               `json:"require_phone_verification"`
	ApplicationTypeID        int64              `json:"application_type_id"`
	Services                 []ServiceUpdBranch `json:"services"`
}

type ServiceUpdBranch struct {
	ID       int64     `json:"id"`
	Settings *Settings `json:"settings,omitempty"`
}

type Settings struct {
	Region              *string              `json:"region,omitempty"`
	QueueURL            *string              `json:"queue_url,omitempty"`
	Arn                 *string              `json:"arn,omitempty"`
	URL                 *string              `json:"url,omitempty"`
	RequestMethod       *string              `json:"request_method,omitempty"`
	CustomRequestParams []CustomRequestParam `json:"custom_request_params,omitempty"`
}

type CustomRequestParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CreateApplicationTypePageReq struct {
	Translations     []Translation     `json:"translations"`
	ApplicationTypes []ApplicationType `json:"application_types"`
}

type ApplicationTypePageReq struct {
	ID       int64 `json:"id"`
	Position int64 `json:"position"`
}

type TranslationPage struct {
	Data     string `json:"data"`
	Language string `json:"language"`
}

type UpdateApplicationTypePageReq struct {
	ID               int64             `json:"id"`
	Translations     []Translation     `json:"translations"`
	ApplicationTypes []ApplicationType `json:"application_types"`
}

type CopyPageSuperAdmin struct {
	With                []string `json:"with"`
	ID                  int64    `json:"id"`
	ToApplicationTypeID int64    `json:"to_application_type_id"`
	Position            int64    `json:"position"`
}

type UpdatePositionsSuperAdminReq struct {
	Positions []Position `json:"positions"`
}

type Position struct {
	ApplicationTypeID       int64 `json:"application_type_id"`
	MoveToApplicationTypeID int64 `json:"move_to_application_type_id"`
	PageID                  int64 `json:"page_id"`
	Position                int64 `json:"position"`
}

type CreateSuperAdminSectionReq struct {
	PageID       int64         `json:"page_id"`
	Position     int64         `json:"position"`
	Translations []Translation `json:"translations"`
}

type UpdateSuperAdminSectionReq struct {
	ID           int64         `json:"id"`
	Translations []Translation `json:"translations"`
	Pages        []Page        `json:"pages"`
}

type PageUpdateSection struct {
	ID       int64 `json:"id"`
	Position int64 `json:"position"`
}

type CopySuperAdminSectionReq struct {
	With        []string `json:"with"`
	ID          int64    `json:"id"`
	ToSectionID int64    `json:"to_section_id"`
	Position    int64    `json:"position"`
}

type UpdateSuperAdminSectionPositionReq struct {
	Positions []PositionSection `json:"positions"`
}

type PositionSection struct {
	PageID       int64 `json:"page_id"`
	MoveToPageID int64 `json:"move_to_page_id"`
	SectionID    int64 `json:"section_id"`
	Position     int64 `json:"position"`
}

type GetListSuperAdminQuestions struct {
	With     []string         `json:"with"`
	Search   string           `json:"search"`
	Filters  FiltersQuestions `json:"filters"`
	OrderBy  OrderBy          `json:"order_by"`
	Paginate Paginate         `json:"paginate"`
}

type FiltersQuestions struct {
	QuestionType string `json:"question_type"`
	UpdatedAt    FromTo `json:"updated_at"`
	CreatedAt    FromTo `json:"created_at"`
}

type FromTo struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type OrderByQuestions struct {
	ID           CreatedAt `json:"id"`
	QuestionType CreatedAt `json:"question_type"`
	UpdatedAt    CreatedAt `json:"updated_at"`
	CreatedAt    CreatedAt `json:"created_at"`
}

type CreatedAt struct {
	Direction string `json:"direction"`
	Priority  int64  `json:"priority"`
}

type Paginate struct {
	PerPage int64 `json:"per_page"`
	Page    int64 `json:"page"`
}

type CreateSuperAdminQuestion struct {
	QuestionTypeID     int64                  `json:"question_type_id"`
	RuleSetID          int64                  `json:"rule_set_id"`
	IsOptional         int64                  `json:"is_optional"`
	IsEncrypted        int64                  `json:"is_encrypted"`
	HasMultipleAnswers int64                  `json:"has_multiple_answers"`
	Translations       []Translation          `json:"translations"`
	Options            []OptionCreateQuestion `json:"options"`
	SystemIdentifier   string                 `json:"system_identifier"`
	SectionID          int64                  `json:"section_id"`
	Position           int64                  `json:"position"`
	With               []string               `json:"with"`
}

type OptionCreateQuestion struct {
	Position         int64         `json:"position"`
	SystemIdentifier string        `json:"system_identifier"`
	Description      *Alert        `json:"description,omitempty"`
	Translations     []Translation `json:"translations"`
	Alert            *Alert        `json:"alert,omitempty"`
}

type Alert struct {
	Translations []Translation `json:"translations"`
}

type UpdateSuperAdminQuestionReq struct {
	With               []string               `json:"with"`
	ID                 int64                  `json:"id"`
	DependentQuestions []DependentQuestion    `json:"dependent_questions"`
	Options            []OptionUpdateQuestion `json:"options"`
}

type DependentQuestion struct {
	ParentOptionID      int64 `json:"parent_option_id"`
	DependentQuestionID int64 `json:"dependent_question_id"`
}

type OptionUpdateQuestion struct {
	ID               int64             `json:"id"`
	SystemIdentifier string            `json:"system_identifier"`
	Position         int64             `json:"position"`
	Translations     []Translation     `json:"translations"`
	DependentOptions []DependentOption `json:"dependent_options"`
}

type DependentOption struct {
	DependentOptionID int64 `json:"dependent_option_id"`
}

type CopySuperAdminQuestion struct {
	With        []string `json:"with"`
	ID          int64    `json:"id"`
	ToSectionID int64    `json:"to_section_id"`
	Position    int64    `json:"position"`
}

type UpdateSuperAdminQuestionPosition struct {
	Positions []QuestionPosition `json:"positions"`
}

type QuestionPosition struct {
	SectionID       int64 `json:"section_id"`
	MoveToSectionID int64 `json:"move_to_section_id"`
	QuestionID      int64 `json:"question_id"`
	Position        int64 `json:"position"`
}
