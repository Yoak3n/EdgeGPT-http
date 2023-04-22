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
}
