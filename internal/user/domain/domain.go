package domain

import (
	"time"
	"xyz-finance-api/internal/user/entity"
)

// domain
type User struct {
	ID              string
	Email           string
	Password        string
	ConfirmPassword string
	Nik             string
	FullName        string
	LegalName       string
	BirthPlace      string
	BirthDate       string
	KtpPhoto        string
	SelfiePhoto     string
	Salary          int
	Role            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// mapper
func UserDomainToUserEntity(userDomain User) entity.User {
	return entity.User{
		ID:              userDomain.ID,
		Email:           userDomain.Email,
		Password:        userDomain.Password,
		Nik:             userDomain.Nik,
		FullName:        userDomain.FullName,
		LegalName:       userDomain.LegalName,
		BirthPlace:      userDomain.BirthPlace,
		BirthDate:       userDomain.BirthDate,
		KtpPhoto:        userDomain.KtpPhoto,
		SelfiePhoto:     userDomain.SelfiePhoto,
		Salary:          userDomain.Salary,
		Role:            userDomain.Role,
		CreatedAt:       userDomain.CreatedAt,
		UpdatedAt:       userDomain.UpdatedAt,
	}
}

func UserEntityToUserDomain(userEntity entity.User) User {
	return User{
		ID:          userEntity.ID,
		Email:       userEntity.Email,
		Password:    userEntity.Password,
		Nik:         userEntity.Nik,
		FullName:    userEntity.FullName,
		LegalName:   userEntity.LegalName,
		BirthPlace:  userEntity.BirthPlace,
		BirthDate:   userEntity.BirthDate,
		KtpPhoto:    userEntity.KtpPhoto,
		SelfiePhoto: userEntity.SelfiePhoto,
		Salary:      userEntity.Salary,
		Role:        userEntity.Role,
		CreatedAt:   userEntity.CreatedAt,
		UpdatedAt:   userEntity.UpdatedAt,
	}
}
