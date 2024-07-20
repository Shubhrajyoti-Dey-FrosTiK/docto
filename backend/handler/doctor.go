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
// @Summary Creates a Doctor in the database
// @Description Takes input the fields and creates a doctor.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /doctor/create [get]
func (handler *Handler) HandleCreateDoctor(ctx *fiber.Ctx) error {
	var request interfaces.CreateDoctorRequest
	ctx.BodyParser(&request)

	doctor, err := controller.CreateDoctor(handler.DB, &request)

	if err != nil {
		return err
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "DOCTOR_CREATED", mapper.DoctorModelToCreateDoctorResponse(doctor), nil))
}

// LoginDoctor godoc
// @Summary Logs in a doctor
// @Description Takes input the email and password and returns a token
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /doctor/login [get]
func (handler *Handler) HandleLoginDoctor(ctx *fiber.Ctx) error {
	var request interfaces.LoginDoctorRequest
	ctx.BodyParser(&request)

	doctor, err := controller.GetDoctorByEmail(handler.DB, request.Email)

	if err != nil {
		return err
	}

	if doctor == nil {
		return errors.New("doctor not found")
	}

	if !util.CheckPasswordHash(request.Password, doctor.Password) {
		return errors.New("password does not match")
	}

	token, err := auth.GenerateToken(true, false, doctor.ID)

	if err != nil {
		return err
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "DOCTOR_LOGGED_IN", interfaces.LoginDoctorResponse{
		Token: token,
	}, nil))
}

// HandleSearchDoctors godoc
// @Summary Searches doctors in the database
// @Description Takes input the searchQuery as a query param and searches it for @id, @name and @email
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /doctor/search [get]
func (handler *Handler) HandleSearchDoctors(ctx *fiber.Ctx) error {

	searchQuery := ctx.Query("searchQuery", "")

	doctors, err := controller.SearchDoctors(handler.DB, searchQuery)

	if err != nil {
		return err
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "DOCTORS_FETCHED", mapper.DoctorsModelToSearchDoctorsResponse(doctors), nil))
}

// HandleGetDoctorWithPatientPopulated godoc
// @Summary Gets doctor with patients
// @Description Takes input the fields and creates a doctor.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /doctor/populated [get]
func (handler *Handler) HandleGetDoctorWithPatientPopulated(ctx *fiber.Ctx) error {
	auth := &auth.Auth{Ctx: ctx}
	doctorId, err := auth.GetId()

	if err != nil {
		return err
	}

	doctor, err := controller.GetDoctorWithPatientsPopulated(handler.DB, *doctorId)

	if err != nil {
		return err
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "DOCTOR_FETCHED", doctor, nil))
}

// CreateDoctor godoc
// @Summary Creates a Doctor in the database
// @Description Takes input the fields and creates a doctor.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /doctor/populated [get]
func (handler *Handler) HandleGetDoctorWithConnectionPopulated(ctx *fiber.Ctx) error {
	auth := &auth.Auth{Ctx: ctx}
	patientId, err := auth.GetId()
	doctorId, err := strconv.ParseUint(ctx.Query("doctorId", ""), 10, 64)

	if err != nil {
		return err
	}

	doctor, connected := controller.GetDoctorWithConnectionPopulated(handler.DB, uint(doctorId), *patientId)

	if err != nil {
		return err
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "DOCTOR_FETCHED", mapper.CreateGetDoctorByConnectionResponse(doctor, connected), nil))
}

// UploadFile godoc
// @Summary Uploads a file and attactches it to the doctor
// @Description Uploads a file and attactches it to the doctor
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /doctor/upload [get]
func (handler *Handler) HandleDoctorUploadFile(ctx *fiber.Ctx) error {
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

	doctor, err := controller.CreateFilesForDoctor(handler.DB, *doctorId, files)

	if err != nil {
		return err
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "DOCTOR_FILE_UPLOADED", mapper.FileModelsToFiles(&doctor.Files), nil))
}

// HandleGetAssociatedUsersForDoctors godoc
// @Summary Gets patient associated with doctors
// @Description Takes input the fields and creates a doctor.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /doctor/populated [get]
func (handler *Handler) HandleGetAssociatedUsersForDoctors(ctx *fiber.Ctx) error {
	auth := &auth.Auth{Ctx: ctx}
	doctorId, err := auth.GetId()

	if err != nil {
		return err
	}

	var patients []models.Patient

	controller.GetPatientsByDoctorId(handler.DB, &patients, *doctorId)

	if err != nil {
		return err
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "PATIENTS_FETCHED", &interfaces.GetAssociatedUsersResponse{
		Users: *mapper.PatientModelsToUserMapper(&patients),
	}, nil))
}

// HandleGetAssociatedFilesForDoctors godoc
// @Summary Gets files associated with doctors
// @Description Takes input the fields and creates a doctor.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /doctor/populated [get]
func (handler *Handler) HandleGetAssociatedFilesForDoctors(ctx *fiber.Ctx) error {
	auth := &auth.Auth{Ctx: ctx}
	doctorId, err := auth.GetId()

	if err != nil {
		return err
	}

	files := controller.GetFilesByDoctorId(handler.DB, *doctorId)

	return ctx.JSON(interfaces.GetGenericResponse(true, "FILES_FETCHED", mapper.FileModelsToFiles(files), nil))
}
