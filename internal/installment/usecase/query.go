package usecase

import (
	"errors"
	"xyz-finance-api/internal/installment/domain"
	"xyz-finance-api/internal/installment/repository"
	"xyz-finance-api/pkg/constant"
)

type installmentQueryUsecase struct {
	installmentCommandRepository repository.InstallmentCommandRepositoryInterface
	installmentQueryRepository   repository.InstallmentQueryRepositoryInterface
}

func NewInstallmentQueryUsecase(icr repository.InstallmentCommandRepositoryInterface, iqr repository.InstallmentQueryRepositoryInterface) InstallmentQueryUsecaseInterface {
	return &installmentQueryUsecase{
		installmentCommandRepository: icr,
		installmentQueryRepository:   iqr,
	}
}

func (iqu *installmentQueryUsecase) GetInstallmentByID(id string) (domain.Installment, error) {
	if id == "" {
		return domain.Installment{}, errors.New(constant.ERROR_ID_INVALID)
	}

	installment, err := iqu.installmentQueryRepository.GetInstallmentByID(id)
	if err != nil {
		return domain.Installment{}, errors.New(constant.ERROR_DATA_EMPTY)
	}

	return installment, nil
}

func (iqu *installmentQueryUsecase) GetAllInstallments(userID, transactionID string) ([]domain.Installment, error) {
	if userID == "" || transactionID == "" {
		return nil, errors.New(constant.ERROR_ID_INVALID)
	}

	installments, err := iqu.installmentQueryRepository.GetAllInstallments(userID, transactionID)
	if err != nil {
		return nil, errors.New(constant.ERROR_DATA_EMPTY)
	}

	return installments, nil
}

