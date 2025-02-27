package usecase

import (
	"errors"
	"time"
	"xyz-finance-api/internal/loan/domain"
	"xyz-finance-api/internal/loan/repository"
	userRepository "xyz-finance-api/internal/user/repository"
	"xyz-finance-api/pkg/constant"
	"xyz-finance-api/pkg/validator"
)

type loanCommandUsecase struct {
	loanCommandRepository repository.LoanCommandRepositoryInterface
	loanQueryRepository   repository.LoanQueryRepositoryInterface
	userQueryRepository   userRepository.UserQueryRepositoryInterface
}

func NewLoanCommandUsecase(lcr repository.LoanCommandRepositoryInterface, lqr repository.LoanQueryRepositoryInterface, uqr userRepository.UserQueryRepositoryInterface) LoanCommandUsecaseInterface {
	return &loanCommandUsecase{
		loanCommandRepository: lcr,
		loanQueryRepository:   lqr,
		userQueryRepository:   uqr,
	}
}

func (lcu *loanCommandUsecase) CreateLoan(loan domain.Loan) (domain.Loan, error) {
	errEmpty := validator.IsDataEmpty([]string{"user_id", "tenor"}, loan.UserID, loan.Tenor)
	if errEmpty != nil {
		return domain.Loan{}, errEmpty
	}

	user, errUser := lcu.userQueryRepository.GetUserByID(loan.UserID)
	if errUser != nil {
		return domain.Loan{}, errors.New(constant.ERROR_ID_NOTFOUND)
	}
	limitAmount := validator.CalculateLoanLimit(user.Salary, loan.Tenor)
	if limitAmount == 0 {
		return domain.Loan{}, errors.New("invalid tenor")
	}

	loan.LimitAmount = limitAmount
	loan.Status = "valid" 
	loan.CreatedAt = time.Now()
	loan.UpdatedAt = time.Now()

	loanEntity, errCreate := lcu.loanCommandRepository.CreateLoan(loan)
	if errCreate != nil {
		return domain.Loan{}, errCreate
	}

	return loanEntity, nil
}

func (lcu *loanCommandUsecase) UpdateLoanStatusByID(id string, loan domain.Loan) (domain.Loan, error) {

	existingLoan, err := lcu.loanQueryRepository.GetLoanByID(id)
	if err != nil {
		return domain.Loan{}, errors.New(constant.ERROR_ID_NOTFOUND)
	}

	if loan.Status != "valid" && loan.Status != "invalid" {
		return domain.Loan{}, errors.New("invalid status update request")
	}

	existingLoan.Status = loan.Status
	existingLoan.UpdatedAt = time.Now()


	updatedLoan, err := lcu.loanCommandRepository.UpdateLoanStatusByID(id, existingLoan)
	if err != nil {
		return domain.Loan{}, err
	}

	return updatedLoan, nil
}

