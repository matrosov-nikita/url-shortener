// Package redis provides implementation for Cacher interface.
package redis

import (
	"github.com/go-redis/redis"
	"github.com/matrosov-nikita/url-shortener/storage"
)

const redisKey = "key"

// New creates redis cacher instance.
func New(redisURL string) (storage.Cacher, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)

	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &cacher{client}, nil
}

type cacher struct{ client *redis.Client }

func (c *cacher) GetUniqueKey() (int64, error) {
	res, err := c.client.Get(redisKey).Int64()
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (c *cacher) IncrUniqueKey() error {
	_, err := c.client.Incr(redisKey).Result()
	return err
}
