package redisDB

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
	"userservice/internal/config"
	"userservice/internal/utils"
)

var ctx = context.Background()

// InitRedis initializes the redis client
func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.EnvConfigs.RedisAddr,
		DB:   config.EnvConfigs.RedisDB,
	})

	utils.LogInfo("Redis client initialized")

	return rdb
}

// CloseRedis closes the redis connection
func CloseRedis(rdb *redis.Client) {
	err := rdb.Close()
	utils.LogErr("Error while closing redis connection", err)
	utils.LogInfo("Redis connection closed")
}

type CookieData struct {
	UserId       uint64    `json:"user_id"`
	SessionToken string    `json:"session_token"`
	Key          string    `json:"key"`
	Timestamp    time.Time `json:"timestamp"`
}

// InsertCookieData inserts the cookie data to redis
func InsertCookieData(data *CookieData) error {

	rdb := InitRedis()
	defer CloseRedis(rdb)

	// Convert data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		utils.LogErr("Error while converting cookie data to JSON", err)
		return err
	}

	// Insert data to redis
	err = rdb.Set(ctx, data.Key, jsonData, 2*time.Minute).Err()
	if err != nil {
		utils.LogErr("Error while converting cookie data to JSON", err)
		return err
	}

	utils.LogInfo("Cookie data inserted to redis")

	return nil
}

// GetCookieData retrieves the cookie data from redis using the specified key
func GetCookieData(key string) (*CookieData, error) {
	rdb := InitRedis()
	defer CloseRedis(rdb)

	// Fetch data from redis using the key
	jsonData, err := rdb.Get(ctx, key).Result()
	if err != nil {
		utils.LogErr("Error while fetching cookie data from redis", err)
		return nil, err
	}

	var data CookieData
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		utils.LogErr("Error while unmarshalling cookie data", err)
		return nil, err
	}

	utils.LogInfo("Cookie data retrieved from redis")

	return &data, nil
}
