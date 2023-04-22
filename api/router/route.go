package router

import (
	"github.com/Yoak3n/EdgeGPT-http/internal/gpt"
	"github.com/gin-gonic/gin"
	"log"
)

func responseQuestion(c *gin.Context) {
	var question string
	q, ok := c.Get("question")
	if ok {
		question = q.(string)
		if question == "" {
			log.Println("blank question")
		}
	}

	bot, ok := c.Get("bot")
	if ok {
		b := bot.(*gpt.EdgeBot)
		switch question {
		case "reset":
			handleReset(b, c)
		default:
			handleAnswer(b, c, question)
		}
	} else {
		log.Fatalln("can't invoke any bot")
	}
}
