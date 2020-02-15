package config

import (
	"cappuccino/utils"
	"log"
)


type AppConfig struct {
	Version    string      `json:"version"`
	Server     server      `json:"server"`
	Database   database    `json:"database"`
	Static     static      `json:"static"`
	App        app         `json:"app"`
	Jwt        jwt         `json:"jwt"`
}

type server struct {
	Port  string `json:"port"`
}

type app struct {
	ApiPrefix string `json:"apiPrefix"`
}

type database struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Schema   string `json:"schema"`
	OpenLog  bool  `json:"openLog"`
}

type static struct {
	Dir string `json:"dir"`
	Url string `json:"url"`
}

type jwt struct {
	Expires int64 `json:"expires"`
}

var conf = &AppConfig{}

func init() {
	var err error
	err = utils.LoadOrStoreConfig("config/config.json",&conf)
	if err != nil {
		log.Fatalf("Fail to parse 'config/config.json': %v", err)
	}
}


func GetAppConfig() *AppConfig {
	return conf
}
