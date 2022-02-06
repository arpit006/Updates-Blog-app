package datastore

import (
	"context"
	"log"
)

var ctx = context.Background()

func AddToRedis(key string, value string) {
	client := GetRedisClientFactory()
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Printf("Error saving data to redis [%s : %s]. Error is %s", key, value, err)
		panic(err)
	}
}

func GetFromRedis(key string) (string, error) {
	client := GetRedisClientFactory()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Printf("Error retreiving value from Redis for key: [%s]. Error is %s", key, err)
		//panic(err)
		return "", err
	}
	return val, nil
}

func LPushToRedis(key string, value string) {
	client := GetRedisClientFactory()
	_, err := client.LPush(ctx, key, value).Result()
	if err != nil {
		log.Printf("Error pushing range to Redis for [%s: %s]. Error is %s", key, value, err)
		panic(err)
	}
}

func GetRangeFromRedis(key string, start int64, end int64) ([]string, error) {
	client := GetRedisClientFactory()
	result, err := client.LRange(ctx, key,  start, end).Result()
	if err != nil {
		log.Printf("Error retreiving range from Redis for key: [%s]. Error is %s", key, err)
		return nil, err
	}
	return result, nil
}

func SaveBytesToRedis(key string, value []byte) {
	client := GetRedisClientFactory()
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Printf("Error Saving bytes to Redis for [%s:%s]. Error is %s", key, string(value), err)
		panic(err)
	}
}

func GetBytesFromRedis(key string) ([]byte, error) {
	client := GetRedisClientFactory()
	bytes, err := client.Get(ctx, key).Bytes()
	if err != nil {
		log.Printf("Error Getting bytes to Redis for Key : [%s]. Error is %s", key, err)
		return nil, err
	}
	return bytes, nil
}
