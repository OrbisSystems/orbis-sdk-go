package news

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/OrbisSystems/orbis-sdk-go/interfaces/mock"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New(nil, nil))
}

func TestNews_GetByFilter(t *testing.T) {
	var (
		id, _ = uuid.Parse("6edba5f5-4a5e-4ec4-8b73-4c98f1341203")
		tt    = time.Date(2023, 10, 10, 9, 9, 1, 0, time.UTC)

		expResp = model.NewsResponse{
			News: []model.News{
				{
					ID:             id,
					SourceID:       2,
					Author:         "Ad",
					Created:        tt,
					Updated:        tt,
					Title:          "T1",
					Teaser:         "De",
					Body:           "dewre",
					Url:            "deee",
					IsPressRelease: false,
				},
			},
			Count: 1,
		}

		rawReq = `{"symbol":null,"author":null,"start_date":null,"end_date":null,"channels":null,"tags":null,"press_release":null,"headline":null,"paging":{"limit":1}}`

		rawResponse = `{"news":[{"id":"6edba5f5-4a5e-4ec4-8b73-4c98f1341203","source_id":2,"author":"Ad","created":"2023-10-10T09:09:01.000Z","updated":"2023-10-10T09:09:01.000Z","title":"T1","teaser":"De","body":"dewre","url":"deee","is_press_release":false,"image":null,"channels":null,"symbols":null,"tags":null}],"count":1}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.NewsFilterRequest
		hasErr bool
		fn     func(ctx context.Context) *News
	}{
		{
			name:   "success",
			input:  model.NewsFilterRequest{Paging: model.Paging{Limit: 1}},
			hasErr: false,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightNewsFilter, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &News{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  model.NewsFilterRequest{Paging: model.Paging{Limit: 1}},
			hasErr: true,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightNewsFilter, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &News{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  model.NewsFilterRequest{Paging: model.Paging{Limit: 1}},
			hasErr: true,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightNewsFilter, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &News{
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
			resp, err := n.GetByFilter(ctx, tc.input)

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

func TestNews_GetByID(t *testing.T) {
	var (
		id, _ = uuid.Parse("6edba5f5-4a5e-4ec4-8b73-4c98f1341203")
		tt    = time.Date(2023, 10, 10, 9, 9, 1, 0, time.UTC)

		expResp = model.News{
			ID:             id,
			SourceID:       2,
			Author:         "Ad",
			Created:        tt,
			Updated:        tt,
			Title:          "T1",
			Teaser:         "De",
			Body:           "dewre",
			Url:            "deee",
			IsPressRelease: false,
		}

		rawReq = `{"id":"6edba5f5-4a5e-4ec4-8b73-4c98f1341203"}`

		rawResponse = `{"id":"6edba5f5-4a5e-4ec4-8b73-4c98f1341203","source_id":2,"author":"Ad","created":"2023-10-10T09:09:01.000Z","updated":"2023-10-10T09:09:01.000Z","title":"T1","teaser":"De","body":"dewre","url":"deee","is_press_release":false,"image":null,"channels":null,"symbols":null,"tags":null}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.NewsRequest
		hasErr bool
		fn     func(ctx context.Context) *News
	}{
		{
			name:   "success",
			input:  model.NewsRequest{ID: id.String()},
			hasErr: false,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightNewsByID, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &News{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  model.NewsRequest{ID: id.String()},
			hasErr: true,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightNewsByID, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &News{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  model.NewsRequest{ID: id.String()},
			hasErr: true,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightNewsByID, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &News{
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
			resp, err := n.GetByID(ctx, tc.input)

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

func TestNews_GetAvailableSymbols(t *testing.T) {
	var (
		expResp = []string{"AAPL", "FB"}

		rawResponse = `["AAPL", "FB"]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		hasErr bool
		fn     func(ctx context.Context) *News
	}{
		{
			name:   "success",
			hasErr: false,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, model.URLInsightBase+model.URLInsightNewsSymbols, nil).Return(httpResponse, nil)

				return &News{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			hasErr: true,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, model.URLInsightBase+model.URLInsightNewsSymbols, nil).Return(httpResponse, nil)

				return &News{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			hasErr: true,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, model.URLInsightBase+model.URLInsightNewsSymbols, nil).Return(nil, testErr)

				return &News{
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
			resp, err := n.GetAvailableSymbols(ctx)

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

func TestNews_GetAvailableAuthors(t *testing.T) {
	var (
		symbol = "AAPL"

		expResp = []string{"Aut1", "Aut2"}

		rawResponse = `["Aut1", "Aut2"]`

		rawReq = `{"symbol":"AAPL"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  *string
		hasErr bool
		fn     func(ctx context.Context) *News
	}{
		{
			name:   "success/with_symbol",
			input:  &symbol,
			hasErr: false,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightNewsAuthors, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &News{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "success/no_symbol",
			input:  nil,
			hasErr: false,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightNewsAuthors, nil, nil).Return(httpResponse, nil)

				return &News{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  &symbol,
			hasErr: true,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightNewsAuthors, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &News{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  &symbol,
			hasErr: true,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightNewsAuthors, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &News{
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
			resp, err := n.GetAvailableAuthors(ctx, tc.input)

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

func TestNews_GetAvailableChannels(t *testing.T) {
	var (
		symbol = "AAPL"

		expResp = []string{"Ch1", "Ch2"}

		rawResponse = `["Ch1", "Ch2"]`

		rawReq = `{"symbol":"AAPL"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  *string
		hasErr bool
		fn     func(ctx context.Context) *News
	}{
		{
			name:   "success/with_symbol",
			input:  &symbol,
			hasErr: false,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightNewsChannels, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &News{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "success/no_symbol",
			input:  nil,
			hasErr: false,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightNewsChannels, nil, nil).Return(httpResponse, nil)

				return &News{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  &symbol,
			hasErr: true,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightNewsChannels, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &News{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  &symbol,
			hasErr: true,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightNewsChannels, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &News{
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
			resp, err := n.GetAvailableChannels(ctx, tc.input)

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

func TestNews_GetAvailableTags(t *testing.T) {
	var (
		symbol = "AAPL"

		expResp = []string{"Tg1", "Tg2"}

		rawResponse = `["Tg1", "Tg2"]`

		rawReq = `{"symbol":"AAPL"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  *string
		hasErr bool
		fn     func(ctx context.Context) *News
	}{
		{
			name:   "success/with_symbol",
			input:  &symbol,
			hasErr: false,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightNewsTags, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &News{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "success/no_symbol",
			input:  nil,
			hasErr: false,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightNewsTags, nil, nil).Return(httpResponse, nil)

				return &News{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  &symbol,
			hasErr: true,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightNewsTags, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &News{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  &symbol,
			hasErr: true,
			fn: func(ctx context.Context) *News {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightNewsTags, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &News{
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
			resp, err := n.GetAvailableTags(ctx, tc.input)

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
