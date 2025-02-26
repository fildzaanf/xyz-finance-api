package dto

import (
	"time"
	"xyz-finance-api/internal/transaction/domain"
)

// request
type (
	TransactionRequest struct {
		LoanID    string `json:"loan_id" form:"loan_id"`
		AssetName string `json:"asset_name" form:"asset_name"`
		OTRPrice  int    `json:"otr_price" form:"otr_price"`
	}
)

// response
type (
	TransactionResponse struct {
		ID                string    `json:"id"`
		LoanID            string    `json:"loan_id"`
		AssetName         string    `json:"asset_name"`
		TotalAmount       int       `json:"total_amount"`
		OTRPrice          int       `json:"otr_price"`
		AdminFee          int       `json:"admin_fee"`
		Interest          int       `json:"interest"`
		InstallmentAmount int       `json:"installment_amount"`
		Status            string    `json:"status"`
		CreatedAt         time.Time `json:"created_at"`
	}
)

// mapper - request
func TransactionRequestToTransactionDomain(request TransactionRequest) domain.Transaction {
	return domain.Transaction{
		LoanID:    request.LoanID,
		AssetName: request.AssetName,
		OTRPrice:  request.OTRPrice,
	}
}

// mapper - response
func TransactionDomainToTransactionResponse(response domain.Transaction) TransactionResponse {
	return TransactionResponse{
		ID:                response.ID,
		LoanID:            response.LoanID,
		AssetName:         response.AssetName,
		TotalAmount:       response.TotalAmount,
		OTRPrice:          response.OTRPrice,
		AdminFee:          response.AdminFee,
		Interest:          response.Interest,
		InstallmentAmount: response.InstallmentAmount,
		Status:            response.Status,
		CreatedAt:         response.CreatedAt,
	}
}
