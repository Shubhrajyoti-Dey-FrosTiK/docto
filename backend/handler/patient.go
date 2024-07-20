package handler

import (
	"docto/auth"
	"docto/controller"
	"docto/interfaces"
	"docto/mapper"
	"docto/models"
	"docto/util"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// CreateDoctor godoc
// @Summary Creates a Patient in the database
// @Description Takes input the fields and creates a patient.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /patient/create [get]
func (handler *Handler) HandleCreatePatient(ctx *fiber.Ctx) error {
	var request interfaces.CreatePatientRequest
	ctx.BodyParser(&request)

	patient, err := controller.CreatePatient(handler.DB, &request)

	if err != nil {
		return err
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "PATIENT_CREATED", mapper.PatientModelToCreatePatientResponse(patient), nil))
}

// LoginPatient godoc
// @Summary Logs in a patient
// @Description Takes input the email and password and returns a token
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /doctor/login [get]
func (handler *Handler) HandleLoginPatient(ctx *fiber.Ctx) error {
	var request interfaces.LoginPatientRequest
	ctx.BodyParser(&request)

	doctor, err := controller.GetPatientByEmail(handler.DB, request.Email)

	if err != nil {
		return err
	}

	if doctor == nil {
		return errors.New("patient not found")
	}

	if !util.CheckPasswordHash(request.Password, doctor.Password) {
		return errors.New("password does not match")
	}

	token, err := auth.GenerateToken(false, true, doctor.ID)

	if err != nil {
		return err
	}

	ctx.JSON(interfaces.GetGenericResponse(true, "PATIENT_LOGGED_IN", interfaces.LoginDoctorResponse{
		Token: token,
	}, nil))

	return nil
}

// HandleSearchPatients godoc
// @Summary Searches patients in the database
// @Description Takes input the searchQuery as a query param and searches it for @id, @name and @email
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /doctor/search [get]
func (handler *Handler) HandleSearchPatients(ctx *fiber.Ctx) error {

	searchQuery := ctx.Query("searchQuery", "")

	patients, err := controller.SearchPatients(handler.DB, searchQuery)

	if err != nil {
		return err
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "DOCTORS_FETCHED", mapper.PatientsModelToSearchDoctorsResponse(patients), nil))
}

// HandlePatientUploadFile godoc
// @Summary Uploads a file and attactches it to the patient
// @Description Uploads a file and attactches it to the patient
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /doctor/upload [get]
func (handler *Handler) HandlePatientUploadFile(ctx *fiber.Ctx) error {
	auth := &auth.Auth{Ctx: ctx}
	doctorId, err := auth.GetId()

	if err != nil {
		return err
	}

	form, _ := ctx.MultipartForm()

	fileReaders, found := form.File["file"]

	if !found {
		return errors.New("no file found")
	}

	files, err := controller.UploadFiles(ctx, handler.S3, fileReaders)

	if err != nil {
		return err
	}

	patient, err := controller.CreateFilesForPatient(handler.DB, *doctorId, files)

	if err != nil {
		return err
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "FILE_UPLOADED", mapper.FileModelsToFiles(&patient.Files), nil))
}

// HandleGetPatientWithConnectionPopulated godoc
// @Summary Creates a Doctor in the database
// @Description Takes input the fields and creates a doctor.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /patient/populated [get]
func (handler *Handler) HandleGetPatientWithConnectionPopulated(ctx *fiber.Ctx) error {
	auth := &auth.Auth{Ctx: ctx}
	doctorId, err := auth.GetId()
	patientId, err := strconv.ParseUint(ctx.Query("patientId", ""), 10, 64)

	if err != nil {
		return err
	}

	patient, connected := controller.GetPatientWithConnectionPopulated(handler.DB, *doctorId, uint(patientId))

	if err != nil {
		return err
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "DOCTOR_FETCHED", mapper.CreateGetPatientByConnectionResponse(patient, connected), nil))
}

// HandleGetAssociatedUsersForDoctors godoc
// @Summary Gets patient associated with doctors
// @Description Takes input the fields and creates a doctor.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /doctor/populated [get]
func (handler *Handler) HandleGetAssociatedUsersForPatients(ctx *fiber.Ctx) error {
	auth := &auth.Auth{Ctx: ctx}
	patientId, err := auth.GetId()

	if err != nil {
		return err
	}

	var doctors []models.Doctor

	controller.GetDoctorsByPatientId(handler.DB, &doctors, *patientId)

	if err != nil {
		return err
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "DOCTORS_FETCHED", interfaces.GetAssociatedUsersResponse{
		Users: *mapper.DoctorModelsToUserMapper(&doctors),
	}, nil))
}

// HandleGetAssociatedFilesForPatients godoc
// @Summary Gets files associated with patient
// @Description Takes input the fields and creates a doctor.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /doctor/populated [get]
func (handler *Handler) HandleGetAssociatedFilesForPatients(ctx *fiber.Ctx) error {
	auth := &auth.Auth{Ctx: ctx}
	patientId, err := auth.GetId()

	if err != nil {
		return err
	}

	files := controller.GetFilesByPatientId(handler.DB, *patientId)

	return ctx.JSON(interfaces.GetGenericResponse(true, "FILES_FETCHED", mapper.FileModelsToFiles(files), nil))
}
