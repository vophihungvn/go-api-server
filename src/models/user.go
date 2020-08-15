package models

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
