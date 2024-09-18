package main

import (
	"github.com/sofyandamha/go-todo-list/database"
	"github.com/sofyandamha/go-todo-list/initializers"
	"github.com/sofyandamha/go-todo-list/routes"
)

func main() {
	initializers.VarEnv()
	database.InitDb()
	routes.Router()

}
