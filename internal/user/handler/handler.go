package handler

import (
	"net/http"
	"xyz-finance-api/internal/user/usecase"
	"xyz-finance-api/internal/user/dto"
	"xyz-finance-api/pkg/cloud"
	"xyz-finance-api/pkg/constant"
	"xyz-finance-api/pkg/middleware"
	"xyz-finance-api/pkg/response"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userCommandUsecase usecase.UserCommandUsecaseInterface
	userQueryUsecase   usecase.UserQueryUsecaseInterface
}

func NewUserHandler(ucu usecase.UserCommandUsecaseInterface, uqu usecase.UserQueryUsecaseInterface) *userHandler {
	return &userHandler{
		userCommandUsecase: ucu,
		userQueryUsecase:   uqu,
	}
}

// Query
func (uh *userHandler) GetUserByID(c echo.Context) error {
	userIDParam := c.Param("user_id")
	if userIDParam == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(constant.ERROR_ID_NOTFOUND))
	}

	tokenUserID, role, errExtract := middleware.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(errExtract.Error()))
	}

	if role != constant.USER {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}

	if userIDParam != tokenUserID {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}

	user, errGetID := uh.userQueryUsecase.GetUserByID(userIDParam)
	if errGetID != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(errGetID.Error()))
	}

	userResponse := dto.UserDomainToUserDataResponse(user)

	return c.JSON(http.StatusOK, response.SuccessResponse(constant.SUCCESS_RETRIEVED, userResponse))
}

// Command
func (uh *userHandler) RegisterUser(c echo.Context) error {
	userRequest := dto.UserRegisterRequest{}

	errBind := c.Bind(&userRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(errBind.Error()))
	}

	ktpPhoto, errKTP := c.FormFile("ktp_photo")
	if errKTP != nil && errKTP != http.ErrMissingFile {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(constant.ERROR_UPLOAD_IMAGE))
	}

	if ktpPhoto != nil {
		ktpURL, errUpload := cloud.UploadImageToS3(ktpPhoto)
		if errUpload != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse(constant.ERROR_UPLOAD_IMAGE_S3))
		}
		userRequest.KtpPhoto = ktpURL
	}

	selfiePhoto, errFile := c.FormFile("selfie_photo")
	if errFile != nil && errFile != http.ErrMissingFile {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(constant.ERROR_UPLOAD_IMAGE))
	}

	if selfiePhoto != nil {
		selfieURL, errUpload := cloud.UploadImageToS3(selfiePhoto)
		if errUpload != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse(constant.ERROR_UPLOAD_IMAGE_S3))
		}
		userRequest.SelfiePhoto = selfieURL
	}

	userDomain := dto.UserRegisterRequestToUserDomain(userRequest)

	registeredUser, errRegister := uh.userCommandUsecase.RegisterUser(userDomain, ktpPhoto, selfiePhoto)
	if errRegister != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(errRegister.Error()))
	}

	userResponse := dto.UserDomainToUserRegisterResponse(registeredUser)

	return c.JSON(http.StatusCreated, response.SuccessResponse(constant.SUCCESS_REGISTER, userResponse))
}

func (uh *userHandler) LoginUser(c echo.Context) error {
	userRequest := dto.UserLoginRequest{}

	errBind := c.Bind(&userRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(errBind.Error()))
	}

	LoginUser, token, errLogin := uh.userCommandUsecase.LoginUser(userRequest.Email, userRequest.Password)
	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(errLogin.Error()))
	}

	userResponse := dto.UserDomainToUserLoginResponse(LoginUser, token)

	return c.JSON(http.StatusOK, response.SuccessResponse(constant.SUCCESS_LOGIN, userResponse))
}
