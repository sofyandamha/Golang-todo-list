package models

import (
	"github.com/sofyandamha/go-todo-list/database"
	"gorm.io/gorm"
)

type Post struct {
	ID        uint32 `gorm:"primary_key"`
	Title     string `gorm:"size:255;not null" json:"title"`
	Deskripsi string `gorm:"size:255;not null" json:"deskripsi"`
	gorm.Model
}

func (Post Post) SavePost() (*Post, error) {
	err := database.Db.Create(&Post).Error

	if err != nil {
		return &Post, err
	}
	return &Post, nil
}

func CreatePost(Post *Post) (err error) {

	err = database.Db.Create(Post).Error

	if err != nil {
		return err
	}

	return nil
}

func GetPost(Post *Post, id int) (err error) {
	err = database.Db.Where("id = ?", id).First(Post).Error
	if err != nil {
		return err
	}
	return nil
}

// Update Post
func UpdatePost(Post *Post) (err error) {
	database.Db.Save(Post)
	return nil
}
