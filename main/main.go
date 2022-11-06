package main

import (
	"flag"
	"hmcalister/api"
	"hmcalister/database"
	"io"
	"log"
	"net/http"
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

func main() {
	var err error
	var routerLogFile io.Writer

	databaseFilepath := flag.String("database", "database.db", "Database file to use. Files stored locally, use a different file for each paper.")
	routerLogFilepath := flag.String("routerLog", "", "File to store router logs. If not set, print router logs to STDOUT.")
	debugMode := flag.Bool("debug", false, "Run application in debug mode. e.g. allow CORS on API, insert some test data...")
	flag.Parse()
	if *debugMode {
		log.Println("DEBUG MODE: ON")
		insertTestData(*databaseFilepath)
	} else {
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		gin.SetMode(gin.ReleaseMode)
		err := database.CreateDatabase(*databaseFilepath)
		checkFatalError(err)
	}

	if *routerLogFilepath == "" {
		routerLogFile = os.Stdout
	} else {
		routerLogFile, err = os.OpenFile(*routerLogFilepath, os.O_RDWR|os.O_CREATE, 0644)
		checkFatalError(err)
	}

	gin.DefaultWriter = routerLogFile
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	if *debugMode {
		router.Use(CORSMiddleware())
	}
	api.SetupAPI(router)

	log.Println("APPLICATION READY: http://localhost:8080/")
	router.Run("localhost:8080")
}
