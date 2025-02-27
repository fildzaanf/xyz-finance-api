package repository

import "xyz-finance-api/internal/payment/domain"

type PaymentCommandRepositoryInterface interface {
	CreatePayment(payment domain.Payment) (domain.Payment, error)
}

type PaymentQueryRepositoryInterface interface {
	// GetAllPayments(transactionID string) ([]domain.Payment, error)
	// GetPaymentByID(id string) (domain.Payment, error)
}
