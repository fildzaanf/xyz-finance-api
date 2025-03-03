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

    tx := tcr.db.Begin()

    if err := tx.Error; err != nil {
        return domain.Transaction{}, err
    }

    result := tx.Create(&transactionEntity)
    if result.Error != nil {
        tx.Rollback() 
        return domain.Transaction{}, result.Error
    }

    if err := tx.Commit().Error; err != nil {
        return domain.Transaction{}, err
    }

    transactionDomain := domain.TransactionEntityToTransactionDomain(transactionEntity)
    return transactionDomain, nil
}

