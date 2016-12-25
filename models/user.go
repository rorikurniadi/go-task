package models

import (
	"github.com/jinzhu/gorm"
)

// User Model
type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Contact  string `json:"contact"`
	Address  string `json:"address"`
	Task     []Task `gorm:"ForeignKey:UserId"`
}
