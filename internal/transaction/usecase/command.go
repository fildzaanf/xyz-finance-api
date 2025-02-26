package usecase

import (
	"errors"
	"time"
	"xyz-finance-api/internal/transaction/domain"
	"xyz-finance-api/internal/transaction/repository"
	"xyz-finance-api/pkg/constant"
	"xyz-finance-api/pkg/validator"
)

type transactionCommandUsecase struct {
	transactionCommandRepository repository.TransactionCommandRepositoryInterface
	transactionQueryRepository   repository.TransactionQueryRepositoryInterface
}

func NewTransactionCommandUsecase(tcr repository.TransactionCommandRepositoryInterface, tqr repository.TransactionQueryRepositoryInterface) TransactionCommandUsecaseInterface {
	return &transactionCommandUsecase{
		transactionCommandRepository: tcr,
		transactionQueryRepository:   tqr,
	}
}

func (tcs *transactionCommandUsecase) CreateTransaction(transaction domain.Transaction) (domain.Transaction, error) {

	errEmpty := validator.IsDataEmpty([]string{"loan_id", "asset_name", "otr_price"}, transaction.LoanID, transaction.AssetName, transaction.OTRPrice)
	if errEmpty != nil {
		return domain.Transaction{}, errEmpty
	}

	loan, errLoan := tcs.transactionQueryRepository.GetLoanByID(transaction.LoanID)
	if errLoan != nil {
		return domain.Transaction{}, errors.New(constant.ERROR_ID_NOTFOUND)
	}

	interestRate := 0.1
	adminFee := 10000
	interest := int(float64(transaction.OTRPrice) * interestRate * float64(loan.Tenor) / 12)

	totalAmount := transaction.OTRPrice + adminFee + interest
	installmentAmount := totalAmount / loan.Tenor

	transaction.TotalAmount = totalAmount
	transaction.AdminFee = adminFee
	transaction.Interest = interest
	transaction.InstallmentAmount = installmentAmount
	transaction.Status = "success"
	transaction.CreatedAt = time.Now()

	transactionEntity, errCreate := tcs.transactionCommandRepository.CreateTransaction(transaction)
	if errCreate != nil {
		return domain.Transaction{}, errCreate
	}

	return transactionEntity, nil
}
