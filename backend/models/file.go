package models

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model

	ID       uint `gorm:"primaryKey"` // id of the file
	Url      string
	Key      string
	FileName string

	DoctorID  uint
	PatientID uint

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}
