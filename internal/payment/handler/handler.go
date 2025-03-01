package handler

import (
	"encoding/json"
	"fmt"
	"io"
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

func (ph *paymentHandler) MidtransWebhook(c echo.Context) error {
	var notification map[string]interface{}

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request body"))
	}

	if err := json.Unmarshal(body, &notification); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid JSON format"))
	}

	orderID, ok := notification["order_id"].(string)
	if !ok {
		fmt.Println("Missing order_id in request")
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("Missing order_id"))
	}

	transactionStatus, ok := notification["transaction_status"].(string)
	if !ok {
		fmt.Println("Missing transaction_status in request")
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("Missing transaction_status"))
	}

	var newStatus string
	switch transactionStatus {
	case "settlement":
		newStatus = "success"
	case "expire":
		newStatus = "expired"
	case "cancel":
		newStatus = "cancel"
	case "deny":
		newStatus = "deny"
	default:
		newStatus = "pending"
	}

	err = ph.paymentCommandUsecase.UpdatePaymentStatus(orderID, newStatus)
	if err != nil {
		fmt.Println("Failed to update payment status:", err)
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to update payment status"))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse("Payment status updated", nil))
}
