package core

import (
	"be-ayaka/config"
	"be-ayaka/internal/core/api"
	"be-ayaka/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func SetupRoutes(app *fiber.App, cfg *config.Config) {
	// middleware request
	app.Use(requestid.New(requestid.Config{
		Header: fiber.HeaderXRequestID,
		ContextKey: "trace_id",
		Generator: func () string {
			return utils.GenerateID("TRACE")
		},
	}))

	// Health Check
	healthHandler := api.NewHealthCheckHandler(cfg)
	// health route
	app.Get("/ping", healthHandler.Ping)
	// ---
}