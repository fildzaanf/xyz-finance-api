package handler

import "github.com/labstack/echo/v4"

type InstallmentHandlerInterface interface {
	// Command
	CreateInstallment(c echo.Context) error
	UpdateInstallmentByID(c echo.Context) error

	// Query
	GetAllInstallments(c echo.Context) error
	GetInstallmentByID(c echo.Context) error
	GetInstallmentByTransactionID(c echo.Context) error
}
