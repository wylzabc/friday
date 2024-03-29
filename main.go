package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wylzabc/friday/add"
	"github.com/wylzabc/friday/multiplication"
	"github.com/wylzabc/friday/subtraction"
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
	router.POST("/add/add", add.Add)
	router.POST("/add/absadd", add.AbsAdd)
	router.POST("/multi/multi", multiplication.Multi)
	router.POST("/sub/sub", subtraction.Sub)
}
func main() {
	InitRouter()
	router.Run() // listen and serve on 0.0.0.0:8080
}
