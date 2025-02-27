package dto

import (
	"time"
	"xyz-finance-api/internal/installment/domain"
)

type Installment struct {
	ID                string    `json:"id" form:"id"`
	TransactionID     string    `json:"transaction_id" form:"transaction_id"`
	InstallmentNumber int       `json:"installment_number" form:"installment_number"`
	Amount            int       `json:"amount" form:"amount"`
	DueDate           time.Time `json:"due_date" form:"due_date"`
	Status            string    `json:"status" form:"status"`
	CreatedAt         time.Time `json:"created_at" form:"-"`
	UpdatedAt         time.Time `json:"updated_at" form:"-"`
}

// request
type (
	InstallmentRequest struct {
		TransactionID string `json:"transaction_id" form:"transaction_id"`
	}

	InstallmentUpdateRequest struct {
		Status string `json:"status" form:"status"`
	}
)

// response
type (
	InstallmentResponse struct {
		ID                string    `json:"id"`
		TransactionID     string    `json:"transaction_id"`
		InstallmentNumber int       `json:"installment_number"`
		Amount            int       `json:"amount" form:"amount"`
		DueDate           time.Time `json:"due_date" form:"due_date"`
		Status            string    `json:"status" form:"status"`
		CreatedAt         time.Time `json:"created_at" form:"-"`
		UpdatedAt         time.Time `json:"updated_at" form:"-"`
	}
)

// mapper - request
func InstallmentRequestToInstallmentDomain(request InstallmentRequest) domain.Installment {
	return domain.Installment{
		TransactionID: request.TransactionID,
	}
}

func InstallmentUpdateRequestToInstallmentDomain(request InstallmentUpdateRequest) domain.Installment {
	return domain.Installment{
		Status: request.Status,
	}
}

// mapper - response
func InstallmentDomainToInstallmentResponse(installment domain.Installment) InstallmentResponse {
	return InstallmentResponse{
		ID:                installment.ID,
		TransactionID:     installment.TransactionID,
		InstallmentNumber: installment.InstallmentNumber,
		Amount:            installment.Amount,
		DueDate:           installment.DueDate,
		Status:            installment.Status,
		CreatedAt:         installment.CreatedAt,
		UpdatedAt:         installment.UpdatedAt,
	}
}

func ListInstallmentDomainToInstallmentResponse(installments []domain.Installment) []InstallmentResponse {
	installmentResponses := []InstallmentResponse{}
	for _, installment := range installments {
		installmentResponse := InstallmentDomainToInstallmentResponse(installment)
		installmentResponses = append(installmentResponses, installmentResponse)
	}
	return installmentResponses
}
