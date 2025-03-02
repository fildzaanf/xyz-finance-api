package usecase

import "xyz-finance-api/internal/payment/domain"

type PaymentCommandUsecaseInterface interface {
	CreatePayment(payment domain.Payment, userID string) (domain.Payment, error)
	UpdatePaymentStatus(installmentID, status string) error
}

type PaymentQueryUsecaseInterface interface {
	GetAllPayments(userID string) ([]domain.Payment, error)
	GetPaymentByID(id string, userID string) (domain.Payment, error)
}
