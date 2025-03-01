package usecase

import (
	"errors"
	"time"

	di "xyz-finance-api/internal/installment/domain"
	repositoryInstallment "xyz-finance-api/internal/installment/repository"
	repositoryLoan "xyz-finance-api/internal/loan/repository"
	"xyz-finance-api/internal/payment/domain"
	"xyz-finance-api/internal/payment/repository"
	"xyz-finance-api/pkg/config"
	"xyz-finance-api/pkg/validator"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type paymentCommandUsecase struct {
	paymentCommandRepository     repository.PaymentCommandRepositoryInterface
	paymentQueryRepository       repository.PaymentQueryRepositoryInterface
	installmentQueryRepository   repositoryInstallment.InstallmentQueryRepositoryInterface
	installmentCommandRepository repositoryInstallment.InstallmentCommandRepositoryInterface
	loanCommandRepository        repositoryLoan.LoanCommandRepositoryInterface
}

func NewPaymentCommandUsecase(pcr repository.PaymentCommandRepositoryInterface, pqr repository.PaymentQueryRepositoryInterface, iqr repositoryInstallment.InstallmentQueryRepositoryInterface, icr repositoryInstallment.InstallmentCommandRepositoryInterface, lcr repositoryLoan.LoanCommandRepositoryInterface) PaymentCommandUsecaseInterface {
	return &paymentCommandUsecase{
		paymentCommandRepository:     pcr,
		paymentQueryRepository:       pqr,
		installmentQueryRepository:   iqr,
		installmentCommandRepository: icr,
		loanCommandRepository:        lcr,
	}
}

func (pcu *paymentCommandUsecase) CreatePayment(payment domain.Payment, userID string) (domain.Payment, error) {
	errEmpty := validator.IsDataEmpty([]string{"installment_id"}, payment.InstallmentID)
	if errEmpty != nil {
		return domain.Payment{}, errEmpty
	}

	installment, errInstallment := pcu.installmentQueryRepository.GetInstallmentByID(payment.InstallmentID)
	if errInstallment != nil {
		return domain.Payment{}, errors.New("installment not found")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		return domain.Payment{}, errors.New("failed to load configuration")
	}

	midtransClient := snap.Client{}
	midtransClient.New(cfg.MIDTRANS.MIDTRANS_SERVER_KEY, midtrans.Sandbox)

	snapRequest := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  installment.ID,
			GrossAmt: int64(installment.Amount),
		},
	}

	snapResponse, errSnap := midtransClient.CreateTransaction(snapRequest)
	if errSnap != nil {
		return domain.Payment{}, errSnap
	}

	payment.GrossAmount = installment.Amount
	payment.PaymentURL = snapResponse.RedirectURL
	payment.Token = snapResponse.Token
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = time.Now()

	paymentEntity, errCreate := pcu.paymentCommandRepository.CreatePayment(payment)
	if errCreate != nil {
		return domain.Payment{}, errCreate
	}

	return paymentEntity, nil
}

func (pcu *paymentCommandUsecase) UpdatePaymentStatus(installmentID, status string) error {
	payment, err := pcu.paymentQueryRepository.GetPaymentByInstallmentID(installmentID)
	if err != nil {
		return errors.New("payment not found")
	}

	switch status {
	case "settlement":
		payment.Status = "success"
	case "expire":
		payment.Status = "expire"
	case "cancel":
		payment.Status = "cancel"
	case "deny":
		payment.Status = "deny"
	default:
		payment.Status = status
	}

	err = pcu.paymentCommandRepository.UpdatePaymentStatus(installmentID, status)
	if err != nil {
		return errors.New("failed to update payment status")
	}

	if payment.Status == "success" {
		_, err = pcu.installmentCommandRepository.UpdateInstallmentStatusByID(payment.InstallmentID, di.Installment{Status: "paid"})
		if err != nil {
			return errors.New("failed to update installment status")
		}


		err = pcu.paymentCommandRepository.UpdateLoanStatus(payment.InstallmentID)
		if err != nil {
			return errors.New("failed to update loan status")
		}
	}

	return nil
}
