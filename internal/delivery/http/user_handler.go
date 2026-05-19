package http

import (
	"be-ayaka/internal/core/customerrors"
	"be-ayaka/internal/core/entity"
	"be-ayaka/internal/core/service"
	"be-ayaka/pkg/logger"
	"be-ayaka/pkg/requestid"
	"be-ayaka/pkg/response"
	"be-ayaka/pkg/validator"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService service.UserService
	validator   validator.Validator
}

func NewUserHandler(userService service.UserService, validator validator.Validator) *UserHandler {
	return &UserHandler{
		userService: userService,
		validator:   validator,
	}
}

// RegisterUser register new user
// @Summary Register new User
// @Description Register new user using email, username, and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body entity.UserRequest true "Payload Register"
// @Success 200 {object} response.Response{nil}
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response "Validation Failed"
// @Router /api/v1/auth/register [post]
func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	requestId := requestid.GetRequestID(c)

	var request entity.UserRequest

	if err := c.BodyParser(&request); err != nil {
		go logger.Log("SYSTEM", "ERROR", "Failed to parse request body: "+err.Error(), requestId)
		return customerrors.ErrBadRequest
	}

	if err := h.validator.Validate(c.Context(), &request); err != nil {
		go logger.Log("SYSTEM", "WARN", "Validation failed: "+err.Error(), requestId)
		return err
	}

	if err := h.userService.Create(c.Context(), &request); err != nil {
		go logger.Log("SYSTEM", "ERROR", "Internal Server Error: "+err.Error(), requestId)
		return err
	}

	go logger.Log("SYSTEM", "INFO", fmt.Sprintf("User %s created successfully", request.Username), requestId)
	return c.Status(fiber.StatusCreated).JSON(
		response.NewSuccessResponse(
			response.StatusSuccess,
			"Success Create User",
			nil,
			requestId,
		),
	)
}

// GetProfile retrieves the profile of the authenticated user
// @Summary Get User Profile
// @Description Retrieve the profile information of the authenticated user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body entity.UserRequest true "Payload Register"
// @Success 200 {object} response.Response{data=entity.User}
// @Failure 404 {object} response.Response "Profile not found"
// @Router /api/v1/auth/profile [get]
func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	requestId, ok := c.Locals("request_id").(string)
	if !ok {
		requestId = c.Get("X-Request-ID", "unknown-request-id")
	}

	userEntity, err := h.userService.GetProfile(c.Context(), userID)
	if err != nil {
		logger.Log("SYSTEM", "ERROR", fmt.Sprintf("Failed to retrieve user profile for userID %s: %v", userID, err), requestId)
		return err
	}

	go logger.Log("SYSTEM", "INFO", fmt.Sprintf("User profile for %s retrieved successfully", userEntity.Username), requestId)
	return c.Status(fiber.StatusOK).JSON(response.NewSuccessResponse(
		response.StatusSuccess,
		"Profile retrieved successfully",
		userEntity,
		requestId,
	))
}
