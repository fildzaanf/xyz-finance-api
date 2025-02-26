package repository

import "xyz-finance-api/internal/transaction/domain"

type TransactionCommandRepositoryInterface interface {
	CreateTransaction(user domain.Transaction) (domain.Transaction, error)
}

type TransactionQueryRepositoryInterface interface {
}
