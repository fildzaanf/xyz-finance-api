package router

import (
	"xyz-finance-api/pkg/middleware"
	"xyz-finance-api/internal/loan/repository"
	"xyz-finance-api/internal/loan/usecase"
	"xyz-finance-api/internal/loan/handler"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	loanQueryRepository := repository.NewLoanQueryRepository(db)
	loanCommandRepository := repository.NewLoanCommandRepository(db)

	loanQueryUsecase := usecase.NewLoanQueryUsecase(loanCommandRepository, loanQueryRepository)
	loanCommandUsecase := usecase.NewLoanCommandUsecase(loanCommandRepository, loanQueryRepository)

	loanHandler := handler.NewLoanHandler(loanCommandUsecase, loanQueryUsecase)

	user := e.Group("/loans")
	user.GET("", loanHandler.GetAllLoans, middleware.JWTMiddleware(false))
}
