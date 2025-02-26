package repository

import (
	"errors"
	"xyz-finance-api/internal/user/domain"
	"xyz-finance-api/internal/user/entity"

	"gorm.io/gorm"
)

type userQueryRepository struct {
	db *gorm.DB
}

func NewUserQueryRepository(db *gorm.DB) UserQueryRepositoryInterface {
	return &userQueryRepository{
		db: db,
	}
}

func (uqr *userQueryRepository) GetUserByID(id string) (domain.User, error) {
	var userEntity entity.User
	result := uqr.db.Where("id = ?", id).First(&userEntity)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, result.Error
	}

	userDomain := domain.UserEntityToUserDomain(userEntity)

	return userDomain, nil
}

func (uqr *userQueryRepository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	result := uqr.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, result.Error
	}

	return user, nil
}
