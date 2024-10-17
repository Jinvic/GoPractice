package routes

import (
	"blog-service/pkg/api/user"
	"blog-service/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func registerUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", user.Register)
		userGroup.POST("/login", user.Login)

		userGroup.Use(middleware.AuthMiddleware())
		userGroup.POST("/logout", user.Logout)
		userGroup.DELETE("/delete", user.Delete)

		userGroup.Use(middleware.AdminMiddleware())
		userGroup.GET("/list", user.List)
	}
}
