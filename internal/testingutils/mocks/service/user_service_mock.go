package mocks

import (
	"be-ayaka/internal/core/entity"
	"context"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Create(ctx context.Context, user *entity.UserRequest) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserService) GetProfile(ctx context.Context, userID string) (*entity.User, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.User), args.Error(1)
	}
	return nil, args.Error(1)
}
