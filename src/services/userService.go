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

//UpdateUser func
func UpdateUser(id string, updatingData map[string]interface{}) (models.User, error) {
	user, err := models.UserRepository.FindByID(id)

	if err != nil {
		return user, err
	}

	models.UserRepository.Update(&user, updatingData)
	return user, nil
}

//DeleteUser func
func DeleteUser(id string) error {
	user, err := models.UserRepository.FindByID(id)

	if err != nil {
		return err
	}

	models.UserRepository.Delete(&user)
	return nil
}
