package usecase

import (
	"errors"
	"xyz-finance-api/internal/transaction/domain"
	"xyz-finance-api/internal/transaction/repository"
	userRepository "xyz-finance-api/internal/user/repository"
	"xyz-finance-api/pkg/constant"
)

type transactionQueryUsecase struct {
	transactionCommandRepository repository.TransactionCommandRepositoryInterface
	transactionQueryRepository   repository.TransactionQueryRepositoryInterface
	userQueryRepository   userRepository.UserQueryRepositoryInterface
}

func NewTransactionQueryUsecase(tqr repository.TransactionQueryRepositoryInterface, tcr repository.TransactionCommandRepositoryInterface, uqr   userRepository.UserQueryRepositoryInterface) TransactionQueryUsecaseInterface {
	return &transactionQueryUsecase{
		transactionQueryRepository: tqr,
		transactionCommandRepository: tcr,
		userQueryRepository: uqr,
	}
}

func (tu *transactionQueryUsecase) GetTransactionByID(transactionID, userID string) (domain.Transaction, error) {
	if transactionID == "" {
		return domain.Transaction{}, errors.New(constant.ERROR_ID_INVALID)
	}

	transaction, err := tu.transactionQueryRepository.GetTransactionByID(transactionID, userID)
	if err != nil {
		return domain.Transaction{}, err
	}

	if transaction.ID != transactionID {
		return domain.Transaction{}, errors.New(constant.ERROR_ID_NOTFOUND)
	}

	user, err := tu.userQueryRepository.GetUserByID(userID)
	if err != nil {
		return domain.Transaction{}, err
	}

	if user.ID != userID {
		return domain.Transaction{}, errors.New(constant.ERROR_ROLE_ACCESS)
	}

	return transaction, nil
}

func (tqu *transactionQueryUsecase) GetAllTransactions(userID string) ([]domain.Transaction, error) {
	if userID == "" {
		return nil, errors.New(constant.ERROR_ID_INVALID)
	}

	transactions, err := tqu.transactionQueryRepository.GetAllTransactions(userID)
	if err != nil {
		return nil, errors.New(constant.ERROR_DATA_EMPTY)
	}

	user, err := tqu.userQueryRepository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	if user.ID != userID {
		return nil, errors.New(constant.ERROR_ROLE_ACCESS)
	}

	return transactions, nil
}
