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

	tx := pcr.db.Begin()
	if err := tx.Error; err != nil {
		return domain.Payment{}, err
	}

	result := tx.Create(&paymentEntity)
	if result.Error != nil {
		tx.Rollback() 
		return domain.Payment{}, result.Error
	}

	if err := tx.Commit().Error; err != nil {
		return domain.Payment{}, err
	}

	paymentDomain := domain.PaymentEntityToPaymentDomain(paymentEntity)

	return paymentDomain, nil
}
func (pcr *paymentCommandRepository) UpdatePaymentStatus(installmentID, status string) error {
	tx := pcr.db.Begin() 
	if err := tx.Error; err != nil {
		return err
	}

	var payment domain.Payment
	if err := tx.Raw("SELECT * FROM payments WHERE installment_id = ? FOR UPDATE", installmentID).
		Scan(&payment).Error; err != nil {
		tx.Rollback() 
		return err
	}

	result := tx.Model(&domain.Payment{}).
		Where("installment_id = ?", installmentID).
		Update("status", status)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("no payment record updated")
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (pcr *paymentCommandRepository) UpdateLoanStatus(installmentID string) error {
	tx := pcr.db.Begin() 
	if err := tx.Error; err != nil {
		return err
	}

	var transactionID string
	err := tx.Raw("SELECT transaction_id FROM installments WHERE id = ? FOR UPDATE", installmentID).
		Scan(&transactionID).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	var unpaidInstallments int64
	err = tx.Raw("SELECT COUNT(*) FROM installments WHERE transaction_id = ? AND status != 'paid' FOR UPDATE", transactionID).
		Scan(&unpaidInstallments).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if unpaidInstallments == 0 {
		var loanID string
		err = tx.Raw("SELECT loan_id FROM transactions WHERE id = ? FOR UPDATE", transactionID).
			Scan(&loanID).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		err = tx.Model(&dl.Loan{}).
			Where("id = ?", loanID).
			Update("status", "valid").Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
