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

func CreatePatient(db *gorm.DB, request *interfaces.CreatePatientRequest) (*models.Patient, error) {
	patient := &models.Patient{
		Name:     request.Name,
		Password: request.Password,
		Email:    request.Email,
	}

	res := db.Create(patient)

	return patient, res.Error
}

func GetPatientByEmail(db *gorm.DB, email string) (*models.Patient, error) {
	var patients []models.Patient
	var patient *models.Patient

	res := db.Limit(1).Model(&models.Patient{}).Find(&patients, "email = ?", email)

	if len(patients) > 0 {
		patient = &patients[0]
	}

	return patient, res.Error
}

func SearchPatients(db *gorm.DB, searchQuery string) (*[]models.Patient, error) {
	var patients []models.Patient

	searchCondition := "( @searchQuery LIKE '%' || LOWER(email) || '%' ) OR ( @searchQuery LIKE '%' || LOWER(name) || '%' )"

	if _, err := strconv.ParseInt(searchQuery, 10, 64); err == nil {
		searchCondition += " OR ( id = " + searchQuery + " )"
	}

	res := db.Select("id, name, email").Model(&models.Doctor{}).Where(searchCondition, sql.Named("searchQuery", strings.ToLower(searchQuery))).Find(&patients)

	return &patients, res.Error
}

func CreateFilesForPatient(db *gorm.DB, patientId uint, files *[]models.File) (*models.Patient, error) {
	patient := models.Patient{ID: patientId}
	patientFiles := *files

	err := db.Model(&patient).Association(constants.ASSOCIATION_FILE).Append(patientFiles)

	return &patient, err
}

func GetPatientWithConnectionPopulated(db *gorm.DB, doctorId uint, patientId uint) (*models.Patient, bool) {
	var result models.Patient
	var wg sync.WaitGroup
	doesConnectionExist := false

	wg.Add(2)
	go GetPatientById(db, doctorId, &result, &wg)
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

func GetPatientById(db *gorm.DB, id uint, result *models.Patient, wg *sync.WaitGroup) {
	defer wg.Done()
	var patients []models.Patient

	db.Limit(1).Model(&models.Patient{}).Find(&patients, "id = ?", id)

	if len(patients) > 0 {
		*result = patients[0]
	}
}

func GetPatientsByDoctorId(db *gorm.DB, results *[]models.Patient, doctorId uint) {
	db.Preload("Doctors").Find(&results, "id=?", doctorId)
}

func GetFilesByPatientId(db *gorm.DB, patientId uint) *[]models.File {
	var patients []models.Patient
	var files []models.File

	db.Model(&models.Patient{}).Preload(constants.ASSOCIATION_FILE).First(&patients, patientId)

	if len(patients) > 0 {
		files = patients[0].Files
	}

	return &files
}
