package usecase

import "xyz-finance-api/internal/installment/domain"

type InstallmentCommandUsecaseInterface interface {
	CreateInstallment(installment domain.Installment, userID string) ([]domain.Installment, error)
	UpdateInstallmentStatusByID(installmentID string, installment domain.Installment, userID string) (domain.Installment, error)
}

type InstallmentQueryUsecaseInterface interface {
	GetAllInstallments(userID string) ([]domain.Installment, error)
	GetInstallmentByID(instalmentID, userID string) (domain.Installment, error)
}
