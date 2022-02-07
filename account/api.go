package account

import (
	"os"
	"time"

	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/http"
	"github.com/OrbisSystems/orbis-sdk-go/models/account"
	"github.com/OrbisSystems/orbis-sdk-go/storage"
)

type API interface {
	Login
	ApiKeys
	PhoneVerification
	Captcha
	User
	Application
	Documents
	Countries
	States
	Devices
	Settings
	Branches
	Dashboard
	Locales
	StockLending
	WebHook
	SuperAdmin
}

type Login interface {
	Login(req account.LoginRequest) (account.Login, error)
	Logout() error
	Refresh(errChan chan error)
}

type ApiKeys interface {
	List(with []string) (interface{}, error)
	Create(with []string) (interface{}, error)
	Delete(ID string, with []string) (interface{}, error)
}

type PhoneVerification interface {
	SendPhoneVerification(typeVerification string) (interface{}, error)
	VerifyPhone(code string) (interface{}, error)
}

type Captcha interface {
	GetCaptcha() (interface{}, error)
}

type User interface {
	GetUserList(userReq account.GetUsers) (interface{}, error)
	GetUser(ID int64, with []string) (interface{}, error)
	CreateUser(createUser account.CreateUser) (interface{}, error)
	UpdateUser(createUser account.CreateUser) (interface{}, error)
	EnableBranchToUser(userID string, branchID string) (interface{}, error)
}

type Application interface {
	GetAvailableAppStatuses() (interface{}, error)
	GetApplicationsList(req account.ListApplicationsReq) (interface{}, error)
	GetApplication(applicationID int64, with []string) (interface{}, error)
	GetApplicationTypes(with []string) (interface{}, error)
	GetApplicationForm(applicationTypeID int64, with []string) (interface{}, error)
	CreateApplication(userID int64, with []string) (interface{}, error)
	Answer(applicationID, questionID string, data os.File) (interface{}, error)
	AnswerMultiply(req account.AnswerMultiple) (interface{}, error)
	DeleteAnswer(applicationID, answerID int64) (interface{}, error)
	Submit(applicationID int64) (interface{}, error)
	EnableResubmit(int64 int64) (interface{}, error)
	BackGroundCheck(applicationID int64, with []string) (interface{}, error)
	ComplianceStatusUpdate(applicationID int64, status string, with []string) (interface{}, error)
	DeleteApplication(applicationID int64) (interface{}, error)
	NoteApplication(applicationID int64, note string, with []string) (interface{}, error)
	GetHistoryOfApplication(applicationHistoryID int64, with []string) (interface{}, error)
	AccountExist(IDNumber string) (interface{}, error)
	ApplicationVerifyDocument(applicationID int64, applicantType string, with []string) (interface{}, error)
	Custodian
}

type Custodian interface {
	Investigation
	SendToCustodian(applicationID string, with []string) (interface{}, error)
}

type Investigation interface {
	UploadDocument(applicationID, name string, data os.File) (interface{}, error)
	Appeal(investigationID int64, action, notes string) (interface{}, error)
}

type Documents interface {
	GenerateQRCode(applicationID int64) (interface{}, error)
	GetDocumentQuestions(applicationID int64, with []string) (interface{}, error)
	GetDocumentAnswers(applicationID int64, with []string) (interface{}, error)
	UpdateDocumentStatus(req account.UpdateDocumentReq) (interface{}, error)
	GetDocument(applicationID, documentID int64) (interface{}, error)
}

type Countries interface {
	GetListCountries(with []string) (interface{}, error)
}

type States interface {
	GetListStates() (interface{}, error)
}

type Devices interface {
	RegisterDevice(name, token string) (interface{}, error)
}

type Settings interface {
	GetListSettings() (interface{}, error)
	GetSetting(key string) (interface{}, error)
	SetSetting(key, value string) (interface{}, error)
}

type Branches interface {
	GetListBranches(with []string) (interface{}, error)
	GetBranch(ID string) (interface{}, error)
}

type Dashboard interface {
	GetDashboardData(interval string, with []string) (interface{}, error)
}

type Locales interface {
	GetListLocales() (interface{}, error)
	ChangeLocale(locale account.Locale) (interface{}, error)
}

type WebHook interface {
	CustodianStatusUpdate(accountNumber, status string) (interface{}, error)
	GenerateTwiMLForVoice(message string) (interface{}, error)
	GlobalID
	CustodianExchange
}

type GlobalID interface {
	UpdateBackgroundCheckService(req account.UpdateServicesReq) (interface{}, error)
}

type CustodianExchange interface {
	UpdateWatchmenCompleteStatus(req account.WatchmenCompleteReq) (interface{}, error)
	PushALEs(req account.PushALEsReq) (interface{}, error)
	PushSketches(req account.PushSketchesReq) (interface{}, error)
}

type StockLending interface {
	GetPositions() (interface{}, error)
	GetAccountPortfolio() (interface{}, error)
}

type SuperAdmin interface {
	SuperAdminApplicationType
	SuperAdminApplicationType
	FirmsSuperAdmin
	SuperAdminBranches
	SuperAdminUsers
	SuperAdminPages
	SuperAdminSections
	SuperAdminQuestions
	SuperAdminOptions
	SuperAdminServices
	SuperAdminServiceTypes
	SuperAdminQuestionTypes
	SuperAdminRulesAdmin
	SuperAdminDocumentTypes
	SuperAdminApexForm
	SuperAdminCustodianExchange
}

type SuperAdminApplicationType interface {
	GetApplicationTypeList(req account.ApplicationTypeListReq) (interface{}, error)
	GetApplicationTypeByID(ID int64, with []string) (interface{}, error)
	CreateApplicationType(req account.CreateApplicationTypeReq) (interface{}, error)
	UpdateApplicationType(req account.UpdateApplicationTypeReq) (interface{}, error)
	CopyApplicationType(req account.CopyApplicationTypeReq) (interface{}, error)
	DeleteApplicationType(int64 int64) (interface{}, error)
	PushApplicationToEnv(req account.PushToEnvReq) (interface{}, error)
	ApplyChanges(req account.ApplicationTypeApplyChanges) (interface{}, error)

	Version
}

type FirmsSuperAdmin interface {
	GetListFirms(with []string) (interface{}, error)
	GetFirm(ID int64, with []string) (interface{}, error)
	CreateFirm(name string) (interface{}, error)
	UpdateFirm(ID int64, name string) (interface{}, error)
	DeleteFirm(ID int64) (interface{}, error)
}

type SuperAdminBranches interface {
	GetSuperAdminListBranches(firmID int64, with []string) (interface{}, error)
	GetSuperAdminBranch(ID int64) (interface{}, error)
	CreateSuperAdminBranch(req account.CreateBranchSuperAdminReq) (interface{}, error)
	UpdateSuperAdminBranch(req account.UpdateBranchSuperAdminReq) (interface{}, error)
	DeleteSuperAdminBranch(ID int64) (interface{}, error)
}

type SuperAdminUsers interface {
	AddToFirm(userID, firmID int64, with []string) (interface{}, error)
	RemoveFromFirm(userID, firmID int64, with []string) (interface{}, error)
	SetDefault(userID, firmID int64, with []string) (interface{}, error)
}

type SuperAdminPages interface {
	GetListPages(with []string) (interface{}, error)
	CreatePage(req account.CreateApplicationTypePageReq) (interface{}, error)
	UpdatePage(req account.UpdateApplicationTypePageReq) (interface{}, error)
	CopyPage(req account.CopyPageSuperAdmin) (interface{}, error)
	DeletePage(ID int64) (interface{}, error)
	UpdatePagePosition(req account.UpdatePositionsSuperAdminReq) (interface{}, error)
}

type SuperAdminSections interface {
	GetListSections(with []string) (interface{}, error)
	CreateSection(req account.CreateSuperAdminSectionReq) (interface{}, error)
	UpdateSection(req account.UpdateSuperAdminSectionReq) (interface{}, error)
	CopySection(req account.CopySuperAdminSectionReq) (interface{}, error)
	DeleteSection(ID int64) (interface{}, error)
	UpdateSectionPosition(req account.UpdateSuperAdminSectionPositionReq) (interface{}, error)
}

type SuperAdminQuestions interface {
	GetListQuestions(req account.GetListSuperAdminQuestions) (interface{}, error)
	GetWhereHasOptionsQuestion(applicationTypeID int64, with []string) (interface{}, error)
	CreateQuestion(req account.CreateSuperAdminQuestion) (interface{}, error)
	UpdateQuestion(req account.UpdateSuperAdminQuestionReq) (interface{}, error)
	CopyQuestion(req account.CopySuperAdminQuestion) (interface{}, error)
	DeleteQuestion(ID int64) (interface{}, error)
	UpdateQuestionPosition(req account.UpdateSuperAdminQuestionPosition) (interface{}, error)
}

type SuperAdminOptions interface {
	GetOptionsList(with []string) (interface{}, error)
}

type SuperAdminServices interface {
	GetServicesList(with []string) (interface{}, error)
}

type SuperAdminServiceTypes interface {
	GetServiceTypesList(with []string) (interface{}, error)
}

type SuperAdminQuestionTypes interface {
	GetServiceQuestionsList(with []string) (interface{}, error)
}

type SuperAdminRulesAdmin interface {
	GetRules(with []string) (interface{}, error)
}

type SuperAdminDocumentTypes interface {
	GetDocumentTypesList(with []string) (interface{}, error)
}

type SuperAdminApexForm interface {
	GetApexFormList(with []string) (interface{}, error)
}

type SuperAdminCustodianExchange interface {
	GetCustodianExchangeFormList(with []string) (interface{}, error)
}

type Version interface {
	SaveNew(req account.VersionReq) (interface{}, error)
	RevertToVersion(req account.RevertVersionReq) (interface{}, error)
}

type Client struct {
	auth.API
	client http.Client

	config        config.OrbisConfig
	tickerRefresh *time.Timer
}

// NewClient returns new API with defaults storage.Storage and http.OrbisClient
func NewClient(config config.OrbisConfig) API {
	return NewConfiguredClient(storage.NewInMemory(), http.DefaultClient, config, http.DefaultAuthKey)
}

// NewConfiguredClient returns new API
func NewConfiguredClient(storage storage.Storage,
	client http.Client, config config.OrbisConfig, key string) API {
	return &Client{
		API:           auth.NewAuth(storage, key),
		client:        client,
		config:        config,
		tickerRefresh: time.NewTimer(30 * time.Minute),
	}
}
