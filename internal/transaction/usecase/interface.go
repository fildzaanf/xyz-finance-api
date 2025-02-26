package usecase

import "xyz-finance-api/internal/transaction/domain"

type TransactionCommandUsecaseInterface interface {
	CreateTransaction(transaction domain.Transaction) (domain.Transaction, error)
}

type TransactionQueryUsecaseInterface interface {
	
}
