package database

import (
	"errors"
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

func UpdateStudentByStudentCode(studentCode string, updatedStudent models.Student) error {
	if updatedStudent.StudentCode != studentCode {
		return errors.New("cannot change StudentCode")
	}
	// Save for SOME reason doesn't correctly update with string primary key
	result := conn.Model(&updatedStudent).Select("*").Updates(updatedStudent)

	return result.Error
}

func DeleteStudentByStudentCode(StudentCode string) error {
	result := conn.Delete(&models.Student{}, "StudentCode = ?", StudentCode)
	return result.Error
}
