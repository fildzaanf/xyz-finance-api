package handler

import (
	"net/http"
	"xyz-finance-api/pkg/constant"
	"xyz-finance-api/pkg/response"
	"xyz-finance-api/internal/loan/usecase"
	"xyz-finance-api/internal/loan/dto"

	"github.com/labstack/echo/v4"
)

type loanHandler struct {
	loanCommandUsecase usecase.LoanCommandUsecaseInterface
	loanQueryUsecase   usecase.LoanQueryUsecaseInterface
}

func NewLoanHandler(lcu usecase.LoanCommandUsecaseInterface, lqu usecase.LoanQueryUsecaseInterface) *loanHandler {
	return &loanHandler{
		loanCommandUsecase: lcu,
		loanQueryUsecase:   lqu,
	}
}

// Query
func (lh *loanHandler) GetAllLoans(c echo.Context) error {
	loans, err := lh.loanQueryUsecase.GetAllLoans()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
	}

	if len(loans) == 0 {
		return c.JSON(http.StatusOK, response.SuccessResponse(constant.ERROR_DATA_EMPTY, nil))
	}

	responses := dto.ListLoanDomainToLoanResponse(loans)

	return c.JSON(http.StatusOK, response.SuccessResponse(constant.SUCCESS_RETRIEVED, responses))
}
