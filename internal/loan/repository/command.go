package repository

import (
	"errors"
	"time"
	"xyz-finance-api/internal/loan/domain"
	"xyz-finance-api/pkg/constant"

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

func (lcr *loanCommandRepository) UpdateLoanStatusByID(id string, loan domain.Loan) (domain.Loan, error) {

	var existingLoan domain.Loan
	err := lcr.db.Where("id = ?", id).First(&existingLoan).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Loan{}, errors.New(constant.ERROR_ID_NOTFOUND)
		}
		return domain.Loan{}, err
	}


	loan.UpdatedAt = time.Now()
	result := lcr.db.Model(&domain.Loan{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     loan.Status,
			"updated_at": loan.UpdatedAt,
		})

	if result.Error != nil {
		return domain.Loan{}, result.Error
	}

	if result.RowsAffected == 0 {
		return domain.Loan{}, errors.New("failed to update loan")
	}

	return loan, nil
}
