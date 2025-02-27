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
	Amount            int        `gorm:"not null"`
	DueDate           time.Time `gorm:"type:date;not null"`
	Status            string    `gorm:"type:varchar(50);not null"`
	CreatedAt         time.Time `gorm:"type:timestamptz;default:current_timestamp"`
	UpdatedAt         time.Time `gorm:"type:timestamptz;default:current_timestamp on update current_timestamp"`
}

// hooks
func (u *Installment) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	u.ID = UUID.String()

	return nil
}
