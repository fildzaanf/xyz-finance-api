package repository

import "gorm.io/gorm"

type paymentQueryRepository struct {
	db *gorm.DB
}

func NewPaymentQueryRepository(db *gorm.DB) PaymentQueryRepositoryInterface {
	return &paymentQueryRepository{
		db: db,
	}
}
