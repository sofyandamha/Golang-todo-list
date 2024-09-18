package models

import (
	"time"

	"github.com/sofyandamha/go-todo-list/database"
	"gorm.io/gorm"
)

type Task struct {
	ID         uint      `gorm:"primary_key"`
	Name       string    `gorm:"size:255; not null" json:"name"`
	Due        time.Time `gorm:"column:due_date" json:"-"`
	Priority   string    `gorm:"size:255;not null" json:"priority"`
	Status     string    `gorm:"size:255;not null" json:"status"`
	Deskripsi  string    `gorm:"size:255;not null" json:"deskripsi"`
	AssigneeID int
	Assignee   Assignee `gorm:"ForeignKey:ID"`
	gorm.Model
}

func (Task Task) SaveTask() (*Task, error) {
	err := database.Db.Create(&Task).Error

	if err != nil {
		return &Task, err
	}
	return &Task, nil
}

func CreateTask(Task *Task) (err error) {

	err = database.Db.Create(Task).Error

	if err != nil {
		return err
	}

	return nil
}

func GetTask(Task *Task, id int) (err error) {
	err = database.Db.Where("id = ?", id).First(Task).Error
	if err != nil {
		return err
	}
	return nil
}

// Update Task
func UpdateTask(Task *Task) (err error) {
	database.Db.Save(Task)
	return nil
}
