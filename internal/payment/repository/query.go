package repository

import (
	"errors"
	"xyz-finance-api/internal/payment/domain"
	"xyz-finance-api/internal/payment/entity"

	"gorm.io/gorm"
)

type paymentQueryRepository struct {
	db *gorm.DB
}

func NewPaymentQueryRepository(db *gorm.DB) PaymentQueryRepositoryInterface {
	return &paymentQueryRepository{
		db: db,
	}
}

func (pqr *paymentQueryRepository) GetPaymentByInstallmentID(installmentID string) (domain.Payment, error) {
	var payment domain.Payment
	result := pqr.db.Where("installment_id = ?", installmentID).First(&payment)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.Payment{}, errors.New("payment not found")
		}
		return domain.Payment{}, result.Error
	}

	return payment, nil
}

func (pr *paymentQueryRepository) GetPaymentByID(paymentID, userID string) (domain.Payment, error) {
	var payment entity.Payment
	result := pr.db.
		Joins("JOIN installments ON installments.id = payments.installment_id").
		Joins("JOIN transactions ON transactions.id = installments.transaction_id").
		Joins("JOIN loans ON loans.id = transactions.loan_id").
		Where("payments.id = ? AND loans.user_id = ?", paymentID, userID). 
		Select("payments.*").
		First(&payment)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.Payment{}, errors.New("payment not found")
		}
		return domain.Payment{}, result.Error
	}

	return domain.PaymentEntityToPaymentDomain(payment), nil
}

func (pr *paymentQueryRepository) GetAllPayments(userID string) ([]domain.Payment, error) {
	var payments []entity.Payment
	result := pr.db.
		Joins("JOIN installments ON installments.id = payments.installment_id").
		Joins("JOIN transactions ON transactions.id = installments.transaction_id").
		Joins("JOIN loans ON loans.id = transactions.loan_id").
		Where("loans.user_id = ?", userID).
		Select("payments.*").
		Find(&payments)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("no payments found")
		}
		return nil, result.Error
	}

	return domain.ListPaymentEntityToPaymentDomain(payments), nil
}