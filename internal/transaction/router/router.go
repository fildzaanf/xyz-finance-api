package router

import (
	repositoryLoan "xyz-finance-api/internal/loan/repository"
	"xyz-finance-api/internal/transaction/handler"
	"xyz-finance-api/internal/transaction/repository"
	"xyz-finance-api/internal/transaction/usecase"
	"xyz-finance-api/pkg/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TransactionRoutes(transaction *echo.Group, db *gorm.DB) {
	transactionQueryRepository := repository.NewTransactionQueryRepository(db)
	transactionCommandRepository := repository.NewTransactionCommandRepository(db)
	loanQueryRepository := repositoryLoan.NewLoanQueryRepository(db)

	transactionQueryUsecase := usecase.NewTransactionQueryUsecase(transactionCommandRepository, transactionQueryRepository)
	transactionCommandUsecase := usecase.NewTransactionCommandUsecase(transactionCommandRepository, transactionQueryRepository, loanQueryRepository)

	transactionHandler := handler.NewTransactionHandler(transactionCommandUsecase, transactionQueryUsecase)

	transaction.POST("", transactionHandler.CreateTransaction, middleware.JWTMiddleware(false))
}
