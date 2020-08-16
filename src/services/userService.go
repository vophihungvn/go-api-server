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
func CreateUser(user models.User) models.User {

	return models.UserRepository.Create(&user)
}
