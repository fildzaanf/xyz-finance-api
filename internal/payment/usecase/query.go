package usecase

import (
	"errors"
	"xyz-finance-api/internal/payment/domain"
	"xyz-finance-api/internal/payment/repository"
	userRepository "xyz-finance-api/internal/user/repository"
	"xyz-finance-api/pkg/constant"
)

type paymentQueryUsecase struct {
	paymentCommandRepository repository.PaymentCommandRepositoryInterface
	paymentQueryRepository   repository.PaymentQueryRepositoryInterface
	userQueryRepository   userRepository.UserQueryRepositoryInterface
}

func NewPaymentQueryUsecase(pqr repository.PaymentQueryRepositoryInterface, pcr repository.PaymentCommandRepositoryInterface, uqr   userRepository.UserQueryRepositoryInterface) PaymentQueryUsecaseInterface {
	return &paymentQueryUsecase{
		paymentQueryRepository:   pqr,
		paymentCommandRepository: pcr,
		userQueryRepository: uqr,
	}
}

func (pqu *paymentQueryUsecase) GetPaymentByID(paymentID, userID string) (domain.Payment, error) {
	if paymentID == "" {
		return domain.Payment{}, errors.New(constant.ERROR_ID_INVALID)
	}

	payment, err := pqu.paymentQueryRepository.GetPaymentByID(paymentID, userID) 
	if err != nil {
		return domain.Payment{}, err
	}

	if payment.ID != paymentID {
		return domain.Payment{}, errors.New(constant.ERROR_ID_NOTFOUND)
	}

	user, err := pqu.userQueryRepository.GetUserByID(userID)
	if err != nil {
		return domain.Payment{}, err
	}

	if user.ID != userID {
		return domain.Payment{}, errors.New(constant.ERROR_ROLE_ACCESS)
	}

	return payment, nil
}

func (pqu *paymentQueryUsecase) GetAllPayments(userID string) ([]domain.Payment, error) {
	if userID == "" {
		return nil, errors.New(constant.ERROR_ID_INVALID)
	}

	payments, err := pqu.paymentQueryRepository.GetAllPayments(userID)
	if err != nil {
		return nil, errors.New(constant.ERROR_DATA_EMPTY)
	}

	user, err := pqu.userQueryRepository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	if user.ID != userID {
		return nil, errors.New(constant.ERROR_ROLE_ACCESS)
	}

	return payments, nil
}
