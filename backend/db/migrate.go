package db

import (
	"docto/models"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Doctor{}, &models.Patient{}, &models.File{})
}
