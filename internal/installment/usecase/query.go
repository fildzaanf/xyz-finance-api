package usecase

import (
	"errors"
	"xyz-finance-api/internal/installment/domain"
	"xyz-finance-api/internal/installment/repository"
	userRepository "xyz-finance-api/internal/user/repository"
	"xyz-finance-api/pkg/constant"
)

type installmentQueryUsecase struct {
	installmentCommandRepository repository.InstallmentCommandRepositoryInterface
	installmentQueryRepository   repository.InstallmentQueryRepositoryInterface
	userQueryRepository          userRepository.UserQueryRepositoryInterface
}

func NewInstallmentQueryUsecase(icr repository.InstallmentCommandRepositoryInterface, iqr repository.InstallmentQueryRepositoryInterface, uqr  userRepository.UserQueryRepositoryInterface) InstallmentQueryUsecaseInterface {
	return &installmentQueryUsecase{
		installmentCommandRepository: icr,
		installmentQueryRepository:   iqr,
		userQueryRepository: uqr,
	}
}

func (iu *installmentQueryUsecase) GetInstallmentByID(installmentID, userID string) (domain.Installment, error) {
    if installmentID == "" {
        return domain.Installment{}, errors.New(constant.ERROR_ID_INVALID)
    }

    installment, err := iu.installmentQueryRepository.GetInstallmentByID(installmentID, userID)
    if err != nil {
        return domain.Installment{}, err
    }

    if installment.TransactionID == "" {
        return domain.Installment{}, errors.New("transactionID cannot be empty")
    }

    if installment.ID != installmentID {
        return domain.Installment{}, errors.New(constant.ERROR_ID_NOTFOUND)
    }

    user, err := iu.userQueryRepository.GetUserByID(userID)
    if err != nil {
        return domain.Installment{}, err
    }

    if user.ID != userID {
        return domain.Installment{}, errors.New(constant.ERROR_ROLE_ACCESS)
    }

    return installment, nil
}

func (iqu *installmentQueryUsecase) GetAllInstallments(userID string) ([]domain.Installment, error) {
	if userID == "" {
		return nil, errors.New(constant.ERROR_ID_INVALID)
	}

	installments, err := iqu.installmentQueryRepository.GetAllInstallments(userID)
	if err != nil {
		return nil, errors.New(constant.ERROR_DATA_EMPTY)
	}

	user, err := iqu.userQueryRepository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	if user.ID != userID {
		return nil, errors.New(constant.ERROR_ROLE_ACCESS)
	}

	return installments, nil
}

