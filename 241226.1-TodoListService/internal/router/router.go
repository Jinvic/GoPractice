package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	InitTodoRouter(r)
	InitSwaggerRouter(r)
	return r
}
