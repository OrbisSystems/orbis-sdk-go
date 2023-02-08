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
)

// --------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------
// 						Insight	URL
// --------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------

const (
	URLInsightBase = "/insight-api/api/v1"

	// *******************************************
	// 		NEWS
	// *******************************************

	URLInsightNewsFilter         = "/news/filter"
	URLInsightNewsByID           = "/news/id"
	URLInsightNewsSymbolSubjects = "/news/symbol_subjects"
	URLInsightNewsTaxonomy       = "/news/taxonomy"
	URLInsightNewsRelevance      = "/news/relevance"
	URLInsightNewsSymbols        = "/news/symbols"
	URLInsightNewsSources        = "/news/sources"
	URLInsightNewsLang           = "/news/lang"

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
)
