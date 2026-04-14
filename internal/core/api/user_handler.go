package api

import (
	"be-ayaka/internal/core/service"
	"be-ayaka/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	requestId, ok := c.Locals("request_id").(string)
	if !ok {
		requestId = "unknown-trace-id"
	}

	userEntity, err := h.userService.GetProfile(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.NewErrorResponse(
			response.DataNotFound,
			"Profile not found",
			requestId,
		))
	}

	return c.Status(fiber.StatusOK).JSON(response.NewSuccessResponse(
		response.StatusSuccess,
		"Profile retrieved successfully",
		userEntity,
		requestId,
	))
}
