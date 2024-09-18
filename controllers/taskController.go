package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sofyandamha/go-todo-list/models"
	"gorm.io/gorm"
)

func Task(c *gin.Context) {
	var Task models.Task
	c.ShouldBindBodyWithJSON(&Task)

	err := models.CreateTask(&Task)
	if err != nil {

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// task := models.Task{
	// 	Name:       Task.Name,
	// 	Due_date:   Task.Due_date,
	// 	Priority:   Task.Priority,
	// 	Status:     Task.Status,
	// 	Deskripsi:  Task.Deskripsi,
	// 	AssigneeID: Task.AssigneeID,
	// }

	// savedUser, err := task.SaveTask()

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"user": Task})

}

func GetTask(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var task models.Task
	err := models.GetTask(&task, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"Error": "Data Tidak Ditemukan"})
			return
		}

		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Ditemukan", "data": task})
}
