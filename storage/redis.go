package storage

import (
	"fmt"

	"github.com/go-redis/redis"
)

// New return redis storage instance
func New(addr, password string) (*redist, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       0,        // use default DB
	})

	fmt.Println(client)

	pong, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	fmt.Println(pong, err)
	return &redist{client}, nil
}

func (r *redist) GetUniqueKey() (int, error) {
	const defaultKeyVal = 125
	_, err := r.client.Get("key").Result()
	if err != nil {
		r.client.Set("key", defaultKeyVal, 0)
		return defaultKeyVal, nil
	}

	res := r.client.Incr("key")
	return int(res.Val()), nil
}

// Add new item to database
func (r *redist) Add(url string) error {
	return nil
}

type redist struct{ client *redis.Client }
