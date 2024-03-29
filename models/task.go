package models

import (
	"time"

	"gorm.io/gorm"
)				
type Task struct {
	gorm.Model
	UserId			uint
	User            User
	Title 			string	`gorm:"not null"`
	Description		string
	DueDate       	time.Time
	Completed		bool `gorm:"default:false"`

}