package model

type SymbolLogosResponse struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Logo             string `json:"logo"`
	LogoOriginal     string `json:"logo_original"`
	LogoNormal       string `json:"logo_normal"`
	LogoThumbnail    string `json:"logo_thumbnail"`
	LogoSquare       string `json:"logo_square"`
	LogoSquareStrict string `json:"logo_square_strict"`
}

type SymbolSocialsResponse struct {
	Symbol             string `json:"symbol"`
	Name               string `json:"name"`
	Twitter            string `json:"twitter"`
	TwitterLink        string `json:"twitter_link"`
	Facebook           string `json:"facebook"`
	FacebookLink       string `json:"facebook_link"`
	Linkedin           string `json:"linkedin"`
	LinkedinLink       string `json:"linkedin_link"`
	Instagram          string `json:"instagram"`
	InstagramLink      string `json:"instagram_link"`
	Youtube            string `json:"youtube"`
	YoutubeLink        string `json:"youtube_link"`
	YoutubeChannel     string `json:"youtube_channel"`
	YoutubeChannelLink string `json:"youtube_channel_link"`
}

type MultipleSymbolLogosRequest struct {
	Symbols []string `json:"symbols"`
}

type Conversion string

const (
	ConversionOriginal             Conversion = "original"
	ConversionDefault              Conversion = "default"
	ConversionNormal               Conversion = "normal"
	ConversionThumbnail            Conversion = "thumbnail"
	ConversionSquare               Conversion = "square"
	ConversionSquareStrict         Conversion = "square_strict"
	ConversionSquareStrictOrNormal Conversion = "square_strict_or_normal"
	ConversionBestMatch            Conversion = "best_match"
)

type SymbolLogoConvertedRequest struct {
	Symbol     string     `json:"symbol"`
	Conversion Conversion `json:"conversion" enums:"original,default,normal,thumbnail,square,square_strict,square_strict_or_normal,best_match"`
}

type MultipleCryptoLogosRequest struct {
	Symbols []string `json:"symbols"`
}
