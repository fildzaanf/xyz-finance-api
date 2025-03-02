package repository

import "xyz-finance-api/internal/payment/domain"

type PaymentCommandRepositoryInterface interface {
	CreatePayment(payment domain.Payment) (domain.Payment, error)
	UpdatePaymentStatus(installmentID, status string) error
	UpdateLoanStatus(installmentID string) error
}

type PaymentQueryRepositoryInterface interface {
	GetAllPayments(userID string) ([]domain.Payment, error)
	GetPaymentByID(id string, userID string) (domain.Payment, error)
	GetPaymentByInstallmentID(installmentID string) (domain.Payment, error)
}
