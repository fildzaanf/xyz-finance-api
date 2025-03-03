package repository

import (
	"errors"
	"time"
	"xyz-finance-api/internal/installment/domain"
	"xyz-finance-api/pkg/constant"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type installmentCommandRepository struct {
	db *gorm.DB
}

func NewInstallmentCommandRepository(db *gorm.DB) InstallmentCommandRepositoryInterface {
	return &installmentCommandRepository{
		db: db,
	}
}
func (icr *installmentCommandRepository) CreateInstallment(installment domain.Installment, userID string) (domain.Installment, error) {
	tx := icr.db.Begin()

	var transactionID string
	err := tx.Table("transactions").
		Select("id").
		Where("loan_id = (SELECT id FROM loans WHERE user_id = ?)", userID).
		Scan(&transactionID).Error

	if err != nil {
		tx.Rollback()
		return domain.Installment{}, errors.New("failed to fetch transaction ID")
	}

	var existingInstallment domain.Installment
	err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("transaction_id = ?", transactionID).
		First(&existingInstallment).Error

	if err == nil {
		tx.Rollback()
		return domain.Installment{}, errors.New("installment already exists for this transaction")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return domain.Installment{}, err
	}

	installmentEntity := domain.InstallmentDomainToInstallmentEntity(installment)

	result := tx.Create(&installmentEntity)
	if result.Error != nil {
		tx.Rollback()
		return domain.Installment{}, result.Error
	}

	installmentDomain := domain.InstallmentEntityToInstallmentDomain(installmentEntity)

	tx.Commit()
	return installmentDomain, nil
}

func (icr *installmentCommandRepository) UpdateInstallmentStatusByID(installmentID string, installment domain.Installment) (domain.Installment, error) {
	tx := icr.db.Begin() 

	var existingInstallment domain.Installment
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}). 
		Where("id = ?", installmentID).
		First(&existingInstallment).Error

	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Installment{}, errors.New(constant.ERROR_ID_NOTFOUND)
		}
		return domain.Installment{}, err
	}

	installment.UpdatedAt = time.Now()
	result := tx.Model(&domain.Installment{}).Where("id = ?", installmentID).
		Updates(map[string]interface{}{
			"status":     installment.Status,
			"updated_at": installment.UpdatedAt,
		})

	if result.Error != nil {
		tx.Rollback()
		return domain.Installment{}, result.Error
	}

	if result.RowsAffected == 0 {
		tx.Rollback()
		return domain.Installment{}, errors.New("failed to update installment")
	}

	tx.Commit() 
	return installment, nil
}
