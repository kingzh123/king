package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `gorm:"size:255" json:"name"`
	Age   uint   `json:"age"`
	Email string
}

func (User) TableName() string {
	return "users"
}
