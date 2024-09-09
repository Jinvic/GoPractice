package main

import (
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	// 1. **创建一个简单的 REST API**
	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"title": "ping",
			"msg1":  "成功",
			"msg2":  "pong",
		})
	})
	route.POST("/ping", func(ctx *gin.Context) {
		var json struct {
			Title   string `json:"title"`
			Message string `json:"msg"`
		}

		if err := ctx.ShouldBindBodyWithJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"title": "ping",
			"msg":   "received messages: " + json.Message,
		})
	})

	route.Run(":8080")
}
