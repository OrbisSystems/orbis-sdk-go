package account

type Empty struct {
}

// Send Phone Verification
//
// POST {{domain}}/api/auth/verifications/phone/send
//
// # Send phone verification
type SendPhoneVerification struct {
	Status            bool                           `json:"status"`
	HTTPStatusMessage string                         `json:"http_status_message"`
	Data              interface{}                    `json:"data"`
	Messages          *SendPhoneVerificationMessages `json:"messages,omitempty"`
}

type SendPhoneVerificationMessages struct {
	Token []string `json:"token"`
}

// Verify Phone
//
// POST {{domain}}/api/auth/verifications/phone/verify
//
// # Veryfy user's phone using code sent to his phone
type VerifyPhone struct {
	Status            bool                           `json:"status"`
	HTTPStatusMessage string                         `json:"http_status_message"`
	Messages          *SendPhoneVerificationMessages `json:"messages,omitempty"`
	Data              interface{}                    `json:"data"`
}

// Get Captcha
//
// POST {{domain}}/api/captcha/get
//
// # Get Captcha for using on registration
type GetCAPTCHA struct {
	Status            *bool           `json:"status,omitempty"`
	HTTPStatusMessage *string         `json:"http_status_message,omitempty"`
	Data              *GetCAPTCHAData `json:"data,omitempty"`
	Token             *string         `json:"token,omitempty"`
	Img               *string         `json:"img,omitempty"`
}

type GetCAPTCHAData struct {
	Token string `json:"token"`
	Img   string `json:"img"`
}

// Get Brokers List
//
// POST {{domain}}/api/brokers/list
//
// # Brokers list for registration
type GetBrokersList struct {
	Status            bool          `json:"status"`
	HTTPStatusMessage string        `json:"http_status_message"`
	Data              []TypeElement `json:"data"`
}

type TypeElement struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Get Sources List
//
// POST {{domain}}/api/sources/list
//
// # Sources list for registration
type GetSourcesList struct {
	Status            bool                  `json:"status"`
	HTTPStatusMessage string                `json:"http_status_message"`
	Data              []GetSourcesListDatum `json:"data"`
}

type GetSourcesListDatum struct {
	ID     int64  `json:"id"`
	Source string `json:"source"`
}

// Register
//
// POST {{domain}}/api/auth/register
//
// # Register request
//
// To only validate single input, just send the input with `only_validate: true` and that
// specific field will be validated.
//
// #### Available includes
// * user
type Register struct {
	Status            bool              `json:"status"`
	HTTPStatusMessage string            `json:"http_status_message"`
	Data              *RegisterData     `json:"data,omitempty"`
	Messages          *RegisterMessages `json:"messages,omitempty"`
}

type RegisterData struct {
	Token string `json:"token"`
}

type RegisterMessages struct {
	Email []string `json:"email"`
}

// Login
//
// POST {{domain}}/api/auth/login
//
// # Login request
//
// #### Available includes
// * user
type Login struct {
	Status bool       `json:"status"`
	Login  LoginClass `json:"login"`
}

type LoginClass struct {
	Token Token `json:"token"`
	User  User  `json:"user"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresAt   string `json:"expires_at"`
}

type User struct {
	ID         int64         `json:"id"`
	FirstName  string        `json:"first_name"`
	LastName   string        `json:"last_name"`
	Email      string        `json:"email"`
	Phone      string        `json:"phone"`
	RepCode    string        `json:"rep_code"`
	Role       Role          `json:"role"`
	Locale     LocaleElement `json:"locale"`
	Initiator  []interface{} `json:"initiator"`
	Timestamps Timestamps    `json:"timestamps"`
}

type LocaleElement struct {
	ID     int64  `json:"id"`
	Locale Locale `json:"locale"`
	Name   string `json:"name"`
}

type Role struct {
	Name string `json:"name"`
}

type Timestamps struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Login With QR Code Token
//
// POST {{domain}}/api/auth/login-with-qr-code
//
// # Login request
//
// #### Available includes
// * user
type LoginWithQRCodeToken struct {
	Status            bool                     `json:"status"`
	HTTPStatusMessage string                   `json:"http_status_message"`
	Data              LoginWithQRCodeTokenData `json:"data"`
}

type LoginWithQRCodeTokenData struct {
	Token string  `json:"token"`
	Role  *string `json:"role,omitempty"`
}

// Forgot Password
//
// POST {{domain}}/api/auth/forgot
//
// # Forgot password request
type ForgotPassword struct {
	Status            bool        `json:"status"`
	HTTPStatusMessage string      `json:"http_status_message"`
	Data              interface{} `json:"data"`
}

// Reset Password
//
// POST {{domain}}/api/auth/reset
//
// # Reset password request
type ResetPassword struct {
	Status            bool        `json:"status"`
	HTTPStatusMessage string      `json:"http_status_message"`
	Data              interface{} `json:"data"`
}

// Logout
//
// POST {{domain}}/api/auth/logout
//
// # Logout request
type Logout struct {
	Status            bool     `json:"status"`
	HTTPStatusMessage string   `json:"http_status_message"`
	Data              []string `json:"data"`
}

// Get
//
// POST {{domain}}/api/applications/get
//
// # List of applications
//
// #### Available includes
// * current_status
// * type
// * user
// * initiator
// * parent
// * advisor_questionnaire
// * statuses
// * documents
// * answers
// * docusign_envelopes
// * background_checks
// * timestamps
type Get struct {
	Status            bool    `json:"status"`
	HTTPStatusMessage string  `json:"http_status_message"`
	Data              GetData `json:"data"`
}

type GetData struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Pages []Page `json:"pages"`
}

type Page struct {
	ID           int64         `json:"id"`
	Translations []Translation `json:"translations"`
	Sections     []Section     `json:"sections"`
}

type Section struct {
	ID           int64             `json:"id"`
	Translations []Translation     `json:"translations"`
	Questions    []QuestionElement `json:"questions"`
}

type QuestionElement struct {
	ID              int64            `json:"id"`
	IsOptional      int64            `json:"is_optional"`
	Answer          *QuestionAnswer  `json:"answer"`
	Translations    []Translation    `json:"translations"`
	ParentQuestions []ParentQuestion `json:"parent_questions"`
	RuleSet         RuleSet          `json:"rule_set"`
	Type            TypeElement      `json:"type"`
	DocumentType    interface{}      `json:"document_type"`
	Options         []Option         `json:"options"`
	Errors          *QuestionErrors  `json:"errors,omitempty"`
}

type QuestionAnswer struct {
	ID       int64  `json:"id"`
	OptionID int64  `json:"option_id"`
	Data     string `json:"data"`
}

type QuestionErrors struct {
	OptionID []string `json:"option_id,omitempty"`
	Data     []string `json:"data,omitempty"`
}

type Option struct {
	ID            int64          `json:"id"`
	Translations  []Translation  `json:"translations"`
	ParentOptions []ParentOption `json:"parent_options"`
}

type ParentOption struct {
	ID                int64 `json:"id"`
	QuestionID        int64 `json:"question_id"`
	ParentOptionID    int64 `json:"parent_option_id"`
	DependentOptionID int64 `json:"dependent_option_id"`
}

type Translation struct {
	Data     string `json:"data"`
	Language Locale `json:"language"`
}

type ParentQuestion struct {
	ID          int64 `json:"id"`
	ParentID    int64 `json:"parent_id"`
	DependentID int64 `json:"dependent_id"`
	OptionID    int64 `json:"option_id"`
}

type RuleSet struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Rules []Rule `json:"rules"`
}

type Rule struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Attribute string  `json:"attribute"`
	Values    []Value `json:"values"`
}

type Value struct {
	ID    int64  `json:"id"`
	Value string `json:"value"`
}

// Answer
//
// POST {{domain}}/api/applications/answer
//
// # Answer the Question
//
// #### Available includes
// * application
// * question
// * option
// * timestamps
type Answer struct {
	Status bool         `json:"status"`
	Answer AnswerAnswer `json:"answer"`
}

type AnswerAnswer struct {
	ID       int64          `json:"id"`
	Data     string         `json:"data"`
	Question AnswerQuestion `json:"question"`
	Option   []interface{}  `json:"option"`
}

type AnswerQuestion struct {
	ID                 int64         `json:"id"`
	QuestionTypeID     int64         `json:"question_type_id"`
	RuleSetID          int64         `json:"rule_set_id"`
	ApexFormKeyID      interface{}   `json:"apex_form_key_id"`
	DocumentTypeID     interface{}   `json:"document_type_id"`
	HasMultipleAnswers int64         `json:"has_multiple_answers"`
	IsOptional         int64         `json:"is_optional"`
	IsEncrypted        int64         `json:"is_encrypted"`
	SystemIdentifier   string        `json:"system_identifier"`
	Translations       []Translation `json:"translations"`
}

// Answer Multiple
//
// POST {{domain}}/api/applications/answer-multiple
//
// # Answer Multiple Questions
type AnswerMultiple struct {
	Status  bool    `json:"status"`
	Answers Answers `json:"answers"`
}

type Answers struct {
	Success []Success `json:"success"`
	Error   []Error   `json:"error"`
}

type Error struct {
	QuestionID int64       `json:"question_id"`
	Errors     ErrorErrors `json:"errors"`
}

type ErrorErrors struct {
	Data []string `json:"data"`
}

type Success struct {
	QuestionID int64         `json:"question_id"`
	Answer     SuccessAnswer `json:"answer"`
}

type SuccessAnswer struct {
	ID            int64       `json:"id"`
	ApplicationID int64       `json:"application_id"`
	QuestionID    int64       `json:"question_id"`
	OptionID      interface{} `json:"option_id"`
	Data          string      `json:"data"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
}

// Submit
//
// POST {{domain}}/api/applications/submit
//
// # Create Application
//
// #### Available includes
// * current_status
// * type
// * user
// * initiator
// * parent
// * advisor_questionnaire
// * statuses
// * documents
// * answers
// * docusign_envelopes
// * background_checks
// * timestamps
type Submit struct {
	Status            bool           `json:"status"`
	HTTPStatusMessage string         `json:"http_status_message"`
	Messages          *MessagesUnion `json:"messages"`
}

type MessagesMessages struct {
	Questions []QuestionElement `json:"questions"`
}

// Delete
//
// POST {{domain}}/api/applications/delete
type Delete struct {
	Status            bool           `json:"status"`
	HTTPStatusMessage string         `json:"http_status_message"`
	Messages          DeleteMessages `json:"messages"`
}

type DeleteMessages struct {
	ApplicationID []string `json:"application_id"`
}

// Add Note
//
// POST {{domain}}/api/applications/add-note
type AddNote struct {
	Status            bool           `json:"status"`
	HTTPStatusMessage string         `json:"http_status_message"`
	Messages          DeleteMessages `json:"messages"`
}

// Get Application History By ID
//
// POST {{domain}}/api/applications/get-history
type GetApplicationHistoryByID struct {
	Status            bool           `json:"status"`
	HTTPStatusMessage string         `json:"http_status_message"`
	Messages          DeleteMessages `json:"messages"`
}

// Account Exists
//
// POST {{domain}}/api/applications/account-exists
type AccountExists struct {
	Status            bool              `json:"status"`
	HTTPStatusMessage string            `json:"http_status_message"`
	Data              AccountExistsData `json:"data"`
}

type AccountExistsData struct {
	AccountExists bool `json:"account_exists"`
}

// List
//
// POST {{domain}}/api/locales/list
//
// # List of available locales
type List struct {
	Status            bool            `json:"status"`
	HTTPStatusMessage string          `json:"http_status_message"`
	Data              []LocaleElement `json:"data"`
}

type Locale string

const (
	En Locale = "en"
)

type MessagesUnion struct {
	MessagesMessages *MessagesMessages
	StringArray      []string
}
