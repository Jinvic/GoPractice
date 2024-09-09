package main

import (
	_ "fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.LoadHTMLGlob("./templates/*")
	route.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

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

	// 2. **处理静态文件**
	route.Static("/static", "./static")
	// route.LoadHTMLFiles("./templates/gin_doc.html")
	route.GET("/gin_doc", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "gin_doc.html", nil)
	})

	// 3. **模板渲染**
	route.GET("/time", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "time.html", gin.H{
			"title": "当前时间",
			"time":  time.Now().String(),
		})
	})

	route.Run(":8080")
}
