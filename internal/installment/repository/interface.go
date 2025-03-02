package repository

import "xyz-finance-api/internal/installment/domain"

type InstallmentCommandRepositoryInterface interface {
	CreateInstallment(installment domain.Installment, userID string) (domain.Installment, error)
	UpdateInstallmentStatusByID(installmentID string, installment domain.Installment) (domain.Installment, error)
}

type InstallmentQueryRepositoryInterface interface {
	GetAllInstallments(userID string) ([]domain.Installment, error)
	GetInstallmentByID(installmentID, userID string) (domain.Installment, error)
	GetInstallmentByTransactionID(transactionID string) ([]domain.Installment, error)
	CountInstallmentsByTransactionID(transactionID string) (int, error)
}
