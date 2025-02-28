package handler

import "github.com/labstack/echo/v4"

type LoanHandlerInterface interface {
	// Command
	CreateLoan(c echo.Context) error
	UpdateLoanStatusByID(c echo.Context) error
	// Query
	GetAllLoans(c echo.Context) error
	GetLoanByID(c echo.Context) error
	GetInstallmentByTransactionID(c echo.Context) error
}
