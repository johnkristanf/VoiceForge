package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"crypto/tls"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

type RedisMethod interface {
	CacheSet(interface{}, string) error
	CacheGet(string, interface{}) error
	CacheDelete(string) error
}

var ctx = context.Background()

func RedisConfig(redisURL string) (*Redis, error) {

	u, err := url.Parse(redisURL)
	if err != nil {
		return nil, err
	}

	var client *redis.Client

	connectToRedis := func() (*redis.Client, error) {

		host := u.Hostname()
	    port := u.Port()
	    password, _ := u.User.Password()
		username := u.User.Username()

		config := &tls.Config{
			InsecureSkipVerify: true,
			ServerName: host,
		}
	

		return redis.NewClient(&redis.Options{
			Addr:      fmt.Sprintf("%s:%s", host, port),
			Username:  username,
			Password:  password, 
			DB:        0,
			TLSConfig: config,
			DialTimeout: 60 * time.Second,  
		    ReadTimeout: 60 * time.Second,
		}), nil
	}

	client, err = connectToRedis()
	if err != nil {
		return nil, err
	}

	// TEST CONNECTION TO THE REDIS SERVER
	pong, err := client.Ping(ctx).Result(); 
	if err != nil {
		return nil, err
	}

	log.Println("CONNECT NIMAL", pong)


	go func() {

		for {

			if _, err := client.Ping(ctx).Result(); err != nil {
				log.Println("CONNECTION LOST ATTEMPTING TO RECONNECT....")

				newClient, reConnErr := connectToRedis()

				if reConnErr != nil {
					log.Printf("FAILED TO CONNECT TO THE REDIS SERVER: %v \n", reConnErr)
					time.Sleep(5 * time.Second)
					continue
				}

				client = newClient

				log.Println("Successfully reconnected to Redis.")

			}

			time.Sleep(1 * time.Minute)
		}

	}()

	return &Redis{
		client: client,
	}, nil
}

func (r *Redis) CacheSet(cachedData interface{}, cacheKey string) error {

	jsonData, err := json.Marshal(cachedData)
	if err != nil {
		return err
	}

	if err := r.client.Set(ctx, cacheKey, jsonData, 10 * time.Second).Err(); err != nil {
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
