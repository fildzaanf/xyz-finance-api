package usecase

import (
	"errors"
	"testing"

	du "xyz-finance-api/internal/user/domain"
	"xyz-finance-api/internal/installment/domain"

	"xyz-finance-api/pkg/constant"
	mocks "xyz-finance-api/test"

	"github.com/stretchr/testify/assert"
)

func TestGetInstallmentByID(t *testing.T) {
	mockInstallmentQueryRepo := new(mocks.InstallmentQueryRepositoryInterface)
	mockInstallmentCommandRepo := new(mocks.InstallmentCommandRepositoryInterface)
	mockUserQueryRepo := new(mocks.UserQueryRepositoryInterface)
	installmentQueryUsecase := NewInstallmentQueryUsecase(mockInstallmentCommandRepo, mockInstallmentQueryRepo, mockUserQueryRepo)

	t.Run("Valid GetInstallmentByID", func(t *testing.T) {
		installmentID := "installment-id"
		userID := "user-id"

		mockInstallmentQueryRepo.On("GetInstallmentByID", installmentID, userID).Return(domain.Installment{ID: installmentID, TransactionID: "transaction-id"}, nil)
		mockUserQueryRepo.On("GetUserByID", userID).Return(du.User{ID: userID}, nil)

		installment, err := installmentQueryUsecase.GetInstallmentByID(installmentID, userID)

		mockInstallmentQueryRepo.AssertExpectations(t)
		mockUserQueryRepo.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, installmentID, installment.ID)
	})

	t.Run("Invalid GetInstallmentByID (Empty ID)", func(t *testing.T) {
		installmentID := ""
		userID := "user-id"

		_, err := installmentQueryUsecase.GetInstallmentByID(installmentID, userID)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_ID_INVALID)

		mockInstallmentQueryRepo.AssertExpectations(t)
		mockUserQueryRepo.AssertExpectations(t)
	})

	t.Run("GetInstallmentByID TransactionID Empty", func(t *testing.T) {
		installmentID := "installment-id"
		userID := "user-id"

		mockInstallmentQueryRepo.On("GetInstallmentByID", installmentID, userID).Return(domain.Installment{ID: installmentID, TransactionID: ""}, nil)

		_, err := installmentQueryUsecase.GetInstallmentByID(installmentID, userID)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "transactionID cannot be empty")

		mockInstallmentQueryRepo.AssertExpectations(t)
	})

	t.Run("GetInstallmentByID Not Found", func(t *testing.T) {
		installmentID := "invalid-id"
		userID := "user-id"

		mockInstallmentQueryRepo.On("GetInstallmentByID", installmentID, userID).Return(domain.Installment{}, errors.New(constant.ERROR_ID_NOTFOUND))

		_, err := installmentQueryUsecase.GetInstallmentByID(installmentID, userID)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_ID_NOTFOUND)

		mockInstallmentQueryRepo.AssertExpectations(t)
	})

	t.Run("GetInstallmentByID User Not Found", func(t *testing.T) {
		installmentID := "installment-id"
		userID := "user-id"

		mockInstallmentQueryRepo.On("GetInstallmentByID", installmentID, userID).Return(domain.Installment{ID: installmentID, TransactionID: "transaction-id"}, nil)
		mockUserQueryRepo.On("GetUserByID", userID).Return(du.User{}, errors.New(constant.ERROR_DATA_EMPTY))

		_, err := installmentQueryUsecase.GetInstallmentByID(installmentID, userID)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_DATA_EMPTY)

		mockInstallmentQueryRepo.AssertExpectations(t)
	})
}


func TestGetAllInstallments(t *testing.T) {
	mockInstallmentQueryRepo := new(mocks.InstallmentQueryRepositoryInterface)
	mockInstallmentCommandRepo := new(mocks.InstallmentCommandRepositoryInterface)
	mockUserQueryRepo := new(mocks.UserQueryRepositoryInterface)
	installmentQueryUsecase := NewInstallmentQueryUsecase(mockInstallmentCommandRepo, mockInstallmentQueryRepo, mockUserQueryRepo)

	t.Run("Valid GetAllInstallments", func(t *testing.T) {
		userID := "user-id"
		installments := []domain.Installment{
			{ID: "installment-1", TransactionID: "transaction-1"},
			{ID: "installment-2", TransactionID: "transaction-2"},
		}

		mockInstallmentQueryRepo.On("GetAllInstallments", userID).Return(installments, nil)
		mockUserQueryRepo.On("GetUserByID", userID).Return(du.User{ID: userID}, nil)

		result, err := installmentQueryUsecase.GetAllInstallments(userID)

		mockInstallmentQueryRepo.AssertExpectations(t)
		mockUserQueryRepo.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(result))
	})

	t.Run("Invalid GetAllInstallments (Empty ID)", func(t *testing.T) {
		userID := ""

		_, err := installmentQueryUsecase.GetAllInstallments(userID)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_ID_INVALID)

		mockInstallmentQueryRepo.AssertExpectations(t)
		mockUserQueryRepo.AssertExpectations(t)
	})

	t.Run("GetAllInstallments Data Not Found", func(t *testing.T) {
		userID := "user-id"

		mockInstallmentQueryRepo.On("GetAllInstallments", userID).Return(nil, errors.New(constant.ERROR_DATA_EMPTY))
		mockUserQueryRepo.On("GetUserByID", userID).Return(du.User{ID: userID}, nil)

		_, err := installmentQueryUsecase.GetAllInstallments(userID)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_DATA_EMPTY)

		mockInstallmentQueryRepo.AssertExpectations(t)
	})

	t.Run("GetAllInstallments User Not Found", func(t *testing.T) {
		userID := "user-id"

		mockInstallmentQueryRepo.On("GetAllInstallments", userID).Return([]domain.Installment{}, nil)
		mockUserQueryRepo.On("GetUserByID", userID).Return(du.User{}, errors.New(constant.ERROR_DATA_EMPTY))

		_, err := installmentQueryUsecase.GetAllInstallments(userID)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_DATA_EMPTY)

		mockInstallmentQueryRepo.AssertExpectations(t)
	})
}
