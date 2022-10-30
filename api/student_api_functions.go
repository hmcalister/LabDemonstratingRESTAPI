package api

import (
	"hmcalister/database"
	"hmcalister/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createStudent(c *gin.Context) {
	var newStudent models.Student
	var err error
	err = c.BindJSON(&newStudent)
	checkError(err)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	err = database.CreateStudent(newStudent)
	checkError(err)
	if err != nil {
		// This is PROBABLY the right request, although
		// TODO fix up status handling later
		c.IndentedJSON(http.StatusConflict, nil)
		return
	}

	c.IndentedJSON(http.StatusCreated, newStudent)
}

func getAllStudents(c *gin.Context) {
	var allStudents []models.Student
	var err error
	allStudents, err = database.GetAllStudents()
	checkError(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, allStudents)
}

func getStudentByStudentCode(c *gin.Context) {
	studentCode := c.Param("studentCode")
	student, err := database.GetStudentByStudentCode(studentCode)
	checkError(err)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, nil)
	} else {
		c.IndentedJSON(http.StatusOK, student)
	}
}

func deleteStudentByStudentCode(c *gin.Context) {
	studentCode := c.Param("studentCode")
	err := database.DeleteStudentByStudentCode(studentCode)
	checkError(err)
	if err != nil {
		// Again, PROBABLY correct status...
		c.IndentedJSON(http.StatusNoContent, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}
