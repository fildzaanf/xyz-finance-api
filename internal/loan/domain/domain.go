package domain

import (
	"time"
	"xyz-finance-api/internal/loan/entity"
)

// domain
type Loan struct {
	ID             string
	UserID         string
	Tenor          int
	LimitAmount    int
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// mapper
func LoanDomainToLoanEntity(loanDomain Loan) entity.Loan {
	return entity.Loan{
		ID:             loanDomain.ID,
		UserID:         loanDomain.UserID,
		Tenor:          loanDomain.Tenor,
		LimitAmount:    loanDomain.LimitAmount,
		Status:         entity.LoanStatus(loanDomain.Status),
		CreatedAt:      loanDomain.CreatedAt,
		UpdatedAt:      loanDomain.UpdatedAt,
	}
}

func LoanEntityToLoanDomain(loanEntity entity.Loan) Loan {
	return Loan{
		ID:             loanEntity.ID,
		UserID:         loanEntity.UserID,
		Tenor:          loanEntity.Tenor,
		LimitAmount:    loanEntity.LimitAmount,
		Status:         string(loanEntity.Status),
		CreatedAt:      loanEntity.CreatedAt,
		UpdatedAt:      loanEntity.UpdatedAt,
	}
}

func ListLoanDomainToLoanEntity(loanDomains []Loan) []entity.Loan {
	listLoanEntities := []entity.Loan{}
	for _, loan := range loanDomains {
		loanEntity := LoanDomainToLoanEntity(loan)
		listLoanEntities = append(listLoanEntities, loanEntity)
	}
	return listLoanEntities
}

func ListLoanEntityToLoanDomain(loanEntities []entity.Loan) []Loan {
	listLoanDomains := []Loan{}
	for _, loan := range loanEntities {
		loanDomain := LoanEntityToLoanDomain(loan)
		listLoanDomains = append(listLoanDomains, loanDomain)
	}
	return listLoanDomains
}
