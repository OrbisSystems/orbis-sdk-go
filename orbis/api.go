package orbis

import (
	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/http"
	models "github.com/OrbisSystems/orbis-sdk-go/models/orbis-api"
	"github.com/OrbisSystems/orbis-sdk-go/storage"
	"io"
	"time"
)

type API interface {
	Order
	Quote
	Research
	User
	Advisory
	Device
	Transfer
	Document
	Miscellaneou
	BranchOrbis
	LegalDocument
}

// Order

type Order interface {
	Equity
	Option
	Bond
	Crypto
	MutualFund

	Listing(listingType string, loadQuotes bool, accountID, account, symbol *string) (io.ReadCloser, error)
	ListingForModel(listingType, listingModel string, loadQuotes bool) (io.ReadCloser, error)
	ListingForBranch(listingType string) (io.ReadCloser, error)
	Status(orderRef string, loadQuotes bool, accountID, account *string) (io.ReadCloser, error)
	CancelRequest(req models.CancelReplaceReq) (io.ReadCloser, error)
	Cancellation(req models.CancellationReq) (io.ReadCloser, error)
}

type Equity interface {
	PreviewEquities(req models.EquitiesRequest) (io.ReadCloser, error)
	PlacementEquities(req models.EquitiesRequest) (io.ReadCloser, error)
	BasketPreviewEquities(req models.EquitiesRequest) (io.ReadCloser, error)
	BasketPlacementEquities(req models.EquitiesRequest) (io.ReadCloser, error)
	AdvisorPlacementEquities(req models.AdvisorRequest) (io.ReadCloser, error)
	AdvisorPreviewEquities(req models.AdvisorRequest) (io.ReadCloser, error)
}

type Option interface {
	PreviewOptions(req models.OptionsPreviewReq) (io.ReadCloser, error)
	PlacementOptions(req models.OptionsPreviewReq) (io.ReadCloser, error)
	MultilegStrategyOptions(req models.MultilegStrategyReq) (io.ReadCloser, error)
	PreviewMultilegOptions(req models.PreviewMultilegReq) (io.ReadCloser, error)
	PlacementMultilegOptions(req models.MultilegStrategyReq) (io.ReadCloser, error)
	ReplacePreviewOptions(req models.ReplacePreviewReq) (io.ReadCloser, error)
	ReplaceMultilegOptions(req models.ReplacePreviewReq) (io.ReadCloser, error)
}

type Bond interface {
	PlacementBonds(req models.BondPlacementReq) (io.ReadCloser, error)
	QuoteRequestBonds(req models.QuoteBondPlacementReq) (io.ReadCloser, error)
	QuoteAcceptanceBonds(req models.BondQuoteAcceptanceReq) (io.ReadCloser, error)
}

type Crypto interface {
	OrderPreviewCrypto(req models.OrderCryptoReq) (io.ReadCloser, error)
	OrderPlaceCrypto(req models.OrderCryptoReq) (io.ReadCloser, error)
}

type MutualFund interface {
	GetMutualFundsList() ([]models.QuoteACAT, error)
	GetMutualFundDetails(symbol string) (models.MutualFundDetails, error)
	CheckMutualFund(symbol string) (io.ReadCloser, error)
	MutualFundsPlacement(req models.MutualFundsRequest) (io.ReadCloser, error)
	MutualFundPreview(req models.MutualFundsRequest) (io.ReadCloser, error)
}

// Quote

type Quote interface {
	GetEquity(symbol string) (models.EquityElement, error)
	SearchQuote(criteria string, limit int, loadQuotes bool) (models.TopStockQuote, error)
	GetOptions(symbols []string, greeks bool) (io.ReadCloser, error)
	GetOptionsGreeks(symbols []string, greeks bool) (io.ReadCloser, error)
	GetOptionsPricingType(underlying, symbol *string, expectedPx *float64) (io.ReadCloser, error)
	GetBonds(cusip []string) (io.ReadCloser, error)
	GetHistorical(req models.HistoricalRequest) (models.Historical, error)
	GetIntradayChartData(symbol *string, rsiSpa *float64, from, to *int64, rowVolumes *bool) (models.IntradayChartData, error)
	GetEquityTicks(symbol *string, startTime, endTime *time.Time) (io.ReadCloser, error)
	GetShortability(symbol string) (models.GetShortability, error)
	GetTopStocksQuote(category, market, order *string, count *int64) (models.TopStocks, error)
	GetOptionChainDates(symbol string, symbols []string) (io.ReadCloser, error)
	GetOptionChain(symbol, date string) (io.ReadCloser, error)
	GetOptionChainGreek(symbol, date string, symbols []string) (io.ReadCloser, error)
}

// Research

type Research interface {
	TipRank
	ETF
	ADRS
	CorporateAction
	News
	AlphaTool
	IPO

	GetCompanyResearch(symbol string) (models.CompanyResearch, error)
	// ??
	GetOptionRatios(symbols []string) (io.ReadCloser, error)
	GetOptionRadar(filter string) (models.OptionRadar, error)
	GetIndexComponents(index string, loadQuotes *bool) (models.IndexComponents, error)
	GetStockUpgrades(symbols []string) (io.ReadCloser, error)
	GetOwnershipDetails(symbols []string, filter, sort string, count, page int64) (models.OwnershipDetails, error)
	// ????
	GetOwnerDetails(ids []string) (io.ReadCloser, error)
	GetEarningCalendar(excludeQuotes, loadQuotes *bool, days *int64, symbols []string) (models.EarningsCalendar, error)
	GetEarningCalendarForPortfolio(excludeQuotes, loadQuotes *bool, days *int64) (models.EarningsCalendarForPortfolio, error)
	GetFundamentals(typeFund, symbol string) (models.Fundamentals, error)
	GetFundamentalTypes() ([]string, error)
	Screener(req models.ScreenerResearch) (models.Screener, error)
	ScreenerV2(req models.ScreenerResearch) (models.Screener, error)
	GetIndustryHeatMap(industryType string) (models.IndustryHeatMap, error)
	GetIndustryList(industrySector string) (models.IndustryList, error)
	GetSymbolHeatMap(amountLimit *float64, numSymbols, industryCode *int64, exchangeFilter string) (models.SymbolHeatMap, error)
	GetIndexOverview() (io.ReadCloser, error)
	GetHistoricalData(symbol string, dates []string) (models.HistoricalData, error)
	GetTopQuotesByCategory(category, market, order string) (io.ReadCloser, error)
}

type TipRank interface {
	GetLiveFeed(count int64) (models.LiveFeedResponse, error)
	GetAnalystPortfolio(expertUUID string, count int64, loadQuotes bool) (models.GetAnalystResponse, error)
	GetAnalystProfile(expertUUID string) (models.AnalystProfileResponse, error)
	GetTopAnalyst(count int64) (models.GetTopAnalystResponse, error)
	GetRatingsBySymbol(symbol string, count int64) (models.RatingBySymbolResponse, error)
	GetInsider(symbols []string) (models.InsiderResponse, error)
	GetInsiderTransactions(symbols []string, count int64) (models.InsiderTransactionsResponse, error)
	GetSimilarStocks(symbol string) ([]string, error)
	Overview(symbols []string) (models.OverviewResponse, error)
	GetTopStocksTipRanks(symbols []string) (models.GetTopStocksResponse, error)
	GetConsensusBySector() (models.ConsensusBySector, error)
	GetConsensusByBlogger(symbol string) (models.ConsensusByBloggerResponse, error)
	GetTrending(count, period int64) (models.TrendingResponse, error)
	GetHedgeFunds(symbol string) (models.HedgeFundsResponse, error)
	GetStockSentiment(symbol string) (models.StockSentimentResponse, error)
	GetBloggerRatings(symbol string) (models.BloggerRatingsResponse, error)
}

type ETF interface {
	GetInitScreener() (models.InitScreener, error)
	GetScreener(req models.ScreenerRequest) (models.ScrenerResponse, error)
	GetTopTenETF(inverseType, leverageType *string) (models.Top10, error)
	GetDetails(symbol *string, loadQuotes *bool) (models.Details, error)
	GetBreakDownBySector(duration *string, displaySize, exposure *int64) (models.BreakdownBySector, error)
}

type ADRS interface {
	GetTopTenADRS(country *string, loadQuotes, loadEarningReleases, loadUpgradesDowngrades bool) (models.Top10, error)
	GetDefaults() ([]string, error)
	SearchAdrS(country *string, loadQuotes, loadEarningReleases, loadUpgradesDowngrades bool) (models.AddrSearch, error)
}

type CorporateAction interface {
	GetTypes() ([]string, error)
	GetPaymentTypes() ([]string, error)
	SearchCorporateActions(symbol *string, dateFrom, dateTo *time.Time, types []string, usePayDate *bool) (
		[]models.CorporatesActionsSearch, error)
}

type News interface {
	GetList(start, count int64, filter string) (io.ReadCloser, error)
	GetNewsBySymbol(start, count int64, symbol string) (io.ReadCloser, error)
	GetSEC(start, count int64, symbol string) (io.ReadCloser, error)
	LookupNews(newsID string) (io.ReadCloser, error)
}

type AlphaTool interface {
	GetAlphaTracker(benchmark, rangE *string, symbols []string, period *int64, setEMA *bool) ([]models.AlphaTracker, error)
	GetSeasonality(benchmark *string, symbols []string, years int64) ([]models.AlphaTracker, error)
	GetWeaklySeasonality(benchmark *string, symbols []string, years int64) ([]models.AlphaTracker, error)
}

type IPO interface {
	GetRecent(loadQuotes *bool, time, sort *string) (models.GetIPOS, error)
	GetUpcoming() (models.GetIPOS, error)
	GetPerformance(loadQuotes *bool, time, sort *string) (models.PerformanceIPOS, error)
}

type LegalDocument interface {
	GetDocumentClassifications() (io.ReadCloser, error)
	GetDocumentListing(classification *string) (io.ReadCloser, error)
	GetPolicyValues(policyType string) (io.ReadCloser, error)
	GetDocumentContent(ID *string) (io.ReadCloser, error)
	UpdatePolicy(policyType string, req models.UpdatePolicyRequest) (io.ReadCloser, error)
}

type BranchOrbis interface {
	GetRtd() (io.ReadCloser, error)
	GetRtds() (io.ReadCloser, error)
	GetInventory() (io.ReadCloser, error)
	GetRtdsTotal() (io.ReadCloser, error)
	GetBranchPortfolio(req models.PortfolioRequest) (io.ReadCloser, error)
	GetOptionExpirations() (io.ReadCloser, error)
	GetOptionExpirationDetails(req models.OptionExpirationsDetailsRequest) (io.ReadCloser, error)
}

type Miscellaneou interface {
	GetMarketDate() (io.ReadCloser, error)
	GetDatesMarket() (io.ReadCloser, error)
	GetLastOpen(market string) (io.ReadCloser, error)
	GetDatesCheck(market *string, date *time.Time) (io.ReadCloser, error)
	GetAdvisoryInfo() (io.ReadCloser, error)
	UpdateAdvisoryInfo(req models.AdvisoryInfoUpdateRequest) (io.ReadCloser, error)
	AvatarUpload(req models.AvatarUploadRequest) (io.ReadCloser, error)
	Upload(req models.UploadRequest) (io.ReadCloser, error)
	UserAvatarUpload(ID string, req models.AvatarUploadRequest) (io.ReadCloser, error)
	GetAvatar() (io.ReadCloser, error)
	GetUserAvatar(ID string) (io.ReadCloser, error)
	FilesDownload(tag, ID string) (io.ReadCloser, error)
	GetFilesList(tag, fileName string) (io.ReadCloser, error)
	FilesAllowedExts() (io.ReadCloser, error)
	GetErrorCodeListing() (io.ReadCloser, error)
	Ping() (io.ReadCloser, error)
}

type Document interface {
	ConfirmDocument(account string, date *time.Time) (models.ConfirmDocumentResponse, error)
	ConfirmListDocuments(account string, beginDate, endDate *time.Time) (models.ListOfConfirms, error)
	GetServiceDocs(req models.ServiceDocsRequest) (models.ServiceDocsResponse, error)
	GetAvailableDocumentTypes() ([]string, error)
	GetAvailableDocumentTypesAndCodes() (models.AvailableDocumentTypesAndCodes, error)
	Download(account, key string) (io.ReadCloser, error)
}

type Device interface {
	DeactivateDevice(req models.DeactivateDeviceBody) (models.DeactivateDeviceResponse, error)
}

// Ask about WS
type Streaming interface {
	NotificationLink(typeLink string, req models.NotificationLinkRequest) (io.ReadCloser, error)
	NotificationUnLink(typeLink string, req models.NotificationUnLinkRequest) (io.ReadCloser, error)
	GetNotificationLinks(typeLink string) (io.ReadCloser, error)
}

// Ask about WS
type StreamingQuotes interface {
}

type Transfer interface {
	Bank
	Position
	Journal
	Wire
	ACAT
	ACH
	AccountVerification
	Relationship

	GetTransferDetails(account, transferID string) (io.ReadCloser, error)
	GetTransferStatus(account string, req models.TransferStatusRequest) (models.TransferStatus, error)
	GetTransferStatusModel(model int64) (models.TransferStatus, error)
	GetTransferStatusBranch() (models.TransferStatus, error)
	GetAvailableAmount(account string) (models.AmountAvailable, error)
	GetLegalDocumentTypes() (io.ReadCloser, error)
	GetLegalDocument(ID string) (io.ReadCloser, error)
	CancelTransfer(req models.CancelTransferRequest) (io.ReadCloser, error)
}

type Bank interface {
	FundTransfer(direction string, req models.FundTransferRequest) (io.ReadCloser, error)
}

type Position interface {
	TransferRequest(req models.TransferRequest) (io.ReadCloser, error)
}

type Journal interface {
	JournalRequest(req models.JournalRequest) (io.ReadCloser, error)
	JournalSearch(req models.JournalSearchRequest) (io.ReadCloser, error)
}

type Wire interface {
	Outgoing(req models.WireRequest) (io.ReadCloser, error)
}

type ACAT interface {
	SearchACAT(req models.SearchRequest) (io.ReadCloser, error)
	CancelACAT(req models.CancelACATRequest) (io.ReadCloser, error)
	InitiateACAT(req models.InitiateACATRequest) (io.ReadCloser, error)
	BulkSchedule(req models.InitiateACATScheduleRequest) (io.ReadCloser, error)
}

type ACH interface {
	DepositACH(req models.ACHDepositRequest) (models.Deposit, error)
	DepositIRA(req models.DepositIRARequest) (models.Deposit, error)
	Withdraw(req models.WithdrawRequest) (models.Withdraw, error)
}

type AccountVerification interface {
	InitiateAccountVerification(subject string) (models.InitiatePost, error)
	CompleteAccount(req models.PlaidRequest) (models.Complete, error)
	GetListBankAccounts(req models.PlaidRequest) (models.ListBankAccounts, error)
	GetListBankAccount(req models.PlaidRequest) (models.ListBankAccount, error)
	Import(req models.ImportInstantRequest) (io.ReadCloser, error)

	Plaid
}

type Plaid interface {
	InitiatePost(newValidation bool) (models.InitiatePost, error)
	CompletePlaid(req models.PlaidRequest) (models.Complete, error)
	GetListBankAccountsPlaid(req models.PlaidRequest) (models.ListBankAccounts, error)
	GetListBankAccountPlaid(req models.PlaidRequest) (models.ListBankAccount, error)
}

type Relationship interface {
	GetListRelationship(account string) (io.ReadCloser, error)
	GetPaymentMethods(country string) (io.ReadCloser, error)
	GetPaymentMethodsViaAddress(req models.BankCodesViaAddressRequest) (io.ReadCloser, error)
	GetBankCodes(country string) (io.ReadCloser, error)
	GetBankCodesViaAddress(req models.BankCodesViaAddressRequest) (io.ReadCloser, error)
	Create(req models.CreateTransferRequest) (models.Create, error)
	GetWireInformation(req models.RelationshipInformationRequest) (io.ReadCloser, error)
	BankingRelationship(req models.RelationshipInformationRequest) (io.ReadCloser, error)
	Approve(req models.ApproveTransferRequest) (models.Approve, error)
	Cancel(entryID string, req models.CancelRelationshipTransferRequest) (models.Cancel, error)
	Rename(entryID, nickname string) (models.Rename, error)
}

type Advisory interface {
	Model
	Charge
	UserAccount
	Group
	Analytics
	Allocation
}

type Charge interface {
	SearchCharge(req models.SearchChargesRequest) (io.ReadCloser, error)
	ApplyCharge(req models.ApplyChargeRequest) (io.ReadCloser, error)
	CancelCharge(chargeRef string) (io.ReadCloser, error)
}

type Model interface {
	GetModelListing(loadQuotes, loadAccounts bool) (io.ReadCloser, error)
	GetComponentTags() (io.ReadCloser, error)
	GetBuyingPower(currency, modelID string) (io.ReadCloser, error)
	GetRealtimeBalance(modelID string) (io.ReadCloser, error)
	GetRealtimeBalanceHistory(modelID string, dateFrom, dateTo *time.Time, page, count int64) (io.ReadCloser, error)
	GetAdjustments(modelID, status string, loadQuotes bool) (io.ReadCloser, error)
	GetAdjustmentPreview(adjustmentID string) (io.ReadCloser, error)
	GetMemberAccounts(modelID string, loadRtb, loadRtbHistory bool) (io.ReadCloser, error)
	OrphanedAccounts() (io.ReadCloser, error)
	GetMemberAccountStatistic(modelID string) (io.ReadCloser, error)
	GetMemberRealtimeBalances(modelID string) (models.MemberRealtimeBalances, error)
	GetPortfolio(modelID string, req models.ModelPortfolioRequest) (models.Portfolio, error)
	GetAllModelBuyingPower(fCurrency, sCurrencyQuery string) (io.ReadCloser, error)
	GetPositionLookup(modelID, symbol string) (io.ReadCloser, error)
	GetPositionLookupOptions(modelID, symbol string) (io.ReadCloser, error)
	UpdateAdjustment(action string, req models.UpdateAdjustmentRequest) (io.ReadCloser, error)
	CreateModel(req models.CreateModelRequest) (io.ReadCloser, error)
	UpdateModel(req models.UpdateModelRequest) (io.ReadCloser, error)
	DeleteModel(ID int64) (io.ReadCloser, error)
	UpdateComponent(req models.UpdateComponentRequest) (io.ReadCloser, error)
	LinkOfMembersAccounts(req models.LinkOfMembersRequest) (io.ReadCloser, error)
	UnlinkMemberAccount(req models.LinkOfMembersRequest) (io.ReadCloser, error)
	Allocations(req models.AllocationModelRequest) (io.ReadCloser, error)
}

type UserAccount interface {
	GetUsersList() (io.ReadCloser, error)
	GetUserAccounts(ID int64, loadRtb bool) (io.ReadCloser, error)
	GetAccountStats() (io.ReadCloser, error)
	GetAccountNotes(accountID int64) (io.ReadCloser, error)
	GetNotes(userID int64) (io.ReadCloser, error)
	UserNotes(action string, req models.UserNotesRequest) (io.ReadCloser, error)
	AccountNotes(action string, req models.AccountNotesRequest) (io.ReadCloser, error)
}

type Group interface {
	ManageGroup(action string, req models.ManageGroupRequest) (io.ReadCloser, error)
	GetGroupsList() (io.ReadCloser, error)
	GetGroupAccount(groupsID string) (io.ReadCloser, error)
	GetGroupPortfolio(req models.PortfolioGroupRequest) (models.GroupPortfoliosGroupID, error)
	GetGroupRtb(groupID string) (io.ReadCloser, error)
	GetGroupRtbs(groupID string) (io.ReadCloser, error)
}

type Analytics interface {
	GetOverexposureAnalysis(overexposureType, scope, ID string, threshold string) (io.ReadCloser, error)
	GetRealtimeBalanceBreakDown(modelID string, bucketSize int64) (io.ReadCloser, error)
	GetModelPerformance(modelID string, performanceRange int64) (io.ReadCloser, error)
	GetModelDriftReport(modelID, threshold string) (io.ReadCloser, error)
}

type Allocation interface {
	SearchAllocation() (io.ReadCloser, error)
	GetDetailsAllocation(allocationRef string) (io.ReadCloser, error)
	CancelAllocation(allocationRef string) (io.ReadCloser, error)
	ValidateAllocation(allocationRef string) (io.ReadCloser, error)
	ScheduleAllocation(req models.ScheduleAllocationRequest) (io.ReadCloser, error)
	TriggerAllocation(allocationRef string) (io.ReadCloser, error)
}

type User interface {
	GetUserInformation(loadRtb bool, accountName, accountNumber *string) (io.ReadCloser, error)
	GetAccountInformation(accountNumber *string, accountID *int64) (io.ReadCloser, error)
	GetAccountPortfolio(req models.AccountPortfolioRequest) (io.ReadCloser, error)
	GetFullAccountPortfolio(req models.FullAccountPortfolioRequest) (io.ReadCloser, error)
	GetAccountTradingLimit(accountNumber, currency string, accountID int64) (models.AccountTradingLimit, error)
	// What kind of account should we use
	GetAccountBalance(accountNumber, currency string) (io.ReadCloser, error)
	GetAccountBalanceHistory(req models.AccountHistoryRequest) (io.ReadCloser, error)
	GetAccountPortfolioHistory(req models.AccountHistoryRequest) (io.ReadCloser, error)
	GetMappedAccountPortfolioHistory(req models.AccountHistoryRequest) (io.ReadCloser, error)
	GetAccountHistory(req models.AccountHistoryRequest) (models.AccountHistory, error)
	GetPosition(symbol, accountNumber *string, accountID int64, option bool) (io.ReadCloser, error)
	PositionSearch(symbols []string, loadQuotes bool) (io.ReadCloser, error)
	GetUserWatchlist(loadQuotes bool) (models.UserWatchlist, error)
	GetPortfolioEarnings(account models.UserAccount) (models.UserPortfolioEarnings, error)
	GetUserPreferences() (models.UserPreferences, error)
	// What kind of body
	UserPreferences(action string, req map[string]string) (io.ReadCloser, error)
	CreateUserWatchlist(req models.CreateWatchlistRequest) (io.ReadCloser, error)
	RenameUserWatchlist(req models.RenameUserWatchlistRequest) (io.ReadCloser, error)
	DeleteUserWatchlist(req models.DeleteUserWatchlistRequest) (io.ReadCloser, error)
	AddEntryUserWatchlist(req models.EntryUserWatchlistRequest) (io.ReadCloser, error)
	RemoveEntryUserWatchlist(symbol string, req models.EntryUserWatchlistRequest) (io.ReadCloser, error)
	FirebaseUserRegister(token string) (io.ReadCloser, error)

	Watchlist
	UserDevice
	UserAgreement
	TradingLimitManagement
}

type Watchlist interface {
	CreateWatchlist(req models.CreateWatchlistRequest) (io.ReadCloser, error)
	RenameWatchlist(req models.RenameWatchlistRequest) (io.ReadCloser, error)
	DeleteWatchlist(req models.DeleteWatchlistRequest) (io.ReadCloser, error)
	AddEntryToWatchlist(req models.EntryToWatchlistRequest) (io.ReadCloser, error)
	DeleteEntryFromWatchlist(req models.EntryToWatchlistRequest) (io.ReadCloser, error)
}

type UserDevice interface {
	RegisterUserDevice() (io.ReadCloser, error)
	DeRegisterUserDevice(token string) (io.ReadCloser, error)
	FirebaseRegister(req models.FirebaseRegisterRequest) (io.ReadCloser, error)
}

type UserAgreement interface {
	GetAvailableAgreements() (io.ReadCloser, error)
	GetAgreement(version, code string) (io.ReadCloser, error)
	UnsignedAgreements() (io.ReadCloser, error)
	AttestStatus(req models.AttestStatusRequest) (io.ReadCloser, error)
	SignAgreement(req models.SignAgreementRequest) (io.ReadCloser, error)
	GetCryptoAgreement() (io.ReadCloser, error)
	SignCryptoAgreement(req models.SignCryptoAgreementRequest) (io.ReadCloser, error)
}

type TradingLimitManagement interface {
	UpdateTradingLimit(req models.TradingLimitManagementRequest) (io.ReadCloser, error)
}

type Client struct {
	auth.API
	client http.Client

	config config.OrbisConfig
}

// NewClient returns new API with defaults storage.Storage and http.OrbisClient
func NewClient(config config.OrbisConfig) API {
	return NewConfiguredClient(storage.NewInMemory(), http.DefaultClient, config, http.DefaultAuthKey)
}

// NewConfiguredClient returns new API
func NewConfiguredClient(storage storage.Storage,
	client http.Client, config config.OrbisConfig, key string) API {
	return &Client{
		API:    auth.NewAuth(storage, key),
		client: client,
		config: config,
	}
}
