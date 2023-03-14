package ws

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/OrbisSystems/orbis-sdk-go/config"
	sdk "github.com/OrbisSystems/orbis-sdk-go/interface"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

var (
	ErrUnknownSubscriptionType = errors.New("unknown subscription type")
)

type Client struct {
	auth     sdk.Auth
	url      string
	newsConn *websocket.Conn
}

func New(cfg config.Config, auth sdk.Auth) *Client {
	return &Client{
		url:  wrapWS(cfg.Host),
		auth: auth,
	}
}

func (c *Client) Subscribe(ctx context.Context, subscriptionType model.SubscriptionType) (chan model.News, error) {
	switch subscriptionType {
	case model.NewsSubscription:
		return c.subscribeNews(ctx)
	default:
		return nil, ErrUnknownSubscriptionType
	}
}

func (c *Client) subscribeNews(ctx context.Context) (chan model.News, error) {
	if c.newsConn == nil {
		tkn, err := c.auth.GetToken(ctx)
		if err != nil {
			return nil, err
		}

		headers := http.Header{}
		headers.Add("Authorization", fmt.Sprintf("Bearer %s", tkn.AccessToken))

		cli, _, err := websocket.DefaultDialer.Dial(c.url+model.WSInsightNews, headers)
		if err != nil {
			return nil, err
		}
		c.newsConn = cli
	}

	outputCh := make(chan model.News, 100)
	go func() {
		for {
			var msg model.News
			err := c.newsConn.ReadJSON(&msg)
			if err != nil {
				log.Errorf("error while getting message via ws: %v", err)
			}

			outputCh <- msg
		}
	}()

	return outputCh, nil
}

func wrapWS(hostname string) string {
	return fmt.Sprintf("ws://%s", hostname)
}
