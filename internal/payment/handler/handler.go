package handler

import (
	"net/http"
	"xyz-finance-api/internal/payment/dto"
	"xyz-finance-api/internal/payment/usecase"
	"xyz-finance-api/pkg/constant"
	"xyz-finance-api/pkg/middleware"
	"xyz-finance-api/pkg/response"

	"github.com/labstack/echo/v4"
)

type paymentHandler struct {
	paymentCommandUsecase usecase.PaymentCommandUsecaseInterface
	paymentQueryUsecase   usecase.PaymentQueryUsecaseInterface
}

func NewPaymentHandler(pcu usecase.PaymentCommandUsecaseInterface, pqu usecase.PaymentQueryUsecaseInterface) *paymentHandler {
	return &paymentHandler{
		paymentCommandUsecase: pcu,
		paymentQueryUsecase:   pqu,
	}
}

func (ph *paymentHandler) CreatePayment(c echo.Context) error {

	tokenUserID, role, errExtract := middleware.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(errExtract.Error()))
	}

	if role != constant.USER {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}

	paymentRequest := dto.PaymentRequest{}
	if errBind := c.Bind(&paymentRequest); errBind != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(errBind.Error()))
	}

	paymentDomain := dto.PaymentRequestToPaymentDomain(paymentRequest)

	createdPayment, errCreated := ph.paymentCommandUsecase.CreatePayment(paymentDomain, tokenUserID)
	if errCreated != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(errCreated.Error()))
	}

	paymentResponse := dto.PaymentDomainToPaymentResponse(createdPayment)

	return c.JSON(http.StatusCreated, response.SuccessResponse(constant.SUCCESS_CREATED, paymentResponse))
}
