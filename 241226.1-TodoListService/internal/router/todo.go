package router

import (
	"todolist/internal/api"

	"github.com/gin-gonic/gin"
)

func InitTodoRouter(r *gin.Engine) {
	todo := r.Group("/todo")
	{
		todo.GET("/list", api.TodoApi.GetList)
		todo.POST("/", api.TodoApi.Create)
		todo.GET("/:id", api.TodoApi.Get)
		todo.PUT("/", api.TodoApi.Update)
		todo.DELETE("/:id", api.TodoApi.Delete)
	}
}
