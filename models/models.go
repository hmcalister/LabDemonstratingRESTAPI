package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"Username" gorm:"primaryKey;autoIncrement:false;unique:true"`
	Password string `json:"Password"`
}

type Student struct {
	gorm.Model
	StudentCode string `json:"StudentCode" gorm:"primaryKey;column:StudentCode;unique:true"`
	FirstName   string `json:"FirstName"`
	MiddleNames string `json:"MiddleNames"`
	LastName    string `json:"LastName"`
}

type Lab struct {
	gorm.Model
	ID          int    `json:"LabID" gorm:"primaryKey"`
	LabName     string `json:"LabName"`
	Description string `json:"Description"`
	Points      int    `json:"Points"`
}

type LabCompletion struct {
	gorm.Model
	StudentCode string  `json:"StudentCode" gorm:"column:StudentCode;primaryKey;autoIncrement:false;"`
	Student     Student `gorm:"foreignKey:StudentCode"`
	LabID       int     `json:"LabID" gorm:"column:LabID;primaryKey;autoIncrement:false"`
	Lab         Lab     `gorm:"foreignKey:LabID"`
	// Timestamp   time.Time `json:"Timestamp"`
}
