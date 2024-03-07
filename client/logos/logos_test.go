package logos

import (
	"bytes"
	"context"
	"errors"
	"fmt"
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

func TestLogos_SymbolLogos(t *testing.T) {
	var (
		expResp = model.SymbolLogosResponse{
			Symbol:           "1",
			Name:             "2",
			Logo:             "3",
			LogoOriginal:     "4",
			LogoNormal:       "5",
			LogoThumbnail:    "6",
			LogoSquare:       "7",
			LogoSquareStrict: "8",
		}

		rawResponse = `{"symbol":"1","name":"2","logo":"3","logo_original":"4","logo_normal":"5","logo_thumbnail":"6","logo_square":"7","logo_square_strict":"8"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  string
		hasErr bool
		fn     func(ctx context.Context, symbol string) *Logos
	}{
		{
			name:   "success",
			input:  "AAPL",
			hasErr: false,
			fn: func(ctx context.Context, symbol string) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightLogosSymbolLogos, symbol), nil).Return(httpResponse, nil)

				return &Logos{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  "AAPL",
			hasErr: true,
			fn: func(ctx context.Context, symbol string) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightLogosSymbolLogos, symbol), nil).Return(httpResponse, nil)

				return &Logos{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  "AAPL",
			hasErr: true,
			fn: func(ctx context.Context, symbol string) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightLogosSymbolLogos, symbol), nil).Return(nil, testErr)

				return &Logos{
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
			f := tc.fn(ctx, tc.input)
			resp, err := f.SymbolLogos(ctx, tc.input)

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

func TestLogos_SocialSymbolLogos(t *testing.T) {
	var (
		expResp = model.SymbolSocialsResponse{
			Symbol:             "1",
			Name:               "2",
			Twitter:            "3",
			TwitterLink:        "4",
			Facebook:           "5",
			FacebookLink:       "6",
			Linkedin:           "7",
			LinkedinLink:       "8",
			Instagram:          "9",
			InstagramLink:      "0",
			Youtube:            "12",
			YoutubeLink:        "13",
			YoutubeChannel:     "14",
			YoutubeChannelLink: "43",
		}

		rawResponse = `{"symbol":"1","name":"2","twitter":"3","twitter_link":"4","facebook":"5","facebook_link":"6","linkedin":"7","linkedin_link":"8","instagram":"9","instagram_link":"0","youtube":"12","youtube_link":"13","youtube_channel":"14","youtube_channel_link":"43"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  string
		hasErr bool
		fn     func(ctx context.Context, symbol string) *Logos
	}{
		{
			name:   "success",
			input:  "AAPL",
			hasErr: false,
			fn: func(ctx context.Context, symbol string) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightLogosSocialSymbolLogos, symbol), nil).Return(httpResponse, nil)

				return &Logos{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  "AAPL",
			hasErr: true,
			fn: func(ctx context.Context, symbol string) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightLogosSocialSymbolLogos, symbol), nil).Return(httpResponse, nil)

				return &Logos{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  "AAPL",
			hasErr: true,
			fn: func(ctx context.Context, symbol string) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightLogosSocialSymbolLogos, symbol), nil).Return(nil, testErr)

				return &Logos{
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
			f := tc.fn(ctx, tc.input)
			resp, err := f.SocialSymbolLogos(ctx, tc.input)

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

func TestLogos_DirectSymbolLogo(t *testing.T) {
	var (
		expResp = io.NopCloser(strings.NewReader("a,b,c,d"))

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  string
		hasErr bool
		fn     func(ctx context.Context, symbol string) *Logos
	}{
		{
			name:   "success",
			input:  "AAPL",
			hasErr: false,
			fn: func(ctx context.Context, symbol string) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("a,b,c,d"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightLogosDirectSymbolLogos, symbol), nil).Return(httpResponse, nil)

				return &Logos{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  "AAPL",
			hasErr: true,
			fn: func(ctx context.Context, symbol string) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightLogosDirectSymbolLogos, symbol), nil).Return(nil, testErr)

				return &Logos{
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
			f := tc.fn(ctx, tc.input)
			resp, err := f.DirectSymbolLogo(ctx, tc.input)

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

func TestLogos_CryptoSymbolLogo(t *testing.T) {
	var (
		expResp = model.SymbolLogosResponse{
			Symbol:           "1",
			Name:             "2",
			Logo:             "3",
			LogoOriginal:     "4",
			LogoNormal:       "5",
			LogoThumbnail:    "6",
			LogoSquare:       "7",
			LogoSquareStrict: "8",
		}

		rawResponse = `{"symbol":"1","name":"2","logo":"3","logo_original":"4","logo_normal":"5","logo_thumbnail":"6","logo_square":"7","logo_square_strict":"8"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  string
		hasErr bool
		fn     func(ctx context.Context, symbol string) *Logos
	}{
		{
			name:   "success",
			input:  "AAPL",
			hasErr: false,
			fn: func(ctx context.Context, symbol string) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightLogosCryptoSymbolLogo, symbol), nil).Return(httpResponse, nil)

				return &Logos{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  "AAPL",
			hasErr: true,
			fn: func(ctx context.Context, symbol string) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightLogosCryptoSymbolLogo, symbol), nil).Return(httpResponse, nil)

				return &Logos{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  "AAPL",
			hasErr: true,
			fn: func(ctx context.Context, symbol string) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightLogosCryptoSymbolLogo, symbol), nil).Return(nil, testErr)

				return &Logos{
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
			f := tc.fn(ctx, tc.input)
			resp, err := f.CryptoSymbolLogo(ctx, tc.input)

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

func TestLogos_DirectCryptoSymbolLogo(t *testing.T) {
	var (
		expResp = io.NopCloser(strings.NewReader("a,b,c,d"))

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  string
		hasErr bool
		fn     func(ctx context.Context, symbol string) *Logos
	}{
		{
			name:   "success",
			input:  "AAPL",
			hasErr: false,
			fn: func(ctx context.Context, symbol string) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("a,b,c,d"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightLogosDirectCryptoSymbolLogos, symbol), nil).Return(httpResponse, nil)

				return &Logos{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  "AAPL",
			hasErr: true,
			fn: func(ctx context.Context, symbol string) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				cli.EXPECT().Get(ctx, fmt.Sprintf("%s?symbol=%s", model.URLInsightBase+model.URLInsightLogosDirectCryptoSymbolLogos, symbol), nil).Return(nil, testErr)

				return &Logos{
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
			f := tc.fn(ctx, tc.input)
			resp, err := f.DirectCryptoSymbolLogo(ctx, tc.input)

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

func TestLogos_MultiSymbolLogos(t *testing.T) {
	var (
		expResp = []model.SymbolLogosResponse{
			{
				Symbol:           "1",
				Name:             "2",
				Logo:             "3",
				LogoOriginal:     "4",
				LogoNormal:       "5",
				LogoThumbnail:    "6",
				LogoSquare:       "7",
				LogoSquareStrict: "8",
			},
		}

		rawReq = `{"symbols":["AAPL"]}`

		rawResponse = `[{"symbol":"1","name":"2","logo":"3","logo_original":"4","logo_normal":"5","logo_thumbnail":"6","logo_square":"7","logo_square_strict":"8"}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.MultipleSymbolLogosRequest
		hasErr bool
		fn     func(ctx context.Context) *Logos
	}{
		{
			name:   "success",
			input:  model.MultipleSymbolLogosRequest{Symbols: []string{"AAPL"}},
			hasErr: false,
			fn: func(ctx context.Context) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightLogosMultiSymbolLogos, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Logos{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  model.MultipleSymbolLogosRequest{Symbols: []string{"AAPL"}},
			hasErr: true,
			fn: func(ctx context.Context) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightLogosMultiSymbolLogos, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Logos{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  model.MultipleSymbolLogosRequest{Symbols: []string{"AAPL"}},
			hasErr: true,
			fn: func(ctx context.Context) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightLogosMultiSymbolLogos, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Logos{
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
			f := tc.fn(ctx)
			resp, err := f.MultiSymbolLogos(ctx, tc.input)

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

func TestLogos_ConvertedSymbolLogo(t *testing.T) {
	var (
		expResp = io.NopCloser(strings.NewReader("a,b,c,d"))

		rawReq = `{"symbol":"AAPL","conversion":"thumbnail"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.SymbolLogoConvertedRequest
		hasErr bool
		fn     func(ctx context.Context) *Logos
	}{
		{
			name:   "success",
			input:  model.SymbolLogoConvertedRequest{Symbol: "AAPL", Conversion: model.ConversionThumbnail},
			hasErr: false,
			fn: func(ctx context.Context) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("a,b,c,d"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightLogosConvertedSymbolLogos, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Logos{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  model.SymbolLogoConvertedRequest{Symbol: "AAPL", Conversion: model.ConversionThumbnail},
			hasErr: true,
			fn: func(ctx context.Context) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightLogosConvertedSymbolLogos, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Logos{
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
			f := tc.fn(ctx)
			resp, err := f.ConvertedSymbolLogo(ctx, tc.input)

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

func TestLogos_MultipleCryptoSymbolLogo(t *testing.T) {
	var (
		expResp = []model.SymbolLogosResponse{
			{
				Symbol:           "1",
				Name:             "2",
				Logo:             "3",
				LogoOriginal:     "4",
				LogoNormal:       "5",
				LogoThumbnail:    "6",
				LogoSquare:       "7",
				LogoSquareStrict: "8",
			},
		}

		rawReq = `{"symbols":["AAPL"]}`

		rawResponse = `[{"symbol":"1","name":"2","logo":"3","logo_original":"4","logo_normal":"5","logo_thumbnail":"6","logo_square":"7","logo_square_strict":"8"}]`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.MultipleCryptoLogosRequest
		hasErr bool
		fn     func(ctx context.Context) *Logos
	}{
		{
			name:   "success",
			input:  model.MultipleCryptoLogosRequest{Symbols: []string{"AAPL"}},
			hasErr: false,
			fn: func(ctx context.Context) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader(rawResponse))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightLogosMultipleCryptoSymbolLogo, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Logos{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/unmarshal",
			input:  model.MultipleCryptoLogosRequest{Symbols: []string{"AAPL"}},
			hasErr: true,
			fn: func(ctx context.Context) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("x"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightLogosMultipleCryptoSymbolLogo, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Logos{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  model.MultipleCryptoLogosRequest{Symbols: []string{"AAPL"}},
			hasErr: true,
			fn: func(ctx context.Context) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightLogosMultipleCryptoSymbolLogo, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Logos{
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
			f := tc.fn(ctx)
			resp, err := f.MultipleCryptoSymbolLogo(ctx, tc.input)

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

func TestLogos_ConvertedCryptoSymbolLogo(t *testing.T) {
	var (
		expResp = io.NopCloser(strings.NewReader("a,b,c,d"))

		rawReq = `{"symbol":"AAPL","conversion":"thumbnail"}`

		testErr = errors.New("process error")
	)

	testCases := []struct {
		name   string
		input  model.SymbolLogoConvertedRequest
		hasErr bool
		fn     func(ctx context.Context) *Logos
	}{
		{
			name:   "success",
			input:  model.SymbolLogoConvertedRequest{Symbol: "AAPL", Conversion: model.ConversionThumbnail},
			hasErr: false,
			fn: func(ctx context.Context) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				r := io.NopCloser(strings.NewReader("a,b,c,d"))
				httpResponse := &http.Response{
					StatusCode: http.StatusOK,
					Body:       r,
				}

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightLogosConvertedCryptoSymbolLogo, bytes.NewBuffer(bb), nil).Return(httpResponse, nil)

				return &Logos{
					cli:    cli,
					logger: logrus.New(),
				}
			},
		},
		{
			name:   "err/cli",
			input:  model.SymbolLogoConvertedRequest{Symbol: "AAPL", Conversion: model.ConversionThumbnail},
			hasErr: true,
			fn: func(ctx context.Context) *Logos {
				ctrl := gomock.NewController(t)
				cli := mock.NewMockHTTPClient(ctrl)

				bb := []byte(rawReq)

				cli.EXPECT().Post(ctx, model.URLInsightBase+model.URLInsightLogosConvertedCryptoSymbolLogo, bytes.NewBuffer(bb), nil).Return(nil, testErr)

				return &Logos{
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
			f := tc.fn(ctx)
			resp, err := f.ConvertedCryptoSymbolLogo(ctx, tc.input)

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
