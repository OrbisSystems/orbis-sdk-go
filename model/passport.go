package model

type ArticlesRequest struct {
	Language       string              `json:"language"` //
	Type           ArticlesRequestType `json:"type" enums:"education,non-education"`
	Symbol         string              `json:"symbol"`
	ReleasedBefore int64               `json:"released_before"`
	ReleasedAfter  int64               `json:"released_after"`
	Tag            string              `json:"tag"`
	TagGroup       string              `json:"tag_group"`
	Count          uint64              `json:"count"`
}

type ArticlesRequestType string

const (
	ArticlesRequestTypeEducation    ArticlesRequestType = "education"
	ArticlesRequestTypeNonEducation ArticlesRequestType = "non-education"
)

type Article struct {
	ArticleID          int                  `json:"article_id"`
	Tags               []string             `json:"tags"`
	Symbols            []string             `json:"symbols"`
	Author             ArticleAuthor        `json:"author"`
	ReleaseDate        string               `json:"release_date"`
	ReleaseStamp       int                  `json:"release_stamp"`
	UpdatedDate        string               `json:"updated_date"`
	Language           string               `json:"language"`
	AvailableLanguages []string             `json:"available_languages"`
	Title              string               `json:"title"`
	Permalink          string               `json:"permalink"`
	URL                string               `json:"url"`
	Summary            string               `json:"summary"`
	LongSummary        string               `json:"long_summary"`
	FirstParagraph     string               `json:"first_paragraph"`
	Body               []ArticleBody        `json:"body"`
	BodyRaw            interface{}          `json:"body_raw"`
	Types              []string             `json:"types"`
	FeaturedImage      ArticleFeaturedImage `json:"featured_image"`
}

type ArticleAuthor struct {
	ID        int    `json:"id"`
	Avatar    string `json:"avatar"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
}

type ArticleBody struct {
	Type string `json:"type"`
	Body string `json:"body"`
}

type ArticleFeaturedImageInfo struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type ArticleFeaturedImage struct {
	Type          string                   `json:"type"`
	ID            int                      `json:"id"`
	Copyright     string                   `json:"copyright"`
	Description   string                   `json:"description"`
	MainImage     ArticleFeaturedImageInfo `json:"main_image"`
	RectangleCrop ArticleFeaturedImageInfo `json:"rectangle_crop"`
	SquareCrop    ArticleFeaturedImageInfo `json:"square_crop"`
	OriginalImage ArticleFeaturedImageInfo `json:"original_image"`
}

type NewsfeedRequest struct {
	Language       string              `json:"language"` //
	Type           ArticlesRequestType `json:"type" enums:"education,non-education"`
	Symbol         string              `json:"symbol"`
	ReleasedBefore int64               `json:"released_before"`
	ReleasedAfter  int64               `json:"released_after"`
	Tag            string              `json:"tag"`
	TagGroup       string              `json:"with_tag_group"`
	Count          uint64              `json:"count"`
}

type Newsfeed struct {
	ID          int         `json:"id"`
	LanguageID  int         `json:"language_id"`
	NewsItemID  int         `json:"news_item_id"`
	UserID      int         `json:"user_id"`
	Permalink   interface{} `json:"permalink"`
	Title       string      `json:"title"`
	Summary     string      `json:"summary"`
	CreatedAt   string      `json:"created_at"`
	UpdatedAt   string      `json:"updated_at"`
	ReleaseDate string      `json:"release_date"`
	Symbols     []string    `json:"symbols"`
	Types       []string    `json:"types,omitempty"`
	Tags        []string    `json:"tags,omitempty"`
}

type ArticleByIDRequest struct {
	Language string `json:"language"` //
	ID       int    `json:"id"`       //
}

type SearchArticleRequest struct {
	Language string            `json:"language"` //
	S        string            `json:"s"`        //
	Type     ArticleSearchType `json:"type" enums:"all,tag,primary,symbol"`
}

type ArticleSearchType string

const (
	ArticleSearchTypeAll     ArticleSearchType = "all"
	ArticleSearchTypeTag     ArticleSearchType = "tag"
	ArticleSearchTypePrimary ArticleSearchType = "primary"
	ArticleSearchTypeSymbol  ArticleSearchType = "symbol"
)

type AuthorProfileRequest struct {
	Language string `json:"language"` //
	ID       string `json:"id"`       //
}

type AuthorProfileResponse struct {
	Author   AuthorInfo `json:"author"`
	Articles []*Article `json:"articles"`
}

type AuthorInfo struct {
	ID        int    `json:"id"`
	Avatar    string `json:"avatar"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
}

type MostPopularTagsRequest struct {
	Language string                   `json:"language"` //
	Time     MostPopularTagsTimeType  `json:"time" enums:"week,month,year"`
	Type     MostPopularTagsTypesType `json:"type" enums:"symbol,primary,tags"` //
}

type MostPopularTagsTimeType string

const (
	MostPopularTagsTimeTypeWeek  MostPopularTagsTimeType = "week"
	MostPopularTagsTimeTypeMonth MostPopularTagsTimeType = "month"
	MostPopularTagsTimeTypeYear  MostPopularTagsTimeType = "year"
)

type MostPopularTagsTypesType string

const (
	MostPopularTagsTypesTypeSymbol  MostPopularTagsTypesType = "symbol"
	MostPopularTagsTypesTypePrimary MostPopularTagsTypesType = "primary"
	MostPopularTagsTypesTypeTags    MostPopularTagsTypesType = "tags"
)

type TagShortInfo struct {
	Symbol   string `json:"symbol"`
	Articles int    `json:"articles"`
}
