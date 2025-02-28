package repository

import "xyz-finance-api/internal/installment/domain"

type InstallmentCommandRepositoryInterface interface {
	CreateInstallment(installment domain.Installment) (domain.Installment, error)
	UpdateInstallmentStatusByID(id string, installment domain.Installment) (domain.Installment, error)
}

type InstallmentQueryRepositoryInterface interface {
	GetAllInstallments(userID, transactionID string) ([]domain.Installment, error)
	GetInstallmentByID(id string) (domain.Installment, error)
	GetInstallmentByTransactionID(transactionID string) ([]domain.Installment, error)
}
