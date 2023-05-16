package client

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

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
