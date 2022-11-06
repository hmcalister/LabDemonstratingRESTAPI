package database

import (
	"hmcalister/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var conn *gorm.DB

func databaseSetup(conn *gorm.DB) {
	conn.AutoMigrate(&models.User{})
	conn.AutoMigrate(&models.Student{})
	conn.AutoMigrate(&models.Lab{})
	conn.AutoMigrate(&models.LabCompletion{})
}

func CreateDatabase(database_file string) error {
	var err error
	conn, err = gorm.Open(sqlite.Open("./"+database_file), &gorm.Config{})
	if err != nil {
		return err
	}

	databaseSetup(conn)

	return nil
}
