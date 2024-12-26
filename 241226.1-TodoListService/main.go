package main

import (
	"todolist/internal/router"
)

func main() {
	
	r := router.InitRouter()
	r.Run()
}
