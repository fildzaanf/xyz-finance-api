package repository

import (
	"gorm.io/gorm"
)

type loanCommandRepository struct {
	db *gorm.DB
}

func NewLoanCommandRepository(db *gorm.DB) LoanCommandRepositoryInterface {
	return &loanCommandRepository{
		db: db,
	}
}
