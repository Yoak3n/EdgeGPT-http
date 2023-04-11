package middleware

import (
	"edgegpt-http/internal/gpt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Created at 2023/4/11 16:58
// Created by Yoake

func QueueInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Param("id")
	}
}
func RequestSource(source string) gin.HandlerFunc {
	return func(c *gin.Context) {
		bot := gpt.NewConv()
		c.Set("bot", bot)
		header := c.GetHeader("Request-Source")
		if header == source {
			c.Next()
		} else {
			log.Println(header)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request source"})
		}
	}
}

func RequestQuestion(question string) {

}
