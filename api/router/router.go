package router

// Created at 2023/4/10 14:48
// Created by Yoake
import (
	"github.com/Yoak3n/EdgeGPT-http/api/middleware"
	"github.com/Yoak3n/EdgeGPT-http/internal/gpt"
	"github.com/Yoak3n/EdgeGPT-http/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var R *gin.Engine

func responseQuestion(c *gin.Context) {
	var question string
	q, ok := c.Get("question")
	if ok {
		question = q.(string)
		if question == "" {
			log.Fatal("blank question")
		}
	}
	bot, ok := c.Get("bot")
	style, _ := c.Get("style")
	if ok {
		b := bot.(*gpt.EdgeBot)
		b.OnQuestion(question)
		answer := utils.FormatAnswer(b.Answer)
		log.Printf("%sRecived :%s", b.Session, answer)
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"style":   style,
			"message": answer,
		})
	} else {
		log.Fatalln("can't invoke any bot")
	}
}

func init() {
	R = gin.Default()
	R.POST("/chat", middleware.RequestSource("0.0.0.0"), responseQuestion)
}
