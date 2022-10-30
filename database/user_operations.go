package database

import (
	"hmcalister/models"
)

func CreateUser(newUser models.User) error {
	result := conn.Create(&newUser)
	return result.Error
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := conn.Find(&users)
	return users, result.Error
}

func GetUserByUsername(Username string) (models.User, error) {
	var user models.User
	result := conn.First(&user, "Username = ?", Username)
	return user, result.Error
}

func UpdateUserPassword(Username string, newPassword string) error {
	result := conn.Model(&models.User{}).Where("Username = ?", Username).Update("Password", newPassword)
	return result.Error
}

func DeleteUserByUsername(Username string) error {
	result := conn.Delete(&models.User{}, "Username = ?", Username)
	return result.Error
}
