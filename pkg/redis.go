package pkg

import (
	"cappuccino/config"
	"github.com/go-redis/redis"
	"github.com/happierall/l"
)

var redisClient *redis.Client

func InitRedis()  {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.Admin.Redis.Host,
		Password: config.Admin.Redis.Password,
		DB:       5,
	})
	l.Warn("redis 初始化成功")
}


// GetInstance 获取redis实例
func GetRedisInstance() *redis.Client {
	return redisClient
}

// Close 关闭redis
func Close() {
	_ = redisClient.Close()
	redisClient = nil
}