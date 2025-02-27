package router

import (
	"xyz-finance-api/pkg/middleware"
	userRepository "xyz-finance-api/internal/user/repository"
	"xyz-finance-api/internal/loan/repository"
	"xyz-finance-api/internal/loan/usecase"
	"xyz-finance-api/internal/loan/handler"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)
func LoanRoutes(loan *echo.Group, db *gorm.DB) {
	loanQueryRepository := repository.NewLoanQueryRepository(db)
	loanCommandRepository := repository.NewLoanCommandRepository(db)
	userQueryRepository := userRepository.NewUserQueryRepository(db)

	loanQueryUsecase := usecase.NewLoanQueryUsecase(loanCommandRepository, loanQueryRepository)
	loanCommandUsecase := usecase.NewLoanCommandUsecase(loanCommandRepository, loanQueryRepository, userQueryRepository)

	loanHandler := handler.NewLoanHandler(loanQueryUsecase, loanCommandUsecase)

	loan.GET("", loanHandler.GetAllLoans, middleware.JWTMiddleware(false))
	loan.GET("/:id", loanHandler.GetLoanByID, middleware.JWTMiddleware(false))
	loan.POST("", loanHandler.CreateLoan,  middleware.JWTMiddleware(false))
	loan.PUT("/:id", loanHandler.UpdateLoanStatusByID, middleware.JWTMiddleware(false))
}
