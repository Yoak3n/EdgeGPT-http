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
		if style == "" {
			style = "bing-b"
		} else {
			err = errors.New(`please input ["bing-c","bing-b","bing-p"]rather than another style to create a conversation`)
		}
	}

	return
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-    Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func RequestSource() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "get raw data failed"})
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
	}
}
