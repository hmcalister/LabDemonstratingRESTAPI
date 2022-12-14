package api

import (
	"log"

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

func SetupAPI(router *gin.Engine) {

	// Student endpoints
	router.POST("/students", createStudent)
	router.GET("/students", getAllStudents)
	router.GET("/students/:studentCode", getStudentByStudentCode)
	router.PUT("/students/:studentCode", updateStudentByStudentCode)
	router.DELETE("/students/:studentCode", deleteStudentByStudentCode)

	// Lab endpoints
	router.POST("/labs", createLab)
	router.GET("/labs", getAllLabs)
	router.GET("/labs/:labID", getLabByLabID)
	router.PUT("/labs/:labID", updateLabByLabID)
	router.DELETE("/labs/:labID", deleteLabByLabID)

	// Lab completion endpoints
	router.POST("/labCompletions", createLabCompletion)
	router.GET("/labCompletions", getAllLabCompletions)
	router.GET("/labCompletions/:studentCode", getAllLabCompletionsByStudentCode)
	router.GET("/labCompletions/:studentCode/:labID", getLabCompletionByStudentCodeAndLabID)
	router.DELETE("/labCompletions/:studentCode/:labID", deleteLabCompletionByStudentCodeAndLabID)
}
