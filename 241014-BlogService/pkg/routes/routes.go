package routes

import "github.com/gin-gonic/gin"

func InitRoutes(router *gin.Engine) {
	registerUserRoutes(router)
	registerArticleRoutes(router)
}
