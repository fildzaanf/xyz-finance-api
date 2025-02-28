package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// entity
type Installment struct {
	ID                string    `gorm:"type:varchar(36);primaryKey"`
	TransactionID     string    `gorm:"type:varchar(36);not null"`
	InstallmentNumber int       `gorm:"not null"`
	Amount            int       `gorm:"not null"`
	DueDate           time.Time 
	Status            string    `gorm:"type:varchar(50);not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time 
}

// hooks
func (u *Installment) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	u.ID = UUID.String()

	return nil
}
