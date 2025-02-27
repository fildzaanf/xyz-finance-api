package usecase

import (
	"errors"
	"xyz-finance-api/internal/transaction/domain"
	"xyz-finance-api/internal/transaction/repository"
	"xyz-finance-api/pkg/constant"
)

type transactionQueryUsecase struct {
	transactionCommandRepository repository.TransactionCommandRepositoryInterface
	transactionQueryRepository   repository.TransactionQueryRepositoryInterface
}

func NewTransactionQueryUsecase(tqr repository.TransactionQueryRepositoryInterface, tcr repository.TransactionCommandRepositoryInterface) TransactionQueryUsecaseInterface {
	return &transactionQueryUsecase{
		transactionQueryRepository: tqr,
		transactionCommandRepository: tcr,
	}
}

func (tqu *transactionQueryUsecase) GetTransactionByID(id string) (domain.Transaction, error) {
	if id == "" {
		return domain.Transaction{}, errors.New(constant.ERROR_ID_INVALID)
	}

	transaction, err := tqu.transactionQueryRepository.GetTransactionByID(id)
	if err != nil {
		return domain.Transaction{}, errors.New(constant.ERROR_DATA_EMPTY)
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

	return transactions, nil
}
