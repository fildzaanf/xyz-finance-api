package repository

import (
	"errors"

	"xyz-finance-api/internal/transaction/domain"
	"xyz-finance-api/internal/transaction/entity"

	"gorm.io/gorm"
)

type transactionQueryRepository struct {
	db *gorm.DB
}

func NewTransactionQueryRepository(db *gorm.DB) TransactionQueryRepositoryInterface {
	return &transactionQueryRepository{
		db: db,
	}
}

func (tqr *transactionQueryRepository) GetTransactionByID(id string) (domain.Transaction, error) {
	var transactionEntity entity.Transaction
	result := tqr.db.Where("id = ?", id).First(&transactionEntity)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.Transaction{}, errors.New("transaction not found")
		}
		return domain.Transaction{}, result.Error
	}

	transactionDomain := domain.TransactionEntityToTransactionDomain(transactionEntity)

	return transactionDomain, nil
}

func (tqr *transactionQueryRepository) GetAllTransactions(userID string) ([]domain.Transaction, error) {
	var transactionEntities []entity.Transaction

	result := tqr.db.Where("user_id = ?", userID).Find(&transactionEntities)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(transactionEntities) == 0 {
		return nil, errors.New("no transactions found for this user")
	}

	var transactions []domain.Transaction
	for _, transactionEntity := range transactionEntities {
		transactions = append(transactions, domain.TransactionEntityToTransactionDomain(transactionEntity))
	}

	return transactions, nil
}
