package repository

import (
	"errors"
	"time"
	"xyz-finance-api/internal/installment/domain"
	"xyz-finance-api/pkg/constant"

	"gorm.io/gorm"
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
	installmentEntity := domain.InstallmentDomainToInstallmentEntity(installment)

	result := icr.db.Create(&installmentEntity)
	if result.Error != nil {
		return domain.Installment{}, result.Error
	}

	installmentDomain := domain.InstallmentEntityToInstallmentDomain(installmentEntity)

	return installmentDomain, nil
}

func (icr *installmentCommandRepository) UpdateInstallmentStatusByID(installmentID string, installment domain.Installment) (domain.Installment, error) {

	var existingInstallment domain.Installment
	err := icr.db.Where("id = ?", installmentID).First(&existingInstallment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Installment{}, errors.New(constant.ERROR_ID_NOTFOUND)
		}
		return domain.Installment{}, err
	}

	installment.UpdatedAt = time.Now()
	result := icr.db.Model(&domain.Installment{}).Where("id = ?", installmentID).
		Updates(map[string]interface{}{
			"status":     installment.Status,
			"updated_at": installment.UpdatedAt,
		})

	if result.Error != nil {
		return domain.Installment{}, result.Error
	}

	if result.RowsAffected == 0 {
		return domain.Installment{}, errors.New("failed to update installment")
	}

	return installment, nil
}
