package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID          string    `gorm:"type:varchar(36);primaryKey"`
	LoanID      string    `gorm:"type:varchar(36);not null"`
	AssetName   string    `gorm:"type:varchar(100);not null"`
	TotalAmount int       `gorm:"not null"`
	Tenor       int       `gorm:"not null"`
	OTRPrice    int       `gorm:"not null"`
	AdminFee    int       `gorm:"not null"`
	Interest    int       `gorm:"not null"`
	Status      string    `gorm:"type:varchar(50);not null"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

// hooks
func (u *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	u.ID = UUID.String()

	return nil
}
