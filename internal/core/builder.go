package core

import (
	adapterRepo "be-ayaka/internal/adapter/repository"
	"be-ayaka/internal/core/api"
	"be-ayaka/internal/core/service"

	"gorm.io/gorm"
)

func BuildUserHandler(db *gorm.DB) *api.UserHandler {
	// adapter
	userRepo := adapterRepo.NewUserRepoPostgres(db)
	// service
	userService := service.NewUserService(userRepo)
	// handler
	return api.NewUserHandler(userService)
}