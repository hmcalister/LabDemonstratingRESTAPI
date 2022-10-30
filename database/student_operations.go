package database

import (
	"hmcalister/models"
)

func CreateStudent(newStudent models.Student) error {
	result := conn.Create(&newStudent)
	return result.Error
}

func GetAllStudents() ([]models.Student, error) {
	var Students []models.Student
	result := conn.Find(&Students)
	return Students, result.Error
}

func GetStudentByStudentCode(StudentCode string) (models.Student, error) {
	var Student models.Student
	result := conn.First(&Student, "StudentCode = ?", StudentCode)
	return Student, result.Error
}

func DeleteStudentByStudentCode(StudentCode string) error {
	result := conn.Delete(&models.Student{}, "StudentCode = ?", StudentCode)
	return result.Error
}
