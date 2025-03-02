package repository

import "xyz-finance-api/internal/transaction/domain"

type TransactionCommandRepositoryInterface interface {
	CreateTransaction(user domain.Transaction) (domain.Transaction, error)
}

type TransactionQueryRepositoryInterface interface {
	GetAllTransactions(userID string) ([]domain.Transaction, error)
	GetTransactionByID(transactionID, userID string) (domain.Transaction, error)
}
