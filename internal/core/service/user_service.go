package service

import (
	"errors"

	"be-ayaka/internal/core/entity"
	"be-ayaka/internal/core/repository"
)

// UserService defines the interface for user-related business logic
type UserService interface {
	GetProfile(userID string) (*entity.User, error)
}

// userServiceImpl is the concrete implementation of UserService
type userServiceImpl struct {
	userRepo repository.UserRepository // Add other repositories if needed
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{
		userRepo: repo,
	}
}

func (s *userServiceImpl) GetProfile(userID string) (*entity.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}