package storage

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

const (
	defaultKeyForStore = "go-sdk-orbis-token-key"
)

type Config struct {
	Host     string
	Port     int
	Password string
	DB       int

	KeyForStore string
}

// Client is a redis client for using for storing token data.
// It implements interface Storage.
// You can use it instead of default in-memory storage.
type Client struct {
	cli *redis.Client

	keyForStore string
}

func NewRedisStorage(cfg Config) (*Client, error) {
	cli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := cli.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	if cfg.KeyForStore == "" {
		cfg.KeyForStore = defaultKeyForStore
	}

	return &Client{
		cli:         cli,
		keyForStore: cfg.KeyForStore,
	}, nil
}

func (c *Client) Store(ctx context.Context, data []byte) error {
	log.Trace("Redis: method Store called")

	return c.cli.Set(ctx, c.keyForStore, data, 0).Err()
}

func (c *Client) Get(ctx context.Context) ([]byte, error) {
	log.Trace("Redis: method Get called")

	return c.cli.Get(ctx, c.keyForStore).Bytes()
}
