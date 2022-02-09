package datastore

import (
	"log"
)

var ctx = ContextFactory()

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

func Incr() (int64, error) {
	client := GetRedisClientFactory()
	return client.Incr(ctx, "user:next-id").Result()
}


func HSetStrStr(hash, key, value string) (int64, error){
	client := GetRedisClientFactory()
	id, err := client.HSet(ctx, hash, key, value).Result()
	if err != nil {
		log.Printf("Error in HSET Redis for hash: [%s] key: [%s] value: [%s]. Error is %s", hash, key, value, err)
		return -1, err
	}
	return id, nil
}

func HGetStrStr(hash, key string) (string, error) {
	client := GetRedisClientFactory()
	res, err := client.HGet(ctx, hash, key).Result()
	if err != nil {
		log.Printf("HGET Redis Error for hash: [%s] key: [%s]. Error is %s", hash, key, err)
		return "", err
	}
	return res, nil
}

func HGetStrBytesArr(hash, key string) ([]byte, error) {
	client := GetRedisClientFactory()
	res, err := client.HGet(ctx, hash, key).Bytes()
	if err != nil {
		log.Printf("HGET Redis Error for hash: [%s] key: [%s]. Error is %s", hash, key, err)
		return nil, err
	}
	return res, nil
}

func HGetStrInt(hash, key string) (int64, error) {
	client := GetRedisClientFactory()
	res, err := client.HGet(ctx, hash, key).Int64()
	if err != nil {
		log.Printf("HGET Redis Error for hash: [%s] key: [%s]. Error is %s", hash, key, err)
		return -1, err
	}
	return res, nil
}
