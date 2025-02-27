package repository

import (
	"errors"
	"xyz-finance-api/internal/installment/domain"

	"gorm.io/gorm"
)

type installmentQueryRepository struct {
	db *gorm.DB
}

func NewInstallmentQueryRepository(db *gorm.DB) InstallmentQueryRepositoryInterface {
	return &installmentQueryRepository{
		db: db,
	}
}

func (iqr *installmentQueryRepository) GetAllInstallments(userID, transactionID string) ([]domain.Installment, error) {
	var installments []domain.Installment

	result := iqr.db.Where("user_id = ? AND transaction_id = ?", userID, transactionID).Find(&installments)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(installments) == 0 {
		return nil, errors.New("no installments found for this transaction")
	}

	return installments, nil
}


func (iqr *installmentQueryRepository) GetInstallmentByID(id string) (domain.Installment, error) {
	var installment domain.Installment

	result := iqr.db.Where("id = ?", id).First(&installment)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.Installment{}, errors.New("installment not found")
		}
		return domain.Installment{}, result.Error
	}

	return installment, nil
}
