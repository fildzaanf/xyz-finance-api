package usecase

import (
	"errors"
	"testing"
	"time"

	"xyz-finance-api/internal/installment/domain"
	dt "xyz-finance-api/internal/transaction/domain"
	du "xyz-finance-api/internal/user/domain"
	"xyz-finance-api/pkg/constant"
	mocks "xyz-finance-api/test"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateInstallment(t *testing.T) {
	mockInstallmentCommandRepo := new(mocks.InstallmentCommandRepositoryInterface)
	mockInstallmentQueryRepo := new(mocks.InstallmentQueryRepositoryInterface)
	mockTransactionQueryRepo := new(mocks.TransactionQueryRepositoryInterface)
	mockUserQueryRepo := new(mocks.UserQueryRepositoryInterface)

	installmentCommandUsecase := NewInstallmentCommandUsecase(mockInstallmentCommandRepo, mockInstallmentQueryRepo, mockTransactionQueryRepo, mockUserQueryRepo)

	t.Run("Create Installment Success", func(t *testing.T) {
		userID := "user123"
		request := domain.Installment{
			TransactionID: "txn123",
		}
	
		mockTransactionQueryRepo.On("GetTransactionByID", request.TransactionID, userID).Return(dt.Transaction{
			ID:          "txn123",
			TotalAmount: 1000,
			Tenor:       5,
			Status:      "success",
		}, nil).Once()
	
		mockUserQueryRepo.On("GetUserByID", userID).Return(du.User{ID: userID}, nil).Once()
	
		mockInstallmentQueryRepo.On("CountInstallmentsByTransactionID", request.TransactionID).Return(0, nil).Once()
	
		mockInstallmentCommandRepo.On("CreateInstallment", mock.Anything, userID).Return(domain.Installment{
			TransactionID:     request.TransactionID,
			InstallmentNumber: 1,
			Amount:            200,
			Status:            "unpaid",
			CreatedAt:         time.Now(),
			DueDate:           time.Now().AddDate(0, 1, 0),
		}, nil).Once()  
	

		installments, err := installmentCommandUsecase.CreateInstallment(request, userID)
	
		assert.Nil(t, err)
	
		assert.Equal(t, 5, len(installments))

		mockInstallmentCommandRepo.AssertExpectations(t)
		mockInstallmentQueryRepo.AssertExpectations(t)
		mockTransactionQueryRepo.AssertExpectations(t)
		mockUserQueryRepo.AssertExpectations(t)
	})
	

	t.Run("Create Installment With Invalid Transaction", func(t *testing.T) {
		userID := "user123"
		request := domain.Installment{
			TransactionID: "txn123",
		}

		mockTransactionQueryRepo.On("GetTransactionByID", request.TransactionID, userID).Return(dt.Transaction{}, errors.New(constant.ERROR_ID_NOTFOUND))

		_, err := installmentCommandUsecase.CreateInstallment(request, userID)

		assert.Error(t, err)
		mockTransactionQueryRepo.AssertExpectations(t)
	})

	t.Run("Create Installment With Invalid User", func(t *testing.T) {
		userID := "user123"
		request := domain.Installment{
			TransactionID: "txn123",
		}

		mockTransactionQueryRepo.On("GetTransactionByID", request.TransactionID, userID).Return(dt.Transaction{
			ID:          "txn123",
			TotalAmount: 1000,
			Tenor:       5,
			Status:      "success",
		}, nil)

		mockUserQueryRepo.On("GetUserByID", userID).Return(du.User{}, errors.New(constant.ERROR_ROLE_ACCESS))

		_, err := installmentCommandUsecase.CreateInstallment(request, userID)

		assert.Error(t, err)
		mockUserQueryRepo.AssertExpectations(t)
	})

	t.Run("Installment Count Exceeds Tenor Limit", func(t *testing.T) {
		userID := "user123"
		request := domain.Installment{
			TransactionID: "txn123",
		}

		mockTransactionQueryRepo.On("GetTransactionByID", request.TransactionID, userID).Return(dt.Transaction{
			ID:           "txn123",
			TotalAmount:  1000,
			Tenor:        3,
			Status:       "success",
		}, nil)

		mockInstallmentQueryRepo.On("CountInstallmentsByTransactionID", request.TransactionID).Return(3, nil)

		_, err := installmentCommandUsecase.CreateInstallment(request, userID)

		assert.Error(t, err)
		mockInstallmentQueryRepo.AssertExpectations(t)
	})
}

func TestUpdateInstallmentStatusByID(t *testing.T) {
	mockInstallmentCommandRepo := new(mocks.InstallmentCommandRepositoryInterface)
	mockInstallmentQueryRepo := new(mocks.InstallmentQueryRepositoryInterface)
	mockUserQueryRepo := new(mocks.UserQueryRepositoryInterface)

	installmentCommandUsecase := NewInstallmentCommandUsecase(mockInstallmentCommandRepo, mockInstallmentQueryRepo, nil, mockUserQueryRepo)

	t.Run("Valid Update Installment Status", func(t *testing.T) {
		installmentID := "inst123"
		userID := "user123"
		installment := domain.Installment{
			Status: "paid",
		}

		mockInstallmentQueryRepo.On("GetInstallmentByID", installmentID, userID).Return(domain.Installment{
			ID:     installmentID,
			Status: "unpaid",
		}, nil)

		mockUserQueryRepo.On("GetUserByID", userID).Return(du.User{ID: userID}, nil)

		mockInstallmentCommandRepo.On("UpdateInstallmentStatusByID", installmentID, mock.Anything).Return(installment, nil)

		updatedInstallment, err := installmentCommandUsecase.UpdateInstallmentStatusByID(installmentID, installment, userID)

		assert.Nil(t, err)
		assert.Equal(t, "paid", updatedInstallment.Status)
		mockInstallmentQueryRepo.AssertExpectations(t)
		mockUserQueryRepo.AssertExpectations(t)
		mockInstallmentCommandRepo.AssertExpectations(t)
	})

	t.Run("Update Installment Status with Invalid Status", func(t *testing.T) {
		installmentID := "inst123"
		userID := "user123"
		installment := domain.Installment{
			Status: "invalid",
		}

		_, err := installmentCommandUsecase.UpdateInstallmentStatusByID(installmentID, installment, userID)

		assert.Error(t, err)
	})

	t.Run("Update Installment Status with Invalid Installment ID", func(t *testing.T) {
		installmentID := "inst123"
		userID := "user123"
		installment := domain.Installment{
			Status: "paid",
		}

		mockInstallmentQueryRepo.On("GetInstallmentByID", installmentID, userID).Return(domain.Installment{}, errors.New(constant.ERROR_ID_NOTFOUND))

		_, err := installmentCommandUsecase.UpdateInstallmentStatusByID(installmentID, installment, userID)

		assert.Error(t, err)
	})

	t.Run("Update Installment Status with Invalid User", func(t *testing.T) {
		installmentID := "inst123"
		userID := "user123"
		installment := domain.Installment{
			Status: "paid",
		}

		mockInstallmentQueryRepo.On("GetInstallmentByID", installmentID, userID).Return(domain.Installment{
			ID:     installmentID,
			Status: "unpaid",
		}, nil)

		mockUserQueryRepo.On("GetUserByID", userID).Return(du.User{}, errors.New(constant.ERROR_ROLE_ACCESS))

		_, err := installmentCommandUsecase.UpdateInstallmentStatusByID(installmentID, installment, userID)

		assert.Error(t, err)
	})
}
