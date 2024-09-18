package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sofyandamha/go-todo-list/controllers"
	"github.com/sofyandamha/go-todo-list/middleware"
)

func Router() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World Golang nih brey",
		})
	})
	authRoutes := router.Group("/auth/v1")
	// registration route
	authRoutes.POST("/register", controllers.Register)
	// login route
	authRoutes.POST("/login", controllers.Login)

	adminRoutes := router.Group("/admin")

	adminRoutes.POST("/users/task", controllers.Task)
	adminRoutes.GET("/task/:id", controllers.GetTask)

	adminRoutes.Use(middleware.JWTAuth())
	adminRoutes.GET("/users", controllers.GetUsers)
	adminRoutes.GET("/user/:id", controllers.GetUser)
	adminRoutes.PUT("/user/:id", controllers.UpdateUser)
	adminRoutes.POST("/user/role", controllers.CreateRole)
	adminRoutes.GET("/user/roles", controllers.GetRoles)
	adminRoutes.PUT("/user/role/:id", controllers.UpdateRole)

	router.Run(":3035")
}
