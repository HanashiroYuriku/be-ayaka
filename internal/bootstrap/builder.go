package bootstrap

import (
	"be-ayaka/config"
	"be-ayaka/internal/adapter/repository"
	adapterRepo "be-ayaka/internal/adapter/repository"
	"be-ayaka/internal/core/service"
	"be-ayaka/internal/delivery/http"
	"be-ayaka/pkg/hash"
	"be-ayaka/pkg/validator"

	"gorm.io/gorm"
)

type Handlers struct {
	User *http.UserHandler
}

func BuildAllDependencies(db *gorm.DB, cfg *config.Config) *Handlers {
	// === validator
	validator := validator.NewGoValidator(db)
	// === tx manager
	txManager := repository.NewTxManager(db)
	// === hash service
	hashService := hash.NewBcryptHash()

	// adapter
	userRepo := adapterRepo.NewUserRepo(db)
	// service
	userService := service.NewUserService(userRepo, hashService, txManager)
	// handler
	return &Handlers{
		User: http.NewUserHandler(userService, validator),
	}
}
