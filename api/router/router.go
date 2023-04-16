package router

// Created at 2023/4/10 14:48
// Created by Yoake
import (
	"fmt"
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
	style, _ := c.Get("style")
	bot, ok := c.Get("bot")
	if ok {
		b := bot.(*gpt.EdgeBot)
		b.OnQuestion(question)
		answer := utils.FormatAnswer(b.Answer)
		log.Printf("%sRecived :%s", b.Session, answer)
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"style":   style,
			"message": b.Answer.Text(),
			"count":   fmt.Sprintf("%d / %d", b.Answer.NumUserMessages(), b.Answer.MaxNumUserMessages()),
		})
	} else {
		log.Fatalln("can't invoke any bot")
	}
}

func init() {
	R = gin.Default()
	R.Use(middleware.Cors())
	R.POST("/chat", middleware.RequestSource(), responseQuestion)
}
