package router

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Yoak3n/EdgeGPT-http/internal/database"
	"github.com/Yoak3n/EdgeGPT-http/internal/gpt"
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

func sessionQuery(c *gin.Context) {
	query := c.Param("session")
	result := database.GetSomeoneAllMessages(query)
	handleQuery(result, c)
}
