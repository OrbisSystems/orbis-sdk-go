// nolint:gosec
package model

// --------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------
// 						Account	URL
// --------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------

const (
	URLB2BLoginByEmail  = "/uma/api/v1/auth/b2b/login/basic"
	URLB2BLoginByAPIKey = "/uma/api/v1/auth/b2b/login/api-keys"
	URLB2BCreateAPIKey  = "/uma/api/v1/b2b/api-keys"
	URLB2BRefreshToken  = "/uma/api/v1/auth/b2b/refresh"
	URLB2BGetUserByID   = "/uma/api/v1/users/b2b"
)

// --------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------
// 						Insight	URL
// --------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------

const (
	URLInsightBase = "/insight-api/api/v1"

	WSInsightNews = "/insight-ws/ws/news"

	// *******************************************
	// 		NEWS
	// *******************************************

	URLInsightNewsFilter   = "/news/filter"
	URLInsightNewsByID     = "/news/id"
	URLInsightNewsSymbols  = "/news/symbols"
	URLInsightNewsAuthors  = "/news/authors"
	URLInsightNewsChannels = "/news/channels"
	URLInsightNewsTags     = "/news/tags"

	// *******************************************
	// 		LOGOS
	// *******************************************

	URLInsightLogosSymbolLogos               = "/logos/symbol-logos"
	URLInsightLogosSocialSymbolLogos         = "/logos/social-symbol-logos"
	URLInsightLogosDirectSymbolLogos         = "/logos/direct-symbol-logo"
	URLInsightLogosCryptoSymbolLogo          = "/logos/crypto-symbol-logo"
	URLInsightLogosDirectCryptoSymbolLogos   = "/direct-crypto-symbol-logo"
	URLInsightLogosMultiSymbolLogos          = "/multiple-symbol-logos"
	URLInsightLogosConvertedSymbolLogos      = "/converted-symbol-logo"
	URLInsightLogosMultipleCryptoSymbolLogo  = "/multiple-crypto-symbol-logo"
	URLInsightLogosConvertedCryptoSymbolLogo = "/converted-crypto-symbol-logo"

	// *******************************************
	// 		PASSPORT
	// *******************************************

	URLInsightPassportArticles        = "/passport/articles"
	URLInsightPassportNewsFeed        = "/passport/newsfeed"
	URLInsightPassportArticleByID     = "/passport/article-by-id"
	URLInsightPassportSearchArticle   = "/passport/search-article"
	URLInsightPassportAuthorProfile   = "/passport/author-profile"
	URLInsightPassportMostPopularTags = "/passport/most-popular-tags"

	// *******************************************
	// 		TipRank
	// *******************************************

	URLInsightTipRankAnalystConsensus                = "/tipranks/analyst-consensus"
	URLInsightTipRankAnalystMulti                    = "/tipranks/analyst-multi"
	URLInsightTipRankLiveFeed                        = "/tipranks/live-feed"
	URLInsightTipRankTrendingStocks                  = "/tipranks/trending-stocks"
	URLInsightTipRankAnalystPortfolio                = "/tipranks/analyst-portfolio"
	URLInsightTipRankAnalystProfile                  = "/tipranks/analyst-profile"
	URLInsightTipRankSectorConsensus                 = "/tipranks/sector-consensus"
	URLInsightTipRankBestPerformingExperts           = "/tipranks/best-performing-experts"
	URLInsightTipRankStocksSimilarStocks             = "/tipranks/stocks-similar-stocks"
	URLInsightTipRankAnalystsExpertPictureStore      = "/tipranks/analysts-expert-picture-store"
	URLInsightTipRankSupportedTickers                = "/tipranks/supported-tickers"
	URLInsightTipRankGeneralStockUpdates             = "/tipranks/general-stock-updates"
	URLInsightTipRankInsidersOverview                = "/tipranks/insiders-overview"
	URLInsightTipRankInsidersBestPerformingExperts   = "/tipranks/insiders-best-performing-experts"
	URLInsightTipRankInsidersLiveFeed                = "/tipranks/insiders-live-feed"
	URLInsightTipRankHedgefundsBestPerformingExperts = "/tipranks/hedgefunds-best-performing-experts"

	// *******************************************
	// 		Quotes
	// *******************************************

	URLInsightQuotesEquity   = "/quotes/equity"
	URLInsightQuoteHistory   = "/quote-history"
	URLInsightIntradayQuotes = "/intraday/quotes"

	// *******************************************
	// 		Funds
	// *******************************************

	URLInsightFundsDetails         = "/funds/details"
	URLInsightFundsScreenerFilters = "/funds/screener/filters"
	URLInsightFundsScreener        = "/funds/screener"
	URLInsightFundsSectorScreener  = "/funds/sector-screener"
	URLInsightFundsTop             = "/funds/top"

	// *******************************************
	// 		Research
	// *******************************************

	URLInsightCompanyProfile        = "/company-profile"
	URLInsightQuoteProfile          = "/quote-profile"
	URLInsightSymbolOwnerships      = "/symbol-ownerships"
	URLInsightOwnerships            = "/ownerships"
	URLInsightEarningsCalendar      = "/earnings-calendar"
	URLInsightFundamentals          = "/fundamentals"
	URLInsightStockScreener         = "/stock-screener"
	URLInsightHeatmaps              = "/heatmaps"
	URLInsightIndustriesPerformance = "/industries-performance"
	URLInsightGetMomentumRatioGraph = "/rst"
	URLInsightSeasonality           = "/seasonality"

	// *******************************************
	// 		IPO
	// *******************************************

	URLInsightIPOsUpcoming = "/ipos/upcoming"
	URLInsightIPOsRecent   = "/ipos/recent"

	// *******************************************
	// 		World Market
	// *******************************************

	URLInsightContinents    = "/continents"
	URLInsightRegions       = "/regions"
	URLInsightCountryCodes  = "/country-codes"
	URLInsightGlobalIndexes = "/global-indexes"

	// *******************************************
	// 		Market Dates
	// *******************************************

	URLInsightMarketDatesHistory = "/market-dates/history "
	URLInsightMarketDatesToday   = "/market-dates/today "

	// *******************************************
	// 		Option Greeks
	// *******************************************

	URLInsightOptionGreeksCalculateParams = "/og/calculate-params "
	URLInsightOptionGreeksCalculateMatrix = "/og/calculate-matrix "

	// *******************************************
	// 		HS
	// *******************************************

	URLInsightHSDailySummary  = "/hs/api/v1/resources/finance/daily"
	URLInsightHSWeeklySummary = "/hs/api/v1/resources/finance/weekly"
)
