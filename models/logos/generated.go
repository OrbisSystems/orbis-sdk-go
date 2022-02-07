package logos

// Symbol Logos
//
// GET {{base_url}}/api/symbol/{{symbol}}
//
// Get list of all logos for a specific symbol
type SymbolLogos struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Logo             string `json:"logo"`
	LogoOriginal     string `json:"logo_original"`
	LogoNormal       string `json:"logo_normal"`
	LogoThumbnail    string `json:"logo_thumbnail"`
	LogoSquare       string `json:"logo_square"`
	LogoSquareStrict string `json:"logo_square_strict"`
}
