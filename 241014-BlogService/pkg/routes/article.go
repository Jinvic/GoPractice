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

		authGroup := articleGroup.Group("/", middleware.AuthMiddleware())
		{
			authGroup.POST("/create", article.Create)
			authGroup.GET("/list", article.List)

			ownershipGroup := authGroup.Group("/", middleware.OwnershipMiddleware())
			{
				ownershipGroup.PUT("/edit/:id", article.Edit)
				ownershipGroup.DELETE("/delete/:id", article.Delete)
			}

			authGroup.GET("/list_all", article.ListAll, middleware.AdminMiddleware())
		}
	}
}
