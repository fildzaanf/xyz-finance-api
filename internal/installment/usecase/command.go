package usecase

import (
	"errors"
	"time"

	"xyz-finance-api/internal/installment/domain"
	"xyz-finance-api/internal/installment/repository"
	repositoryTransaction "xyz-finance-api/internal/transaction/repository"
	"xyz-finance-api/pkg/constant"
	"xyz-finance-api/pkg/validator"
)

type installmentCommandUsecase struct {
	installmentCommandRepository repository.InstallmentCommandRepositoryInterface
	installmentQueryRepository   repository.InstallmentQueryRepositoryInterface
	transactionQueryRepository   repositoryTransaction.TransactionQueryRepositoryInterface
}

func NewInstallmentCommandUsecase(icr repository.InstallmentCommandRepositoryInterface, iqr repository.InstallmentQueryRepositoryInterface, tqr repositoryTransaction.TransactionQueryRepositoryInterface) InstallmentCommandUsecaseInterface {
	return &installmentCommandUsecase{
		installmentCommandRepository: icr,
		installmentQueryRepository:   iqr,
		transactionQueryRepository:   tqr,
	}
}

func (ics *installmentCommandUsecase) CreateInstallment(request domain.Installment) ([]domain.Installment, error) {

	errEmpty := validator.IsDataEmpty([]string{"transaction_id"}, request.TransactionID)
	if errEmpty != nil {
		return nil, errEmpty
	}

	transaction, errTransaction := ics.transactionQueryRepository.GetTransactionByID(request.TransactionID)
	if errTransaction != nil {
		return nil, errors.New(constant.ERROR_ID_NOTFOUND)
	}

	if transaction.Status == "failed" {
		return nil, errors.New("transaction failed")
	}

	if transaction.Tenor <= 0 {
		return nil, errors.New("invalid tenor value")
	}


	existingInstallments, errCount := ics.installmentQueryRepository.CountInstallmentsByTransactionID(request.TransactionID)
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

		installmentEntity, errCreate := ics.installmentCommandRepository.CreateInstallment(installment)
		if errCreate != nil {
			return nil, errCreate
		}

		installments = append(installments, installmentEntity)
	}

	return installments, nil
}

func (icu *installmentCommandUsecase) UpdateInstallmentStatusByID(id string, installment domain.Installment) (domain.Installment, error) {
	existingInstallment, err := icu.installmentQueryRepository.GetInstallmentByID(id)
	if err != nil {
		return domain.Installment{}, errors.New(constant.ERROR_ID_NOTFOUND)
	}

	if installment.Status != "paid" && installment.Status != "unpaid" {
		return domain.Installment{}, errors.New("invalid status update request")
	}

	existingInstallment.Status = installment.Status
	existingInstallment.UpdatedAt = time.Now()

	updatedInstallment, err := icu.installmentCommandRepository.UpdateInstallmentStatusByID(id, existingInstallment)
	if err != nil {
		return domain.Installment{}, err
	}

	return updatedInstallment, nil
}
