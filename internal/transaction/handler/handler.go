package handler

import "github.com/labstack/echo/v4"

type TransactionHandlerInterface interface {
	// Command
	CreateTransaction(c echo.Context) error
}
