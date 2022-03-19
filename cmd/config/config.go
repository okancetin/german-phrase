package config

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/okancetin/german-phrase/cmd/api/cmd/entity"
	"time"
)

type RedisConfig struct {
	host     string
	password string
	db       int
	expires  time.Duration
}

func NewRedisClient(host, password string, exp time.Duration) *RedisConfig {
	return &RedisConfig{
		host:     host,
		password: password,
		expires:  exp,
	}
}

func (redisConfig *RedisConfig) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     redisConfig.host,
		Password: redisConfig.password,
		DB:       redisConfig.db,
	})
}

func (redisConfig *RedisConfig) Get(key string) *entity.Phrase {
	client := redisConfig.getClient()
	val, err := client.Get(key).Result()

	if err != nil {
		return nil
	}

	phrase := entity.Phrase{}
	err = json.Unmarshal([]byte(val), &phrase)
	if err != nil {
		panic(err)
	}
	return &phrase
}
