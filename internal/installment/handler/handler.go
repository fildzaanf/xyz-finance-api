package handler

import (
	"net/http"
	"xyz-finance-api/internal/installment/dto"
	"xyz-finance-api/internal/installment/usecase"
	"xyz-finance-api/internal/installment/domain"
	"xyz-finance-api/pkg/constant"
	"xyz-finance-api/pkg/middleware"
	"xyz-finance-api/pkg/response"

	"github.com/labstack/echo/v4"
)

type installmentHandler struct {
	installmentCommandUsecase usecase.InstallmentCommandUsecaseInterface
	installmentQueryUsecase   usecase.InstallmentQueryUsecaseInterface
}

func NewInstallmentHandler(icu usecase.InstallmentCommandUsecaseInterface, iqu usecase.InstallmentQueryUsecaseInterface) *installmentHandler {
	return &installmentHandler{
		installmentCommandUsecase: icu,
		installmentQueryUsecase:   iqu,
	}
}

// command
func (ih *installmentHandler) CreateInstallment(c echo.Context) error {
	var installmentRequest dto.InstallmentRequest

	if err := c.Bind(&installmentRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
	}

	installmentDomain := dto.InstallmentRequestToInstallmentDomain(installmentRequest)

	createdInstallment, err := ih.installmentCommandUsecase.CreateInstallment(installmentDomain)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
	}

	installmentResponse := dto.ListInstallmentDomainToInstallmentResponse(createdInstallment)
	return c.JSON(http.StatusCreated, response.SuccessResponse(constant.SUCCESS_CREATED, installmentResponse))
}

func (ih *installmentHandler) UpdateInstallmentStatusByID(c echo.Context) error {
	installmentID := c.Param("id")
	if installmentID == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(constant.ERROR_ID_NOTFOUND))
	}

	var installmentUpdate dto.InstallmentUpdateRequest
	if err := c.Bind(&installmentUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
	}

	installment := dto.InstallmentUpdateRequestToInstallmentDomain(installmentUpdate)

	_, errUpdate := ih.installmentCommandUsecase.UpdateInstallmentStatusByID(installmentID, installment)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(errUpdate.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(constant.SUCCESS_UPDATED, nil))
}

// query
func (ih *installmentHandler) GetAllInstallments(c echo.Context) error {
	tokenUserID, role, errExtract := middleware.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(errExtract.Error()))
	}

	if role != constant.USER {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}

	transactionID := c.QueryParam("transaction_id")

	var installments []domain.Installment
	var err error
	if transactionID != "" {
		installments, err = ih.installmentQueryUsecase.GetAllInstallments(tokenUserID, transactionID)
	} else {
		installments, err = ih.installmentQueryUsecase.GetAllInstallments(tokenUserID, "")
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(constant.SUCCESS_RETRIEVED, installments))
}



func (ih *installmentHandler) GetInstallmentByID(c echo.Context) error {
	installmentID := c.Param("id")
	if installmentID == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(constant.ERROR_ID_NOTFOUND))
	}

	installment, err := ih.installmentQueryUsecase.GetInstallmentByID(installmentID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse(constant.ERROR_ID_NOTFOUND))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(constant.SUCCESS_RETRIEVED, installment))
}
