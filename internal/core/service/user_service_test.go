package service_test

import (
	"be-ayaka/internal/core/customerrors"
	"be-ayaka/internal/core/service"
	"be-ayaka/internal/delivery/http/dto"
	mocksPkg "be-ayaka/internal/testingutils/mocks/pkg"
	mocksRepo "be-ayaka/internal/testingutils/mocks/repository"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserServiceSuite struct {
	suite.Suite
	mockUserRepo *mocksRepo.MockUserRepo
	mockHash     *mocksPkg.MockHashService
	mockTx       *mocksRepo.MockTxManager
	service      service.UserService
}

// SetupTest
func (s *UserServiceSuite) SetupTest() {
	s.mockUserRepo = new(mocksRepo.MockUserRepo)
	s.mockHash = new(mocksPkg.MockHashService)
	s.mockTx = new(mocksRepo.MockTxManager)

	s.service = service.NewUserService(
		s.mockUserRepo,
		s.mockHash,
		s.mockTx,
	)
}

// Runner utama untuk menjalankan seluruh isi suite
func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UserServiceSuite))
}

// =============================================================================
// CREATE USER SCENARIOS
// =============================================================================

func (s *UserServiceSuite) TestCreate_Success() {
	ctx := context.Background()
	user := &dto.UserRequest{
		Username: "yuriku",
		Email:    "yuriku@mail.com",
		Password: "password123",
	}
	hashedPassword := "HASHEDPASSWORD"

	s.mockHash.On("HashPassword", user.Password).Return(hashedPassword, nil).Once()
	s.mockTx.On("WithTx", ctx, mock.Anything).Return(nil).Once()
	s.mockUserRepo.On("Create", ctx, mock.Anything).Return(nil).Once()

	err := s.service.Create(ctx, user)

	s.NoError(err)
	s.mockHash.AssertExpectations(s.T())
	s.mockTx.AssertExpectations(s.T())
	s.mockUserRepo.AssertExpectations(s.T())
}

func (s *UserServiceSuite) TestCreate_Failed_HashPassword() {
	ctx := context.Background()
	user := &dto.UserRequest{
		Username: "yuriku",
		Email:    "yuriku@mail.com",
		Password: "password123",
	}
	expectedError := customerrors.ErrFailHash
	s.mockHash.On("HashPassword", user.Password).Return("", expectedError).Once()

	err := s.service.Create(ctx, user)

	s.ErrorIs(err, expectedError)
	s.mockTx.AssertNotCalled(s.T(), "WithTx", mock.Anything, mock.Anything)
	s.mockUserRepo.AssertNotCalled(s.T(), "Create", mock.Anything, mock.Anything)
	s.mockHash.AssertExpectations(s.T())
}

func (s *UserServiceSuite) TestCreate_Failed_SaveToUserTable() {
	ctx := context.Background()
	user := &dto.UserRequest{
		Username: "yuriku",
		Email:    "yuriku@mail.com",
		Password: "password123",
	}
	hashedPassword := "HASHEDPASSWORD"
	dbError := errors.New("sql error")

	s.mockHash.On("HashPassword", user.Password).Return(hashedPassword, nil).Once()
	s.mockUserRepo.On("Create", ctx, mock.Anything).Return(dbError).Once()
	s.mockTx.On("WithTx", ctx, mock.Anything).Return(dbError).Once()

	err := s.service.Create(ctx, user)

	s.ErrorIs(err, dbError)
	s.mockHash.AssertExpectations(s.T())
	s.mockUserRepo.AssertExpectations(s.T())
	s.mockTx.AssertExpectations(s.T())
}

func (s *UserServiceSuite) TestCreate_Failed_InsideGenerateAndSendVerif() {
	ctx := context.Background()
	user := &dto.UserRequest{
		Username: "yuriku",
		Email:    "yuriku@mail.com",
		Password: "password123",
	}
	hashedPassword := "HASHEDPASSWORD"
	dbError := errors.New("sql error")

	s.mockHash.On("HashPassword", user.Password).Return(hashedPassword, nil).Once()
	s.mockUserRepo.On("Create", ctx, mock.Anything).Return(nil).Once()
	s.mockTx.On("WithTx", ctx, mock.Anything).Return(dbError).Once()

	err := s.service.Create(ctx, user)

	s.ErrorIs(err, dbError)
	s.mockHash.AssertExpectations(s.T())
	s.mockUserRepo.AssertExpectations(s.T())
	s.mockTx.AssertExpectations(s.T())
}
