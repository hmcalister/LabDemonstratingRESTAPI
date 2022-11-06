package database

import (
	"errors"
	"hmcalister/models"
)

func CreateLab(newLab models.Lab) error {
	result := conn.Create(&newLab)
	return result.Error
}

func GetAllLabs() ([]models.Lab, error) {
	var Labs []models.Lab
	result := conn.Find(&Labs)
	return Labs, result.Error
}

func GetLabByLabID(LabID int) (models.Lab, error) {
	var Lab models.Lab
	result := conn.First(&Lab, "ID = ?", LabID)
	return Lab, result.Error
}

func UpdateLabByLabID(LabID int, newLab models.Lab) error {
	if newLab.ID != LabID {
		return errors.New("cannot update LabID")
	}
	result := conn.Save(&newLab)
	return result.Error
}

func DeleteLabByLabID(LabID int) error {
	result := conn.Delete(&models.Lab{}, "ID = ?", LabID)
	return result.Error
}
