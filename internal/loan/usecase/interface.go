package usecase

import (
	"xyz-finance-api/internal/loan/domain"
)

type LoanCommandUsecaseInterface interface {
}

type LoanQueryUsecaseInterface interface {
	GetAllLoans(userID string) ([]domain.Loan, error)
	GetLoanByID(loanID string)
}
