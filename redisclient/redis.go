package redisclient

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/hanhnguyenduc/config-server/setting"
)

var redisClient *redis.Client

func Setup() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.RedisAddress,
		Password: setting.RedisSetting.RedisPassword, // no password set
		DB:       setting.RedisSetting.RedisDB,       // use default DB
	})
}

func Set(key string, value interface{}, timeOut time.Duration) error {
	err := redisClient.Set(key, value, timeOut).Err()
	if err != nil {
		return err
	}
	return nil
}

func Get(key string) (interface{}, error) {
	val, err := redisClient.Get(key).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}
