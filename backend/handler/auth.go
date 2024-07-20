package handler

import (
	"docto/auth"
	"docto/interfaces"
	"docto/util"

	"github.com/gofiber/fiber/v2"
)

// HandleAuthCheck godoc
// @Summary Returns the status of the user token
// @Description This endpoint can be called to check if the user token is valid as well as to check if the user is a doctor or a patient
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /auth/token/verify [get]
func (h *Handler) HandleVerifyToken(ctx *fiber.Ctx) error {
	authContext := auth.Auth{Ctx: ctx}
	userId, err := authContext.GetId()
	if err != nil {
		return err
	}

	role := util.NilPtr[string]()
	if authContext.IsDoctor() {
		role = util.ToPtr[string]("DOCTOR")
	}
	if authContext.IsPatient() {
		role = util.ToPtr[string]("PATIENT")
	}

	token, err := auth.GenerateToken(authContext.IsDoctor(), authContext.IsPatient(), *userId)
	if err != nil {
		return err
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "TOKEN_VERIFIED", interfaces.TokenVerifyResponse{
		Token:  token,
		Role:   role,
		UserId: userId,
	}, nil))
}
