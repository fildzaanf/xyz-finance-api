package usecase

import (
	"errors"
	"testing"
	"xyz-finance-api/internal/user/domain"
	"xyz-finance-api/pkg/constant"
	mocks "xyz-finance-api/test"

	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	mockUserQueryRepo := new(mocks.UserQueryRepositoryInterface)
	mockUserCommandRepo := new(mocks.UserCommandRepositoryInterface)
	userQueryUsecase := NewUserQueryUsecase(mockUserCommandRepo, mockUserQueryRepo)

	t.Run("Valid GetUserByID", func(t *testing.T) {
		userID := "f0b24674-a9e0-4d59-9f53-eadc03dcba9a"

		mockUserQueryRepo.On("GetUserByID", userID).Return(domain.User{}, nil)

		user, err := userQueryUsecase.GetUserByID(userID)

		mockUserQueryRepo.AssertExpectations(t)
		assert.Nil(t, err)
		assert.NotNil(t, user)
	})

	t.Run("Invalid GetUserByID (Empty ID)", func(t *testing.T) {
		userID := ""

		_, err := userQueryUsecase.GetUserByID(userID)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_ID_INVALID)

		mockUserQueryRepo.AssertExpectations(t)
	})

	t.Run("GetUserByID Data Not Found", func(t *testing.T) {
		userID := "nonexistent-id"

		mockUserQueryRepo.On("GetUserByID", userID).Return(domain.User{}, errors.New(constant.ERROR_DATA_EMPTY))

		user, err := userQueryUsecase.GetUserByID(userID)

		mockUserQueryRepo.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_DATA_EMPTY)
		assert.Equal(t, domain.User{}, user)
	})
}
