package dto

import (
	"time"
	"xyz-finance-api/internal/loan/domain"
)

// response
type (
	LoanResponse struct {
		ID             string    `json:"id"`
		Tenor          int       `json:"tenor"`
		LimitAmount    int       `json:"limit_amount"`
		UsedAmount     int       `json:"used_amount"`
		RemainingLimit int       `json:"remaining_limit"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}
)

// mapper - response
func LoanDomainToLoanResponse(response domain.Loan) LoanResponse {
	return LoanResponse{
		ID:             response.ID,
		Tenor:          response.Tenor,
		LimitAmount:    response.LimitAmount,
		UsedAmount:     response.UsedAmount,
		RemainingLimit: response.RemainingLimit,
		CreatedAt:      response.CreatedAt,
		UpdatedAt:      response.UpdatedAt,
	}
}

func ListLoanDomainToLoanResponse(loans []domain.Loan) []LoanResponse {
	loanResponses := []LoanResponse{}
	for _, loan := range loans {
		loanResponse := LoanDomainToLoanResponse(loan)
		loanResponses = append(loanResponses, loanResponse)
	}
	return loanResponses
}


