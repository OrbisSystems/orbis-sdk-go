package ws

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/OrbisSystems/orbis-sdk-go/config"
	sdk "github.com/OrbisSystems/orbis-sdk-go/interfaces"
	"github.com/OrbisSystems/orbis-sdk-go/model"
)

const (
	retryAttempts = 5
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
		if err := c.connectWithRetry(ctx); err != nil {
			return nil, err
		}
	}

	go c.processNewsFeed(ctx)

	return c.newsOutputCh, nil
}

func (c *Client) connectWithRetry(ctx context.Context) error {
	var err error
	for i := 0; i < retryAttempts; i++ {
		log.Debugf("trying to connect to ws. Attemp #%d", i+1)
		if err = c.connect(ctx); err != nil {
			log.Warnf("cannot connect to ws server. Attempt #%d, err %v", i+1, err)
			time.Sleep(time.Second * 5)
			continue
		}
		break
	}

	if err != nil {
		log.Errorf("cannot connect to ws server: %v", err)
	}

	return err
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
			if err := c.getFeedNews(ctx); err != nil {
				return
			}
		}
	}
}

func (c *Client) getFeedNews(ctx context.Context) error {
	var msg model.News
	err := c.newsConn.ReadJSON(&msg)
	if err != nil {
		if ce, ok := err.(*websocket.CloseError); ok {
			switch ce.Code {
			case websocket.CloseProtocolError, websocket.CloseTLSHandshake:
				log.Errorf("ws error: %v", err)
				return err
			default:
				if err := c.connectWithRetry(ctx); err != nil {
					log.Errorf("can't reconnect to ws server. Terminating. Err : %v", err)
					c.Close()
					return err
				}
			}
		}
		log.Errorf("error while getting message via ws: %v", err)
	}

	c.newsOutputCh <- msg

	return nil
}

func wrapWS(hostname string) string {
	return fmt.Sprintf("wss://%s", hostname)
}
