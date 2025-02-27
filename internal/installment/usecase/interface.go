package usecase

import "xyz-finance-api/internal/installment/domain"

type InstallmentCommandUsecaseInterface interface {
	CreateInstallment(installment domain.Installment) ([]domain.Installment, error)
	UpdateInstallmentStatusByID(id string, installment domain.Installment) (domain.Installment, error)
}

type InstallmentQueryUsecaseInterface interface {
	GetAllInstallments(userID string) ([]domain.Installment, error)
	GetInstallmentByID(id string) (domain.Installment, error)
}
