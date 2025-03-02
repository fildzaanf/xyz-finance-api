package usecase

import (
	"errors"
	"testing"
	"xyz-finance-api/internal/loan/domain"
	"xyz-finance-api/pkg/constant"
	mocks "xyz-finance-api/test"

	"github.com/stretchr/testify/assert"
)

func TestGetLoanByID(t *testing.T) {
	mockLoanQueryRepo := new(mocks.LoanQueryRepositoryInterface)
	mockLoanCommandRepo := new(mocks.LoanCommandRepositoryInterface)
	loanQueryUsecase := NewLoanQueryUsecase(mockLoanCommandRepo, mockLoanQueryRepo)

	t.Run("Valid GetLoanByID", func(t *testing.T) {
		loanID := "loan1234"

		mockLoanQueryRepo.On("GetLoanByID", loanID).Return(domain.Loan{}, nil)

		loan, err := loanQueryUsecase.GetLoanByID(loanID)

		mockLoanQueryRepo.AssertExpectations(t)
		assert.Nil(t, err)
		assert.NotNil(t, loan)
	})

	t.Run("Invalid GetLoanByID (Empty ID)", func(t *testing.T) {
		loanID := ""

		_, err := loanQueryUsecase.GetLoanByID(loanID)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_ID_INVALID)

		mockLoanQueryRepo.AssertExpectations(t)
	})

	t.Run("GetLoanByID Data Not Found", func(t *testing.T) {
		loanID := "nonexistent-loan-id"

		mockLoanQueryRepo.On("GetLoanByID", loanID).Return(domain.Loan{}, errors.New(constant.ERROR_DATA_EMPTY))

		loan, err := loanQueryUsecase.GetLoanByID(loanID)

		mockLoanQueryRepo.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_DATA_EMPTY)
		assert.Equal(t, domain.Loan{}, loan)
	})
}

func TestGetAllLoans(t *testing.T) {
	mockLoanQueryRepo := new(mocks.LoanQueryRepositoryInterface)
	mockLoanCommandRepo := new(mocks.LoanCommandRepositoryInterface)
	loanQueryUsecase := NewLoanQueryUsecase(mockLoanCommandRepo, mockLoanQueryRepo)

	t.Run("Valid GetAllLoans", func(t *testing.T) {
		userID := "user1234"
		mockLoans := []domain.Loan{
			{ID: "loan1", UserID: userID, Tenor: 1, LimitAmount: 1000, Status: "valid"},
			{ID: "loan2", UserID: userID, Tenor: 1, LimitAmount: 2000, Status: "valid"},
		}

		mockLoanQueryRepo.On("GetAllLoans", userID).Return(mockLoans, nil)

		loans, err := loanQueryUsecase.GetAllLoans(userID)

		mockLoanQueryRepo.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, len(loans), 2)
		assert.Equal(t, loans[0].LimitAmount, 1000)
	})

	t.Run("Invalid GetAllLoans (Empty UserID)", func(t *testing.T) {
		userID := ""

		_, err := loanQueryUsecase.GetAllLoans(userID)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_ID_INVALID)

		mockLoanQueryRepo.AssertExpectations(t)
	})

	t.Run("GetAllLoans Data Not Found", func(t *testing.T) {
		userID := "nonexistent-user-id"

		mockLoanQueryRepo.On("GetAllLoans", userID).Return(nil, errors.New(constant.ERROR_DATA_EMPTY))

		loans, err := loanQueryUsecase.GetAllLoans(userID)

		mockLoanQueryRepo.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_DATA_EMPTY)
		assert.Nil(t, loans)
	})
}
