package repository

import (
	"errors"
	"xyz-finance-api/internal/loan/domain"
	"xyz-finance-api/internal/loan/entity"

	"gorm.io/gorm"
)

type loanQueryRepository struct {
	db *gorm.DB
}

func NewLoanQueryRepository(db *gorm.DB) LoanQueryRepositoryInterface {
	return &loanQueryRepository{
		db: db,
	}
}

func (lqr *loanQueryRepository) GetLoanByID(id string) (domain.Loan, error) {
	var loanEntity entity.Loan
	result := lqr.db.Where("id = ?", id).First(&loanEntity)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.Loan{}, errors.New("loan not found")
		}
		return domain.Loan{}, result.Error
	}

	loanDomain := domain.LoanEntityToLoanDomain(loanEntity)

	return loanDomain, nil
}

func (lqr *loanQueryRepository) GetAllLoans(userID string) ([]domain.Loan, error) {
	var loanEntities []entity.Loan

	result := lqr.db.Where("user_id = ?", userID).Find(&loanEntities)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(loanEntities) == 0 {
		return nil, errors.New("no loans found for this user")
	}

	var loans []domain.Loan
	for _, loanEntity := range loanEntities {
		loans = append(loans, domain.LoanEntityToLoanDomain(loanEntity))
	}

	return loans, nil
}

func (lr *loanQueryRepository) GetLoanByUserID(userID string, tenor int) (domain.Loan, error) {
	var loan domain.Loan
	err := lr.db.Where("user_id = ? AND tenor = ?", userID, tenor).First(&loan).Error
	if err != nil {
		return domain.Loan{}, err
	}
	return loan, nil
}
