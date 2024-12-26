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

//	@Summary		获取待办事项列表
//	@Description	获取所有待办事项的列表
//	@Tags			Todo
//	@Success		200	{object}	define.CommonSeccessDataRsp{data=define.GetListRsp}
//	@Failure		500	{object}	define.CommonFailRsp
//	@Router			/todo/list [get]
func (a *todoApi) GetList(c *gin.Context) {

	list := service.TodoService.GetList()
	SuccessData(c, define.GetListRsp{
		List:  *list,
		Total: len(*list),
	})
}

//	@Summary		创建待办事项
//	@Description	创建一个新的待办事项
//	@Tags			Todo
//	@Param			todo	body		define.Todo	true	"待办事项信息"
//	@Success		200		{object}	define.CommonSuccessRsp
//	@Failure		500		{object}	define.CommonFailRsp
//	@Router			/todo [post]
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

//	@Summary		获取待办事项
//	@Description	根据ID获取待办事项
//	@Tags			Todo
//	@Param			id	path		int	true	"待办事项ID"
//	@Success		200	{object}	define.CommonSeccessDataRsp{data=define.Todo}
//	@Failure		500	{object}	define.CommonFailRsp
//	@Router			/todo/{id} [get]
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

//	@Summary		更新待办事项
//	@Description	更新待办事项的信息
//	@Tags			Todo
//	@Param			todo	body		define.Todo	true	"待办事项信息"
//	@Success		200		{object}	define.CommonSuccessRsp
//	@Failure		500		{object}	define.CommonFailRsp
//	@Router			/todo [put]
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

//	@Summary		删除待办事项
//	@Description	根据ID删除待办事项
//	@Tags			Todo
//	@Param			id	path		int	true	"待办事项ID"
//	@Success		200	{object}	define.CommonSuccessRsp
//	@Failure		500	{object}	define.CommonFailRsp
//	@Router			/todo/{id} [delete]
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

//	@Summary		完成待办事项
//	@Description	根据ID标记待办事项为完成
//	@Tags			Todo
//	@Param			id	path		int	true	"待办事项ID"
//	@Success		200	{object}	define.CommonSuccessRsp
//	@Failure		500	{object}	define.CommonFailRsp
//	@Router			/todo/{id} [patch]
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
