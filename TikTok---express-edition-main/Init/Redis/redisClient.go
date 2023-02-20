package Redis

import (
	"douyin/config"
	"github.com/go-redis/redis/v8"
)

// InitRedis 获取redis客户端
func InitRedis() *redis.Client {
	redisDb := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_ADDR,
		Password: config.REDIS_PASSSWORD,
		DB:       config.DB,
	})
	return redisDb
}
