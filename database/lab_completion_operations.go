package database

import (
	"hmcalister/models"
)

func CreateLabCompletion(newLabCompletion models.LabCompletion) error {
	result := conn.Create(&newLabCompletion)
	return result.Error
}

func GetAllLabCompletions() ([]models.LabCompletion, error) {
	var LabCompletions []models.LabCompletion
	result := conn.Find(&LabCompletions)
	return LabCompletions, result.Error
}

func GetAllLabCompletionsByStudentCode(StudentCode string) ([]models.LabCompletion, error) {
	var LabCompletions []models.LabCompletion
	result := conn.Find(&LabCompletions, "StudentCode = ?", StudentCode)
	return LabCompletions, result.Error
}

func GetLabCompletionByStudentCodeAndLabID(StudentCode string, LabID int) (models.LabCompletion, error) {
	var LabCompletion models.LabCompletion
	result := conn.First(&LabCompletion, "StudentCode = ? AND LabID = ?", StudentCode, LabID)
	return LabCompletion, result.Error
}

func DeleteLabCompletionByStudentCodeAndLabID(StudentCode string, LabID int) error {
	result := conn.Delete(&models.LabCompletion{}, "StudentCode = ? AND LabID = ?", StudentCode, LabID)
	return result.Error
}
