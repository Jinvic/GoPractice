package routes

import (
	"blog-service/pkg/api/user"

	"github.com/gin-gonic/gin"
)

func registerUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", user.Register)
		userGroup.POST("/login", user.Login)
		userGroup.POST("/logout", user.Logout)
	}
}
