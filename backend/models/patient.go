package models

import (
	"docto/util"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model

	ID      uint     `gorm:"primaryKey"` // id of the patient
	Name    string   // name of the patient
	Email   string   `gorm:"uniqueIndex"` // email of the patient
	Files   []File   `gorm:"OnDelete:SET ARRAY[]::varchar[]; many2many:patient_files;"`
	Doctors []Doctor `gorm:"OnDelete:SET ARRAY[]::varchar[]; many2many:doctor_patients;"`

	// Sensitive
	Password string // Will be hashed before being inserted

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}

func (patient *Patient) BeforeCreate(tx *gorm.DB) (err error) {
	if patient.Doctors == nil {
		patient.Doctors = []Doctor{}
	}

	patient.Password, err = util.HashPassword(patient.Password)

	if err != nil {
		return err
	}

	return
}
