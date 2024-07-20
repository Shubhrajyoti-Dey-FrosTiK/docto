package handler

import (
	"docto/interfaces"

	"github.com/gofiber/fiber/v2"
)

// HealthCheck godoc
// @Summary Show the health of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func (h *Handler) HealthCheckHandler(ctx *fiber.Ctx) error {
	ctx.JSON(interfaces.GetGenericResponse(true, "SERVER_HEALTHY", nil, nil))
	return nil
}
