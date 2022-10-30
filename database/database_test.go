package database_test

import (
	"hmcalister/database"
	"hmcalister/models"
	"os"
	"testing"
	"time"
)

// SETUP DATABASE

func TestCreateDatabase(t *testing.T) {
	err := database.CreateDatabase("test_database.db")
	if err != nil {
		t.Fatalf("Failed to create Database!")
	}
}

// USER TESTING

func TestAddUser(t *testing.T) {
	var newUser models.User
	var err error
	newUser = models.User{Username: "Hayden"}
	err = database.CreateUser(newUser)
	if err != nil {
		t.Fatalf("Failed to create new User: %v", newUser)
	}
}

func TestGetAllUsers(t *testing.T) {
	_, err := database.GetAllUsers()
	if err != nil {
		t.Errorf("Failed to find all users!")
	}
}

func TestUpdateUserPassword(t *testing.T) {
	newPassword := "Password1234"
	err := database.UpdateUserPassword("Hayden", newPassword)
	if err != nil {
		t.Errorf("Failed to update password!")
	}
	user, _ := database.GetUserByUsername("Hayden")
	if newPassword != user.Password {
		t.Errorf("Password update failed to persist!")
	}
}

// STUDENT TESTING

func TestCreateStudent(t *testing.T) {
	newStudent := models.Student{
		StudentCode: "hayden000",
		FullName:    "Hayden McAlister",
	}

	err := database.CreateStudent(newStudent)
	if err != nil {
		t.Fatalf("Failed to create student!")
	}
}

func TestGetAllStudents(t *testing.T) {
	_, err := database.GetAllStudents()
	if err != nil {
		t.Errorf("Failed to get all students!")
	}
}

func TestGetStudentByStudentCode(t *testing.T) {
	fullName := "Hayden McAlister"
	student, err := database.GetStudentByStudentCode("hayden000")
	if err != nil {
		t.Errorf("Failed to get student by student code!")
	}

	if student.FullName != fullName {
		t.Errorf("Failed to get correct student by code!")
	}
}

// LAB TESTING

func TestCreateLab(t *testing.T) {
	newLab := models.Lab{
		ID:          1,
		LabName:     "Test Lab",
		Description: "Lab for testing",
		Points:      1,
	}

	err := database.CreateLab(newLab)
	if err != nil {
		t.Fatalf("Failed to create new lab!")
	}
}

func TestGetAllLabs(t *testing.T) {
	_, err := database.GetAllLabs()
	if err != nil {
		t.Errorf("Failed to find all labs!")
	}
}

func TestGetLabByLabID(t *testing.T) {
	labName := "Test Lab"
	resultLab, err := database.GetLabByLabID(1)
	if err != nil {
		t.Errorf("Failed to find lab by ID!")
	}
	if resultLab.LabName != labName {
		t.Errorf("Failed to find correct lab by ID!")
	}
}

// LAB COMPLETION

func TestCreateLabCompletion(t *testing.T) {
	newLabCompletion := models.LabCompletion{
		StudentCode: "hayden000",
		LabID:       1,
		Timestamp:   time.Now(),
	}

	err := database.CreateLabCompletion(newLabCompletion)
	if err != nil {
		t.Fatalf("Failed to create new lab completion!")
	}
}

func TestGetAllLabCompletions(t *testing.T) {
	_, err := database.GetAllLabCompletions()
	if err != nil {
		t.Fatalf("Failed to get all lab completions!")
	}
}

func TestGetAllLabCompletionsByStudentCode(t *testing.T) {
	labCompletions, err := database.GetAllLabCompletionsByStudentCode("hayden000")
	if err != nil {
		t.Fatalf("Failed to get lab completions by student code!")
	}
	if len(labCompletions) != 1 {
		t.Fatalf("Failed to get correct number of lab completions!")
	}
}

func TestGetLabCompletionsByStudentCodeAndLabID(t *testing.T) {
	labCompletion, err := database.GetLabCompletionByStudentCodeAndLabID("hayden000", 1)
	if err != nil {
		t.Fatalf("Failed to get lab completions by student code and lab ID!")
	}
	if labCompletion.LabID != 1 {
		t.Fatalf("Failed to get correct lab completion!")
	}
}

// DELETION / CLEANUP TESTS

func TestDeleteLabCompletionByStudentCodeAndLabID(t *testing.T) {
	err := database.DeleteLabCompletionByStudentCodeAndLabID("hayden000", 1)
	if err != nil {
		t.Fatalf("Failed to delete lab completion!")
	}
}

func TestDeleteLabByLabID(t *testing.T) {
	err := database.DeleteLabByLabID(1)
	if err != nil {
		t.Errorf("Failed to delete lab by ID!")
	}
}

func TestDeleteStudentByStudentCode(t *testing.T) {
	err := database.DeleteStudentByStudentCode("hayden000")
	if err != nil {
		t.Errorf("Failed to delete student!")
	}
}

func TestDeleteUser(t *testing.T) {
	err := database.DeleteUserByUsername("Hayden")
	if err != nil {
		t.Errorf("Failed to delete user!")
	}
}

func TestRemoveDatabaseFile(t *testing.T) {
	err := os.Remove("test_database.db")
	if err != nil {
		t.Errorf("Failed to remove database file!")
	}
}
