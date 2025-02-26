package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// entity
type User struct {
	ID          string `gorm:"primarykey"`
	Email       string `gorm:"not null"`
	Password    string `gorm:"not null"`
	Nik         string `gorm:"not null;unique"`
	FullName    string `gorm:"not null"`
	LegalName   string `gorm:"not null"`
	BirthPlace  string `gorm:"not null"`
	BirthDate   string `gorm:"not null"`
	KtpPhoto    string `gorm:"not null"`
	SelfiePhoto string `gorm:"not null"`
	Salary      int    `gorm:"type:bigint;not null"`
	Role        string `gorm:"type:enum('user');default:'user'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// hooks
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	u.ID = UUID.String()

	if u.Role == "" {
		u.Role = "user"
	}

	validRoles := map[string]bool{"user": true}
	if !validRoles[u.Role] {
		return errors.New("invalid role")
	}

	return nil
}

/*
CREATE TYPE role AS ENUM ('user');
*/
