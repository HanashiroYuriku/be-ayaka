package bootstrap

import (
	"be-ayaka/config"
	adapterRepo "be-ayaka/internal/adapter/repository"
	"be-ayaka/internal/core/service"
	"be-ayaka/internal/delivery/http"

	"gorm.io/gorm"
)

type Handlers struct {
	User *http.UserHandler
}

func BuildAllDependencies(db *gorm.DB, cfg *config.Config) *Handlers {
	// adapter
	userRepo := adapterRepo.NewUserRepoPostgres(db)
	// service
	userService := service.NewUserService(userRepo)
	// handler
	return &Handlers{
		User: http.NewUserHandler(userService),
	}
}