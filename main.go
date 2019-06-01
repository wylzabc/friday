package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	//gin.SetMode(gin.DebugMode)
	router = gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
func main() {
	InitRouter()
	router.Run() // listen and serve on 0.0.0.0:8080
}
