package mocks

import (
	"be-ayaka/internal/core/entity"
	"context"

	"github.com/stretchr/testify/mock"
)

// mock user repo
type MockUserRepo struct {
	mock.Mock
}

// func create
func (m *MockUserRepo) Create(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

// func find by id
func (m *MockUserRepo) FindByID(ctx context.Context, id string) (*entity.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.User), args.Error(1)
	}
	return nil, args.Error(1)
}

// func find by email
func (m *MockUserRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.User), args.Error(1)
	}
	return nil, args.Error(1)
}
