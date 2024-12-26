package service

import (
	"time"
	"todolist/internal/dao"
	"todolist/internal/define"
	"todolist/internal/model"
)

type todoService struct{}

var TodoService = &todoService{}

func (s *todoService) GetList() *[]define.TodoListItem {
	list := make([]define.TodoListItem, 0)
	todos := dao.TodoDao.GetList()
	for _, v := range *todos {
		list = append(list, define.TodoListItem{
			ID:        v.ID,
			Title:     v.Title,
			Completed: v.Completed,
		})
	}
	return &list
}

func (s *todoService) Create(title, description string) {
	now := time.Now()
	todo := model.Todo{
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	dao.TodoDao.Create(&todo)
}

func (s *todoService) Get(id int) (*define.Todo, error) {
	todo, err := dao.TodoDao.Get(id)
	if err != nil {
		return nil, err
	}
	return &define.Todo{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   todo.UpdatedAt.Format("2006-01-02 15:04:05"),
		CompletedAt: todo.CompletedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *todoService) Update(id int, title, description string) error {
	todo, err := dao.TodoDao.Get(id)
	if err != nil {
		return err
	}

	now := time.Now()
	todo.Title = title
	todo.Description = description
	todo.UpdatedAt = now
	return dao.TodoDao.Update(todo)
}

func (s *todoService) Delete(id int) error {
	return dao.TodoDao.Delete(id)
}

func (s *todoService) Complete(id int) error {
	todo, err := dao.TodoDao.Get(id)
	if err != nil {
		return err
	}

	now := time.Now()
	todo.Completed = true
	todo.CompletedAt = now
	return dao.TodoDao.Update(todo)
}
