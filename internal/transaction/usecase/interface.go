package usecase

import "xyz-finance-api/internal/transaction/domain"

type TransactionCommandUsecaseInterface interface {
	CreateTransaction(transaction domain.Transaction, userID string) (domain.Transaction, error)
}

type TransactionQueryUsecaseInterface interface {
	GetAllTransactions(userID string) ([]domain.Transaction, error)
	GetTransactionByID(id string) (domain.Transaction, error)
	
}
