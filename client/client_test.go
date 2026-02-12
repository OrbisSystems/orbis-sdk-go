package client

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/OrbisSystems/orbis-sdk-go/config"
	"github.com/OrbisSystems/orbis-sdk-go/interfaces/mock"
)

func TestClient_Close(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ws := mock.NewMockWS(ctrl)
		ws.EXPECT().Close().Return(nil)
		c := &Client{WS: ws}

		err := c.Close()

		assert.NoError(t, err)
	})

	t.Run("err", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ws := mock.NewMockWS(ctrl)
		ws.EXPECT().Close().Return(errors.New("error"))
		c := &Client{WS: ws}

		err := c.Close()

		assert.Error(t, err)
		assert.EqualError(t, err, "error")
	})
}

func TestClient_getBaseURL(t *testing.T) {
	t.Run("with protocol", func(t *testing.T) {
		cfg := config.Config{
			Host:     "example.com",
			Protocol: "http",
		}

		result := getBaseURL(cfg)

		assert.Equal(t, "http://example.com", result)
	})

	t.Run("without protocol", func(t *testing.T) {
		cfg := config.Config{
			Host: "example.com",
		}

		result := getBaseURL(cfg)

		assert.Equal(t, "https://example.com", result)
	})

	t.Run("custom protocol is not acceptable", func(t *testing.T) {
		cfg := config.Config{
			Host:     "api.test.io",
			Protocol: "wss",
		}

		result := getBaseURL(cfg)

		assert.Equal(t, "https://api.test.io", result)
	})
}
