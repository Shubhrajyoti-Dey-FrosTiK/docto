package controller

import (
	"database/sql"
	"docto/constants"
	"docto/interfaces"
	"docto/models"
	"strconv"
	"strings"
	"sync"

	"gorm.io/gorm"
)

func CreateDoctor(db *gorm.DB, request *interfaces.CreateDoctorRequest) (*models.Doctor, error) {
	doctor := &models.Doctor{
		Name:        request.Name,
		Password:    request.Password,
		Designation: request.Designation,
		Headline:    request.Headline,
		Email:       request.Email,
	}

	res := db.Create(doctor)

	return doctor, res.Error
}

func CreateFilesForDoctor(db *gorm.DB, doctorId uint, files *[]models.File) (*models.Doctor, error) {
	doctor := models.Doctor{ID: doctorId}
	doctorFiles := *files

	err := db.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Model(&doctor).Association(constants.ASSOCIATION_FILE).Append(doctorFiles)

	return &doctor, err
}

func GetDoctorWithPatientsPopulated(db *gorm.DB, doctorId uint) (*models.Doctor, error) {
	var doctors []models.Doctor
	var result *models.Doctor

	res := db.Limit(1).Model(&models.Doctor{}).Preload(constants.ASSOCIATION_PATIENT).Find(&doctors, "id = ?", doctorId)

	if len(doctors) > 0 {
		result = &doctors[0]
	}

	return result, res.Error
}

func GetDoctorWithConnectionPopulated(db *gorm.DB, doctorId uint, patientId uint) (*models.Doctor, bool) {
	var result models.Doctor
	var wg sync.WaitGroup
	doesConnectionExist := false

	wg.Add(2)
	go GetDoctorsById(db, doctorId, &result, &wg)
	go func(db *gorm.DB, doctorId uint, patientId uint, doesConnectionExist *bool, wg *sync.WaitGroup) {
		defer wg.Done()
		var result struct {
			Ispresent int
		}

		db.Raw("SELECT COUNT(*) AS ispresent FROM doctor_patients WHERE doctor_id = ? AND patient_id = ?", doctorId, patientId).Scan(&result)

		if result.Ispresent > 0 {
			*doesConnectionExist = true
		}
	}(db, doctorId, patientId, &doesConnectionExist, &wg)
	wg.Wait()

	return &result, doesConnectionExist
}

func GetDoctorByEmail(db *gorm.DB, email string) (*models.Doctor, error) {
	var doctors []models.Doctor
	var result *models.Doctor

	res := db.Limit(1).Model(&models.Doctor{}).Find(&doctors, "email = ?", email)

	if len(doctors) > 0 {
		result = &doctors[0]
	}

	return result, res.Error
}

func GetDoctorsById(db *gorm.DB, id uint, result *models.Doctor, wg *sync.WaitGroup) {
	defer wg.Done()
	var doctors []models.Doctor

	db.Limit(1).Model(&models.Doctor{}).Find(&doctors, "id = ?", id)

	if len(doctors) > 0 {
		*result = doctors[0]
	}
}

func SearchDoctors(db *gorm.DB, searchQuery string) (*[]models.Doctor, error) {
	var doctors []models.Doctor

	searchCondition := "( @searchQuery LIKE '%' || LOWER(email) || '%' ) OR ( @searchQuery LIKE '%' || LOWER(name) || '%' )"

	if _, err := strconv.ParseInt(searchQuery, 10, 64); err == nil {
		searchCondition += " OR ( id = " + searchQuery + " )"
	}

	res := db.Select("id, name, email, designation, headline").Model(&models.Doctor{}).Where(searchCondition, sql.Named("searchQuery", strings.ToLower(searchQuery))).Find(&doctors)

	return &doctors, res.Error
}

func GetDoctorsByPatientId(db *gorm.DB, results *[]models.Doctor, patientId uint) {
	db.Preload("Doctors").Find(&results, "id=?", patientId)
}

func GetFilesByDoctorId(db *gorm.DB, doctorId uint) *[]models.File {
	var doctors []models.Doctor
	var files []models.File

	db.Model(&models.Doctor{}).Preload(constants.ASSOCIATION_FILE).First(&doctors, doctorId)

	if len(doctors) > 0 {
		files = doctors[0].Files
	}

	return &files
}
