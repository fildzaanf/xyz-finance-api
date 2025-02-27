package repository

import "xyz-finance-api/internal/loan/domain"

type LoanCommandRepositoryInterface interface {
	CreateLoan(loan domain.Loan) (domain.Loan, error)
	UpdateLoanStatusByID(id string, loan domain.Loan) (domain.Loan, error)
}

type LoanQueryRepositoryInterface interface {
	GetAllLoans(userID string) ([]domain.Loan, error)
	GetLoanByID(id string) (domain.Loan, error)
}
