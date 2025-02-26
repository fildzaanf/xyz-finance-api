package usecase

import (
	"errors"
	"xyz-finance-api/pkg/constant"
	"xyz-finance-api/internal/user/domain"
	"xyz-finance-api/internal/user/repository"
)

type userQueryUsecase struct {
	userCommandRepository repository.UserCommandRepositoryInterface
	userQueryRepository   repository.UserQueryRepositoryInterface
}

func NewUserQueryUsecase(ucr repository.UserCommandRepositoryInterface, uqr repository.UserQueryRepositoryInterface) UserQueryUsecaseInterface {
	return &userQueryUsecase{
		userCommandRepository: ucr,
		userQueryRepository:   uqr,
	}
}

func (uqs *userQueryUsecase) GetUserByID(id string) (domain.User, error) {
	if id == "" {
		return domain.User{}, errors.New(constant.ERROR_ID_INVALID)
	}

	userDomain, errGetID := uqs.userQueryRepository.GetUserByID(id)
	if errGetID != nil {
		return domain.User{}, errors.New(constant.ERROR_DATA_EMPTY)
	}

	return userDomain, nil
}
