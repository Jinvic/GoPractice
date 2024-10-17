package routes

import (
	"blog-service/pkg/api/article"
	"blog-service/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func registerArticleRoutes(router *gin.Engine) {
	articleGroup := router.Group("/article")
	{
		articleGroup.GET("/view/:id", article.View)

		articleGroup.Use(middleware.AuthMiddleware())
		articleGroup.POST("/create", article.Create)
		articleGroup.GET("/list", article.List)
		{
			articleGroup.Use(middleware.OwnershipMiddleware())
			articleGroup.PUT("/edit/:id", article.Edit)
			articleGroup.DELETE("/delete/:id", article.Delete)
		}
		articleGroup.Use(middleware.AdminMiddleware())
		articleGroup.GET("/list_all", article.ListAll)
	}
}
