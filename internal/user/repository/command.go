package repository

import (
	"errors"
	"mime/multipart"
	"xyz-finance-api/pkg/constant"
	"xyz-finance-api/pkg/crypto"
	"xyz-finance-api/internal/user/entity"
	"xyz-finance-api/internal/user/domain"

	"gorm.io/gorm"
)

type userCommandRepository struct {
	db *gorm.DB
}

func NewUserCommandRepository(db *gorm.DB) UserCommandRepositoryInterface {
	return &userCommandRepository{
		db: db,
	}
}

func (ucr *userCommandRepository) RegisterUser(user domain.User, ktpPhoto *multipart.FileHeader, selfiePhoto *multipart.FileHeader) (domain.User, error) {
	userEntity := domain.UserDomainToUserEntity(user)

	result := ucr.db.Create(&userEntity)
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	userDomain := domain.UserEntityToUserDomain(userEntity)

	return userDomain, nil
}

func (ucr *userCommandRepository) LoginUser(email, password string) (domain.User, error) {
	userEntity := entity.User{}

	result := ucr.db.Where("email = ?", email).First(&userEntity)
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	if errComparePass := crypto.ComparePassword(userEntity.Password, password); errComparePass != nil {
		return domain.User{}, errors.New(constant.ERROR_PASSWORD_INVALID)
	}

	userDomain := domain.UserEntityToUserDomain(userEntity)

	return userDomain, nil
}
