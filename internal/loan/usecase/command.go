package usecase

import "xyz-finance-api/internal/loan/repository"

type loanCommandUsecase struct {
	loanCommandRepository repository.LoanCommandRepositoryInterface
	loanQueryRepository   repository.LoanQueryRepositoryInterface
}

func NewLoanCommandUsecase(lcr repository.LoanCommandRepositoryInterface, lqr repository.LoanQueryRepositoryInterface) LoanCommandUsecaseInterface {
	return &loanCommandUsecase{
		loanCommandRepository: lcr,
		loanQueryRepository:   lqr,
	}
}
