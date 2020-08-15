package models

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type userRepository struct {
	db *gorm.DB
}

// Create user function
func (u User) Create() {
	dbInstance.Save(&u)
}

// GetAllUser function
func (ur userRepository) GetAll() []User {
	users := []User{}
	ur.db.Find(&users)

	return users
}
