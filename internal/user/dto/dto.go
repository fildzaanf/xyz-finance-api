package dto

import (
	"xyz-finance-api/internal/user/domain"
)

// request
type (
	UserRegisterRequest struct {
		Email           string `json:"email" form:"email"`
		Password        string `json:"password" form:"password"`
		ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
		Nik             string `json:"nik" form:"nik"`
		FullName        string `json:"full_name" form:"full_name"`
		LegalName       string `json:"legal_name" form:"legal_name"`
		BirthPlace      string `json:"birth_place" form:"birth_place"`
		BirthDate       string `json:"birth_date" form:"birth_date"`
		KtpPhoto        string `json:"ktp_photo" form:"ktp_photo"`
		SelfiePhoto     string `json:"selfie_photo" form:"selfie_photo"`
		Salary          int    `json:"salary" form:"salary"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}
)

// response
type (
	UserRegisterResponse struct {
		ID string `json:"id"`
	}

	UserLoginResponse struct {
		ID    string `json:"id"`
		Token string `json:"token"`
	}

	UserDataResponse struct {
		ID          string `json:"id"`
		Nik         string `json:"nik"`
		FullName    string `json:"full_name"`
		LegalName   string `json:"legal_name"`
		BirthPlace  string `json:"birth_place"`
		BirthDate   string `json:"birth_date"`
		KtpPhoto    string `json:"ktp_photo"`
		SelfiePhoto string `json:"selfie_photo"`
		Salary      int    `json:"salary"`
	}
)

// mapper - request
func UserRegisterRequestToUserDomain(request UserRegisterRequest) domain.User {
	return domain.User{
		Email:           request.Email,
		Password:        request.Password,
		ConfirmPassword: request.ConfirmPassword,
		Nik:             request.Nik,
		FullName:        request.FullName,
		LegalName:       request.LegalName,
		BirthPlace:      request.BirthPlace,
		BirthDate:       request.BirthDate,
		KtpPhoto:        request.KtpPhoto,
		SelfiePhoto:     request.SelfiePhoto,
		Salary:          request.Salary,
	}
}

func UserLoginRequestToUserDomain(request UserLoginRequest) domain.User {
	return domain.User{
		Email:    request.Email,
		Password: request.Password,
	}
}

// mapper - response
func UserDomainToUserRegisterResponse(response domain.User) UserRegisterResponse {
	return UserRegisterResponse{
		ID: response.ID,
	}
}

func UserDomainToUserLoginResponse(response domain.User, token string) UserLoginResponse {
	return UserLoginResponse{
		ID:    response.ID,
		Token: token,
	}
}

func UserDomainToUserDataResponse(response domain.User) UserDataResponse {
	return UserDataResponse{
		ID:          response.ID,
		Nik:         response.Nik,
		FullName:    response.FullName,
		LegalName:   response.LegalName,
		BirthPlace:  response.BirthPlace,
		BirthDate:   response.BirthDate,
		KtpPhoto:    response.KtpPhoto,
		SelfiePhoto: response.SelfiePhoto,
		Salary:      response.Salary,
	}
}
