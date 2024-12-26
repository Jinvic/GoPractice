package database

import (
	"sync"
	"todolist/internal/model"
)

var Mu sync.Mutex
var RMu sync.RWMutex
var DB = make(map[int]model.Todo)
var Idx = 1
