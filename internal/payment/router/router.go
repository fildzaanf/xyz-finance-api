package router

import (
	repositoryInstallment "xyz-finance-api/internal/installment/repository"
	repositoryLoan "xyz-finance-api/internal/loan/repository"
	"xyz-finance-api/internal/payment/repository"
	"xyz-finance-api/internal/payment/usecase"
	"xyz-finance-api/internal/payment/handler"
	"xyz-finance-api/pkg/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func PaymentRoutes(payment *echo.Group, db *gorm.DB) {

	paymentQueryRepository := repository.NewPaymentQueryRepository(db)
	paymentCommandRepository := repository.NewPaymentCommandRepository(db)
	loanCommandRepository := repositoryLoan.NewLoanCommandRepository(db)
	installmentCommandRepository := repositoryInstallment.NewInstallmentCommandRepository(db)
	installmentQueryRepository := repositoryInstallment.NewInstallmentQueryRepository(db)

	paymentQueryUsecase := usecase.NewPaymentQueryUsecase(paymentQueryRepository, paymentCommandRepository)
	paymentCommandUsecase := usecase.NewPaymentCommandUsecase(paymentCommandRepository, paymentQueryRepository, installmentQueryRepository, installmentCommandRepository, loanCommandRepository)

	paymentHandler := handler.NewPaymentHandler(paymentCommandUsecase, paymentQueryUsecase)

	payment.POST("", paymentHandler.CreatePayment, middleware.JWTMiddleware(false)) 
	payment.POST("/midtrans/webhook", paymentHandler.MidtransWebhook)

}


