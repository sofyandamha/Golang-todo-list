package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sofyandamha/go-todo-list/models"
)

func Post(c *gin.Context) {
	var Post models.Post
	c.ShouldBindJSON(&Post)

	err := models.CreatePost(&Post)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// post := models.Post{
	// 	Title:     Post.Title,
	// 	Deskripsi: Post.Deskripsi,
	// }

	// savedUser, err := post.SavePost()

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusCreated, gin.H{"user": Post})
}
