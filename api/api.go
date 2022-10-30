package api

import (
	"hmcalister/database"
	"hmcalister/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func checkFatalError(err error) {
	if err != nil {
		log.Fatalln("ERROR: ", err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Println("ERROR: ", err)
	}
}

func getAllStudents(c *gin.Context) {
	var allStudents []models.Student
	var err error
	allStudents, err = database.GetAllStudents()
	checkError(err)

	c.IndentedJSON(http.StatusOK, allStudents)
}

func SetupAPI() {

	router := gin.Default()
	router.GET("/students", getAllStudents)

	router.Run("localhost:8080")
}
