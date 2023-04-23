package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Yoak3n/EdgeGPT-http/config"
	"github.com/Yoak3n/EdgeGPT-http/internal/database"
	"github.com/Yoak3n/EdgeGPT-http/internal/gpt"
	"github.com/Yoak3n/EdgeGPT-http/internal/model"
	"github.com/Yoak3n/EdgeGPT-http/pkg/utils"
)

// Created at 2023/4/13 3:12
// Created by Yoake
const (
	success = "success"
	failed  = "failed"
)

func handleAnswer(b *gpt.EdgeBot, c *gin.Context, question string) {
	session, _ := c.Get("session")
	style, _ := c.Get("style")
	if b.End {
		err := b.Reset()
		if err != nil {
			log.Println(err)
		} else {
			log.Println("already reset")
		}
	}

	b.OnQuestion(question)
	answer := utils.FormatAnswer(b.Answer)
	log.Printf("%s Recived:%s", b.Session, answer)

	current := b.Answer.NumUserMessages()
	max := b.Answer.MaxNumUserMessages()
	if current >= max {
		b.End = true
	} else {
		b.End = false
	}
	count := make(map[string]any)
	count["currentNum"] = current
	count["maxNum"] = max
	c.JSON(http.StatusOK, gin.H{
		"status":  success,
		"style":   style,
		"message": b.Answer.Text(),
		"count":   count,
	})
	if config.Preset.Mysql.DBName != "" {
		database.CreateMessage(question, b.Answer.Text(), session.(string))
	}
}

func handleReset(b *gpt.EdgeBot, c *gin.Context) {
	err := b.Reset()
	if err != nil {
		log.Println("reset failed")
		c.JSON(http.StatusFailedDependency, gin.H{
			"status": failed,
			"error":  "reset failed",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  success,
			"message": "reset successfully",
		})
	}
}

func handleQuery(result []model.Data, c *gin.Context) {
	var messages []model.Message
	for _, item := range result {
		messages = append(messages, item.Message)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  success,
		"results": messages,
	})
}
