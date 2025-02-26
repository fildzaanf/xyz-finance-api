package usecase

import "xyz-finance-api/internal/transaction/repository"

type transactionQueryUsecase struct {
	transactionCommandRepository repository.TransactionCommandRepositoryInterface
	transactionQueryRepository   repository.TransactionQueryRepositoryInterface
}

func NewTransactionQueryUsecase(tcr repository.TransactionCommandRepositoryInterface, tqr repository.TransactionQueryRepositoryInterface) TransactionQueryUsecaseInterface {
	return &transactionQueryUsecase{
		transactionCommandRepository: tcr,
		transactionQueryRepository:   tqr,
	}
}
