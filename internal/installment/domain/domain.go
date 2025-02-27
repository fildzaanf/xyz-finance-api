package domain

import (
	"time"
	"xyz-finance-api/internal/installment/entity"
)

type Installment struct {
	ID                string
	TransactionID     string
	InstallmentNumber int
	Amount            int
	DueDate           time.Time
	Status            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func InstallmentDomainToInstallmentEntity(installmentDomain Installment) entity.Installment {
	return entity.Installment{
		ID:                installmentDomain.ID,
		TransactionID:     installmentDomain.TransactionID,
		InstallmentNumber: installmentDomain.InstallmentNumber,
		Amount:            installmentDomain.Amount,
		DueDate:           installmentDomain.DueDate,
		Status:            installmentDomain.Status,
		CreatedAt:         installmentDomain.CreatedAt,
		UpdatedAt:         installmentDomain.UpdatedAt,
	}
}

func InstallmentEntityToInstallmentDomain(installmentEntity entity.Installment) Installment {
	return Installment{
		ID:                installmentEntity.ID,
		TransactionID:     installmentEntity.TransactionID,
		InstallmentNumber: installmentEntity.InstallmentNumber,
		Amount:            installmentEntity.Amount,
		DueDate:           installmentEntity.DueDate,
		Status:            installmentEntity.Status,
		CreatedAt:         installmentEntity.CreatedAt,
		UpdatedAt:         installmentEntity.UpdatedAt,
	}
}

func ListInstallmentDomainToInstallmentEntity(installmentDomains []Installment) []entity.Installment {
	listInstallmentEntities := []entity.Installment{}
	for _, installment := range installmentDomains {
		installmentEntity := InstallmentDomainToInstallmentEntity(installment)
		listInstallmentEntities = append(listInstallmentEntities, installmentEntity)
	}
	return listInstallmentEntities
}

func ListInstallmentEntityToInstallmentDomain(installmentEntities []entity.Installment) []Installment {
	listInstallmentDomains := []Installment{}
	for _, installment := range installmentEntities {
		installmentDomain := InstallmentEntityToInstallmentDomain(installment)
		listInstallmentDomains = append(listInstallmentDomains, installmentDomain)
	}
	return listInstallmentDomains
}
