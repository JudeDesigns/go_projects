package models

import "gorm.io/gorm" //For this line to work, you first need to have initialized your go mod and then run "go get gorm.io/gorm"

type ToDoList struct {
	gorm.Model         //ID, CreatedAt, UpdatedAt, DeletedAt
	Title       string `json:"title" gorm:"not null"`          //Title of the TODo item, cannot be null
	Description string `json:"description"`                    //Description of the ToDo item
	Completed   bool   `json:"completed" gorm:"default:false"` //Whether the ToDo item is completed or not
}

type ToDoRepository struct {
	DB *gorm.DB
}
