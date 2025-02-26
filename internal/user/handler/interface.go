package handler

import "github.com/labstack/echo/v4"

type UserHandlerInterface interface {
	// Query
	GetUserByID(c echo.Context) error

	// Command
	RegisterUser(c echo.Context) error
	LoginUser(c echo.Context) error
}
