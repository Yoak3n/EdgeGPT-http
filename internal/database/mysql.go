package database

import (
	"database/sql"
	"fmt"
	"github.com/Yoak3n/EdgeGPT-http/config"
	"github.com/Yoak3n/EdgeGPT-http/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// Created at 2023/4/10 14:56
// Created by Yoake
var conn *sql.DB
var DB *gorm.DB
var user model.User
var err error

func init() {
	conf := config.Preset.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.MName, conf.MPwd, conf.Port, conf.DBName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("database connected err")
	}
	err = DB.AutoMigrate(model.Data{})
	//err = db.AutoMigrate(model.User{})
	if err != nil {
		log.Panic("create table failed")
	}
	conn, _ = DB.DB()
	conn.SetMaxOpenConns(100)
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(time.Hour)
	log.Println("MySQL already connected")
}

func UserRegister(session string, uid string) {
	DB.Create(&model.User{UID: uid, Session: session})
}

func ReadUser(session string) model.User {
	DB.First(user, "session = ?", session)
	return user
}

func CloseConnect() {
	conn.Close()
}

func CreateMessage(question string, answer string, session string) {
	message := model.Message{
		Question: question,
		Answer:   answer,
		Session:  session,
	}
	DB.Create(&model.Data{Message: message})
	log.Printf("MySQL成功写入一条属于%s的消息", session)
}

func GetSomeoneAllMessages(session string) []model.Data {
	var results []model.Data
	DB.Find(&results).Where("session =?", session)
	return results
}
