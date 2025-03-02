package repository

import (
	"errors"
	"xyz-finance-api/internal/installment/domain"
	"xyz-finance-api/internal/installment/entity"

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

func (ir *installmentQueryRepository) GetAllInstallments(userID string) ([]domain.Installment, error) {
	var installments []entity.Installment
	result := ir.db.
		Joins("JOIN transactions ON transactions.id = installments.transaction_id").
		Joins("JOIN loans ON loans.id = transactions.loan_id").
		Where("loans.user_id = ?", userID).
		Select("installments.*").
		Find(&installments)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("no installments found")
		}
		return nil, result.Error
	}

	return domain.ListInstallmentEntityToInstallmentDomain(installments), nil
}


// func (ir *installmentQueryRepository) GetInstallmentByID(installmentID, userID string) (domain.Installment, error) {
// 	var installment entity.Installment
// 	result := ir.db.
// 		Joins("JOIN transactions ON transactions.id = installments.transaction_id").
// 		Joins("JOIN loans ON loans.id = transactions.loan_id").
// 		Where("installments.id = ? AND loans.user_id = ?", installmentID, userID).
// 		Select("installments.*").
// 		First(&installment)

// 	if result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			return domain.Installment{}, errors.New("installment not found")
// 		}
// 		return domain.Installment{}, result.Error
// 	}

// 	return domain.InstallmentEntityToInstallmentDomain(installment), nil
// }
func (ir *installmentQueryRepository) GetInstallmentByID(installmentID, userID string) (domain.Installment, error) {
	var installment entity.Installment
	result := ir.db.
		Joins("JOIN transactions ON transactions.id = installments.transaction_id").
		Joins("JOIN loans ON loans.id = transactions.loan_id").
		Where("installments.id = ? AND loans.user_id = ?", installmentID, userID).
		Select("installments.*").
		First(&installment)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.Installment{}, errors.New("installment not found")
		}
		return domain.Installment{}, result.Error
	}

	return domain.InstallmentEntityToInstallmentDomain(installment), nil
}


func (repo *installmentQueryRepository) GetInstallmentByTransactionID(transactionID string) ([]domain.Installment, error) {
	var installments []domain.Installment

	result := repo.db.Where("transaction_id = ?", transactionID).Find(&installments)
	if result.Error != nil {
		return nil, result.Error
	}

	return installments, nil
}

func (repo *installmentQueryRepository) CountInstallmentsByTransactionID(transactionID string) (int, error) {
	var count int64
	err := repo.db.Model(&domain.Installment{}).
		Where("transaction_id = ?", transactionID).
		Count(&count).Error

	if err != nil {
		return 0, err
	}
	return int(count), nil
}
