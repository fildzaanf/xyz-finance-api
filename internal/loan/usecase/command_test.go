package usecase

import (
	"errors"
	"testing"
	"xyz-finance-api/internal/loan/domain"
	du "xyz-finance-api/internal/user/domain"
	"xyz-finance-api/pkg/constant"
	"xyz-finance-api/pkg/validator"
	mocks "xyz-finance-api/test"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateLoan(t *testing.T) {
	mockLoanCommandRepo := new(mocks.LoanCommandRepositoryInterface)
	mockLoanQueryRepo := new(mocks.LoanQueryRepositoryInterface)
	mockUserQueryRepo := new(mocks.UserQueryRepositoryInterface)

	loanCommandUsecase := NewLoanCommandUsecase(mockLoanCommandRepo, mockLoanQueryRepo, mockUserQueryRepo)

	t.Run("Create Loan - Missing Required Fields", func(t *testing.T) {
		loan := domain.Loan{
			Tenor:  0,
		}

		_, err := loanCommandUsecase.CreateLoan(loan)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "user_id and tenor cannot be empty")
	})

	t.Run("Create Loan - Invalid Tenor", func(t *testing.T) {
		loan := domain.Loan{
			UserID: "user123",
			Tenor:  5,
		}

		_, err := loanCommandUsecase.CreateLoan(loan)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "the tenor is not valid, it can only be 1, 2, 3, or 6")
	})

	t.Run("Create Loan - Existing Loan for User", func(t *testing.T) {
		loan := domain.Loan{
			UserID: "f0b24674-a9e0-4d59-9f53-eadc03dcba9a",
			Tenor:  1,
		}

		mockLoanQueryRepo.On("GetLoanByUserID", loan.UserID, loan.Tenor).Return(domain.Loan{}, nil)

		_, err := loanCommandUsecase.CreateLoan(loan)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "user already has a loan with this tenor")

		mockLoanQueryRepo.AssertExpectations(t)
	})

	t.Run("Create Loan - User Not Found", func(t *testing.T) {
		loan := domain.Loan{
			UserID: "user123",
			Tenor:  3,
		}

		mockLoanQueryRepo.On("GetLoanByUserID", loan.UserID, loan.Tenor).Return(domain.Loan{}, nil)
		mockUserQueryRepo.On("GetUserByID", loan.UserID).Return(du.User{}, errors.New(constant.ERROR_ID_NOTFOUND))

		_, err := loanCommandUsecase.CreateLoan(loan)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_ID_NOTFOUND)

		mockUserQueryRepo.AssertExpectations(t)
	})

	t.Run("Create Loan - Loan Limit Calculation", func(t *testing.T) {
		loan := domain.Loan{
			UserID: "user123",
			Tenor:  3,
		}

		user := du.User{
			ID:     "user123",
			Salary: 5000,
		}

		mockLoanQueryRepo.On("GetLoanByUserID", loan.UserID, loan.Tenor).Return(domain.Loan{}, nil)
		mockUserQueryRepo.On("GetUserByID", loan.UserID).Return(user, nil)

		limitAmount := validator.CalculateLoanLimit(user.Salary, loan.Tenor)
		assert.Equal(t, limitAmount, 15000)

		mockLoanQueryRepo.AssertExpectations(t)
		mockUserQueryRepo.AssertExpectations(t)
	})

}

func TestUpdateLoanStatusByID(t *testing.T) {
	mockLoanCommandRepo := new(mocks.LoanCommandRepositoryInterface)
	mockLoanQueryRepo := new(mocks.LoanQueryRepositoryInterface)
	mockUserQueryRepo := new(mocks.UserQueryRepositoryInterface)

	loanCommandUsecase := NewLoanCommandUsecase(mockLoanCommandRepo, mockLoanQueryRepo, mockUserQueryRepo)

	t.Run("Update Loan Status - Loan Not Found", func(t *testing.T) {
		loanID := "loan123"
		loan := domain.Loan{
			Status: "invalid",
		}

		mockLoanQueryRepo.On("GetLoanByID", loanID).Return(domain.Loan{}, errors.New(constant.ERROR_ID_NOTFOUND))

		_, err := loanCommandUsecase.UpdateLoanStatusByID(loanID, loan)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_ID_NOTFOUND)

		mockLoanQueryRepo.AssertExpectations(t)
	})

	t.Run("Update Loan Status - Invalid Status", func(t *testing.T) {
		loanID := "loan123"
		loan := domain.Loan{
			Status: "rejected",
		}

		mockLoanQueryRepo.On("GetLoanByID", loanID).Return(domain.Loan{ID: loanID}, nil)

		_, err := loanCommandUsecase.UpdateLoanStatusByID(loanID, loan)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "invalid status update request")

		mockLoanQueryRepo.AssertExpectations(t)
	})

	t.Run("Update Loan Status - Successful Update", func(t *testing.T) {
		loanID := "loan123"
		loan := domain.Loan{
			Status: "valid",
		}

		existingLoan := domain.Loan{
			ID:     loanID,
			Status: "invalid",
		}

		mockLoanQueryRepo.On("GetLoanByID", loanID).Return(existingLoan, nil)
		mockLoanCommandRepo.On("UpdateLoanStatusByID", loanID, mock.Anything).Return(domain.Loan{ID: loanID, Status: "valid"}, nil)

		updatedLoan, err := loanCommandUsecase.UpdateLoanStatusByID(loanID, loan)

		assert.Nil(t, err)
		assert.Equal(t, "valid", updatedLoan.Status)

		mockLoanQueryRepo.AssertExpectations(t)
		mockLoanCommandRepo.AssertExpectations(t)
	})
}
