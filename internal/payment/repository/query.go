package repository

import (
	"errors"
	"xyz-finance-api/internal/payment/domain"

	"gorm.io/gorm"
)

type paymentQueryRepository struct {
	db *gorm.DB
}

func NewPaymentQueryRepository(db *gorm.DB) PaymentQueryRepositoryInterface {
	return &paymentQueryRepository{
		db: db,
	}
}

func (pqr *paymentQueryRepository) GetPaymentByInstallmentID(installmentID string) (domain.Payment, error) {
	var payment domain.Payment
	result := pqr.db.Where("installment_id = ?", installmentID).First(&payment)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.Payment{}, errors.New("payment not found")
		}
		return domain.Payment{}, result.Error
	}

	return payment, nil
}
