package models

import (
	"docto/util"

	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model

	ID          uint      `gorm:"primaryKey"` // id of the doctor
	Name        string    // name of the doctor
	Email       string    `gorm:"uniqueIndex"` // email of the patient
	Designation string    // eg. MBBS etc
	Headline    string    // eg. Chief of surgery at AIMS Delhi
	Patients    []Patient `gorm:"OnDelete:SET ARRAY[]::varchar[]; many2many:doctor_patients;"`
	Files       []File    `gorm:"OnDelete:SET ARRAY[]::varchar[]; many2many:doctor_files;"`

	// Sensitive
	Password string // Will be hashed before being inserted

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}

func (doctor *Doctor) BeforeCreate(tx *gorm.DB) (err error) {
	if doctor.Patients == nil {
		doctor.Patients = []Patient{}
	}

	if doctor.Files == nil {
		doctor.Files = []File{}
	}

	doctor.Password, err = util.HashPassword(doctor.Password)

	if err != nil {
		return err
	}

	return
}
