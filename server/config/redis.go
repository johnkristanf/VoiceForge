package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

var ctx = context.Background()

type RedisMethod interface {
	CacheSet(interface{}, string) error
	CacheGet(string, interface{}) error
	CacheDelete(string) error
}

func RedisConfig(redisURL string) (*Redis, error) {
	u, err := url.Parse(redisURL)
	if err != nil {
		return nil, err
	}

	host := u.Hostname()
	port := u.Port()

	password, _ := u.User.Password()

	var client *redis.Client

	// Function to establish a new connection to Redis
	connectToRedis := func() (*redis.Client, error) {
		return redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: password,
			DB:       0,
			OnConnect: func(ctx context.Context, cn *redis.Conn) error {
				log.Println("Successfully reconnected to Redis.")
				return nil
			},
		}), nil
	}
	

	// Create a new Redis client
	client, err = connectToRedis()
	if err != nil {
		return nil, err
	}

	// Test the connection to Redis
	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &Redis{
		client: client,
	}, nil
}

func (r *Redis) CacheSet(cachedData interface{}, cacheKey string) error {

	jsonData, err := json.Marshal(cachedData)
	if err != nil {
		return err
	}

	if err := r.client.Set(ctx, cacheKey, jsonData, 10*time.Second).Err(); err != nil {
		return err
	}

	return nil
}

func (r *Redis) CacheGet(cacheKey string, value interface{}) error {

	cachedData, err := r.client.Get(ctx, cacheKey).Result()
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(cachedData), value); err != nil {
		return err
	}

	return nil
}

func (r *Redis) CacheDelete(cacheKey string) error {

	if err := r.client.Del(ctx, cacheKey).Err(); err != nil {
		return err
	}

	return nil
}
