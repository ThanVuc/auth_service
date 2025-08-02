package initialize

import (
	"auth_service/global"
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

/*
@Author: Sinh
@Date: 2025/6/1
@Description: Initialize the Redis database connection.
This function uses the go-redis package to create a Redis client.
It constructs the connection string using the configuration from global.Config.Redis.
@Note: The Redis client is stored in global.RedisDb for use throughout the application.
*/
var ctx = context.Background()

func InitRedis() {
	redisConfig := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + strconv.Itoa(redisConfig.Port),
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.DB,       // use default DB
		PoolSize: redisConfig.PoolSize,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}

	global.Logger.Info("Redis connection established successfully", "")
	global.RedisDb = rdb
}

func testRedis() {
	// Example of setting and getting a value
	err := global.RedisDb.Set(ctx, "test_key", "test_value", 0).Err()
	if err != nil {
		fmt.Println("Failed to set test key in Redis:", err)
		return
	}

	val, err := global.RedisDb.Get(ctx, "test_key").Result()
	if err != nil {
		fmt.Println("Failed to get test key from Redis:", err)
		return
	}

	fmt.Println("Value for 'test_key':", val)
}
