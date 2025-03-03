package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID            string    `gorm:"type:varchar(36);primaryKey"`
	InstallmentID string    `gorm:"type:varchar(36);not null"`
	GrossAmount   int       `gorm:"not null"`
	Status        string    `gorm:"type:enum('deny', 'success', 'cancel', 'expire', 'pending');default:'pending'"`
	PaymentURL    string    `gorm:"type:text"`
	Token         string    `gorm:"type:text"`
	CreatedAt     time.Time 
	UpdatedAt     time.Time 
}

// hooks
func (u *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	u.ID = UUID.String()

	return nil
}
