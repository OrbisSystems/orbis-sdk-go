package passport

type MostPopularTagsKeywords []MostPopularTagsKeyword

type Article struct {
	ArticleID          int64         `json:"article_id"`
	Tags               []string      `json:"tags"`
	Symbols            []string      `json:"symbols"`
	Author             Author        `json:"author"`
	ReleaseDate        string        `json:"release_date"`
	ReleaseStamp       int64         `json:"release_stamp"`
	UpdatedDate        string        `json:"updated_date"`
	Language           string        `json:"language"`
	AvailableLanguages []string      `json:"available_languages"`
	Title              string        `json:"title"`
	Permalink          string        `json:"permalink"`
	Summary            string        `json:"summary"`
	LongSummary        string        `json:"long_summary"`
	FirstParagraph     string        `json:"first_paragraph"`
	Body               []Body        `json:"body"`
	Types              []string      `json:"types"`
	FeaturedImage      FeaturedImage `json:"featured_image"`
}

type Author struct {
	ID        int64     `json:"id"`
	Avatar    string    `json:"avatar"`
	Name      string    `json:"name"`
	Biography string    `json:"biography"`
	Articles  []Article `json:"articles"`
}

type Body struct {
	Type string `json:"type"`
	Body string `json:"body"`
}

type FeaturedImage struct {
	Type          string    `json:"type"`
	ID            int64     `json:"id"`
	Copyright     string    `json:"copyright"`
	Description   string    `json:"description"`
	MainImage     MainImage `json:"main_image"`
	RectangleCrop MainImage `json:"rectangle_crop"`
	SquareCrop    MainImage `json:"square_crop"`
	OriginalImage MainImage `json:"original_image"`
}

type MainImage struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type MostPopularTagsKeyword struct {
	Symbol   string `json:"symbol"`
	Articles string `json:"articles"`
}

// JWT Authenticate
//
// POST jwt/authenticate?email&password
//
// Get a JWT token using credentials provided by Orbis.
type JWTAuthenticate struct {
	Token string `json:"token"`
}
