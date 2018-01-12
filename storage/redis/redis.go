// Package redis provides implementation for Cacher interface
package redis

import (
	"github.com/go-redis/redis"
)

// TODO should we pass DB?
// New return redis cacher instance
func New(addr, password string) (*cacher, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &cacher{client}, nil
}

type cacher struct{ client *redis.Client }

func (c *cacher) GetUniqueKey(key string) (int64, error) {
	res, err := c.client.Get(key).Int64()
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (c *cacher) IncrUniqueKey(key string) error {
	_, err := c.client.Incr(key).Result()
	return err
}
