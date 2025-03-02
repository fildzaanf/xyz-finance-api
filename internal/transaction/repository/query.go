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

func (tr *transactionQueryRepository) GetTransactionByID(transactionID, userID string) (domain.Transaction, error) {
	var transaction entity.Transaction
	result := tr.db.
		Joins("JOIN loans ON loans.id = transactions.loan_id").
		Where("transactions.id = ? AND loans.user_id = ?", transactionID, userID).
		Select("transactions.*").
		First(&transaction)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.Transaction{}, errors.New("transaction not found")
		}
		return domain.Transaction{}, result.Error
	}

	return domain.TransactionEntityToTransactionDomain(transaction), nil
}


func (tr *transactionQueryRepository) GetAllTransactions(userID string) ([]domain.Transaction, error) {
	var transactions []entity.Transaction
	result := tr.db.
		Joins("JOIN loans ON loans.id = transactions.loan_id").
		Where("loans.user_id = ?", userID).
		Select("transactions.*").
		Find(&transactions)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("no transactions found")
		}
		return nil, result.Error
	}

	return domain.ListTransactionEntityToTransactionDomain(transactions), nil
}