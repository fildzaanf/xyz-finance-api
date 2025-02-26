package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// entity
type Loan struct {
	ID             string `gorm:"primaryKey"`
	Tenor          int    `gorm:"not null"`
	LimitAmount    int    `gorm:"not null"`
	UsedAmount     int    `gorm:"not null"`
	RemainingLimit int    `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// hooks
func (u *Loan) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	u.ID = UUID.String()

	return nil
}
