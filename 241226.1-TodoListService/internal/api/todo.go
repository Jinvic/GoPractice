package api

import (
	"todolist/internal/define"
	"todolist/internal/service"
	"todolist/pkg/e"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type todoApi struct{}

var TodoApi = &todoApi{}

func (a *todoApi) GetList(c *gin.Context) {

	list := service.TodoService.GetList()
	SuccessData(c, define.GetListRsp{
		List:  *list,
		Total: len(*list),
	})
}

func (a *todoApi) Create(c *gin.Context) {
	todo := define.Todo{}
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		FailError(c, e.PARAM_ERROR)
		return
	}

	service.TodoService.Create(todo.Title, todo.Description)
	Success(c)
}

func (a *todoApi) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := cast.ToIntE(idStr)
	if err != nil {
		FailError(c, e.PARAM_ERROR)
		return
	}

	todo, err := service.TodoService.Get(id)
	if err != nil {
		FailError(c, err)
		return
	}
	SuccessData(c, todo)
}

func (a *todoApi) Update(c *gin.Context) {
	todo := define.Todo{}
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		FailError(c, e.PARAM_ERROR)
		return
	}

	err = service.TodoService.Update(todo.ID, todo.Title, todo.Description)
	if err != nil {
		FailError(c, err)
		return
	}
	Success(c)
}

func (a *todoApi) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := cast.ToIntE(idStr)
	if err != nil {
		FailError(c, e.PARAM_ERROR)
		return
	}

	err = service.TodoService.Delete(id)
	if err != nil {
		FailError(c, err)
		return
	}
	Success(c)
}

func (a *todoApi) Complete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := cast.ToIntE(idStr)
	if err != nil {
		FailError(c, e.PARAM_ERROR)
		return
	}

	err = service.TodoService.Complete(id)
	if err != nil {
		FailError(c, err)
		return
	}
	Success(c)
}
