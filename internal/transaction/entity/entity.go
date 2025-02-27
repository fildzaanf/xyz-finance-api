package entity

import (
	"time"
	el "xyz-finance-api/internal/loan/entity"
	// ep "xyz-finance-api/internal/payment/entity"
)

type Transaction struct {
	ID                string    `gorm:"type:varchar(36);primaryKey"`
	LoanID            string    `gorm:"type:varchar(36);not null"`
	AssetName         string    `gorm:"type:varchar(100);not null"`
	TotalAmount       int       `gorm:"not null"`
	OTRPrice          int       `gorm:"not null"`
	AdminFee          int       `gorm:"not null"`
	Interest          int       `gorm:"not null"`
	InstallmentAmount int       `gorm:"not null"`
	Status            string    `gorm:"type:varchar(50);not null"`
	CreatedAt         time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
	Loan              el.Loan   `gorm:"foreignKey:LoanID"`
	// Payments          []ep.Payment `gorm:"foreignKey:TransactionID"`
}
