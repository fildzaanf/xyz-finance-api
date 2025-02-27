package handler

import (
	"net/http"
	"xyz-finance-api/internal/transaction/dto"
	"xyz-finance-api/internal/transaction/usecase"
	"xyz-finance-api/pkg/constant"
	"xyz-finance-api/pkg/middleware"
	"xyz-finance-api/pkg/response"

	"github.com/labstack/echo/v4"
)

type transactionHandler struct {
	transactionCommandUsecase usecase.TransactionCommandUsecaseInterface
	transactionQueryUsecase   usecase.TransactionQueryUsecaseInterface
}

func NewTransactionHandler(tcu usecase.TransactionCommandUsecaseInterface, tqu usecase.TransactionQueryUsecaseInterface) *transactionHandler {
	return &transactionHandler{
		transactionCommandUsecase: tcu,
		transactionQueryUsecase:   tqu,
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

	transactionDomain := dto.TransactionRequestToTransactionDomain(transactionRequest)

	createdTransaction, errCreated := th.transactionCommandUsecase.CreateTransaction(transactionDomain, tokenUserID)
	if errCreated != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(errCreated.Error()))
	}

	transactionResponse := dto.TransactionDomainToTransactionResponse(createdTransaction)

	return c.JSON(http.StatusCreated, response.SuccessResponse(constant.SUCCESS_CREATED_TRANSACTION, transactionResponse))
}

func (th *transactionHandler) GetAllTransactions(c echo.Context) error {
	tokenUserID, role, errExtract := middleware.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(errExtract.Error()))
	}

	if role != constant.USER {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}

	transactions, err := th.transactionQueryUsecase.GetAllTransactions(tokenUserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(constant.SUCCESS_RETRIEVED, transactions))
}

func (th *transactionHandler) GetTransactionByID(c echo.Context) error {
	transactionID := c.Param("id")
	if transactionID == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(constant.ERROR_ID_NOTFOUND))
	}

	transaction, err := th.transactionQueryUsecase.GetTransactionByID(transactionID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse(constant.ERROR_ID_NOTFOUND))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(constant.SUCCESS_RETRIEVED, transaction))
}
