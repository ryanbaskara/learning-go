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

type getUserSuite struct {
	suite.Suite
	ctx context.Context

	ctrl          *gomock.Controller
	userRepo      *mockusecase.MockUserRepository
	userCacheRepo *mockusecase.MockUserCacheRepository
	usecase       *usecase.Usecase
}

func TestGetUserSuite(t *testing.T) {
	suite.Run(t, new(getUserSuite))
}

func (s *getUserSuite) SetupSubTest() {
	s.ctx = context.Background()
	s.ctrl = gomock.NewController(s.T())

	s.userRepo = mockusecase.NewMockUserRepository(s.ctrl)
	s.userCacheRepo = mockusecase.NewMockUserCacheRepository(s.ctrl)
	s.usecase = usecase.NewUsecase(s.userRepo, s.userCacheRepo)
}

func (s *getUserSuite) TearDownSubTest() {
	s.ctrl.Finish()
}

func (s *getUserSuite) TestGetUser_PositiveCases() {
	s.Run("Successfully get user from mysql", func() {
		mockUser := &entity.User{
			ID:   1,
			Name: "John",
		}

		gomock.InOrder(
			s.userCacheRepo.EXPECT().GetUser(s.ctx, int64(1)).Return(nil, nil),
			s.userRepo.EXPECT().GetUser(s.ctx, int64(1)).Return(mockUser, nil),
			s.userCacheRepo.EXPECT().SetUser(s.ctx, mockUser).Return(nil),
		)

		user, err := s.usecase.GetUser(s.ctx, 1)

		s.Nil(err)
		s.Equal(mockUser.ID, user.ID)
		s.Equal(mockUser.Name, user.Name)
		s.Equal("mysql", user.Source)
	})

	s.Run("Successfully get user from redis", func() {
		mockUser := &entity.User{
			ID:   1,
			Name: "John",
		}

		gomock.InOrder(
			s.userCacheRepo.EXPECT().GetUser(s.ctx, int64(1)).Return(mockUser, nil),
		)

		user, err := s.usecase.GetUser(s.ctx, 1)

		s.Nil(err)
		s.Equal(mockUser.ID, user.ID)
		s.Equal(mockUser.Name, user.Name)
		s.Equal("redis", user.Source)
	})
}

func (s *getUserSuite) TestGetUser_NegativeCases() {
	mockErr := errors.New("mock error")

	s.Run("Error call cache repository get user", func() {
		gomock.InOrder(
			s.userCacheRepo.EXPECT().GetUser(s.ctx, int64(1)).Return(nil, mockErr),
		)

		user, err := s.usecase.GetUser(s.ctx, 1)

		s.Nil(user)
		s.Equal(mockErr, err)
	})

	s.Run("Error call repository get user", func() {
		gomock.InOrder(
			s.userCacheRepo.EXPECT().GetUser(s.ctx, int64(1)).Return(nil, nil),
			s.userRepo.EXPECT().GetUser(s.ctx, int64(1)).Return(nil, mockErr),
		)

		user, err := s.usecase.GetUser(s.ctx, 1)

		s.Nil(user)
		s.Equal(mockErr, err)
	})
}
