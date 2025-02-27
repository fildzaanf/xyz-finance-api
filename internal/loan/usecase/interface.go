package usecase

import (
	"xyz-finance-api/internal/loan/domain"
)

type LoanCommandUsecaseInterface interface {
	CreateLoan(loan domain.Loan) (domain.Loan, error)
}

type LoanQueryUsecaseInterface interface {
	GetAllLoans(userID string) ([]domain.Loan, error)
	GetLoanByID(id string) (domain.Loan, error)
}
