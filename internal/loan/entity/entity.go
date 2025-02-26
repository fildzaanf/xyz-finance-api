package entity

import (
	"time"
	eu "xyz-finance-api/internal/user/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// entity
type Loan struct {
	ID             string           `gorm:"type:varchar(36);primaryKey"`
	UserID         string           `gorm:"type:varchar(36);not null"`
	Tenor          int              `gorm:"not null"`
	LimitAmount    int              `gorm:"not null"`
	UsedAmount     int              `gorm:"default:0"`
	RemainingLimit int              `gorm:"-"`
	CreatedAt      time.Time        `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time        `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
	User           eu.User          `gorm:"foreignKey:UserID"`
}

// hooks
func (u *Loan) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	u.ID = UUID.String()

	return nil
}
