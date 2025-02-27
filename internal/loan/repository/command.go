package repository

import (
	"xyz-finance-api/internal/loan/domain"

	"gorm.io/gorm"
)

type loanCommandRepository struct {
	db *gorm.DB
}

func NewLoanCommandRepository(db *gorm.DB) LoanCommandRepositoryInterface {
	return &loanCommandRepository{
		db: db,
	}
}

func (lcr *loanCommandRepository) CreateLoan(loan domain.Loan) (domain.Loan, error) {
	loanEntity := domain.LoanDomainToLoanEntity(loan)

	result := lcr.db.Create(&loanEntity)
	if result.Error != nil {
		return domain.Loan{}, result.Error
	}

	loanDomain := domain.LoanEntityToLoanDomain(loanEntity)

	return loanDomain, nil
}
