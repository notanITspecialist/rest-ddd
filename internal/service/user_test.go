package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	mockRepository "rest-ddd/internal/mocks/repository"
	"rest-ddd/internal/repository"
)

type UserServiceSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	repo    *mockRepository.MockUserRepository
	service UserService
}

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UserServiceSuite))
}

func (s *UserServiceSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())

	s.repo = mockRepository.NewMockUserRepository(s.ctrl)
	s.service = newUserService(s.repo)
}

func (s *UserServiceSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *UserServiceSuite) TestGetAllUsersSuccess() {
	ctx := context.Background()

	expectedData := []repository.User{
		{Id: "1", FirstName: "John", LastName: "Doe"},
		{Id: "2", FirstName: "John", LastName: "Doe"},
	}
	s.repo.EXPECT().GetAllUsers(ctx).Return(expectedData, nil)

	actual, err := s.service.GetAllUsers(ctx)
	s.NoError(err)
	s.Equal(expectedData, actual)
}

func (s *UserServiceSuite) TestGetAllUsersError() {
	ctx := context.Background()

	expectedError := errors.New("some error")
	s.repo.EXPECT().GetAllUsers(ctx).Return(nil, expectedError)

	var expectedData []repository.User
	actual, err := s.service.GetAllUsers(ctx)
	s.Error(err)
	s.Equal(expectedData, actual)
}

func (s *UserServiceSuite) TestCreateUserSuccess() {
	ctx := context.Background()
	user := repository.User{
		FirstName: "John",
		LastName:  "Doe",
	}
	body := UserCreateData{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	s.repo.EXPECT().CreateUser(ctx, user).Return(nil)

	err := s.service.CreateUser(ctx, body)
	s.NoError(err)
}

func (s *UserServiceSuite) TestCreateUserError() {
	ctx := context.Background()
	user := repository.User{
		FirstName: "John",
		LastName:  "Doe",
	}
	body := UserCreateData{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	expectedError := errors.New("some error")
	s.repo.EXPECT().CreateUser(ctx, user).Return(expectedError)

	err := s.service.CreateUser(ctx, body)
	s.Error(err)
}
