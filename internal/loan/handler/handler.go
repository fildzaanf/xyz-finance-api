package handler

import (
	"net/http"
	"xyz-finance-api/internal/loan/usecase"
	"xyz-finance-api/pkg/constant"
	"xyz-finance-api/pkg/middleware"
	"xyz-finance-api/pkg/response"

	"github.com/labstack/echo/v4"
)

type loanHandler struct {
	loanQueryUsecase usecase.LoanQueryUsecaseInterface
}

func NewLoanHandler(lqu usecase.LoanQueryUsecaseInterface) *loanHandler {
	return &loanHandler{
		loanQueryUsecase: lqu,
	}
}

func (lh *loanHandler) GetAllLoans(c echo.Context) error {
	tokenUserID, role, errExtract := middleware.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(errExtract.Error()))
	}

	if role != constant.USER {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}

	loans, err := lh.loanQueryUsecase.GetAllLoans(tokenUserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(constant.SUCCESS_RETRIEVED, loans))
}

func (lh *loanHandler) GetLoanByID(c echo.Context) error {
	loanID := c.Param("loan_id")
	if loanID == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(constant.ERROR_ID_NOTFOUND))
	}

	tokenUserID, role, errExtract := middleware.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(errExtract.Error()))
	}

	if role != constant.USER {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}

	loan, err := lh.loanQueryUsecase.GetLoanByID(loanID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse(constant.ERROR_ID_NOTFOUND))
	}

	if loan.UserID != tokenUserID {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(constant.SUCCESS_RETRIEVED, loan))
}
