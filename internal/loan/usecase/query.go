package usecase

import (
	"errors"
	"xyz-finance-api/internal/loan/domain"
	"xyz-finance-api/internal/loan/repository"
	userRepository "xyz-finance-api/internal/user/repository"
	paymentRepository "xyz-finance-api/internal/payment/repository"
	"xyz-finance-api/pkg/constant"
	"xyz-finance-api/pkg/validator"
)

type loanQueryUsecase struct {
	loanCommandRepository repository.LoanCommandRepositoryInterface
	loanQueryRepository   repository.LoanQueryRepositoryInterface
	userQueryRepository   userRepository.UserQueryRepositoryInterface
	paymentQueryRepository   paymentRepository.PaymentQueryRepositoryInterface
}

func NewLoanQueryUsecase(lcr repository.LoanCommandRepositoryInterface, lqr repository.LoanQueryRepositoryInterface, uqr userRepository.UserQueryRepositoryInterface, pqr paymentRepository.PaymentQueryRepositoryInterface) LoanQueryUsecaseInterface {
	return &loanQueryUsecase{
		loanCommandRepository: lcr,
		loanQueryRepository:   lqr,
		userQueryRepository: uqr,
		paymentQueryRepository: pqr,
	}
}

func (lqs *loanQueryUsecase) GetAllLoans(userID string) ([]domain.Loan, error) {
	
	user, err := lqs.userQueryRepository.GetUserByID(userID)
	if err != nil {
		return nil, errors.New(constant.ERROR_ID_NOTFOUND)
	}

	loans, err := lqs.loanQueryRepository.GetAllLoans()
	if err != nil {
		return nil, errors.New(constant.ERROR_DATA_EMPTY)
	}

	var loanResponses []domain.Loan

	for _, loan := range loans {
		loan.LimitAmount = validator.CalculateLoanLimit(user.Salary, loan.Tenor)

		payments, err := lqs.paymentQueryRepository.GetPayments(loan.ID)
		if err != nil {
			return nil, err
		}
		loan.UsedAmount = validator.CalculateUsedAmount(payments)

		loan.RemainingLimit = loan.LimitAmount - loan.UsedAmount
		if loan.RemainingLimit < 0 {
			loan.RemainingLimit = 0
		}

		loanResponses = append(loanResponses, loan)
	}

	return loanResponses, nil
}
