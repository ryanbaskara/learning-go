package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ryanbaskara/learning-go/entity"
	"github.com/ryanbaskara/learning-go/usecase"
	mockusecase "github.com/ryanbaskara/learning-go/usecase/mocks"
	"github.com/stretchr/testify/suite"
)

type createUserSuite struct {
	suite.Suite
	ctx context.Context

	ctrl     *gomock.Controller
	userRepo *mockusecase.MockUserRepository
	usecase  *usecase.Usecase

	req *entity.CreateUserRequest
}

func TestCreateUserSuite(t *testing.T) {
	suite.Run(t, new(createUserSuite))
}

func (s *createUserSuite) SetupSubTest() {
	s.ctx = context.Background()
	s.ctrl = gomock.NewController(s.T())

	s.userRepo = mockusecase.NewMockUserRepository(s.ctrl)
	s.usecase = usecase.NewUsecase(s.userRepo, nil)
	s.req = &entity.CreateUserRequest{
		Name:        "John",
		Email:       "john@email.com",
		PhoneNumber: "08999123",
	}
}

func (s *createUserSuite) TearDownSubTest() {
	s.ctrl.Finish()
}

func (s *createUserSuite) TestCreateUser_PositiveCases() {
	s.Run("Successfully crate user", func() {
		gomock.InOrder(
			s.userRepo.EXPECT().CreateUser(s.ctx, gomock.Any()).Do(func(_ context.Context, user *entity.User) {
				user.ID = 1
			}).Return(nil),
		)

		user, err := s.usecase.CreateUser(s.ctx, s.req)

		s.Nil(err)
		s.Equal(int64(1), user.ID)
		s.Equal(s.req.Name, user.Name)
		s.Equal(s.req.Email, user.Email)
		s.Equal(s.req.PhoneNumber, user.PhoneNumber)
	})
}

func (s *createUserSuite) TestCreateUser_NegativeCases() {
	s.Run("Error when validate request", func() {
		s.req.Name = ""

		user, err := s.usecase.CreateUser(s.ctx, s.req)

		s.Nil(user)
		s.Equal("Name is required", err.Error())
	})

	s.Run("Error call repository create user", func() {
		mockErr := errors.New("mock error")

		gomock.InOrder(
			s.userRepo.EXPECT().CreateUser(s.ctx, gomock.Any()).Return(mockErr),
		)

		user, err := s.usecase.CreateUser(s.ctx, s.req)

		s.Nil(user)
		s.Equal(mockErr, err)
	})
}
