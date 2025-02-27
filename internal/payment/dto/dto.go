package dto

import (
	"time"
	"xyz-finance-api/internal/payment/domain"
)

// request
type (
	PaymentRequest struct {
		InstallmentID string `json:"installment_id" form:"installment_id"`
	}
)

// response
type (
	PaymentResponse struct {
		ID            string    `json:"id"`
		InstallmentID string    `json:"installment_id"`
		GrossAmount   int       `json:"gross_amount"`
		Status        string    `json:"status"`
		PaymentURL    string    `json:"payment_url"`
		Token         string    `json:"token"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
	}
)

// mapper - request
func PaymentRequestToPaymentDomain(request PaymentRequest) domain.Payment {
	return domain.Payment{
		InstallmentID: request.InstallmentID,
	}
}

// mapper - response
func PaymentDomainToPaymentResponse(payment domain.Payment) PaymentResponse {
	return PaymentResponse{
		ID:            payment.ID,
		InstallmentID: payment.InstallmentID,
		GrossAmount:   payment.GrossAmount,
		Status:        payment.Status,
		PaymentURL:    payment.PaymentURL,
		Token:         payment.Token,
		CreatedAt:     payment.CreatedAt,
		UpdatedAt:     payment.UpdatedAt,
	}
}

func ListPaymentDomainToPaymentResponse(payments []domain.Payment) []PaymentResponse {
	paymentResponses := make([]PaymentResponse, len(payments))
	for i, payment := range payments {
		paymentResponses[i] = PaymentDomainToPaymentResponse(payment)
	}
	return paymentResponses
}
