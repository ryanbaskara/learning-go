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

	ctrl    *gomock.Controller
	repo    *mockusecase.MockRepository
	usecase *usecase.Usecase
}

func TestGetUserSuite(t *testing.T) {
	suite.Run(t, new(getUserSuite))
}

func (s *getUserSuite) SetupSubTest() {
	s.ctx = context.Background()
	s.ctrl = gomock.NewController(s.T())

	s.repo = mockusecase.NewMockRepository(s.ctrl)
	s.usecase = usecase.NewUsecase(s.repo)
}

func (s *getUserSuite) TearDownSubTest() {
	s.ctrl.Finish()
}

func (s *getUserSuite) TestGetUser_PositiveCases() {
	s.Run("Successfully get user", func() {
		mockUser := &entity.User{
			ID:   1,
			Name: "John",
		}

		gomock.InOrder(
			s.repo.EXPECT().GetUser(s.ctx, int64(1)).Return(mockUser, nil),
		)

		user, err := s.usecase.GetUser(s.ctx, 1)

		s.Nil(err)
		s.Equal(mockUser, user)
	})
}

func (s *getUserSuite) TestGetUser_NegativeCases() {
	s.Run("Error call repository get user", func() {
		mockErr := errors.New("mock error")

		gomock.InOrder(
			s.repo.EXPECT().GetUser(s.ctx, int64(1)).Return(nil, mockErr),
		)

		user, err := s.usecase.GetUser(s.ctx, 1)

		s.Nil(user)
		s.Equal(mockErr, err)
	})
}
