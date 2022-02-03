package datastore

import (
	"github.com/go-redis/redis/v8"
	"log"
)

var REDIS_CLIENT *redis.Client

func GetRedisClientFactory() *redis.Client {
	if REDIS_CLIENT != nil {
		return REDIS_CLIENT
	}
	REDIS_CLIENT = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	log.Println("Redis connection established successfully!.")
	return REDIS_CLIENT
}
