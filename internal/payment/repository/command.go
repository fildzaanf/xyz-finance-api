package repository

import (
	"xyz-finance-api/internal/payment/domain"

	"gorm.io/gorm"
)

type paymentCommandRepository struct {
	db *gorm.DB
}

func NewPaymentCommandRepository(db *gorm.DB) PaymentCommandRepositoryInterface {
	return &paymentCommandRepository{
		db: db,
	}
}

func (pcr *paymentCommandRepository) CreatePayment(payment domain.Payment) (domain.Payment, error) {
	paymentEntity := domain.PaymentDomainToPaymentEntity(payment)

	result := pcr.db.Create(&paymentEntity)
	if result.Error != nil {
		return domain.Payment{}, result.Error
	}

	paymentDomain := domain.PaymentEntityToPaymentDomain(paymentEntity)

	return paymentDomain, nil
}
