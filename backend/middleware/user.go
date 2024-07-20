package middleware

import (
	"docto/auth"
	"docto/constants"
	"docto/interfaces"

	"github.com/gofiber/fiber/v2"
)

/*
Checks if the user calling the endpoint is a doctor or not
This should be called after the jwt auth is done by the middleware and the 'user' locals is populated
*/
func CheckDoctor(ctx *fiber.Ctx) error {
	auth := &auth.Auth{Ctx: ctx}

	if !auth.IsDoctor() {
		return ctx.Status(fiber.StatusForbidden).JSON(interfaces.GetGenericResponse(false, constants.ERROR_USER_NOT_ALLOWED, nil, nil))
	}

	return ctx.Next()
}

/*
Checks if the user calling the endpoint is a patient or not
This should be called after the jwt auth is done by the middleware and the 'user' locals is populated
*/
func CheckPatient(ctx *fiber.Ctx) error {
	auth := &auth.Auth{Ctx: ctx}

	if !auth.IsPatient() {
		return ctx.Status(fiber.StatusForbidden).JSON(interfaces.GetGenericResponse(false, constants.ERROR_USER_NOT_ALLOWED, nil, nil))
	}

	return ctx.Next()
}
