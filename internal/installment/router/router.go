package router

import (
	"xyz-finance-api/internal/installment/repository"
	"xyz-finance-api/internal/installment/usecase"
	repositoryTransaction "xyz-finance-api/internal/transaction/repository"
	"xyz-finance-api/pkg/middleware"
	repositoryUser "xyz-finance-api/internal/user/repository"
	"xyz-finance-api/internal/installment/handler"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InstallmentRoutes(installment *echo.Group, db *gorm.DB) {
	installmentQueryRepository := repository.NewInstallmentQueryRepository(db)
	installmentCommandRepository := repository.NewInstallmentCommandRepository(db)
	transactionQueryRepository := repositoryTransaction.NewTransactionQueryRepository(db)
	userQueryReository := repositoryUser.NewUserQueryRepository(db)

	installmentQueryUsecase := usecase.NewInstallmentQueryUsecase(installmentCommandRepository, installmentQueryRepository, userQueryReository)
	installmentCommandUsecase := usecase.NewInstallmentCommandUsecase(installmentCommandRepository, installmentQueryRepository, transactionQueryRepository, userQueryReository)

	installmentHandler := handler.NewInstallmentHandler(installmentCommandUsecase, installmentQueryUsecase)

	installment.GET("", installmentHandler.GetAllInstallments, middleware.JWTMiddleware(false))
	installment.GET("/:id", installmentHandler.GetInstallmentByID, middleware.JWTMiddleware(false))
	installment.POST("", installmentHandler.CreateInstallment, middleware.JWTMiddleware(false))
	installment.PUT("/:id", installmentHandler.UpdateInstallmentStatusByID, middleware.JWTMiddleware(false))
}
