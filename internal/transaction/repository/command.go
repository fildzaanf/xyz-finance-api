package repository

import (
	"xyz-finance-api/internal/transaction/domain"

	"gorm.io/gorm"
)

type transactionCommandRepository struct {
	db *gorm.DB
}

func NewTransactionCommandRepository(db *gorm.DB) TransactionCommandRepositoryInterface {
	return &transactionCommandRepository{
		db: db,
	}
}

func (tcr *transactionCommandRepository) CreateTransaction(transaction domain.Transaction) (domain.Transaction, error) {
	transactionEntity := domain.TransactionDomainToTransactionEntity(transaction)

	result := tcr.db.Create(&transactionEntity)
	if result.Error != nil {
		return domain.Transaction{}, result.Error
	}

	transactionDomain := domain.TransactionEntityToTransactionDomain(transactionEntity)

	return transactionDomain, nil
}
