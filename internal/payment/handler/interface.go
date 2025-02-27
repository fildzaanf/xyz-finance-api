package handler

import "github.com/labstack/echo/v4"

type TransactionHandlerInterface interface {
	// Command
	CreatePayment(c echo.Context) error

	// Query
	GetAllPayments(c echo.Echo) error
	GetPaymentByID(c echo.Echo) error
}
