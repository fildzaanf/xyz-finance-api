package repository

import (
	"errors"
	"mime/multipart"
	"xyz-finance-api/internal/user/domain"
	"xyz-finance-api/internal/user/entity"
	"xyz-finance-api/pkg/constant"
	"xyz-finance-api/pkg/crypto"

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
	tx := ucr.db.Begin()
    if tx.Error != nil {
        return domain.User{}, tx.Error
    }

    userEntity := domain.UserDomainToUserEntity(user) 

    if err := tx.Create(&userEntity).Error; err != nil {
        tx.Rollback() 
        return domain.User{}, err
    }

    userDomain := domain.UserEntityToUserDomain(userEntity) 

    if err := tx.Commit().Error; err != nil {
        return domain.User{}, err
    }

    return userDomain, nil
}


func (ucr *userCommandRepository) LoginUser(email, password string) (domain.User, error) {
	tx := ucr.db.Begin()

	if tx.Error != nil {
		return domain.User{}, tx.Error
	}

	userEntity := entity.User{}

	result := tx.Where("email = ?", email).First(&userEntity)
	if result.Error != nil {
		tx.Rollback() 
		return domain.User{}, result.Error
	}

	if errComparePass := crypto.ComparePassword(userEntity.Password, password); errComparePass != nil {
		tx.Rollback() 
		return domain.User{}, errors.New(constant.ERROR_PASSWORD_INVALID)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback() 
		return domain.User{}, err
	}

	userDomain := domain.UserEntityToUserDomain(userEntity)

	return userDomain, nil
}
