package middleware

import (
	"errors"
	"github.com/Yoak3n/EdgeGPT-http/internal/gpt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"log"
	"net/http"
)

// Created at 2023/4/11 16:58
// Created by Yoake

var Pool gpt.BotPool

func init() {
	Pool.Workers = make(map[string]*gpt.EdgeBot)
}

func checkInput(r *gjson.Result) (session string, style string, question string, err error) {
	session = r.Get("name").String()
	question = r.Get("question").String()
	style = r.Get("style").String()

	if style != "bing-c" && style != "bing-b" && style != "bing-p" {
		err = errors.New(`please input ["bing-c","bing-b","bing-p"]rather than another style`)
	}

	return
}

func RequestSource(source string) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "get rawdata failed"})
		}
		r := gjson.Parse(string(data))
		session, style, question, err := checkInput(&r)
		log.Printf("%s Send:%s", session, question)
		if err != nil {
			c.Set("error", err)
		}
		c.Set("question", question)
		c.Set("style", style)

		// 查找内存连接池中是否存在相应对话
		value, ok := Pool.Workers[session]
		if ok {
			c.Set("bot", value)
			c.Next()
		} else {
			bot := gpt.NewBot(session, style)
			c.Set("bot", bot)
			Pool.Workers[session] = bot
			c.Next()
		}
		header := c.GetHeader("Request-Source")
		if header == source {
			c.Next()
		} else {
			log.Println(header)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request source"})
		}
	}
}
