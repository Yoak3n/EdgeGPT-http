package gpt

// Created at 2023/4/10 18:50
// Created by Yoake
import (
	"edgegpt-http/config"
	"edgegpt-http/edgegpt"
	"log"
)

type EdgeBot struct {
	text string
	json map[string]interface{}
}

func NewConv() *edgegpt.ChatBot {
	bot := edgegpt.NewChatBot("cookies.json", []map[string]interface{}{}, config.Preset.Proxy)
	err := bot.Init()
	if err != nil {
		log.Println("init bot failed")
	}
	return bot
}
