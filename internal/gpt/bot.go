package gpt

// Created at 2023/4/10 18:50
// Created by Yoake
import (
	"github.com/Yoak3n/EdgeGPT-http/config"
	"github.com/Yoak3n/EdgeGPT-http/edgegpt"
	"github.com/google/uuid"
	"log"
)

type BotPool struct {
	Workers map[string]*EdgeBot
	Size    int
}

type EdgeBot struct {
	UUID    string
	Session string
	Answer  *edgegpt.Answer
	Bot     *edgegpt.ChatBot
	Style   edgegpt.ConversationStyle
	End     bool
}

func NewBot(session string, style string) *EdgeBot {
	botConf := config.Preset.EdgeGPT
	bot := edgegpt.NewChatBot(botConf.CookiePath, botConf.Cookies, botConf.Proxy)
	err := bot.Init()
	if err != nil {
		log.Println("init bot failed")
	}
	var s edgegpt.ConversationStyle
	switch style {
	case "bing-c":
		s = edgegpt.Creative
	case "bing-b":
		s = edgegpt.Balanced
	case "bing-p":
		s = edgegpt.Precise
	}
	// import uuid
	return &EdgeBot{uuid.NewString(), session, nil, bot, s, false}
}

func (e *EdgeBot) OnQuestion(question string) {
	err := e.Bot.Ask(question, e.Style, e.callback)
	if err != nil {
		log.Println("no answer")
	}
}

func (e *EdgeBot) callback(a *edgegpt.Answer) {
	e.Answer = a
}

func (e *EdgeBot) Reset() error {
	err := e.Bot.Reset()
	if err != nil {
		return err
	} else {
		return nil
	}
}

//func (e *EdgeBot) OnAnswer() {
//
//}
