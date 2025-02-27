package usecase

import (
	"errors"
	"xyz-finance-api/internal/loan/domain"
	"xyz-finance-api/internal/loan/repository"
	"xyz-finance-api/pkg/constant"
)

type loanQueryUsecase struct {
	loanCommandRepository repository.LoanCommandRepositoryInterface
	loanQueryRepository   repository.LoanQueryRepositoryInterface
}

func NewLoanQueryUsecase(lcr repository.LoanCommandRepositoryInterface, lqr repository.LoanQueryRepositoryInterface) LoanQueryUsecaseInterface {
	return &loanQueryUsecase{
		loanCommandRepository: lcr,
		loanQueryRepository:   lqr,
	}
}

func (lqu *loanQueryUsecase) GetLoanByID(id string) (domain.Loan, error) {
	if id == "" {
		return domain.Loan{}, errors.New(constant.ERROR_ID_INVALID)
	}

	loanDomain, errGetID := lqu.loanQueryRepository.GetLoanByID(id)
	if errGetID != nil {
		return domain.Loan{}, errors.New(constant.ERROR_DATA_EMPTY)
	}

	return loanDomain, nil
}

func (lqu *loanQueryUsecase) GetAllLoans(userID string) ([]domain.Loan, error) {
	if userID == "" {
		return nil, errors.New(constant.ERROR_ID_INVALID)
	}

	loans, err := lqu.loanQueryRepository.GetAllLoans(userID)
	if err != nil {
		return nil, errors.New(constant.ERROR_DATA_EMPTY)
	}

	return loans, nil
}
