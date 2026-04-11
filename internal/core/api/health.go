package api

import (
	"be-ayaka/config"
	ayaka "be-ayaka/pkg/logger"
	"be-ayaka/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type HealthCheckHandler struct {
	Config *config.Config
}

func NewHealthCheckHandler(cfg *config.Config) *HealthCheckHandler {
	return &HealthCheckHandler{
		Config: cfg,
	}
}

func (h *HealthCheckHandler) Ping(c *fiber.Ctx) error {
	go ayaka.Log("SYSTEM", "INFO", "Ayaka Server is running")

	data := fiber.Map{
		"serverName": h.Config.App.Name,
		"version":     h.Config.App.Version,
		"isHealthy":  true,
	}

	return c.Status(fiber.StatusOK).JSON(response.NewSuccessResponse(data))
}