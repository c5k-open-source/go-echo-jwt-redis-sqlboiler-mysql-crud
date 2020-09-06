package external

import (
	"time"

	"github.com/c5k-open-source/go-echo-jwt-redis-sqlboiler-mysql-crud/pkg/config"
	"github.com/go-redis/redis"
)

func NewSecurityRedisClient() *redis.Client {

	rClient := redis.NewClient(&redis.Options{
		Addr:         config.Config.RedisSecurityHost,
		Password:     "",
		MaxRetries:   3,
		WriteTimeout: 15 * time.Second,
	})

	if err := rClient.Ping().Err(); err != nil {
		panic(err)
	}

	return rClient
}

func NewRedisClient() *redis.Client {

	rClient := redis.NewClient(&redis.Options{
		Addr:         config.Config.RedisHost,
		Password:     "",
		MaxRetries:   3,
		WriteTimeout: 15 * time.Second,
	})

	if err := rClient.Ping().Err(); err != nil {
		panic(err)
	}

	return rClient
}
