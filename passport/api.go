package passport

import (
	"github.com/OrbisSystems/orbis-sdk-go/auth"
	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/http"
	"github.com/OrbisSystems/orbis-sdk-go/models/passport"
	"github.com/OrbisSystems/orbis-sdk-go/storage"
)

type API interface {
	Article
	Authenticate
	SearchArticle
	Newsfeed
	Author
}

type Authenticate interface {
	Auth(email, password string) (passport.JWTAuthenticate, error)
}

type Article interface {
	GetArticles(req passport.ArticleRequest) ([]passport.Article, error)
	GetArticlesBySecret(req passport.ArticleWithSecretRequest) ([]passport.Article, error)
	GetArticleByID(ID, lang string) (passport.Article, error)
	GetArticleByIDSecret(ID, lang string, secret passport.Secret) (passport.Article, error)

	GetMostPopular(typeArticle passport.TypeArticle, time, lang string) (passport.MostPopularTagsKeywords, error)
	GetMostPopularBySecret(typeArticle passport.TypeArticle, time, lang string, secret passport.Secret) (passport.MostPopularTagsKeywords, error)
}

type Author interface {
	GetAuthorProfile(lang, ID string) (passport.Author, error)
	GetAuthorProfileBySecret(lang, ID string, secret passport.Secret) (passport.Author, error)
}

type Newsfeed interface {
	Newsfeed(lang, search string) ([]passport.Article, error)
}

type SearchArticle interface {
	SearchArticle(lang, s string, searchType passport.TypeArticle) ([]passport.Article, error)
	SearchArticleSecretBySecret(lang, s string, searchType passport.TypeArticle, secret passport.Secret) ([]passport.Article, error)
}

type Client struct {
	auth.API
	client http.Client

	config config.OrbisConfig
}

// NewConfiguredClient returns new API
func NewConfiguredClient(storage storage.Storage,
	client http.Client, config config.OrbisConfig, authKey string) API {
	return &Client{
		API:    auth.NewAuth(storage, authKey),
		client: client,
		config: config,
	}
}
