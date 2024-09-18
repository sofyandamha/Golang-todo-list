package main

import (
	"os"

	"github.com/sofyandamha/go-todo-list/database"
	"github.com/sofyandamha/go-todo-list/initializers"
	"github.com/sofyandamha/go-todo-list/models"
)

func main() {
	initializers.VarEnv()
	database.InitDb()
	database.Db.AutoMigrate(&models.Role{})
	database.Db.AutoMigrate(&models.User{})
	database.Db.AutoMigrate(&models.Post{})
	database.Db.AutoMigrate(&models.Task{})
	seedData()
}

// load seed data into the database
func seedData() {
	var roles = []models.Role{{Name: "admin", Description: "Administrator role"}, {Name: "customer", Description: "Authenticated customer role"}, {Name: "anonymous", Description: "Unauthenticated customer role"}}
	var user = []models.User{{
		Username: os.Getenv("ADMIN_USERNAME"),
		Email:    os.Getenv("ADMIN_EMAIL"),
		Password: os.Getenv("ADMIN_PASSWORD"),
		RoleID:   1}, {
		Username: "user",
		Email:    "user@example.com",
		Password: "123456",
		RoleID:   2}}
	var post = []models.Post{{Title: "test", Deskripsi: "Test juga"}}
	database.Db.Save(&roles)
	database.Db.Save(&user)
	database.Db.Save(&post)

}
