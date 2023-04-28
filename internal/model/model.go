package model

// Created at 2023/4/10 21:57
// Created by Yoake
import (
	"gorm.io/gorm"
)

type Data struct {
	gorm.Model
	Message
}

type User struct {
	UID     string `gorm:"UNIQUE"`
	Session string `gorm:"NOT NULL"`
}

type Message struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Session  string `json:"session"`
}
