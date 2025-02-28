package entity

import (
	"time"
	// eu "xyz-finance-api/internal/user/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// entity
type LoanStatus string

const (
	Valid   LoanStatus = "valid"
	Invalid LoanStatus = "invalid"
)

type (
	Loan struct {
		ID          string     `gorm:"type:varchar(36);primaryKey"`
		UserID      string     `gorm:"type:varchar(36);not null"`
		Tenor       int        `gorm:"not null"`
		LimitAmount int        `gorm:"not null"`
		Status      LoanStatus `gorm:"type:enum('valid', 'invalid');default:'valid'"`
		CreatedAt   time.Time  
		UpdatedAt   time.Time  
	}
)

// hooks
func (u *Loan) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	u.ID = UUID.String()

	return nil
}
