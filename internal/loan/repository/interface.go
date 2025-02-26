package repository

import "xyz-finance-api/internal/loan/domain"

type LoanCommandRepositoryInterface interface {
}

type LoanQueryRepositoryInterface interface {
	GetAllLoans() ([]domain.Loan, error)
}
