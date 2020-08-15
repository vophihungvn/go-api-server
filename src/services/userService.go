package services

import (
	models "models"
)

//GetUserList func
func GetUserList() []models.User {
	users := models.UserRepository.GetAll()
	return users
}

// CreateUser func
func CreateUser(username string, password string) models.User {
	user := models.User{Username: username, Password: password}
	user.Create()
	return user
}
