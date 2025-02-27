package domain

import (
	"time"
	"xyz-finance-api/internal/payment/entity"
)

// domain
type Payment struct {
	ID            string
	InstallmentID string
	GrossAmount   int
	Status        string
	PaymentURL    string
	Token         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// mapper
func PaymentDomainToPaymentEntity(paymentDomain Payment) entity.Payment {
	return entity.Payment{
		ID:            paymentDomain.ID,
		InstallmentID: paymentDomain.InstallmentID,
		GrossAmount:   paymentDomain.GrossAmount,
		Status:        paymentDomain.Status,
		PaymentURL:    paymentDomain.PaymentURL,
		Token:         paymentDomain.Token,
		CreatedAt:     paymentDomain.CreatedAt,
		UpdatedAt:     paymentDomain.UpdatedAt,
	}
}

func PaymentEntityToPaymentDomain(paymentEntity entity.Payment) Payment {
	return Payment{
		ID:            paymentEntity.ID,
		InstallmentID: paymentEntity.InstallmentID,
		GrossAmount:   paymentEntity.GrossAmount,
		Status:        paymentEntity.Status,
		PaymentURL:    paymentEntity.PaymentURL,
		Token:         paymentEntity.Token,
		CreatedAt:     paymentEntity.CreatedAt,
		UpdatedAt:     paymentEntity.UpdatedAt,
	}
}

func ListPaymentDomainToPaymentEntity(paymentDomains []Payment) []entity.Payment {
	paymentEntities := make([]entity.Payment, len(paymentDomains))
	for i, payment := range paymentDomains {
		paymentEntities[i] = PaymentDomainToPaymentEntity(payment)
	}
	return paymentEntities
}

func ListPaymentEntityToPaymentDomain(paymentEntities []entity.Payment) []Payment {
	paymentDomains := make([]Payment, len(paymentEntities))
	for i, payment := range paymentEntities {
		paymentDomains[i] = PaymentEntityToPaymentDomain(payment)
	}
	return paymentDomains
}
