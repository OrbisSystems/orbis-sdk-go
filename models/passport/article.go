package passport

type TypeArticle string

var (
	SymbolType  TypeArticle = "symbol"
	PrimaryType TypeArticle = "primary"
	TagsType    TypeArticle = "tags"
)

type ArticleRequestType string

var (
	Educational    ArticleRequestType = "educational"
	NonEducational ArticleRequestType = "non-educational"
)

type ArticleRequest struct {
	Type           ArticleRequestType
	WithSymbol     string
	ReleasedBefore string
	ReleasedAfter  string
	WithTag        string
}
type ArticleWithSecretRequest struct {
	ArticleRequest
	Secret
}

type Secret struct {
	APIKey    string
	APISecret string
}
