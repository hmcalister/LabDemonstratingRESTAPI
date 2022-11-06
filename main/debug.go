package main

import (
	"hmcalister/database"
	"hmcalister/models"
	"os"

	"github.com/gin-gonic/gin"
)

func insertTestData(databaseFilepath string) {
	os.Remove(databaseFilepath)
	err := database.CreateDatabase(databaseFilepath)
	checkFatalError(err)

	newStudents := []models.Student{
		{
			StudentCode: "hayden000",
			FirstName:   "Hayden",
			MiddleNames: "Robert",
			LastName:    "McAlister",
		},
		{
			StudentCode: "alice123",
			FirstName:   "Alice",
			LastName:    "Test",
		},
		{
			StudentCode: "bob456",
			FirstName:   "Bob",
			LastName:    "ORM",
		},
	}

	newLabs := []models.Lab{
		{
			LabName:     "Test Lab",
			Description: "A lab for testing",
			Points:      0,
		},
		{
			LabName:     "Second Lab",
			Description: "Follow up to the first lab",
			Points:      0,
		},
		{
			LabName:     "Final lab",
			Description: "Actually worth something!",
			Points:      1,
		},
	}

	newCompletions := []models.LabCompletion{
		{
			StudentCode: "hayden000",
			LabID:       1,
		},
		{
			StudentCode: "hayden000",
			LabID:       3,
		},
		{
			StudentCode: "bob456",
			LabID:       1,
		},
	}

	for _, student := range newStudents {
		database.CreateStudent(student)
	}

	for _, lab := range newLabs {
		database.CreateLab(lab)
	}

	for _, labCompletion := range newCompletions {
		database.CreateLabCompletion(labCompletion)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
