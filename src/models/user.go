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

// GetAllUser function
func (ur userRepository) GetAll() []User {
	users := []User{}
	ur.db.Find(&users)

	return users
}

// Create user function
func (ur userRepository) Create(user *User) User {

	ur.db.Create(&user)

	return *user
}

// FindByID
func (ur userRepository) FindByID(id string) (User, error) {
	user := User{}

	if err := ur.db.First(&user, "ID = ?", id).Error; err != nil {
		return user, err
	}

	return user, nil
}

// Update user
func (ur userRepository) Update(user *User, updatingData map[string]interface{}) User {
	ur.db.Model(&user).Update(updatingData)

	return *user
}

// Delete user
func (ur userRepository) Delete(user *User) error {
	return ur.db.Delete(user).Error
}
