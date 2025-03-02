package router

import (
	repositoryLoan "xyz-finance-api/internal/loan/repository"
	repositoryUser "xyz-finance-api/internal/user/repository"
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
	loanCommandRepository := repositoryLoan.NewLoanCommandRepository(db)
	userQueryReository := repositoryUser.NewUserQueryRepository(db)

	transactionQueryUsecase := usecase.NewTransactionQueryUsecase(transactionQueryRepository, transactionCommandRepository, userQueryReository)
	transactionCommandUsecase := usecase.NewTransactionCommandUsecase(transactionCommandRepository, transactionQueryRepository, loanQueryRepository, loanCommandRepository)

	transactionHandler := handler.NewTransactionHandler(transactionCommandUsecase, transactionQueryUsecase)

	transaction.POST("", transactionHandler.CreateTransaction, middleware.JWTMiddleware(false))
    transaction.GET("", transactionHandler.GetAllTransactions, middleware.JWTMiddleware(false))
	transaction.GET("/:id", transactionHandler.GetTransactionByID, middleware.JWTMiddleware(false))
}
