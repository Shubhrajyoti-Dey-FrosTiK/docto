package middleware

import (
	"docto/constants"
	"docto/interfaces"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validator = validator.New(validator.WithRequiredStructEnabled())
var idValidPassword = regexp.MustCompile(`^[a-zA-Z0-9@_-]+$`).MatchString

func ValidateBody[Doc any](ctx *fiber.Ctx) error {
	Validator.RegisterValidation("strongpassword", CheckStrongPassword)

	body := new(Doc)
	ctx.BodyParser(&body)

	err := Validator.Struct(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(interfaces.GetGenericResponse(false, constants.ERROR_INCORRENT_BODY, nil, err))
	}
	return ctx.Next()
}

func CheckStrongPassword(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	return len(field) > 2 && idValidPassword(field)
}
