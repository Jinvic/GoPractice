package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()

	route.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusSeeOther, "/index")
	})

	route.GET("/index", func(ctx *gin.Context) {
		route.LoadHTMLGlob("./templates/*")
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	page := route.Group("/page")
	{
		page.GET("/:page_name", func(ctx *gin.Context) {
			route.LoadHTMLGlob("./templates/*")
			page_name := ctx.Param("page_name")
			ctx.HTML(http.StatusOK, page_name+".html", nil)
		})
	}

	route.POST("/create_article", func(ctx *gin.Context) {
		article := Article{
			Title:   ctx.PostForm("title"),
			Content: ctx.PostForm("content"),
			Author:  ctx.PostForm("author"),
		}
		createArticle(&article)
		ctx.Redirect(http.StatusSeeOther, "/index")
	})

	route.GET("/all_articles", func(ctx *gin.Context) {
		route.LoadHTMLGlob("./templates/*")
		articles := getAllArticles()
		ctx.HTML(http.StatusOK, "all_articles.html", gin.H{
			"articles": articles,
		})
	})

	route.GET("/article/:article_id", func(ctx *gin.Context) {
		route.LoadHTMLGlob("./templates/*")
		articleIDstr := ctx.Param("article_id")
		articleID, _ := strconv.Atoi(articleIDstr)
		article := getArticle(uint(articleID))
		ctx.HTML(http.StatusOK, "article.html", gin.H{
			"article": article,
		})
	})

	route.GET("/update_article", func(ctx *gin.Context) {

	})
	route.GET("/delete_article", func(ctx *gin.Context) {

	})
	route.GET("/popular_articles", func(ctx *gin.Context) {

	})

	route.Run(":8080")
}
