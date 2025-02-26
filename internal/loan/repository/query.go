package repository

import (
	"gorm.io/gorm"
	"xyz-finance-api/internal/loan/domain"
	"xyz-finance-api/internal/loan/entity"
)

type loanQueryRepository struct {
	db *gorm.DB
}

func NewLoanQueryRepository(db *gorm.DB) LoanQueryRepositoryInterface {
	return &loanQueryRepository{
		db: db,
	}
}

func (lqr *loanQueryRepository) GetAllLoans() ([]domain.Loan, error) {
	var loanEntities []entity.Loan
	result := lqr.db.Find(&loanEntities)
	if result.Error != nil {
		return nil, result.Error
	}

	var loans []domain.Loan
	for _, loanEntity := range loanEntities {
		loanDomain := domain.LoanEntityToLoanDomain(loanEntity)
		loans = append(loans, loanDomain)
	}

	return loans, nil
}
