package models

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	BaseModel
	Username string `json:"username"`
	Password string `json:"password"`
}

type userRepository struct {
	db *gorm.DB
}

// Create user function
func (ur userRepository) Create(user *User) User {

	ur.db.Create(&user)

	return *user
}

// GetAllUser function
func (ur userRepository) GetAll() []User {
	users := []User{}
	ur.db.Find(&users)

	return users
}
