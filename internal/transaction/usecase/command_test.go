package usecase

import (
	"errors"
	"testing"
	"xyz-finance-api/pkg/constant"
	mocks "xyz-finance-api/test"

	dl "xyz-finance-api/internal/loan/domain"
	"xyz-finance-api/internal/transaction/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransaction(t *testing.T) {
	mockTransactionCmdRepo := new(mocks.TransactionCommandRepositoryInterface)
	mockTransactionQryRepo := new(mocks.TransactionQueryRepositoryInterface)
	mockLoanQryRepo := new(mocks.LoanQueryRepositoryInterface)
	mockLoanCmdRepo := new(mocks.LoanCommandRepositoryInterface)

	usecase := NewTransactionCommandUsecase(mockTransactionCmdRepo, mockTransactionQryRepo, mockLoanQryRepo, mockLoanCmdRepo)

	t.Run("Create Transaction Success", func(t *testing.T) {
		loanID := "123"
		userID := "user-1"
		transaction := domain.Transaction{
			LoanID:    loanID,
			AssetName: "Car",
			OTRPrice:  5000000,
		}
		loan := dl.Loan{
			ID:          loanID,
			UserID:      userID,
			Status:      "valid",
			LimitAmount: 10000000,
			Tenor:       12,
		}
	
		expectedTransaction := transaction
		expectedTransaction.Status = "success" 
	
		mockLoanQryRepo.On("GetLoanByID", loanID).Return(loan, nil)
		mockTransactionCmdRepo.On("CreateTransaction", mock.Anything).Return(expectedTransaction, nil)
		mockLoanCmdRepo.On("UpdateLoanStatusByID", loanID, mock.Anything).Return(loan, nil)
	
		result, err := usecase.CreateTransaction(transaction, userID)
	
		assert.NoError(t, err)
		assert.Equal(t, "success", result.Status) 
		mockLoanQryRepo.AssertExpectations(t)
		mockTransactionCmdRepo.AssertExpectations(t)
		mockLoanCmdRepo.AssertExpectations(t)
	})
	

	t.Run("Create Transaction With Loan Not Found", func(t *testing.T) {
		loanID := "invalid-loan-id"
		userID := "user-1"
		transaction := domain.Transaction{
			LoanID:    loanID,
			AssetName: "Car", 
			OTRPrice:  5000000,
		}
	
		mockLoanQryRepo.On("GetLoanByID", loanID).Return(dl.Loan{}, errors.New(constant.ERROR_ID_NOTFOUND))
	
		result, err := usecase.CreateTransaction(transaction, userID)
	
		assert.Error(t, err)
		assert.Equal(t, constant.ERROR_ID_NOTFOUND, err.Error()) 
		assert.Equal(t, domain.Transaction{}, result)
	
		mockLoanQryRepo.AssertExpectations(t) 
	})
	

	t.Run("Create Transaction With Empty Data", func(t *testing.T) {
		userID := "user1"
		transaction := domain.Transaction{
			LoanID:    "",
			AssetName: "",
			OTRPrice:  0,
		}
	
		_, err := usecase.CreateTransaction(transaction, userID)

		assert.Error(t, err)
		mockTransactionCmdRepo.AssertExpectations(t)
	})
	
}
