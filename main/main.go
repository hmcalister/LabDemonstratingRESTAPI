package main

import (
	"hmcalister/api"
	"hmcalister/database"
	"hmcalister/models"
	"log"
	"os"

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

func insertTestData() {
	err := database.CreateDatabase("database.db")
	checkFatalError(err)

	newStudents := []models.Student{
		{
			StudentCode: "hayden000",
			FullName:    "Hayden McAlister",
		},
		{
			StudentCode: "alice123",
			FullName:    "Alice",
		},
		{
			StudentCode: "bob456",
			FullName:    "Bob",
		},
	}

	for _, student := range newStudents {
		database.CreateStudent(student)
	}
}

func main() {
	os.Remove("database.db")
	insertTestData()
	api.SetupAPI()

	gin.DefaultWriter = routerLogFile
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	api.SetupAPI(router)

	router.Run("localhost:" + fmt.Sprint(*portNumber))
}
