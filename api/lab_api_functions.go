package api

import (
	"hmcalister/database"
	"hmcalister/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createLab(c *gin.Context) {
	var newLab models.Lab
	var err error
	err = c.BindJSON(&newLab)
	checkError(err)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	err = database.CreateLab(newLab)
	checkError(err)
	if err != nil {
		// This is PROBABLY the right request, although
		// TODO fix up status handling later
		c.IndentedJSON(http.StatusConflict, nil)
		return
	}

	c.IndentedJSON(http.StatusCreated, newLab)
}

func getAllLabs(c *gin.Context) {
	var allLabs []models.Lab
	var err error
	allLabs, err = database.GetAllLabs()
	checkError(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, allLabs)
}

func getLabByLabID(c *gin.Context) {
	labID, err := strconv.Atoi(c.Param("labID"))
	checkError(err)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	lab, err := database.GetLabByLabID(labID)
	checkError(err)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, lab)
}

func updateLabByLabID(c *gin.Context) {
	var newLab models.Lab
	var err error

	labID, err := strconv.Atoi(c.Param("labID"))
	checkError(err)

	err = c.BindJSON(&newLab)
	checkError(err)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	newLab.ID = labID

	err = database.UpdateLabByLabID(labID, newLab)
	checkError(err)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, newLab)
}

func deleteLabByLabID(c *gin.Context) {
	labID, err := strconv.Atoi(c.Param("labID"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	err = database.DeleteLabByLabID(labID)
	checkError(err)
	if err != nil {
		// Again, PROBABLY correct status...
		c.IndentedJSON(http.StatusNoContent, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}
