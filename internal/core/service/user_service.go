package service

import (
	"context"

	"be-ayaka/internal/core/customerrors"
	"be-ayaka/internal/core/entity"
	"be-ayaka/internal/core/port"
	"be-ayaka/internal/delivery/http/dto"
	"be-ayaka/pkg/generateid"
	"be-ayaka/pkg/hash"
)

// UserService defines the interface for user-related business logic
type UserService interface {
	Create(ctx context.Context, user *dto.UserRequest) error
	GetProfile(ctx context.Context, userID string) (*dto.UserResponse, error)
}

// userServiceImpl is the concrete implementation of UserService
type userServiceImpl struct {
	userRepo    port.UserRepository
	hashService hash.HashService
	txManager   port.TxManager
	// Add other repositories if needed
}

func NewUserService(repo port.UserRepository, hashService hash.HashService, txManager port.TxManager) UserService {
	return &userServiceImpl{
		userRepo:    repo,
		hashService: hashService,
		txManager:   txManager, //tx manager for handle transaction
	}
}

func (s *userServiceImpl) Create(ctx context.Context, user *dto.UserRequest) error {
	passwordHash, err := s.hashService.HashPassword(user.Password)
	if err != nil {
		return customerrors.ErrFailHash
	}

	userModel := &entity.User{
		Username: user.Username,
		Email:    user.Email,
		Password: passwordHash,
		Role:     "user",
	}
	userModel.ID = generateid.GenerateID("USER")

	// tx manager implementation example
	return s.txManager.WithTx(ctx, func(ctx context.Context) error {
		if err := s.userRepo.Create(ctx, userModel); err != nil {
			return err
		}

		return nil
	})
}

func (s *userServiceImpl) GetProfile(ctx context.Context, userID string) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}, nil
}
