package api

import (
	"hmcalister/database"
	"hmcalister/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createLabCompletion(c *gin.Context) {
	var newLabCompletion models.LabCompletion
	var err error
	err = c.BindJSON(&newLabCompletion)
	checkError(err)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	err = database.CreateLabCompletion(newLabCompletion)
	checkError(err)
	if err != nil {
		// This is PROBABLY the right request, although
		// TODO fix up status handling later
		c.IndentedJSON(http.StatusConflict, nil)
		return
	}

	c.IndentedJSON(http.StatusCreated, newLabCompletion)
}

func getAllLabCompletions(c *gin.Context) {
	var allLabCompletions []models.LabCompletion
	var err error
	allLabCompletions, err = database.GetAllLabCompletions()
	checkError(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, allLabCompletions)
}

func getAllLabCompletionsByStudentCode(c *gin.Context) {
	studentCode := c.Param("studentCode")
	var allLabCompletions []models.LabCompletion
	var err error
	allLabCompletions, err = database.GetAllLabCompletionsByStudentCode(studentCode)
	checkError(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, allLabCompletions)
}

func getLabCompletionByStudentCodeAndLabID(c *gin.Context) {
	studentCode := c.Param("studentCode")
	labID, err := strconv.Atoi(c.Param("labID"))
	checkError(err)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	labCompletion, err := database.GetLabCompletionByStudentCodeAndLabID(studentCode, labID)
	checkError(err)
	if err != nil {
		c.IndentedJSON(http.StatusNoContent, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, labCompletion)
}

func deleteLabCompletionByStudentCodeAndLabID(c *gin.Context) {
	studentCode := c.Param("studentCode")
	labID, err := strconv.Atoi(c.Param("labID"))
	checkError(err)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	err = database.DeleteLabCompletionByStudentCodeAndLabID(studentCode, labID)
	if err != nil {
		c.IndentedJSON(http.StatusNoContent, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}
