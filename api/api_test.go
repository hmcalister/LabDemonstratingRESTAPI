package api_test

// Note! These tests rely heavily on the tests in the database module passing
// Database module functions are called with the expectation of working correctly!

import (
	"bytes"
	"encoding/json"
	"hmcalister/api"
	"hmcalister/database"
	"hmcalister/models"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

const API_URL_PREFIX string = "http://localhost:8080"

func TestStartAPI(t *testing.T) {
	os.Remove("./test_database.db")
	database.CreateDatabase("./test_database.db")

	gin.DefaultWriter, _ = os.Create("api_test.log")
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1", "[::1]"})
	api.SetupAPI(router)

	go router.Run("localhost:8080")

}

// Student tests

func TestCreateStudent(t *testing.T) {
	var err error
	postBody, _ := json.Marshal(map[string]string{
		"StudentCode": "alice123",
		"FullName":    "Alice",
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(API_URL_PREFIX+"/students", "application/json", responseBody)
	if err != nil {
		t.Fatalf("Failed to create student! %v", err)
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to parse response of creating student! %v", err)
	}
}

func TestGetAllStudents(t *testing.T) {
	resp, err := http.Get(API_URL_PREFIX + "/students")
	if err != nil {
		t.Errorf("Failed to get all students! %v", err)
	}
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Failed to parse response body! %v", err)
	}
}

func TestGetStudentByStudentCode(t *testing.T) {
	resp, err := http.Get(API_URL_PREFIX + "/students/alice123")
	if err != nil {
		t.Errorf("Failed to get student! %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Failed to parse response body! %v", err)
	}
	respStudent := models.Student{}
	err = json.Unmarshal([]byte(string(body)), &respStudent)
	if err != nil || respStudent.StudentCode != "alice123" {
		t.Errorf("Failed to get correct student! %v", err)
	}
}

// Lab test

func TestCreateLab(t *testing.T) {
	var err error
	postBody, _ := json.Marshal(map[string]string{
		"LabName":     "Test Lab",
		"Description": "A lab for testing",
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(API_URL_PREFIX+"/labs", "application/json", responseBody)
	if err != nil {
		t.Fatalf("Failed to create lab! %v", err)
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to parse response of creating lab! %v", err)
	}
}

func TestGetAllLabs(t *testing.T) {
	resp, err := http.Get(API_URL_PREFIX + "/labs")
	if err != nil {
		t.Errorf("Failed to get all labs! %v", err)
	}
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Failed to parse response body! %v", err)
	}
}

func TestGetLabByLabID(t *testing.T) {
	resp, err := http.Get(API_URL_PREFIX + "/labs/1")
	if err != nil {
		t.Errorf("Failed to get lab! %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Failed to parse response body! %v", err)
	}
	respLab := models.Lab{}
	err = json.Unmarshal([]byte(string(body)), &respLab)
	if err != nil || respLab.LabName != "Test Lab" {
		t.Errorf("Failed to get correct lab! %v", err)
	}
}

func TestUpdateLabByLabID(t *testing.T) {
	updatedLab := models.Lab{
		Points: 1,
	}
	jsonLab, _ := json.Marshal(&updatedLab)
	responseBody := bytes.NewBuffer(jsonLab)

	req, err := http.NewRequest("PUT",
		API_URL_PREFIX+"/labs/1",
		responseBody)
	if err != nil {
		t.Errorf("Failed to create PUT request! %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("Failed to execute PUT request! %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Failed to correctly update lab! %v", err)
	}
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to parse response of updating lab! %v", err)
	}

	updatedLab, err = database.GetLabByLabID(1)
	if err != nil || updatedLab.Points != 1 {
		t.Errorf("Failed to persist lab update to database! %v", err)
	}
}

// Lab Completion Tests

func TestCreateLabCompletion(t *testing.T) {
	var err error
	newLabCompletion := models.LabCompletion{
		StudentCode: "alice123",
		LabID:       1,
	}
	jsonLab, _ := json.Marshal(&newLabCompletion)
	responseBody := bytes.NewBuffer(jsonLab)

	resp, err := http.Post(API_URL_PREFIX+"/labCompletions", "application/json", responseBody)
	if err != nil {
		t.Fatalf("Failed to create lab completion! %v", err)
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to parse response of creating lab completion! %v", err)
	}
}

func TestGetAllLabCompletions(t *testing.T) {
	resp, err := http.Get(API_URL_PREFIX + "/labCompletions")
	if err != nil {
		t.Errorf("Failed to get all labs! %v", err)
	}
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Failed to parse response body! %v", err)
	}
}

func TestGetAllLabCompletionsByStudentCode(t *testing.T) {
	resp, err := http.Get(API_URL_PREFIX + "/labCompletions/alice123")
	if err != nil {
		t.Errorf("Failed to get lab completions by student code! %v", err)
	}
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Failed to parse response body! %v", err)
	}
}

func TestGetLabCompletionByStudentCodeAndLabID(t *testing.T) {
	resp, err := http.Get(API_URL_PREFIX + "/labCompletions/alice123/1")
	if err != nil {
		t.Errorf("Failed to get lab completions by student code and lab ID! %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Failed to parse response body! %v", err)
	}
	respLabCompletion := models.LabCompletion{}
	err = json.Unmarshal([]byte(string(body)), &respLabCompletion)
	if err != nil ||
		respLabCompletion.StudentCode != "alice123" ||
		respLabCompletion.LabID != 1 {
		t.Errorf("Failed to get correct lab! %v", err)
	}
}

// DELETE tests

func TestDeleteLabCompletion(t *testing.T) {
	req, err := http.NewRequest("DELETE",
		API_URL_PREFIX+"/labCompletions/alice123/1",
		nil)
	if err != nil {
		t.Errorf("Failed to create DELETE request! %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("Failed to execute delete request! %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Failed to correctly delete lab completion! %v", err)
	}
}

func TestDeleteStudentByStudentCode(t *testing.T) {
	req, err := http.NewRequest("DELETE",
		API_URL_PREFIX+"/students/alice123",
		nil)
	if err != nil {
		t.Errorf("Failed to create DELETE request! %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("Failed to execute delete request! %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Failed to correctly delete student! %v", err)
	}
}

func TestDeleteLabByLabID(t *testing.T) {
	req, err := http.NewRequest("DELETE",
		API_URL_PREFIX+"/labs/1",
		nil)
	if err != nil {
		t.Errorf("Failed to create DELETE request! %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("Failed to execute delete request! %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Failed to correctly delete lab! %v", err)
	}
}

func TestCleanup(t *testing.T) {
	os.Remove("./test_database.db")
}
