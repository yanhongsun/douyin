package redisdb

import (
	"douyin/cmd/comment/pack/configdata"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func Init() {
	client := redis.NewClient(&redis.Options{
		Addr:     configdata.RedisDatabaseConfig.Host,
		Password: configdata.RedisDatabaseConfig.Password,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}
