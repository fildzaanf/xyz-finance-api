package router

import (
	repositoryInstallment "xyz-finance-api/internal/installment/repository"
	repositoryLoan "xyz-finance-api/internal/loan/repository"
	"xyz-finance-api/internal/payment/handler"
	"xyz-finance-api/internal/payment/repository"
	"xyz-finance-api/internal/payment/usecase"
	repositoryUser "xyz-finance-api/internal/user/repository"
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
	userQueryReository := repositoryUser.NewUserQueryRepository(db)

	paymentQueryUsecase := usecase.NewPaymentQueryUsecase(paymentQueryRepository, paymentCommandRepository, userQueryReository)
	paymentCommandUsecase := usecase.NewPaymentCommandUsecase(paymentCommandRepository, paymentQueryRepository, installmentQueryRepository, installmentCommandRepository, loanCommandRepository, userQueryReository)

	paymentHandler := handler.NewPaymentHandler(paymentCommandUsecase, paymentQueryUsecase)

	payment.POST("", paymentHandler.CreatePayment, middleware.JWTMiddleware(false))
	payment.POST("/midtrans/webhook", paymentHandler.MidtransWebhook)
	payment.GET("", paymentHandler.GetAllPayments, middleware.JWTMiddleware(false))
	payment.GET("/:id", paymentHandler.GetPaymentByID, middleware.JWTMiddleware(false))

}
