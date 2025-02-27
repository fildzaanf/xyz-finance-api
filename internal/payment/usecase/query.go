package usecase

import "xyz-finance-api/internal/payment/repository"

type paymentQueryUsecase struct {
	paymentCommandRepository repository.PaymentCommandRepositoryInterface
	paymentQueryRepository   repository.PaymentQueryRepositoryInterface
}

func NewPaymentQueryUsecase(pqr repository.PaymentQueryRepositoryInterface, pcr repository.PaymentCommandRepositoryInterface) PaymentQueryUsecaseInterface {
	return &paymentQueryUsecase{
		paymentQueryRepository:   pqr,
		paymentCommandRepository: pcr,
	}
}
