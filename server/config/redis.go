package config

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct{
	client *redis.Client
}

var ctx = context.Background()


type RedisMethod interface{
	CacheSet(interface{}, string) error
	CacheGet(string, interface{}) error
	CacheDelete(string) error
}

func RedisConfig() (*Redis, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", 
		Password: "",               
		DB:       0,              
	})

	return &Redis{
		client: client,
	}, nil
}


func (r *Redis) CacheSet(cachedData interface{}, cacheKey string) error {
	jsonData, err := json.Marshal(cachedData)
	if err != nil{
		return err
	}

	if err := r.client.Set(ctx, cacheKey, jsonData, 10 * time.Second).Err(); err != nil{
		return err
	}

	return nil
}


func (r *Redis) CacheGet(cacheKey string, value interface{}) error {

	cachedData, err := r.client.Get(ctx, cacheKey).Result()
	if err != nil{
		return err
	}

	if err := json.Unmarshal([]byte(cachedData), value); err != nil{
		return err
	}

	return nil
}


func (r *Redis) CacheDelete(cacheKey string) error {

	if err := r.client.Del(ctx, cacheKey).Err(); err != nil{
		return err
	}
	
	return nil
}