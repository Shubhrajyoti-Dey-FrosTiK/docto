package handler

import (
	"docto/auth"
	"docto/controller"
	"docto/interfaces"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// AssignPatient godoc
// @Summary Assign Patient to Doctor
// @Description Takes input the patient_id which has to be assigned
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /assign/patient [get]
func (handler *Handler) HandleAssignPatient(ctx *fiber.Ctx) error {
	var request interfaces.AssignPatientRequest
	ctx.BodyParser(&request)

	auth := auth.Auth{Ctx: ctx}

	doctorId, err := auth.GetId()

	if err != nil {
		return err
	}

	patientId, err := strconv.ParseUint(request.PatientId, 10, 64)

	if err != nil {
		return err
	}

	if err = controller.AssignPatient(handler.DB, *doctorId, uint(patientId)); err != nil {
		return err
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "PATIENT_ASSIGNED", nil, nil))
}

// HandleAssignDoctor godoc
// @Summary Assign Patient to Doctor
// @Description Takes input the patient_id which has to be assigned
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /assign/doctor [get]
func (handler *Handler) HandleAssignDoctor(ctx *fiber.Ctx) error {
	var request interfaces.AssignDoctorRequest
	ctx.BodyParser(&request)

	auth := auth.Auth{Ctx: ctx}

	patientId, err := auth.GetId()

	if err != nil {
		return err
	}

	doctorId, err := strconv.ParseUint(request.DoctorId, 10, 64)

	if err != nil {
		return err
	}

	if err = controller.AssignDoctor(handler.DB, uint(doctorId), *patientId); err != nil {
		return err
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "DOCTOR_ASSIGBED", nil, nil))
}
