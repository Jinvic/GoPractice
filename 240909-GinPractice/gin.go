package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	// 5. **使用中间件**
	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	route.Use(saveRequest)

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

	// 4. **表单提交**
	route.GET("/form", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "form.html", nil)
	})
	route.POST("/post_form", func(ctx *gin.Context) {
		field1 := ctx.PostForm("field1")
		field2 := ctx.PostForm("field2")
		field3 := ctx.PostForm("field3")

		ctx.JSON(http.StatusOK, gin.H{
			"field1": field1,
			"field2": field2,
			"field3": field3,
		})
	})

	// z1. **文件上传**
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	route.MaxMultipartMemory = 8 << 20 // 8 MiB
	route.GET("/file", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "file.html", nil)
	})
	route.POST("/upload_file", func(ctx *gin.Context) {
		var fileList []string
		// 单文件
		singleFile, err := ctx.FormFile("file")
		if err == nil {
			dst := "./files/" + singleFile.Filename
			// 上传文件至指定的完整文件路径
			ctx.SaveUploadedFile(singleFile, dst)
			fileList = append(fileList, singleFile.Filename)
		}

		// Multipart form
		form, err := ctx.MultipartForm()
		if err == nil {
			files := form.File["files"]
			for _, file := range files {
				// 上传文件至指定目录
				dst := "./files/" + file.Filename
				ctx.SaveUploadedFile(file, dst)
				fileList = append(fileList, file.Filename)
			}
		}

		res := "uploaded files:\n"
		if len(fileList) > 0 {
			for _, filename := range fileList {
				res += filename + "\n"
			}
		} else {
			res = "no file uploaded."
		}
		ctx.String(http.StatusOK, res)
	})

	route.Run(":8080")
}

// 5. **使用中间件**
func saveRequest(ctx *gin.Context) {
	ctx.Next()
	// //请求耗时
	// cost := time.Since(start).Milliseconds()
	// //响应状态码
	// responseStatus := c.Writer.Status()
	// //响应 header
	// responseHeader := c.Writer.Header()
	// //响应体大小
	// responseBodySize := c.Writer.Size()
	// //响应体 body
	// responseBody := writer.b.String()

	method := ctx.Request.Method
	url := ctx.Request.URL.String()
	status := strconv.Itoa(ctx.Writer.Status())
	time := time.Now().Format("2006-01-02 15:04:05")

	//OpenFile读取文件，不存在时则创建，使用追加模式
	file, err := os.OpenFile("request.csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println("文件打开失败！")
	}
	defer file.Close()

	//创建写入接口
	WriterCsv := csv.NewWriter(file)
	str := []string{method, url, status, time} //需要写入csv的数据，切片类型

	//写入一条数据，传入数据为切片(追加模式)
	err = WriterCsv.Write(str)
	if err != nil {
		log.Println("WriterCsv写入文件失败")
	}
	WriterCsv.Flush() //刷新，不刷新是无法写入的
	log.Println("数据写入成功...")
}
