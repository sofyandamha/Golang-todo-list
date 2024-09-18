package models

import (
	"github.com/sofyandamha/go-todo-list/database"
	"gorm.io/gorm"
)

type Assignee struct {
	ID   uint32 `gorm:"primary_key"`
	Name string `gorm:"size:255; not null"`
	gorm.Model
}

func (Assignee Assignee) SaveAsg() (*Assignee, error) {
	err := database.Db.Create(&Assignee).Error

	if err != nil {
		return &Assignee, err
	}
	return &Assignee, nil
}

func CreateAssg(Assignee *Assignee) (err error) {

	err = database.Db.Create(Assignee).Error

	if err != nil {
		return err
	}

	return nil
}

func GetAss(Assignee *Assignee, id int) (err error) {
	err = database.Db.Where("id = ?", id).First(Assignee).Error
	if err != nil {
		return err
	}
	return nil
}

// Update Assignee
func UpdateAss(Assignee *Assignee) (err error) {
	database.Db.Save(Assignee)
	return nil
}
