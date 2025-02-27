package entity

import "time"

type PaymentStatus string

const (
	Pending PaymentStatus = "pending"
	Success PaymentStatus = "success"
	Failed  PaymentStatus = "failed"
	Expired PaymentStatus = "expired"
)

type Payment struct {
	ID            string        `gorm:"type:varchar(36);primaryKey"`
	TransactionID string        `gorm:"type:varchar(36);not null"`
	OrderID       string        `gorm:"type:varchar(100);not null;unique"` 
	Amount        int           `gorm:"not null"`
	PaymentType   string        `gorm:"type:varchar(50);not null"` 
	Status        PaymentStatus `gorm:"type:enum('pending', 'success', 'failed', 'expired');default:'pending'"`
	PaymentURL    string        `gorm:"type:text"` 
	CreatedAt     time.Time     `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time     `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`

	// Transaction Transaction `gorm:"foreignKey:TransactionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

