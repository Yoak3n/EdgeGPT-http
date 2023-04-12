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
	Port  int
	Proxy string
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
	Preset.Proxy = viper.GetString("Server.proxy")

	Preset.Mysql.MName = viper.GetString("Mysql.user")
	Preset.Mysql.MPwd = viper.GetString("Mysql.password")
	Preset.Mysql.Port = viper.GetInt("Mysql.port")
	Preset.Mysql.DBName = viper.GetString("Mysql.db_name")

	get := viper.GetString("EdgeGPT.cookies")
	if get != "" {
		var data []map[string]interface{}
		err = json.Unmarshal([]byte(get), &data)
		if err != nil {
			log.Panic("Unmarshal cookies err")
		}
		Preset.EdgeGPT.Cookies = data
	}

	Preset.EdgeGPT.CookiePath = viper.GetString("EdgeGPT.cookiePath")

	log.Println("Configuration loaded successfully!")
}
