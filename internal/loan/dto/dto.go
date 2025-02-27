package dto

import (
	"time"
	"xyz-finance-api/internal/loan/domain"
)

// request
type (
	LoanRequest struct {
		Tenor int `json:"tenor" form:"tenor"`
	}

	LoanUpdateRequest struct {
		Status string `json:"status" form:"status"`
	}
)

// response
type (
	LoanResponse struct {
		ID          string    `json:"id"`
		UserID      string    `json:"user_id"`
		Tenor       int       `json:"tenor"`
		LimitAmount int       `json:"limit_amount"`
		Status      string    `json:"status"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)

// mapper - request
func LoanRequestToLoanDomain(request LoanRequest, userID string) domain.Loan {
	return domain.Loan{
		UserID: userID,
		Tenor:  request.Tenor,
	}
}

func LoanUpdateRequestToLoanDomain(request LoanUpdateRequest) domain.Loan {
	return domain.Loan{
		Status: request.Status,
	}
}

// mapper - response
func LoanDomainToLoanResponse(response domain.Loan) LoanResponse {
	return LoanResponse{
		ID:          response.ID,
		UserID:      response.UserID,
		Tenor:       response.Tenor,
		LimitAmount: response.LimitAmount,
		Status:      response.Status,
		CreatedAt:   response.CreatedAt,
		UpdatedAt:   response.UpdatedAt,
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
