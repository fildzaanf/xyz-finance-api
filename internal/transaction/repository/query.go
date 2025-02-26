package repository

import "gorm.io/gorm"

type transactionQueryRepository struct {
	db *gorm.DB
}

func NewTransactionQueryRepository(db *gorm.DB) TransactionQueryRepositoryInterface {
	return &transactionQueryRepository{
		db: db,
	}
}
