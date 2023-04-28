package router

// Created at 2023/4/10 14:48
// Created by Yoake
import (
	"github.com/Yoak3n/EdgeGPT-http/api/middleware"
	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func init() {
	R = gin.Default()
	R.Use(middleware.Cors())
	R.POST("/chat", middleware.RequestSource(), responseQuestion)
	R.GET("/search/:session", sessionSearch) // 好像没什么必要的查询功能
	R.GET("/query/:question", questionQuery) // 这个功能比上面这个重要多了
}
