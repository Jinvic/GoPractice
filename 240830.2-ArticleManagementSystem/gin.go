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
		page.GET("/create_article", func(ctx *gin.Context) {
			route.LoadHTMLGlob("./templates/*")
			ctx.HTML(http.StatusOK, "create_article.html", nil)
		})
		page.GET("/update_article/:article_id", func(ctx *gin.Context) {
			route.LoadHTMLGlob("./templates/*")
			articleIDstr := ctx.Param("article_id")
			articleID, _ := strconv.Atoi(articleIDstr)
			article := getArticle(uint(articleID))
			ctx.HTML(http.StatusOK, "update_article.html", gin.H{
				"article": article,
			})
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
		// article := getArticle(uint(articleID))
		article := visitArticleR(uint(articleID))
		ctx.HTML(http.StatusOK, "article.html", gin.H{
			"article": article,
		})
	})

	route.POST("/update_article", func(ctx *gin.Context) {

		articleIDstr := ctx.PostForm("id")
		articleIDint, _ := strconv.Atoi(articleIDstr)
		articleID := uint(articleIDint)
		upds := map[string]interface{}{
			"title":   ctx.PostForm("title"),
			"content": ctx.PostForm("content"),
			"author":  ctx.PostForm("author"),
		}
		updateArticle(articleID, upds)
		ctx.Redirect(http.StatusSeeOther, "/article/"+articleIDstr)
	})

	route.GET("/delete_article/:article_id/", func(ctx *gin.Context) {
		articleIDstr := ctx.Param("article_id")
		articleID, _ := strconv.Atoi(articleIDstr)
		deleteArticle(uint(articleID))
		ctx.Redirect(http.StatusSeeOther, "/all_articles")
	})

	route.GET("/popular_articles", func(ctx *gin.Context) {

	})

	route.Run(":8080")
}
