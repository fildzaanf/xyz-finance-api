package usecase

import (
	"errors"
	"fmt"
	"time"
	id "xyz-finance-api/internal/installment/domain"
	"xyz-finance-api/internal/installment/dto"
	repositoryInstallment "xyz-finance-api/internal/installment/repository"
	repositoryLoan "xyz-finance-api/internal/loan/repository"
	"xyz-finance-api/internal/transaction/domain"
	"xyz-finance-api/internal/transaction/repository"
	"xyz-finance-api/pkg/constant"
	"xyz-finance-api/pkg/validator"
)

type transactionCommandUsecase struct {
	transactionCommandRepository repository.TransactionCommandRepositoryInterface
	transactionQueryRepository   repository.TransactionQueryRepositoryInterface
	loanQueryRepository          repositoryLoan.LoanQueryRepositoryInterface
	loanCommandRepository        repositoryLoan.LoanCommandRepositoryInterface
	installmentCommandRepository repositoryInstallment.InstallmentCommandRepositoryInterface
}

func NewTransactionCommandUsecase(tcr repository.TransactionCommandRepositoryInterface, tqr repository.TransactionQueryRepositoryInterface, lqr repositoryLoan.LoanQueryRepositoryInterface, lcr repositoryLoan.LoanCommandRepositoryInterface, icr repositoryInstallment.InstallmentCommandRepositoryInterface) TransactionCommandUsecaseInterface {
	return &transactionCommandUsecase{
		transactionCommandRepository: tcr,
		transactionQueryRepository:   tqr,
		loanQueryRepository:          lqr,
		loanCommandRepository:        lcr,
		installmentCommandRepository: icr,
	}
}

func (tcs *transactionCommandUsecase) CreateTransaction(transaction domain.Transaction, userID string) (domain.Transaction, error) {

	errEmpty := validator.IsDataEmpty([]string{"loan_id", "asset_name", "otr_price"}, transaction.LoanID, transaction.AssetName, transaction.OTRPrice)
	if errEmpty != nil {
		return domain.Transaction{}, errEmpty
	}

	loan, errLoan := tcs.loanQueryRepository.GetLoanByID(transaction.LoanID)
	if errLoan != nil {
		return domain.Transaction{}, errors.New(constant.ERROR_ID_NOTFOUND)
	}

	if loan.UserID != userID {
		return domain.Transaction{}, errors.New(constant.ERROR_ROLE_ACCESS)
	}

	if loan.Status == "invalid" {
		return domain.Transaction{}, errors.New("loan has already been taken")
	}

	if transaction.OTRPrice > loan.LimitAmount {
		return domain.Transaction{}, fmt.Errorf("limit amount exceeded: max allowed is %d", loan.LimitAmount)
	}

	interestRate := 0.1
	adminFee := 10000
	interest := int(float64(transaction.OTRPrice) * interestRate * float64(loan.Tenor) / 12)
	totalAmount := transaction.OTRPrice + adminFee + interest

	transaction.TotalAmount = totalAmount
	transaction.AdminFee = adminFee
	transaction.Interest = interest
	transaction.Status = "success"
	transaction.Tenor = loan.Tenor
	transaction.CreatedAt = time.Now()

	transactionEntity, errCreate := tcs.transactionCommandRepository.CreateTransaction(transaction)
	if errCreate != nil {
		return domain.Transaction{}, errCreate
	}

	loan.Status = "invalid"
	_, errUpdateLoan := tcs.loanCommandRepository.UpdateLoanStatusByID(loan.ID, loan)
	if errUpdateLoan != nil {
		return domain.Transaction{}, errUpdateLoan
	}

	installmentRequest := dto.InstallmentRequest{
		TransactionID: transactionEntity.ID,
	}

	installmentDomain := id.Installment{
		TransactionID: installmentRequest.TransactionID,
	}

	_, errCreateInstallments := tcs.installmentCommandRepository.CreateInstallment(installmentDomain)
	if errCreateInstallments != nil {
		return domain.Transaction{}, errCreateInstallments
	}

	return transactionEntity, nil
}
