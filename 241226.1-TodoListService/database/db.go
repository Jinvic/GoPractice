package database

import (
	"todolist/internal/model"
)

var DB = make(map[int]model.Todo)
