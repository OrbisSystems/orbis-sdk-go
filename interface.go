package orbis_sdk_go

import (
	"context"
	"io"
	"net/http"

	"github.com/OrbisSystems/orbis-sdk-go/model"
)

type Storage interface {
	Store(ctx context.Context, data []byte) error
	Get(ctx context.Context) ([]byte, error)
}

type HTTPClient interface {
	Get(ctx context.Context, url string, headers http.Header) (*http.Response, error)
	Post(ctx context.Context, url string, body io.Reader, headers http.Header) (*http.Response, error)

	Wrapper
}

type Wrapper interface {
	GetBodyAndCheckOK(r *http.Response) (io.ReadCloser, error)
	UnmarshalAndCheckOk(v interface{}, r *http.Response) error
}

type Auth interface {
	SetToken(ctx context.Context, token model.Token) error
	GetToken(ctx context.Context) (model.Token, error)
}

type AccountService interface {
	LoginByEmail(ctx context.Context, req model.LoginByEmailRequest) error
	LoginByAPIKey(ctx context.Context, req model.LoginByAPIKeyRequest) error
	CreateAPIKey(ctx context.Context, req model.CreateAPIKeyRequest) (model.CreateAPIKeyResponse, error)
}

type NewsService interface {
	GetByFilter(ctx context.Context, req model.NewsFilterRequest) ([]model.NewsResponse, error)
	GetByID(ctx context.Context, req model.NewsByIDRequest) (model.NewsResponse, error)
	GetSymbolSubjects(ctx context.Context, req model.SymbolSubjectsRequest) ([]model.SymbolSubjectsResponse, error)
	GetAvailableTaxonomy(ctx context.Context) ([]model.TaxonomyCode, error)
	GetAvailableSubjects(ctx context.Context) ([]string, error)
	GetAvailableSymbols(ctx context.Context) ([]string, error)
	GetAvailableSources(ctx context.Context) ([]string, error)
	GetAvailableLanguages(ctx context.Context) ([]string, error)
}

type LogosService interface {
	SymbolLogos(ctx context.Context, symbol string) (model.SymbolLogosResponse, error)
	SocialSymbolLogos(ctx context.Context, symbol string) (model.SymbolSocialsResponse, error)
	DirectSymbolLogo(ctx context.Context, symbol string) (io.ReadCloser, error)
	CryptoSymbolLogo(ctx context.Context, symbol string) (model.SymbolLogosResponse, error)
	DirectCryptoSymbolLogo(ctx context.Context, symbol string) (io.ReadCloser, error)
	MultiSymbolLogos(ctx context.Context, req model.MultipleSymbolLogosRequest) ([]model.SymbolLogosResponse, error)
	ConvertedSymbolLogo(ctx context.Context, req model.SymbolLogoConvertedRequest) (io.ReadCloser, error)
	MultipleCryptoSymbolLogo(ctx context.Context, req model.MultipleCryptoLogosRequest) ([]model.SymbolLogosResponse, error)
	ConvertedCryptoSymbolLogo(ctx context.Context, req model.SymbolLogoConvertedRequest) (io.ReadCloser, error)
}

type PassportService interface {
	Articles(ctx context.Context, req model.ArticlesRequest) ([]model.Article, error)
	Newsfeed(ctx context.Context, req model.NewsfeedRequest) ([]model.Newsfeed, error)
	ArticleByID(ctx context.Context, req model.ArticleByIDRequest) (model.Article, error)
	SearchArticle(ctx context.Context, req model.SearchArticleRequest) ([]model.Article, error)
	AuthorProfile(ctx context.Context, req model.AuthorProfileRequest) ([]model.AuthorProfileResponse, error)
	MostPopularTags(ctx context.Context, req model.MostPopularTagsRequest) ([]model.TagShortInfo, error)
}

type TipRankService interface {
	AnalystConsensus(ctx context.Context, req model.AnalystConsensusRequest) ([]model.AnalystConsensusResponse, error)
	LatestAnalystRatingsOnStock(ctx context.Context, req model.LatestAnalystRatingsOnStockRequest) ([]model.LatestAnalystRatingsOnStockResponse, error)
	LiveFeed(ctx context.Context, req model.LiveFeedRequest) ([]model.LiveFeedResponse, error)
	TrendingStocks(ctx context.Context, req model.TrendingStocksRequest) ([]model.TrendingStocksResponse, error)
	AnalystPortfolios(ctx context.Context, req model.PortfoliosRequest) ([]model.PortfoliosResponse, error)
	AnalystProfile(ctx context.Context, id string) (model.AnalystProfileResponse, error)
	SectorConsensus(ctx context.Context) (model.SectorConsensusResponse, error)
	BestPerformingExperts(ctx context.Context, num int) ([]model.BestPerformingExpertsResponse, error)
	StocksSimilarStocks(ctx context.Context, ticker string) ([]string, error)
	AnalystsExpertPictureStore(ctx context.Context) (model.AnalystsExpertPictureStoreResponse, error)
	SupportedTickers(ctx context.Context) (model.SupportedTickersResponse, error)
	GeneralStockUpdates(ctx context.Context, utcTime, details string) (model.GeneralStockUpdatesResponse, error)
	InsidersOverview(ctx context.Context, expertUID string) (model.InsidersOverviewResponse, error)
	InsidersBestPerformingExperts(ctx context.Context, num int) ([]model.InsidersBestPerformingExpertsResponse, error)
	InsidersLiveFeed(ctx context.Context, num int, sort string) ([]model.InsidersLiveFeedResponse, error)
	HedgeFundsBestPerformingExperts(ctx context.Context, num int) ([]model.HedgeFundsBestPerformingExpertsResponse, error)
}

type EquityService interface {
	GetQuotesEquityData(ctx context.Context, symbols, quoteType string) ([]model.EquityDataResponse, error)
}

type FundsService interface {
	GetFundDetails(ctx context.Context, symbol string) (model.GetFundDetailsResponse, error)
	GetFundScreenerFilters(ctx context.Context) (model.GetFundScreenerFiltersResponse, error)
	ScreenFunds(ctx context.Context, req model.FundScreenerRequest) (model.FundScreenerResponse, error)
	ScreenSectorFunds(ctx context.Context, req model.FundSectorScreenerRequest) (model.FundScreenerResponse, error)
	GetTopFunds(ctx context.Context, req model.GetTopFundsRequest) ([]string, error)
}
