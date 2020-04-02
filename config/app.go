package config

import (
	"github.com/jinzhu/configor"
	"log"
	"time"
)

var Admin = struct {
	App struct {
		Name        string
		Description string
		Version     string
		ApiPrefix   string
	}
	Server struct {
		Port         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
		RunMode      string
	}
	Db struct {
		Type     string
		Host     string
		User     string
		Password string
		Schema   string
		Port     string
		Log      bool
	}
	Mail struct {
		User     string
		Password string
		Host     string
		Port     string
	}
	Mode struct {
		Type      string
		SecretKey string
	}
	File struct {
		Path   string
		Dir    string
		Max    int64
		Suffix []string
	}
	Jwt struct {
		Expires     int64
		SecretKey   string
		RedisDb     int
		RedisPrefix string
	}
	Redis struct {
		Host     string
		Password string
	}
	Captcha struct {
		Store       string
		Length      int
		Width       int
		Height      int
		RedisDB     int
		RedisPrefix string
	}
}{}

func init() {
	var err error
	err = configor.Load(&Admin, "config/app.yml")
	if err != nil {
		log.Fatalf("Fail to parse 'config/config.yml': %v", err)
	}
}
