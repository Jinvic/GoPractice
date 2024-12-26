package service

import "todolist/internal/define"

type todoService struct{}

var TodoService = &todoService{}

func (s *todoService) GetList() *[]define.TodoListItem {
	return &[]define.TodoListItem{}
}
func (s *todoService) Create(title, description string) {}
func (s *todoService) Get(id int) (define.Todo, error) {
	return define.Todo{}, nil
}
func (s *todoService) Update(id int, title, description string) error {
	return nil
}
func (s *todoService) Delete(id int) error {
	return nil
}
