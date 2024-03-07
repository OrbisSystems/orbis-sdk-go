package passport

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/OrbisSystems/orbis-sdk-go/interfaces/mock"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New(nil, nil))
}

func TestPassport_Articles(t *testing.T) {
	var (
		expResp = []model.Article{
			{
				ArticleID:    1,
				ReleaseStamp: 23,
				Language:     "en",
			},
		}

		rawReq = `{"language":"","type":"","symbol":"AAPL","released_before":0,"released_after":0,"tag":"","tag_group":"","count":0}`

		rawResponse = `[{"article_id": 1, "release_stamp": 23, "language": "en"}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.ArticlesRequest
		hasErr bool
		fn     func(ctx context.Context) *Passport
	}{
		{
			name:   "success",
			input:  model.ArticlesRequest{Symbol: "AAPL"},
			hasErr: false,
			fn: func(ctx context.Context) *Passport {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightPassportArticles, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Passport{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  model.ArticlesRequest{Symbol: "AAPL"},
			hasErr: true,
			fn: func(ctx context.Context) *Passport {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightPassportArticles, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Passport{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  model.ArticlesRequest{Symbol: "AAPL"},
			hasErr: true,
			fn: func(ctx context.Context) *Passport {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightPassportArticles, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Passport{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			n := tc.fn(ctx)
			resp, err := n.Articles(ctx, tc.input)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, expResp, resp)
			}
		})
	}
}

func TestPassport_Newsfeed(t *testing.T) {
	var (
		expResp = []model.Newsfeed{
			{
				ID:         1,
				LanguageID: 2,
				NewsItemID: 3,
			},
		}

		rawReq = `{"language":"","type":"","symbol":"AAPL","released_before":0,"released_after":0,"tag":"","with_tag_group":"","count":0}`

		rawResponse = `[{"id": 1, "language_id": 2, "news_item_id": 3}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.NewsfeedRequest
		hasErr bool
		fn     func(ctx context.Context) *Passport
	}{
		{
			name:   "success",
			input:  model.NewsfeedRequest{Symbol: "AAPL"},
			hasErr: false,
			fn: func(ctx context.Context) *Passport {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightPassportNewsFeed, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Passport{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  model.NewsfeedRequest{Symbol: "AAPL"},
			hasErr: true,
			fn: func(ctx context.Context) *Passport {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightPassportNewsFeed, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Passport{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  model.NewsfeedRequest{Symbol: "AAPL"},
			hasErr: true,
			fn: func(ctx context.Context) *Passport {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightPassportNewsFeed, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Passport{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			n := tc.fn(ctx)
			resp, err := n.Newsfeed(ctx, tc.input)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, expResp, resp)
			}
		})
	}
}

func TestPassport_ArticleByID(t *testing.T) {
	var (
		expResp = model.Article{
			ArticleID:    1,
			ReleaseStamp: 23,
			Language:     "en",
		}

		rawReq = `{"language":"","id":1}`

		rawResponse = `{"article_id": 1, "release_stamp": 23, "language": "en"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.ArticleByIDRequest
		hasErr bool
		fn     func(ctx context.Context) *Passport
	}{
		{
			name:   "success",
			input:  model.ArticleByIDRequest{ID: 1},
			hasErr: false,
			fn: func(ctx context.Context) *Passport {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightPassportArticleByID, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Passport{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  model.ArticleByIDRequest{ID: 1},
			hasErr: true,
			fn: func(ctx context.Context) *Passport {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightPassportArticleByID, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Passport{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  model.ArticleByIDRequest{ID: 1},
			hasErr: true,
			fn: func(ctx context.Context) *Passport {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightPassportArticleByID, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Passport{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			n := tc.fn(ctx)
			resp, err := n.ArticleByID(ctx, tc.input)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, expResp, resp)
			}
		})
	}
}

func TestPassport_SearchArticle(t *testing.T) {
	var (
		expResp = []model.Article{
			{
				ArticleID:    1,
				ReleaseStamp: 23,
				Language:     "en",
			},
		}

		rawReq = `{"language":"","s":"","type":"primary"}`

		rawResponse = `[{"article_id": 1, "release_stamp": 23, "language": "en"}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.SearchArticleRequest
		hasErr bool
		fn     func(ctx context.Context) *Passport
	}{
		{
			name:   "success",
			input:  model.SearchArticleRequest{Type: model.ArticleSearchTypePrimary},
			hasErr: false,
			fn: func(ctx context.Context) *Passport {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightPassportSearchArticle, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Passport{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  model.SearchArticleRequest{Type: model.ArticleSearchTypePrimary},
			hasErr: true,
			fn: func(ctx context.Context) *Passport {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightPassportSearchArticle, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Passport{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  model.SearchArticleRequest{Type: model.ArticleSearchTypePrimary},
			hasErr: true,
			fn: func(ctx context.Context) *Passport {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightPassportSearchArticle, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Passport{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			n := tc.fn(ctx)
			resp, err := n.SearchArticle(ctx, tc.input)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, expResp, resp)
			}
		})
	}
}

func TestPassport_AuthorProfile(t *testing.T) {
	var (
		expResp = []model.AuthorProfileResponse{
			{
				Author: model.AuthorInfo{
					ID:     111,
					Avatar: "Wdqw",
					Name:   "dea",
				},
				Articles: []*model.Article{
					{
						ArticleID:    1,
						ReleaseStamp: 23,
						Language:     "en",
					},
				},
			},
		}

		rawReq = `{"language":"","id":"111"}`

		rawResponse = `[{"author": {"id": 111, "avatar": "Wdqw", "name": "dea"}, "articles": [{"article_id": 1, "release_stamp": 23, "language": "en"}]}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.AuthorProfileRequest
		hasErr bool
		fn     func(ctx context.Context) *Passport
	}{
		{
			name:   "success",
			input:  model.AuthorProfileRequest{ID: "111"},
			hasErr: false,
			fn: func(ctx context.Context) *Passport {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightPassportAuthorProfile, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Passport{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  model.AuthorProfileRequest{ID: "111"},
			hasErr: true,
			fn: func(ctx context.Context) *Passport {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightPassportAuthorProfile, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Passport{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  model.AuthorProfileRequest{ID: "111"},
			hasErr: true,
			fn: func(ctx context.Context) *Passport {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightPassportAuthorProfile, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Passport{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			n := tc.fn(ctx)
			resp, err := n.AuthorProfile(ctx, tc.input)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, expResp, resp)
			}
		})
	}
}

func TestPassport_MostPopularTags(t *testing.T) {
	var (
		expResp = []model.TagShortInfo{
			{
				Symbol:   "AAPL",
				Articles: 10,
			},
			{
				Symbol:   "FB",
				Articles: 23,
			},
		}

		rawReq = `{"language":"","time":"","type":"primary"}`

		rawResponse = `[{"symbol": "AAPL", "articles": 10}, {"symbol": "FB", "articles": 23}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.MostPopularTagsRequest
		hasErr bool
		fn     func(ctx context.Context) *Passport
	}{
		{
			name:   "success",
			input:  model.MostPopularTagsRequest{Type: model.MostPopularTagsTypesTypePrimary},
			hasErr: false,
			fn: func(ctx context.Context) *Passport {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightPassportMostPopularTags, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Passport{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  model.MostPopularTagsRequest{Type: model.MostPopularTagsTypesTypePrimary},
			hasErr: true,
			fn: func(ctx context.Context) *Passport {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightPassportMostPopularTags, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Passport{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  model.MostPopularTagsRequest{Type: model.MostPopularTagsTypesTypePrimary},
			hasErr: true,
			fn: func(ctx context.Context) *Passport {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightPassportMostPopularTags, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Passport{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			n := tc.fn(ctx)
			resp, err := n.MostPopularTags(ctx, tc.input)

			if tc.hasErr {
				assert.Error(t, err)
				assert.Empty(t, resp)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, expResp, resp)
			}
		})
	}
}
