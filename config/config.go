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
	viper.SetConfigType("yml")
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
	viper.SetDefault("EdgeGPT.cookiePath", "cookies.json")
	viper.SetDefault("EdgeGPT.cookies", "")
	viper.SetDefault("EdgeGPT.proxy", "")

	Preset.Server.Port = viper.GetInt("Server.port")
	Preset.Mysql.MName = viper.GetString("MySQL.user")
	Preset.Mysql.MPwd = viper.GetString("MySQL.password")
	Preset.Mysql.Port = viper.GetInt("MySQL.port")
	Preset.Mysql.DBName = viper.GetString("MySQL.db_name")

	Preset.EdgeGPT.CookiePath = viper.GetString("EdgeGPT.cookiePath")
	Preset.EdgeGPT.Proxy = viper.GetString("EdgeGPT.proxy")
	var data []map[string]interface{}
	get := viper.GetString("EdgeGPT.cookies")
	if get != "" {
		err = json.Unmarshal([]byte(get), &data)
		if err != nil {
			log.Panic("Unmarshal cookies err")
		}
		Preset.EdgeGPT.Cookies = data
	}

	log.Println("Configuration loaded successfully!")

	// 监控配置文件
	viper.WatchConfig()
}
