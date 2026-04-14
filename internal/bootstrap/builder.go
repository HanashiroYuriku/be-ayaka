package bootstrap

import (
	adapterRepo "be-ayaka/internal/adapter/repository"
	"be-ayaka/internal/delivery/http"
	"be-ayaka/internal/core/service"

	"gorm.io/gorm"
)

func BuildUserHandler(db *gorm.DB) *http.UserHandler {
	// adapter
	userRepo := adapterRepo.NewUserRepoPostgres(db)
	// service
	userService := service.NewUserService(userRepo)
	// handler
	return http.NewUserHandler(userService)
}