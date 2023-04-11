package router

// Created at 2023/4/10 14:48
// Created by Yoake
import (
	"edgegpt-http/api/middleware"
	"edgegpt-http/edgegpt"
	"edgegpt-http/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"log"
	"net/http"
)

var R *gin.Engine

func responseQuestion(c *gin.Context) {
	b, _ := c.GetRawData()
	r := gjson.Parse(string(b))
	session := r.Get("name").String()
	question := r.Get("question").String()
	log.Println(session)
	bot, ok := c.Get("bot")
	if ok {
		b := bot.(*edgegpt.ChatBot)
		answer, err := b.Ask(question, edgegpt.Balanced)
		if err != nil {
			log.Fatalln("edgegpt no answer")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"messages": utils.FormatAnswer(answer),
			})
		}
	}
}
func init() {
	R = gin.Default()
	R.POST("/chat", middleware.RequestSource("127.0.0.1"), responseQuestion)
}
