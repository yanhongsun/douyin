package redisdb

import "github.com/go-redis/redis"

var RedisClient *redis.Client

func Init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "gorm",
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}
