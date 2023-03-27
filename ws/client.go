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
	sdk.Auth
	url string

	newsConn *websocket.Conn

	newsOutputCh chan model.News
	closeCh      chan struct{}
}

func New(cfg config.Config, auth sdk.Auth) *Client {
	return &Client{
		url:          wrapWS(cfg.Host),
		Auth:         auth,
		newsOutputCh: make(chan model.News, 100),
		closeCh:      make(chan struct{}),
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

func (c *Client) Close() error {
	close(c.newsOutputCh)
	close(c.closeCh)

	return c.newsConn.Close()
}

func (c *Client) subscribeNews(ctx context.Context) (chan model.News, error) {
	if c.newsConn == nil {
		if err := c.connect(ctx); err != nil {
			return nil, err
		}
	}

	go c.processNewsFeed(ctx)

	return c.newsOutputCh, nil
}

func (c *Client) connect(ctx context.Context) error {
	tkn, err := c.GetToken(ctx)
	if err != nil {
		return err
	}

	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", tkn.AccessToken))

	cli, _, err := websocket.DefaultDialer.Dial(c.url+model.WSInsightNews, headers)
	if err != nil {
		return err
	}
	c.newsConn = cli

	return nil
}

func (c *Client) processNewsFeed(ctx context.Context) {
	for {
		select {
		case <-c.closeCh:
			return
		default:
			c.getFeedNews(ctx)
		}
	}
}

func (c *Client) getFeedNews(ctx context.Context) {
	var msg model.News
	err := c.newsConn.ReadJSON(&msg)
	if err != nil {
		if ce, ok := err.(*websocket.CloseError); ok {
			switch ce.Code {
			case websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseNoStatusReceived:
				log.Info("normal ws connection closing.")
			case websocket.CloseProtocolError, websocket.CloseAbnormalClosure, websocket.CloseInternalServerErr:
				// reconnect
				log.Warn("abnormal ws connection closing. Reconnect...")
				if err := c.connect(ctx); err != nil {
					log.Errorf("can't reconnect to ws server. Terminating. Err : %v", err)
					c.Close()
					return
				}
			}
			return
		}
		log.Errorf("error while getting message via ws: %v", err)
	}

	c.newsOutputCh <- msg
}

func wrapWS(hostname string) string {
	return fmt.Sprintf("wss://%s", hostname)
}
