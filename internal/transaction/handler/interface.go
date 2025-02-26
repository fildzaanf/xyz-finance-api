package handler

import (
	"net/http"
	"xyz-finance-api/internal/transaction/dto"
	"xyz-finance-api/pkg/constant"
	"xyz-finance-api/pkg/middleware"
	"xyz-finance-api/pkg/response"
	"xyz-finance-api/internal/transaction/usecase"
	loanUsecase "xyz-finance-api/internal/loan/usecase"

	"github.com/labstack/echo/v4"
)

type transactionHandler struct {
	transactionCommandUsecase usecase.TransactionCommandUsecaseInterface
	transactionQueryUsecase   usecase.TransactionQueryUsecaseInterface
	loanQueryUsecase          loanUsecase.LoanQueryUsecaseInterface
}

func NewTransactionHandler(tcu usecase.TransactionCommandUsecaseInterface, tqu usecase.TransactionQueryUsecaseInterface, lqu  loanUsecase.LoanQueryUsecaseInterface) *transactionHandler {
	return &transactionHandler{
		transactionCommandUsecase: tcu,
		transactionQueryUsecase:   tqu,
		loanQueryUsecase: lqu,
	}
}

func (th *transactionHandler) CreateTransaction(c echo.Context) error {

	tokenUserID, role, errExtract := middleware.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(errExtract.Error()))
	}

	if role != constant.USER {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}

	transactionRequest := dto.TransactionRequest{}

	errBind := c.Bind(&transactionRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(errBind.Error()))
	}

	loan, err := th.loanQueryUsecase.GetLoanByID(transactionRequest.LoanID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse(constant.ERROR_ID_NOTFOUND))
	}

	if loan.UserID != tokenUserID {
        return c.JSON(http.StatusUnauthorized, response.ErrorResponse(constant.ERROR_ROLE_ACCESS))
    }

	transactionDomain := dto.TransactionRequestToTransactionDomain(transactionRequest)

    createdTransaction, errCreated := th.transactionCommandUsecase.CreateTransaction(transactionDomain)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, response.ErrorResponse(errCreated.Error()))
    }

    transactionResponse := dto.TransactionDomainToTransactionResponse(createdTransaction)
	
    return c.JSON(http.StatusCreated, response.SuccessResponse(constant.SUCCESS_CREATED_TRANSACTION, transactionResponse))
}
