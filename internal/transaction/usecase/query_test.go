package usecase

import (
	"errors"
	"testing"
	"xyz-finance-api/internal/transaction/domain"
	du "xyz-finance-api/internal/user/domain"
	"xyz-finance-api/pkg/constant"
	mocks "xyz-finance-api/test"

	"github.com/stretchr/testify/assert"
)

func TestGetTransactionByID(t *testing.T) {
	mockTransactionQueryRepo := new(mocks.TransactionQueryRepositoryInterface)
	mockTransactionCommandRepo := new(mocks.TransactionCommandRepositoryInterface)
	mockUserQueryRepo := new(mocks.UserQueryRepositoryInterface)
	transactionQueryUsecase := NewTransactionQueryUsecase(mockTransactionQueryRepo, mockTransactionCommandRepo, mockUserQueryRepo)

	t.Run("Valid GetTransactionByID", func(t *testing.T) {
		transactionID := "transaction1234"
		userID := "user1234"

		mockTransactionQueryRepo.On("GetTransactionByID", transactionID, userID).Return(domain.Transaction{ID: transactionID}, nil)
		mockUserQueryRepo.On("GetUserByID", userID).Return(du.User{ID: userID}, nil)

		transaction, err := transactionQueryUsecase.GetTransactionByID(transactionID, userID)

		mockTransactionQueryRepo.AssertExpectations(t)
		mockUserQueryRepo.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, transaction.ID, transactionID)
	})

	t.Run("Invalid GetTransactionByID (Empty ID)", func(t *testing.T) {
		transactionID := ""
		userID := "user1234"

		_, err := transactionQueryUsecase.GetTransactionByID(transactionID, userID)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_ID_INVALID)
	})

	t.Run("GetTransactionByID Transaction Not Found", func(t *testing.T) {
		transactionID := "nonexistent-transaction-id"
		userID := "user1234"

		mockTransactionQueryRepo.On("GetTransactionByID", transactionID, userID).Return(domain.Transaction{}, errors.New(constant.ERROR_ID_NOTFOUND))

		transaction, err := transactionQueryUsecase.GetTransactionByID(transactionID, userID)

		mockTransactionQueryRepo.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_ID_NOTFOUND)
		assert.Equal(t, domain.Transaction{}, transaction)
	})

	t.Run("GetTransactionByID User Not Found", func(t *testing.T) {
		transactionID := "transaction1234"
		userID := "user1234"

		mockTransactionQueryRepo.On("GetTransactionByID", transactionID, userID).Return(domain.Transaction{ID: transactionID}, nil)
		mockUserQueryRepo.On("GetUserByID", userID).Return(du.User{}, errors.New(constant.ERROR_ROLE_ACCESS))

		_, err := transactionQueryUsecase.GetTransactionByID(transactionID, userID)

		mockTransactionQueryRepo.AssertExpectations(t)
		mockUserQueryRepo.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_ROLE_ACCESS)
	})
}

func TestGetAllTransactions(t *testing.T) {
	mockTransactionQueryRepo := new(mocks.TransactionQueryRepositoryInterface)
	mockTransactionCommandRepo := new(mocks.TransactionCommandRepositoryInterface)
	mockUserQueryRepo := new(mocks.UserQueryRepositoryInterface)
	transactionQueryUsecase := NewTransactionQueryUsecase(mockTransactionQueryRepo, mockTransactionCommandRepo, mockUserQueryRepo)

	t.Run("Valid GetAllTransactions", func(t *testing.T) {
		userID := "user1234"
		mockTransactions := []domain.Transaction{
			{ID: "transaction1", LoanID: "loanid1", AssetName: "AssetName", TotalAmount: 1000, Tenor: 1,  OTRPrice: 1000,  AdminFee: 10000, Interest: 1000, Status: "completed"},
			{ID: "transaction2", LoanID: "loanid2", AssetName: "AssetName", TotalAmount: 2000, Tenor: 1, OTRPrice: 1000, AdminFee: 10000, Interest: 1000, Status: "completed"},
		} 
		
		mockTransactionQueryRepo.On("GetAllTransactions", userID).Return(mockTransactions, nil)
		mockUserQueryRepo.On("GetUserByID", userID).Return(du.User{ID: userID}, nil)

		transactions, err := transactionQueryUsecase.GetAllTransactions(userID)

		mockTransactionQueryRepo.AssertExpectations(t)
		mockUserQueryRepo.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Len(t, transactions, 2)
	})

	t.Run("Invalid GetAllTransactions (Empty UserID)", func(t *testing.T) {
		userID := ""

		_, err := transactionQueryUsecase.GetAllTransactions(userID)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_ID_INVALID)
	})

	t.Run("GetAllTransactions Data Not Found", func(t *testing.T) {
		userID := "nonexistent-user-id"

		mockTransactionQueryRepo.On("GetAllTransactions", userID).Return(nil, errors.New(constant.ERROR_DATA_EMPTY))

		transactions, err := transactionQueryUsecase.GetAllTransactions(userID)

		mockTransactionQueryRepo.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_DATA_EMPTY)
		assert.Nil(t, transactions)
	})

	t.Run("GetAllTransactions User Not Found", func(t *testing.T) {
		userID := "user1234"

		mockTransactionQueryRepo.On("GetAllTransactions", userID).Return([]domain.Transaction{}, nil)
		mockUserQueryRepo.On("GetUserByID", userID).Return(du.User{}, errors.New(constant.ERROR_ROLE_ACCESS))

		_, err := transactionQueryUsecase.GetAllTransactions(userID)

		mockTransactionQueryRepo.AssertExpectations(t)
		mockUserQueryRepo.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_ROLE_ACCESS)
	})
}
