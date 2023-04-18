package config

import (
	"encoding/json"
	"github.com/spf13/viper"
	"log"
)

// Created at 2023/4/10 15:04
// Created by Yoake

type Configuration struct {
	Server
	Mysql
	Redis
	EdgeGPT
}

type Server struct {
	Port int
}
type Mysql struct {
	MName  string
	MPwd   string
	DBName string
	Port   int
}
type Redis struct {
	RName string
	RPwd  string
}
type EdgeGPT struct {
	Cookies    []map[string]interface{}
	CookiePath string
	Proxy      string
}

var Preset Configuration

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("config not exists")
		} else {
			log.Println("read config error")
		}

	}

	Preset.Server.Port = viper.GetInt("Server.port")
	Preset.Mysql.MName = viper.GetString("Mysql.user")
	Preset.Mysql.MPwd = viper.GetString("Mysql.password")
	Preset.Mysql.Port = viper.GetInt("Mysql.port")
	Preset.Mysql.DBName = viper.GetString("Mysql.db_name")

	Preset.EdgeGPT.setDefaultConf("EdgeGPT.cookiePath", "cookies.json")
	Preset.EdgeGPT.setDefaultConf("EdgeGPT.cookies", "")
	Preset.EdgeGPT.setDefaultConf("EdgeGPT.proxy", "http://127.0.0.1:7890")
	Preset.EdgeGPT.Proxy = viper.GetString("EdgeGPT.proxy")
	log.Println("Configuration loaded successfully!")

}

func (e *EdgeGPT) setDefaultConf(key string, value string) {
	get := viper.GetString(key)
	if get != "" {
		switch key {
		case "EdgeGPT.cookiePath":
			e.CookiePath = get
		case "EdgeGPT.cookies":
			var data []map[string]interface{}
			err := json.Unmarshal([]byte(get), &data)
			if err != nil {
				log.Panic("Unmarshal cookies err")
			}
			e.Cookies = data
		case "EdgeGPT.proxy":
			e.Proxy = get
		}
	} else {
		switch key {
		case "EdgeGPT.cookiePath":
			e.CookiePath = value
		case "EdgeGPT.cookies":
			if value != "" {
				var data []map[string]interface{}
				err := json.Unmarshal([]byte(value), &data)
				if err != nil {
					log.Panic("Unmarshal cookies err")
				}
				e.Cookies = data
			} else {
				e.Cookies = nil
			}

		case "EdgeGPT.proxy":
			e.CookiePath = value
		}
	}

}
