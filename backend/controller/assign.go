package controller

import (
	"docto/constants"
	"docto/models"

	"gorm.io/gorm"
)

func AssignPatient(db *gorm.DB, doctorId uint, patientId uint) error {
	return db.Session(&gorm.Session{FullSaveAssociations: false}).Model(&models.Doctor{ID: doctorId}).Association(constants.ASSOCIATION_PATIENT).Append(&models.Patient{ID: patientId})
}

func AssignDoctor(db *gorm.DB, doctorId uint, patientId uint) error {
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Model(&models.Patient{ID: patientId}).Association(constants.ASSOCIATION_DOCTOR).Append(&models.Doctor{ID: doctorId})
}
