package handler

import "github.com/labstack/echo/v4"

type LoanHandlerInterface interface {
	// Query
	GetAllLoans(c echo.Context) error
}
