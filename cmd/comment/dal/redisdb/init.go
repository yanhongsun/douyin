package redisdb

import (
	"context"
	"douyin/cmd/comment/pack/configdata"

	"github.com/go-redis/redis/v8"
	hook "github.com/imcvampire/opentracing-goredisv8"
	"github.com/opentracing/opentracing-go"
)

var RedisClient *redis.Client

func Init() {
	client := redis.NewClient(&redis.Options{
		Addr:     configdata.RedisDatabaseConfig.Host,
		Password: configdata.RedisDatabaseConfig.Password,
	})

	h := hook.NewHook(hook.WithTracer(opentracing.GlobalTracer()))
	client.AddHook(h)

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}
