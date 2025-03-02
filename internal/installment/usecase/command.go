package usecase

import (
	"errors"
	"time"

	"xyz-finance-api/internal/installment/domain"
	"xyz-finance-api/internal/installment/repository"
	repositoryTransaction "xyz-finance-api/internal/transaction/repository"
	"xyz-finance-api/pkg/constant"
	"xyz-finance-api/pkg/validator"
	userRepository "xyz-finance-api/internal/user/repository"
)

type installmentCommandUsecase struct {
	installmentCommandRepository repository.InstallmentCommandRepositoryInterface
	installmentQueryRepository   repository.InstallmentQueryRepositoryInterface
	transactionQueryRepository   repositoryTransaction.TransactionQueryRepositoryInterface
	userQueryRepository          userRepository.UserQueryRepositoryInterface
}

func NewInstallmentCommandUsecase(icr repository.InstallmentCommandRepositoryInterface, iqr repository.InstallmentQueryRepositoryInterface, tqr repositoryTransaction.TransactionQueryRepositoryInterface, uqr  userRepository.UserQueryRepositoryInterface) InstallmentCommandUsecaseInterface {
	return &installmentCommandUsecase{
		installmentCommandRepository: icr,
		installmentQueryRepository:   iqr,
		transactionQueryRepository:   tqr,
		userQueryRepository: uqr,
	}
}

func (icu *installmentCommandUsecase) CreateInstallment(request domain.Installment, userID string) ([]domain.Installment, error) {

	errEmpty := validator.IsDataEmpty([]string{"transaction_id"}, request.TransactionID)
	if errEmpty != nil {
		return nil, errEmpty
	}

	transaction, errTransaction := icu.transactionQueryRepository.GetTransactionByID(request.TransactionID, userID)
	if errTransaction != nil {
		return nil, errors.New(constant.ERROR_ID_NOTFOUND)
	}


	user, err := icu.userQueryRepository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	if user.ID != userID {
		return nil, errors.New(constant.ERROR_ROLE_ACCESS)
	}

	if transaction.ID != request.TransactionID {
		return nil, errors.New(constant.ERROR_ID_NOTFOUND)
	}
	
	if transaction.Status == "failed" {
		return nil, errors.New("transaction failed")
	}

	if transaction.Tenor <= 0 {
		return nil, errors.New("invalid tenor value")
	}


	existingInstallments, errCount := icu.installmentQueryRepository.CountInstallmentsByTransactionID(request.TransactionID)
	if errCount != nil {
		return nil, errCount
	}

	if existingInstallments >= transaction.Tenor {
		return nil, errors.New("installment count exceeds tenor limit")
	}


	amountPerInstallment := transaction.TotalAmount / transaction.Tenor
	installments := make([]domain.Installment, 0, transaction.Tenor)

	
	for i := 1; i <= transaction.Tenor; i++ {
		createdAt := time.Now()

		dueDate := createdAt.AddDate(0, i, 0)

		installment := domain.Installment{
			TransactionID:     request.TransactionID,
			InstallmentNumber: i,
			Amount:            amountPerInstallment,
			Status:            "unpaid",
			CreatedAt:         createdAt,
			DueDate:           dueDate,
		}

		installmentEntity, errCreate := icu.installmentCommandRepository.CreateInstallment(installment, userID)
		if errCreate != nil {
			return nil, errCreate
		}

		installments = append(installments, installmentEntity)
	}

	return installments, nil
}

func (icu *installmentCommandUsecase) UpdateInstallmentStatusByID(installmentID string, installment domain.Installment, userID string) (domain.Installment, error) {
	existingInstallment, err := icu.installmentQueryRepository.GetInstallmentByID(installmentID, userID)
	if err != nil {
		return domain.Installment{}, errors.New(constant.ERROR_ID_NOTFOUND)
	}

	if existingInstallment.ID != installmentID {
		return domain.Installment{}, errors.New(constant.ERROR_ID_NOTFOUND)
	}

	user, err := icu.userQueryRepository.GetUserByID(userID)
	if err != nil {
		return domain.Installment{}, err
	}

	if user.ID != userID {
		return domain.Installment{}, errors.New(constant.ERROR_ROLE_ACCESS)
	}


	if installment.Status != "paid" && installment.Status != "unpaid" {
		return domain.Installment{}, errors.New("invalid status update request")
	}

	existingInstallment.Status = installment.Status
	existingInstallment.UpdatedAt = time.Now()

	updatedInstallment, err := icu.installmentCommandRepository.UpdateInstallmentStatusByID(installmentID, existingInstallment)
	if err != nil {
		return domain.Installment{}, err
	}

	return updatedInstallment, nil
}
