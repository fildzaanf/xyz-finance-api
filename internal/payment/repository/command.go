package repository

import (
	"errors"
	dl "xyz-finance-api/internal/loan/domain"
	"xyz-finance-api/internal/payment/domain"

	"gorm.io/gorm"
)

type paymentCommandRepository struct {
	db *gorm.DB
}

func NewPaymentCommandRepository(db *gorm.DB) PaymentCommandRepositoryInterface {
	return &paymentCommandRepository{
		db: db,
	}
}

func (pcr *paymentCommandRepository) CreatePayment(payment domain.Payment) (domain.Payment, error) {
	paymentEntity := domain.PaymentDomainToPaymentEntity(payment)

	result := pcr.db.Create(&paymentEntity)
	if result.Error != nil {
		return domain.Payment{}, result.Error
	}

	paymentDomain := domain.PaymentEntityToPaymentDomain(paymentEntity)

	return paymentDomain, nil
}
func (pcr *paymentCommandRepository) UpdatePaymentStatus(installmentID, status string) error {
	result := pcr.db.Model(&domain.Payment{}).
		Where("installment_id = ?", installmentID).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no payment record updated")
	}

	return nil
}

func (pcr *paymentCommandRepository) UpdateLoanStatus(installmentID string) error {
	var transactionID string
	err := pcr.db.Table("installments").
		Select("transaction_id").
		Where("id = ?", installmentID).
		Scan(&transactionID).Error
	if err != nil {
		return err
	}

	var unpaidInstallments int64
	err = pcr.db.Table("installments").
		Where("transaction_id = ? AND status != ?", transactionID, "paid").
		Count(&unpaidInstallments).Error
	if err != nil {
		return err
	}

	if unpaidInstallments == 0 {
		var loanID string
		err = pcr.db.Table("transactions").
			Select("loan_id").
			Where("id = ?", transactionID).
			Scan(&loanID).Error
		if err != nil {
			return err
		}

		err = pcr.db.Model(&dl.Loan{}).
			Where("id = ?", loanID).
			Update("status", "valid").Error
		if err != nil {
			return err
		}
	}

	return nil
}
