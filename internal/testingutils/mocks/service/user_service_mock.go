package mocks

import (
	"be-ayaka/internal/delivery/http/dto"
	"context"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Create(ctx context.Context, user *dto.UserRequest) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserService) GetProfile(ctx context.Context, userID string) (*dto.UserResponse, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) != nil {
		return args.Get(0).(*dto.UserResponse), args.Error(1)
	}
	return nil, args.Error(1)
}
