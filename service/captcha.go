package service

import (
	"cappuccino/config"
	"github.com/LyricTian/captcha"
	"github.com/LyricTian/captcha/store"
	"github.com/go-redis/redis"
)

// InitCaptcha 初始化图形验证码
func InitCaptcha() {
	cfg := config.Admin.Captcha
	if cfg.Store == "redis" {
		rc := config.Admin.Redis
		captcha.SetCustomStore(store.NewRedisStore(&redis.Options{
			Addr:     rc.Host,
			Password: rc.Password,
			DB:       cfg.RedisDB,
		}, captcha.Expiration, nil, cfg.RedisPrefix))
	}
}
