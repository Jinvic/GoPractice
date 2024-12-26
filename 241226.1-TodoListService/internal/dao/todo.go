package dao

import (
	db "todolist/database"
	"todolist/internal/model"
	"todolist/pkg/e"
)

type todoDao struct{}

var TodoDao = &todoDao{}

func (d *todoDao) GetList() *[]model.Todo {
	list := make([]model.Todo, 0)
	db.RMu.RLock()
	defer db.RMu.RUnlock()
	for _, v := range db.DB {
		list = append(list, v)
	}
	return &list
}

func (d *todoDao) Create(todo *model.Todo) {
	db.Mu.Lock()
	defer db.Mu.Unlock()
	todo.ID = db.Idx
	db.DB[db.Idx] = *todo
	db.Idx++
}

func (d *todoDao) Get(id int) (*model.Todo, error) {
	db.RMu.RLock()
	defer db.RMu.RUnlock()
	todo, ok := db.DB[id]
	if !ok {
		return nil, e.DATA_NOT_EXIST
	}
	return &todo, nil
}

func (d *todoDao) Update(todo *model.Todo) error {
	db.Mu.Lock()
	defer db.Mu.Unlock()
	_, ok := db.DB[todo.ID]
	if !ok {
		return e.DATA_NOT_EXIST
	}
	db.DB[todo.ID] = *todo
	return nil
}

func (d *todoDao) Delete(id int) error {
	db.Mu.Lock()
	defer db.Mu.Unlock()
	_, ok := db.DB[id]
	if !ok {
		return e.DATA_NOT_EXIST
	}
	delete(db.DB, id)
	return nil
}
