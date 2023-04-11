package database

import (
	"database/sql"
	"edgegpt-http/config"
	"edgegpt-http/internal/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Conn *sql.DB

// Created at 2023/4/10 14:56
// Created by Yoake
func init() {
	conf := config.Preset.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.MName, conf.MPwd, conf.Port, conf.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("database connected err")
	}
	err = db.AutoMigrate(model.Data{})
	if err != nil {
		log.Panic("create table failed")
	}
	Conn, _ = db.DB()
	Conn.SetMaxOpenConns(5)
	Conn.SetMaxIdleConns(2)

}
