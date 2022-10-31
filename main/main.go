package main

import (
	"flag"
	"fmt"
	"hmcalister/api"
	"hmcalister/database"
	"hmcalister/models"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
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

func insertTestData(databaseFilepath string) {
	os.Remove(databaseFilepath)
	err := database.CreateDatabase(databaseFilepath)
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
	var err error
	var routerLogFile io.Writer

	databaseFilepath := flag.String("database", "database.db", "Database file to use. Defaults to 'database.db'.")
	routerLogFilepath := flag.String("routerLog", "", "File to store router logs. If not set, print router logs to STDOUT.")
	portNumber := flag.Int("port", 80, "Port to serve webserver and API on. Defaults to 80.")
	flag.Parse()
	if *routerLogFilepath == "" {
		routerLogFile = os.Stdout
	} else {
		routerLogFile, err = os.OpenFile(*routerLogFilepath, os.O_RDWR|os.O_CREATE, 0644)
		checkFatalError(err)
	}

	insertTestData(*databaseFilepath)

	gin.DefaultWriter = routerLogFile
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	api.SetupAPI(router)

	router.Run("localhost:" + fmt.Sprint(*portNumber))
}
